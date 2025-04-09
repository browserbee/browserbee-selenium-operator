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
	v1 "github.com/browserbee/browserbee-selenium-operator/api/selenium-test-case/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SeleniumTestSuiteSpec defines the desired state of SeleniumTestSuite
type SeleniumTestSuiteSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	TestCases []v1.SeleniumTestCase `json:"testCases,omitempty"`
}

// SeleniumTestSuiteStatus defines the observed state of SeleniumTestSuite
type SeleniumTestSuiteStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// SeleniumTestSuite is the Schema for the seleniumtestsuites API
type SeleniumTestSuite struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SeleniumTestSuiteSpec   `json:"spec,omitempty"`
	Status SeleniumTestSuiteStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SeleniumTestSuiteList contains a list of SeleniumTestSuite
type SeleniumTestSuiteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SeleniumTestSuite `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SeleniumTestSuite{}, &SeleniumTestSuiteList{})
}
