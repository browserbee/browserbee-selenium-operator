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

type Browser string

const (
	// BrowserChromium represents the Chrome browser.
	BrowserChromium Browser = "chromium"
	// BrowserFirefox represents the Firefox browser.
	BrowserFirefox Browser = "firefox"
)

// SeleniumNodeSpec defines the desired state of SeleniumNode
type SeleniumNodeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// GridRef is a reference to the Selenium Grid configuration.
	HubRef seleniumhubv1.SeleniumHubRef `json:"hubRef,omitempty"`

	// Browser is the type of browser to run on this node.
	// +kubebuilder:validation:Enum=chromium;firefox
	// +kubebuilder:default=chromium
	Browser Browser `json:"browser,omitempty"`

	// Version is the version of the browser node.
	// +kubebuilder:default="latest"
	Version string `json:"version"`

	// Replicas is the number of node replicas to run for this browser type.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:default=1
	Replicas int32 `json:"replicas,omitempty"`

	// ScreenWidth is the width of the browser window.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:default=1920
	ScreenWidth int32 `json:"screenWidth,omitempty"`

	// ScreenHeight is the height of the browser window.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:default=1080
	ScreenHeight int32 `json:"screenHeight,omitempty"`

	// ScreenDepth is the color depth of the browser window.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:default=24
	ScreenDepth int32 `json:"screenDepth,omitempty"`

	// ScreenDPI is the DPI (dots per inch) of the browser window.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:default=96
	ScreenDPI int32 `json:"screenDPI,omitempty"`

	// MaxSessions is the maximum number of concurrent sessions this node can handle.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:default=1
	MaxSessions int32 `json:"maxSessions,omitempty"`

	// SessionTimeout is the timeout for a session in seconds.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:default=60
	SessionTimeout int32 `json:"sessionTimeout,omitempty"`

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
