package manifest

import "k8s.io/apimachinery/pkg/api/resource"

// Worload describes the supported resource declarations of a workload.
type Worload struct {
	APIVersion        string      `json:"apiVersion"`
	Kind              string      `json:"kind"`
	Name              string      `json:"name"`
	Namespace         string      `json:"namespace"`
	Replicas          int32       `json:"replicas"`
	ReplicasDefaulted bool        `json:"replicasDefaulted"`
	Containers        []Container `json:"containers"`
}

// Container describes resource reuests declared by a regular container.
type Container struct {
	Name          string             `json:"name"`
	CPURequest    *resource.Quantity `json:"cpu,omitempty"`
	MemoryRequest *resource.Quantity `json:"memory,omitempty"`
}

// HasCompleteRequests reports whether every container declares both CPU and memory requests.
func (w Worload) HasCompleteRequests() bool {
	if len(w.Containers) == 0 {
		return false
	}

	for _, c := range w.Containers {
		if c.CPURequest == nil || c.MemoryRequest == nil {
			return false
		}
	}
	return true
}
