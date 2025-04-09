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

package seleniumstandalone

import (
	"context"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/util/retry"
	"reflect"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	seleniumstandalonev1 "github.com/browserbee/browserbee-selenium-operator/api/selenium-standalone/v1"
)

// SeleniumStandaloneReconciler reconciles a SeleniumStandalone object
type SeleniumStandaloneReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=selenium-standalone.browserbee.io,resources=seleniumstandalones,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=selenium-standalone.browserbee.io,resources=seleniumstandalones/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=selenium-standalone.browserbee.io,resources=seleniumstandalones/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the SeleniumStandalone object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.2/pkg/reconcile
func (r *SeleniumStandaloneReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	var standalone seleniumstandalonev1.SeleniumStandalone
	if err := r.Get(ctx, req.NamespacedName, &standalone); err != nil {
		logger.Error(err, "Failed to fetch Selenium standalone node")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	logger.Info("Reconciling SeleniumStandAlone", "name", standalone.Name)
	if err := r.reconcileDeployment(ctx, &standalone); err != nil {
		logger.Error(err, "Failed to reconcile Deployment")
		return ctrl.Result{}, err
	}

	if err := r.reconcileService(ctx, &standalone); err != nil {
		logger.Error(err, "Failed to reconcile Service")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func commonLabels(standalone *seleniumstandalonev1.SeleniumStandalone) map[string]string {
	return map[string]string{
		"app":       standalone.Name,
		"namespace": standalone.Namespace,
	}
}

func (r *SeleniumStandaloneReconciler) reconcileDeployment(ctx context.Context, standalone *seleniumstandalonev1.SeleniumStandalone) error {
	replicas := int32(standalone.Spec.Replicas)
	labels := commonLabels(standalone)

	desired := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      standalone.Name,
			Namespace: standalone.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{MatchLabels: labels},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{Labels: labels},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "selenium-standalone",
							Image: standalone.Spec.Image,
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

	if err := ctrl.SetControllerReference(standalone, desired, r.Scheme); err != nil {
		return err
	}

	return applyDeployment(ctx, r.Client, desired)
}

func (r *SeleniumStandaloneReconciler) reconcileService(ctx context.Context, standalone *seleniumstandalonev1.SeleniumStandalone) error {
	labels := commonLabels(standalone)

	desired := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      standalone.Name,
			Namespace: standalone.Namespace,
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

	if err := ctrl.SetControllerReference(standalone, desired, r.Scheme); err != nil {
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
func (r *SeleniumStandaloneReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&seleniumstandalonev1.SeleniumStandalone{}).
		Complete(r)
}
