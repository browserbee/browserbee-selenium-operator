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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SeleniumWorkflowSpec defines the desired state of SeleniumWorkflow
type SeleniumWorkflowSpec struct {
	Name        string    `json:"name,omitempty"`
	WebDriver   WebDriver `json:"webDriver,omitempty"`
	Actions     []Action  `json:"actions,omitempty"`
	Description string    `json:"description,omitempty"`
	GridRef     string    `json:"gridRef,omitempty"` // Reference to a resource containing the Selenium Grid URL
}

// WebDriver contains configuration for the Selenium WebDriver
type WebDriver struct {
	Browser      string            `json:"browser,omitempty"`
	Headless     bool              `json:"headless,omitempty"`
	ImplicitWait int32             `json:"implicitWait,omitempty"`
	WindowSize   WindowSize        `json:"windowSize,omitempty"`
	RemoteURL    string            `json:"remoteUrl,omitempty"`
	Capabilities map[string]string `json:"capabilities,omitempty"`
}

// WindowSize specifies the dimensions of the browser window
type WindowSize struct {
	Width  int32 `json:"width,omitempty"`
	Height int32 `json:"height,omitempty"`
}

// Action represents a single step in the workflow
type Action struct {
	Name      string   `json:"name,omitempty"`
	Action    string   `json:"action,omitempty"`
	Target    string   `json:"target,omitempty"`
	Selector  Selector `json:"selector,omitempty"`
	Value     string   `json:"value,omitempty"`
	Expected  string   `json:"expected,omitempty"`
	Condition string   `json:"condition,omitempty"`
	Fallback  Fallback `json:"fallback,omitempty"`
	Retry     Retry    `json:"retry,omitempty"`
}

// Fallback defines alternative selectors for finding elements
type Fallback struct {
	Selectors []Selector `json:"selectors,omitempty"`
}

// Selector specifies a strategy and value for locating elements
type Selector struct {
	Strategy string `json:"strategy,omitempty"`
	Value    string `json:"value,omitempty"`
}

// Retry defines the retry strategy for actions
type Retry struct {
	MaxRetries int32 `json:"maxRetries,omitempty"`
	Delay      int32 `json:"delay,omitempty"`
}

// SeleniumWorkflowStatus defines the observed state of SeleniumWorkflow
type SeleniumWorkflowStatus struct {
	LastScheduleTime   *metav1.Time `json:"lastScheduleTime,omitempty"`
	LastSuccessfulTime *metav1.Time `json:"lastSuccessfulTime,omitempty"`
	LastFailedTime     *metav1.Time `json:"lastFailedTime,omitempty"`
	ExecutionCount     int32        `json:"executionCount,omitempty"`
	SuccessCount       int32        `json:"successCount,omitempty"`
	FailureCount       int32        `json:"failureCount,omitempty"`
	CurrentStatus      string       `json:"currentStatus,omitempty"`
	Message            string       `json:"message,omitempty"`
	LogURL             string       `json:"logUrl,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// SeleniumWorkflow is the Schema for the seleniumworkflows API
type SeleniumWorkflow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SeleniumWorkflowSpec   `json:"spec,omitempty"`
	Status SeleniumWorkflowStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SeleniumWorkflowList contains a list of SeleniumWorkflow
type SeleniumWorkflowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SeleniumWorkflow `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SeleniumWorkflow{}, &SeleniumWorkflowList{})
}
