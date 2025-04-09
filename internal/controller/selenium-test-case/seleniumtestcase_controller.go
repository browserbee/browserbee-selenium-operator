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

package seleniumtestcase

import (
	"context"
	"fmt"
	v1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	seleniumtestcasev1 "github.com/browserbee/browserbee-selenium-operator/api/selenium-test-case/v1"
)

// SeleniumTestCaseReconciler reconciles a SeleniumTestCase object
type SeleniumTestCaseReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=selenium-test-case.browserbee.io,resources=seleniumtestcases,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=selenium-test-case.browserbee.io,resources=seleniumtestcases/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=selenium-test-case.browserbee.io,resources=seleniumtestcases/finalizers,verbs=update
//+kubebuilder:rbac:groups=batch,resources=cronjobs,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the SeleniumTestCase object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.2/pkg/reconcile
// Reconcile creates or updates a CronJob to run the browserbee-runner image
func (r *SeleniumTestCaseReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	var testcase seleniumtestcasev1.SeleniumTestCase
	if err := r.Get(ctx, req.NamespacedName, &testcase); err != nil {
		if kerrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	testcase.Spec.WebDriver.RemoteURL = fmt.Sprintf("http://%s:%s.svc.cluster.local:4444", testcase.Spec.HubRef.Name, testcase.Spec.HubRef.Namespace)

	// 🔄 Marshal the Spec to JSON
	specJSON, err := json.Marshal(testcase.Spec)
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("failed to marshal test case spec: %w", err)
	}

	cronJob := &v1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("run-%s", testcase.Name),
			Namespace: testcase.Namespace,
		},
		Spec: v1.CronJobSpec{
			Schedule: testcase.Spec.Schedule, // can be made configurable later
			JobTemplate: v1.JobTemplateSpec{
				Spec: v1.JobSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							RestartPolicy: corev1.RestartPolicyNever,
							ImagePullSecrets: []corev1.LocalObjectReference{
								{
									Name: "browserbee-ghcr-secret",
								},
							},
							Containers: []corev1.Container{
								{
									Name:            "selenium-test-runner",
									Image:           "ghcr.io/browserbee/selenium-test-runner/runner:latest",
									ImagePullPolicy: corev1.PullAlways,
									Command: []string{
										"./selenium-test-runner",
									},
									Args: []string{
										fmt.Sprintf("--spec=%s", string(specJSON)),
									},
								},
							},
						},
					},
				},
			},
		},
	}

	if err := ctrl.SetControllerReference(&testcase, cronJob, r.Scheme); err != nil {
		return ctrl.Result{}, err
	}

	var existing v1.CronJob
	err = r.Get(ctx, client.ObjectKeyFromObject(cronJob), &existing)
	if err != nil {
		if kerrors.IsNotFound(err) {
			logger.Info("Creating new CronJob", "name", cronJob.Name)
			return ctrl.Result{}, r.Create(ctx, cronJob)
		}
		return ctrl.Result{}, err
	}

	existing.Spec = cronJob.Spec
	return ctrl.Result{}, r.Update(ctx, &existing)
}

// SetupWithManager sets up the controller with the Manager.
func (r *SeleniumTestCaseReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&seleniumtestcasev1.SeleniumTestCase{}).
		Owns(&v1.CronJob{}).
		Complete(r)
}
