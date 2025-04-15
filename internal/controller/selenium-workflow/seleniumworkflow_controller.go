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

package seleniumworkflow

import (
	"context"
	"encoding/json"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"

	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	seleniumworkflowv1 "github.com/browserbee/browserbee-selenium-operator/api/selenium-workflow/v1"
)

// SeleniumWorkflowReconciler reconciles a SeleniumWorkflow object
type SeleniumWorkflowReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=selenium-workflow.browserbee.io,resources=seleniumworkflows,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=selenium-workflow.browserbee.io,resources=seleniumworkflows/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=selenium-workflow.browserbee.io,resources=seleniumworkflows/finalizers,verbs=update
//+kubebuilder:rbac:groups=batch,resources=cronjobs,verbs=create;get;list;update;delete;watch
//+kubebuilder:rbac:groups=batch,resources=jobs,verbs=create;get;list;update;delete;watch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the SeleniumWorkflow object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.2/pkg/reconcile
func (r *SeleniumWorkflowReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	// Fetch the SeleniumWorkflow instance
	workflow := &seleniumworkflowv1.SeleniumWorkflow{}
	if err := r.Get(ctx, req.NamespacedName, workflow); err != nil {
		if errors.IsNotFound(err) {
			log.Info("SeleniumWorkflow resource not found. Ignoring since object must be deleted.")
			return ctrl.Result{}, nil
		}
		log.Error(err, "Failed to get SeleniumWorkflow")
		return ctrl.Result{}, err
	}

	// Resolve GridRef to the actual Selenium Grid URL
	if workflow.Spec.GridRef != "" {
		configMap := &v1.ConfigMap{}
		err := r.Get(ctx, client.ObjectKey{Name: workflow.Spec.GridRef, Namespace: workflow.Namespace}, configMap)
		if err != nil {
			log.Error(err, "Failed to fetch ConfigMap for GridRef")
			return ctrl.Result{}, err
		}
		gridURL, exists := configMap.Data["gridUrl"]
		if !exists {
			log.Error(errors.NewNotFound(schema.GroupResource{Group: "", Resource: "ConfigMap"}, "gridUrl key not found"), "Invalid GridRef")
			return ctrl.Result{}, errors.NewNotFound(schema.GroupResource{Group: "", Resource: "ConfigMap"}, "gridUrl key not found")
		}
		workflow.Spec.WebDriver.RemoteURL = gridURL
	}

	// Serialize the SeleniumWorkflow spec to JSON
	workflowSpecJSON, err := json.Marshal(workflow.Spec)
	if err != nil {
		log.Error(err, "Failed to serialize SeleniumWorkflow spec to JSON")
		return ctrl.Result{}, err
	}

	// Define the CronJob
	cronJob := &batchv1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      workflow.Name + "-cronjob",
			Namespace: workflow.Namespace,
		},
		Spec: batchv1.CronJobSpec{
			Schedule: "*/5 * * * *", // Example schedule: every 5 minutes
			JobTemplate: batchv1.JobTemplateSpec{
				Spec: batchv1.JobSpec{
					Template: v1.PodTemplateSpec{
						Spec: v1.PodSpec{
							Containers: []v1.Container{
								{
									Name:  "selenium-test-runner",
									Image: "ghcr.io/browserbee/selenium-test-runner/runner:latest",
									Env: []v1.EnvVar{
										{
											Name:  "SELENIUM_WORKFLOW_SPEC",
											Value: string(workflowSpecJSON),
										},
										{
											Name:  "SELENIUM_GRID_URL",
											Value: workflow.Spec.WebDriver.RemoteURL, // Pass Selenium Grid URL
										},
									},
								},
							},
							RestartPolicy: v1.RestartPolicyOnFailure,
						},
					},
				},
			},
		},
	}

	// Check if the CronJob already exists
	existingCronJob := &batchv1.CronJob{}
	err = r.Get(ctx, client.ObjectKey{Name: cronJob.Name, Namespace: cronJob.Namespace}, existingCronJob)
	if err != nil && errors.IsNotFound(err) {
		// Create the CronJob if it doesn't exist
		if err := r.Create(ctx, cronJob); err != nil {
			log.Error(err, "Failed to create CronJob")
			return ctrl.Result{}, err
		}
		log.Info("CronJob created successfully")
	} else if err != nil {
		log.Error(err, "Failed to get CronJob")
		return ctrl.Result{}, err
	} else {
		// Update the CronJob if it already exists
		existingCronJob.Spec = cronJob.Spec
		if err := r.Update(ctx, existingCronJob); err != nil {
			log.Error(err, "Failed to update existing CronJob")
			return ctrl.Result{}, err
		}
		log.Info("CronJob updated successfully")
	}

	// Update the SeleniumWorkflow status
	workflow.Status.CurrentStatus = "Scheduled"
	workflow.Status.Message = "CronJob has been scheduled."
	if err := r.Status().Update(ctx, workflow); err != nil {
		log.Error(err, "Failed to update SeleniumWorkflow status")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *SeleniumWorkflowReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&seleniumworkflowv1.SeleniumWorkflow{}).
		Complete(r)
}
