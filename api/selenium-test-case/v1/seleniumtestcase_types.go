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

type HubRef struct {
	// Name of the hub
	Name string `json:"name,omitempty"`
	// Namespace of the hub
	Namespace string `json:"namespace,omitempty"`
}

// SeleniumTestCaseSpec defines the desired state of SeleniumTestCase
type SeleniumTestCaseSpec struct {
	HubRef      HubRef        `json:"hubRef,omitempty"`
	Schedule    string        `json:"schedule,omitempty"`
	Description string        `json:"description,omitempty"`
	WebDriver   WebDriverSpec `json:"webDriver,omitempty"`
	Steps       []Step        `json:"steps,omitempty"`
}

type WebDriverSpec struct {
	Browser      string            `json:"browser,omitempty"`
	Headless     bool              `json:"headless,omitempty"`
	ImplicitWait int               `json:"implicitWait,omitempty"`
	WindowSize   WindowSize        `json:"windowSize,omitempty"`
	RemoteURL    string            `json:"remoteURL,omitempty"`
	Capabilities map[string]string `json:"capabilities,omitempty"`
}

type WindowSize struct {
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}

type Retries struct {
	Attempts int `json:"attempts,omitempty"`
	Delay    int `json:"delay,omitempty"`
}

type Step struct {
	Name      string      `json:"name,omitempty"`
	Action    string      `json:"action,omitempty"`
	Target    string      `json:"target,omitempty"`
	Selector  *Selector   `json:"selector,omitempty"`
	Fallbacks []*Selector `json:"fallbacks,omitempty"`
	Retries   Retries     `json:"retries,omitempty"`
	Value     string      `json:"value,omitempty"`
	Expected  string      `json:"expected,omitempty"`
	Condition string      `json:"condition,omitempty"`
}

type Selector struct {
	Strategy string `json:"strategy,omitempty"`
	Value    string `json:"value,omitempty"`
}

// SeleniumTestCaseStatus defines the observed state of SeleniumTestCase
type SeleniumTestCaseStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// SeleniumTestCase is the Schema for the seleniumtestcases API
type SeleniumTestCase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SeleniumTestCaseSpec   `json:"spec,omitempty"`
	Status SeleniumTestCaseStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SeleniumTestCaseList contains a list of SeleniumTestCase
type SeleniumTestCaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SeleniumTestCase `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SeleniumTestCase{}, &SeleniumTestCaseList{})
}
