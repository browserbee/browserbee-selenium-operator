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
	seleniumcommonv1 "github.com/browserbee/browserbee-selenium-operator/api/selenium-common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// LogLevel defines the logging level for the Selenium Hub.
// +kubebuilder:validation:Enum=DEBUG;INFO;WARN;ERROR;FATAL
type LogLevel string

const (
	// LogLevelDebug represents the debug log level.
	LogLevelDebug LogLevel = "DEBUG"
	// LogLevelInfo represents the info log level.
	LogLevelInfo LogLevel = "INFO"
	// LogLevelWarn represents the warning log level.
	LogLevelWarn LogLevel = "WARN"
	// LogLevelError represents the error log level.
	LogLevelError LogLevel = "ERROR"
	// LogLevelFatal represents the fatal log level.
	LogLevelFatal LogLevel = "FATAL"
)

// SeleniumHubRef represents a reference to a Selenium Hub.
type SeleniumHubRef struct {
	// Name is the name of the Selenium Hub.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[a-zA-Z0-9-]+$`
	Name string `json:"name"`
	// Namespace is the namespace where the Selenium Hub is deployed.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[a-zA-Z0-9-]+$`
	Namespace string `json:"namespace"`
}

// SeleniumHubSpec defines the desired state of SeleniumHub
type SeleniumHubSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// GridRef is a reference to the Selenium Grid configuration.
	GridRef seleniumcommonv1.GridRef `json:"gridRef,omitempty"`

	// +kubebuilder:validation:Required
	// +kubebuilder:default="4.21.0"
	Version string `json:"version,omitempty"`

	// EnableTracing enables tracing for the Selenium Hub.
	// +kubebuilder:default=false
	EnableTracing bool `json:"enableTracing,omitempty"`

	// Replicas indicates how many Hub instances to run in parallel.
	// Typically, you only need 1 Hub replica for Selenium Grid,
	// but you can scale horizontally if your use case supports it.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:default=1
	Replicas int32 `json:"replicas,omitempty"`

	// LogLevel specifies the logging level for the Selenium Hub.
	// Valid values are "DEBUG", "INFO", "WARN", "ERROR", and "FATAL".
	LogLevel LogLevel `json:"logLevel,omitempty"`
}

// SeleniumHubStatus defines the observed state of SeleniumHub
type SeleniumHubStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// SeleniumHub is the Schema for the seleniumhubs API
type SeleniumHub struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SeleniumHubSpec   `json:"spec,omitempty"`
	Status SeleniumHubStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SeleniumHubList contains a list of SeleniumHub
type SeleniumHubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SeleniumHub `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SeleniumHub{}, &SeleniumHubList{})
}
