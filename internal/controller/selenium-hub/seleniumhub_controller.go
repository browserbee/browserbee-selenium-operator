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

package seleniumhub

import (
	"context"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/util/retry"
	"reflect"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	seleniumhubv1 "github.com/browserbee/browserbee-selenium-operator/api/selenium-hub/v1"
)

// SeleniumHubReconciler reconciles a SeleniumHub object
type SeleniumHubReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=selenium-hub.browserbee.io,resources=seleniumhubs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=selenium-hub.browserbee.io,resources=seleniumhubs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=selenium-hub.browserbee.io,resources=seleniumhubs/finalizers,verbs=update
//+kubebuilder:rbac:groups="",resources=services;pods,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the SeleniumHub object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.2/pkg/reconcile
func (r *SeleniumHubReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	var hub seleniumhubv1.SeleniumHub
	if err := r.Get(ctx, req.NamespacedName, &hub); err != nil {
		logger.Error(err, "Failed to fetch SeleniumHub")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	logger.Info("Reconciling SeleniumHub", "name", hub.Name)

	if err := r.reconcileDeployment(ctx, &hub); err != nil {
		logger.Error(err, "Failed to reconcile Deployment")
		return ctrl.Result{}, err
	}
	if err := r.reconcileService(ctx, &hub); err != nil {
		logger.Error(err, "Failed to reconcile Service")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func commonLabels(hub *seleniumhubv1.SeleniumHub) map[string]string {
	return map[string]string{
		"app":            hub.Name,
		"grid-name":      hub.Spec.GridRef.Name,
		"grid-namespace": hub.Spec.GridRef.Namespace,
	}
}

func (r *SeleniumHubReconciler) reconcileDeployment(ctx context.Context, hub *seleniumhubv1.SeleniumHub) error {
	replicas := int32(hub.Spec.Replicas)
	labels := commonLabels(hub)

	desired := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      hub.Name,
			Namespace: hub.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{MatchLabels: labels},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{Labels: labels},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "selenium-hub",
							Image: hub.Spec.Image,
							Ports: []corev1.ContainerPort{
								{ContainerPort: 4442},
								{ContainerPort: 4443},
								{ContainerPort: 4444},
							},
						},
					},
				},
			},
		},
	}

	if err := ctrl.SetControllerReference(hub, desired, r.Scheme); err != nil {
		return err
	}

	return applyDeployment(ctx, r.Client, desired)
}

func (r *SeleniumHubReconciler) reconcileService(ctx context.Context, hub *seleniumhubv1.SeleniumHub) error {
	labels := commonLabels(hub)

	desired := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      hub.Name,
			Namespace: hub.Namespace,
		},
		Spec: corev1.ServiceSpec{
			Selector: labels,
			Ports: []corev1.ServicePort{
				{Name: "web", Port: 4444, TargetPort: intstr.FromInt(4444)},
				{Name: "publish", Port: 4442, TargetPort: intstr.FromInt(4442)},
				{Name: "subscribe", Port: 4443, TargetPort: intstr.FromInt(4443)},
			},
			Type: corev1.ServiceTypeClusterIP,
		},
	}

	if err := ctrl.SetControllerReference(hub, desired, r.Scheme); err != nil {
		return err
	}

	return applyService(ctx, r.Client, desired)
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

func applyService(ctx context.Context, c client.Client, desired *corev1.Service) error {
	return retry.RetryOnConflict(retry.DefaultBackoff, func() error {
		var existing corev1.Service
		err := c.Get(ctx, client.ObjectKeyFromObject(desired), &existing)
		if err != nil {
			if errors.IsNotFound(err) {
				return c.Create(ctx, desired)
			}
			return err
		}
		// Only update if needed
		if reflect.DeepEqual(existing.Spec.Ports, desired.Spec.Ports) &&
			reflect.DeepEqual(existing.Spec.Selector, desired.Spec.Selector) &&
			existing.Spec.Type == desired.Spec.Type {
			return nil
		}
		existing.Spec.Ports = desired.Spec.Ports
		existing.Spec.Selector = desired.Spec.Selector
		existing.Spec.Type = desired.Spec.Type
		return c.Update(ctx, &existing)
	})
}

// SetupWithManager sets up the controller with the Manager.
func (r *SeleniumHubReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&seleniumhubv1.SeleniumHub{}).
		Complete(r)
}
