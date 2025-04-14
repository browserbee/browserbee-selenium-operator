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

// SeleniumStandaloneSpec defines the desired state of SeleniumStandalone
type SeleniumStandaloneSpec struct {
	Image            string            `json:"image,omitempty"`   // e.g. selenium/standalone-chrome
	Version          string            `json:"version,omitempty"` // e.g. 120.0
	Replicas         int32             `json:"replicas,omitempty"`
	BrowserName      string            `json:"browserName,omitempty"`      // e.g. chrome, firefox
	ScreenResolution string            `json:"screenResolution,omitempty"` // e.g. 1920x1080
	Headless         bool              `json:"headless,omitempty"`
	Port             int               `json:"port,omitempty"`        // default 4444
	StartupArgs      []string          `json:"startupArgs,omitempty"` // Additional args to pass to the container
	Env              map[string]string `json:"env,omitempty"`         // Env vars for container
	Resources        ResourceSpec      `json:"resources,omitempty"`   // CPU/Memory limits
}

// ResourceSpec defines the resource specifications.
type ResourceSpec struct {
	Requests ResourceRequirements `json:"requests,omitempty"`
	Limits   ResourceRequirements `json:"limits,omitempty"`
}

// ResourceRequirements defines the resource requirements.
type ResourceRequirements struct {
	CPU    string `json:"cpu,omitempty"`    // e.g. "500m"
	Memory string `json:"memory,omitempty"` // e.g. "512Mi"
}

// SeleniumStandaloneStatus defines the observed state of SeleniumStandalone
type SeleniumStandaloneStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// SeleniumStandalone is the Schema for the seleniumstandalones API
type SeleniumStandalone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SeleniumStandaloneSpec   `json:"spec,omitempty"`
	Status SeleniumStandaloneStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SeleniumStandaloneList contains a list of SeleniumStandalone
type SeleniumStandaloneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SeleniumStandalone `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SeleniumStandalone{}, &SeleniumStandaloneList{})
}
