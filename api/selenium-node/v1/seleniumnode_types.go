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
	seleniumhubv1 "github.com/browserbee/browserbee-selenium-operator/api/selenium-hub/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SeleniumNodeSpec defines the desired state of SeleniumNode
type SeleniumNodeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// GridRef is a reference to the Selenium Grid configuration.
	HubRef seleniumhubv1.SeleniumHubRef `json:"hubRef,omitempty"`

	// Image is the container image for this browser node.
	// e.g., "selenium/node-chrome:4.8.3-20230321"
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^.+:.*$`
	Image string `json:"image"`

	// Replicas is the number of node replicas to run for this browser type.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:default=1
	Replicas int32 `json:"replicas,omitempty"`

	// Resources (optional) might define CPU/memory requests & limits (not shown here),
	// environment variables, or other container configs for the node.
	// This is omitted for simplicity, but can be added as needed.
	// Resources corev1.ResourceRequirements `json:"resources,omitempty"`
}

// SeleniumNodeStatus defines the observed state of SeleniumNode
type SeleniumNodeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// SeleniumNode is the Schema for the seleniumnodes API
type SeleniumNode struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SeleniumNodeSpec   `json:"spec,omitempty"`
	Status SeleniumNodeStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SeleniumNodeList contains a list of SeleniumNode
type SeleniumNodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SeleniumNode `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SeleniumNode{}, &SeleniumNodeList{})
}
