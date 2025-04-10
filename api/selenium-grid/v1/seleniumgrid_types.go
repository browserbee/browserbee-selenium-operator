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
	v1seleniumhub "github.com/browserbee/browserbee-selenium-operator/api/selenium-hub/v1"
	v1seleniumnode "github.com/browserbee/browserbee-selenium-operator/api/selenium-node/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SeleniumGridSpec defines the desired state of SeleniumGrid
type SeleniumGridSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Hub defines the Selenium Hub configuration (image, replicas, ports, etc.).
	Hub v1seleniumhub.SeleniumHubSpec `json:"hub"`
	// Nodes is a list of node configurations (Chrome, Firefox, etc.).
	// This allows you to define multiple node types in a single CR.
	// +kubebuilder:validation:MinItems=1
	Nodes []v1seleniumnode.SeleniumNodeSpec `json:"nodes,omitempty"`
}

// SeleniumGridStatus defines the observed state of SeleniumGrid
type SeleniumGridStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// SeleniumGrid is the Schema for the seleniumgrids API
type SeleniumGrid struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SeleniumGridSpec   `json:"spec,omitempty"`
	Status SeleniumGridStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SeleniumGridList contains a list of SeleniumGrid
type SeleniumGridList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SeleniumGrid `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SeleniumGrid{}, &SeleniumGridList{})
}
