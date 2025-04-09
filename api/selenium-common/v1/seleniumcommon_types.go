package v1

type GridRef struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace,omitempty"`
}
