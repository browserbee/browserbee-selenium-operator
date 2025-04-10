/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package seleniumgrid

import (
	"context"
	"fmt"
	seleniumhubv1 "github.com/browserbee/browserbee-selenium-operator/api/selenium-hub/v1"
	seleniumnodev1 "github.com/browserbee/browserbee-selenium-operator/api/selenium-node/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	seleniumgridv1 "github.com/browserbee/browserbee-selenium-operator/api/selenium-grid/v1"
)

// SeleniumGridReconciler reconciles a SeleniumGrid object
type SeleniumGridReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=selenium-grid.browserbee.io,resources=seleniumgrids,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=selenium-grid.browserbee.io,resources=seleniumgrids/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=selenium-grid.browserbee.io,resources=seleniumgrids/finalizers,verbs=update
//+kubebuilder:rbac:groups=selenium-hub.browserbee.io,resources=seleniumhubs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=selenium-node.browserbee.io,resources=seleniumnodes,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the SeleniumGrid object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.2/pkg/reconcile
func (r *SeleniumGridReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	var grid seleniumgridv1.SeleniumGrid
	if err := r.Get(ctx, req.NamespacedName, &grid); err != nil {
		logger.Error(err, "unable to fetch SeleniumGrid")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	logger.Info("Reconciling SeleniumGrid", "name", grid.Name)

	if err := r.reconcileHubCR(ctx, &grid); err != nil {
		logger.Error(err, "failed to reconcile hub CR")
		return ctrl.Result{}, err
	}

	for i, node := range grid.Spec.Nodes {
		if err := r.reconcileNodeCR(ctx, &grid, node, i); err != nil {
			logger.Error(err, "failed to reconcile node CR", "index", i)
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

func (r *SeleniumGridReconciler) reconcileHubCR(ctx context.Context, grid *seleniumgridv1.SeleniumGrid) error {
	hubCR := &seleniumhubv1.SeleniumHub{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-hub", grid.Name),
			Namespace: grid.Namespace,
		},
		Spec: grid.Spec.Hub,
	}

	if err := ctrl.SetControllerReference(grid, hubCR, r.Scheme); err != nil {
		return err
	}

	var existing seleniumhubv1.SeleniumHub
	err := r.Get(ctx, client.ObjectKeyFromObject(hubCR), &existing)
	if err != nil {
		if errors.IsNotFound(err) {
			return r.Create(ctx, hubCR)
		}
		return err
	}

	existing.Spec = hubCR.Spec
	return r.Update(ctx, &existing)
}

func (r *SeleniumGridReconciler) reconcileNodeCR(ctx context.Context, grid *seleniumgridv1.SeleniumGrid, nodeCfg seleniumnodev1.SeleniumNodeSpec, index int) error {
	nodeCR := &seleniumnodev1.SeleniumNode{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-node-%d", grid.Name, index),
			Namespace: grid.Namespace,
		},
		Spec: seleniumnodev1.SeleniumNodeSpec{
			Version:  nodeCfg.Version,
			Replicas: nodeCfg.Replicas,
			HubRef: seleniumhubv1.SeleniumHubRef{
				Name:      fmt.Sprintf("%s-hub", grid.Name),
				Namespace: grid.Namespace,
			},
		},
	}

	if err := ctrl.SetControllerReference(grid, nodeCR, r.Scheme); err != nil {
		return err
	}

	var existing seleniumnodev1.SeleniumNode
	err := r.Get(ctx, client.ObjectKeyFromObject(nodeCR), &existing)
	if err != nil {
		if errors.IsNotFound(err) {
			return r.Create(ctx, nodeCR)
		}
		return err
	}

	existing.Spec = nodeCR.Spec
	return r.Update(ctx, &existing)
}

// SetupWithManager sets up the controller with the Manager.
func (r *SeleniumGridReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&seleniumgridv1.SeleniumGrid{}).
		Complete(r)
}
