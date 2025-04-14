// Package v1 contains API Schema definitions for the selenium-common v1 API group
// +kubebuilder:object:generate=true
// +groupName=selenium-common.browserbee.io
package v1

// GridRef represents a reference to a Selenium Grid.
type GridRef struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace,omitempty"`
}
