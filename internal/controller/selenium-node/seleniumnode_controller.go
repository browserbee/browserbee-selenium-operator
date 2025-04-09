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

package seleniumnode

import (
	"context"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
	"reflect"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	seleniumnodev1 "github.com/browserbee/browserbee-selenium-operator/api/selenium-node/v1"
)

// SeleniumNodeReconciler reconciles a SeleniumNode object
type SeleniumNodeReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=selenium-node.browserbee.io,resources=seleniumnodes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=selenium-node.browserbee.io,resources=seleniumnodes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=selenium-node.browserbee.io,resources=seleniumnodes/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the SeleniumNode object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.2/pkg/reconcile
func (r *SeleniumNodeReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	var node seleniumnodev1.SeleniumNode
	if err := r.Get(ctx, req.NamespacedName, &node); err != nil {
		logger.Error(err, "Unable to fetch SeleniumNode")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	logger.Info("Reconciling SeleniumNode", "name", node.Name)

	if err := r.reconcileDeployment(ctx, &node); err != nil {
		logger.Error(err, "Failed to reconcile Deployment")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func commonLabels(node *seleniumnodev1.SeleniumNode) map[string]string {
	return map[string]string{
		"app":            node.Name,
		"grid-name":      node.Spec.HubRef.Name,
		"grid-namespace": node.Spec.HubRef.Namespace,
	}
}

func (r *SeleniumNodeReconciler) reconcileDeployment(ctx context.Context, node *seleniumnodev1.SeleniumNode) error {
	replicas := int32(node.Spec.Replicas)
	labels := commonLabels(node)

	deploy := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      node.Name,
			Namespace: node.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{MatchLabels: labels},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{Labels: labels},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "selenium-node",
							Image: node.Spec.Image,
							Ports: []corev1.ContainerPort{
								{ContainerPort: 5555},
							},
							Env: []corev1.EnvVar{
								{Name: "SE_EVENT_BUS_HOST", Value: fmt.Sprintf("%s", node.Spec.HubRef.Name)},
								{Name: "SE_EVENT_BUS_PUBLISH_PORT", Value: "4442"},
								{Name: "SE_EVENT_BUS_SUBSCRIBE_PORT", Value: "4443"},
							},
						},
					},
				},
			},
		},
	}

	if err := ctrl.SetControllerReference(node, deploy, r.Scheme); err != nil {
		return err
	}

	return applyDeployment(ctx, r.Client, deploy)
}

func applyDeployment(ctx context.Context, c client.Client, desired *appsv1.Deployment) error {
	return retry.RetryOnConflict(retry.DefaultBackoff, func() error {
		var existing appsv1.Deployment
		err := c.Get(ctx, client.ObjectKeyFromObject(desired), &existing)
		if err != nil {
			if errors.IsNotFound(err) {
				return c.Create(ctx, desired)
			}
			return err
		}
		if reflect.DeepEqual(existing.Spec, desired.Spec) {
			return nil
		}
		existing.Spec = desired.Spec
		return c.Update(ctx, &existing)
	})
}

// SetupWithManager sets up the controller with the Manager.
func (r *SeleniumNodeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&seleniumnodev1.SeleniumNode{}).
		Complete(r)
}
