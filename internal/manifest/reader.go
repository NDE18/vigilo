package manifest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8syaml "k8s.io/apimachinery/pkg/util/yaml"
)

// Read decodes supported workload manifests.
//
// Empty documents are skipped. Unsupported resources cause an error.
// If any document fails, no partial result is returned.
func Read(r io.Reader) ([]Worload, error) {
	decoder := k8syaml.NewYAMLOrJSONDecoder(r, 4096)
	var workloads []Worload

	for document := 1; ; document++ {
		var raw json.RawMessage

		if err := decoder.Decode(&raw); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, fmt.Errorf("document %d: decode manifest: %w", document, err)
		}

		if len(raw) == 0 || string(raw) == "null" {
			continue
		}

		// Process the raw JSON message and convert it to a workload
		workload, err := decodeDeployment(raw)
		if err != nil {
			return nil, fmt.Errorf("document %d: %w", document, err)
		}
		workloads = append(workloads, workload)
	}

	if len(workloads) == 0 {
		return nil, errors.New("No workload manifests found")
	}

	return workloads, nil
}

func decodeDeployment(raw json.RawMessage) (Worload, error) {
	var typeMeta metav1.TypeMeta

	if err := json.Unmarshal(raw, &typeMeta); err != nil {
		return Worload{}, fmt.Errorf("Decode resource type: %w", err)
	}

	if typeMeta.APIVersion != "apps/v1" || typeMeta.Kind != "Deployment" {
		return Worload{}, fmt.Errorf(
			"Unsupported resource %q with apiVersion %q: only apps/v1 Deployment is supported",
			typeMeta.Kind,
			typeMeta.APIVersion,
		)
	}

	var deployment appsv1.Deployment

	if err := json.Unmarshal(raw, &deployment); err != nil {
		return Worload{}, fmt.Errorf("Decode Deployment: %w", err)
	}

	return extractDeployment(&deployment)
}

func extractDeployment(deployment *appsv1.Deployment) (Worload, error) {
	if deployment.Name == "" {
		return Worload{}, errors.New("metadata.name is required")
	}

	podsSpec := deployment.Spec.Template.Spec

	if len(podsSpec.Containers) == 0 {
		return Worload{}, fmt.Errorf(
			"Deployment  %q: at least one container is required",
			deployment.Name,
		)
	}

	if len(podsSpec.InitContainers) > 0 {
		return Worload{}, fmt.Errorf(
			"Deployment  %q: init containers are not supported yet",
			deployment.Name,
		)
	}

	if podsSpec.Resources != nil {
		return Worload{}, fmt.Errorf(
			"Deployment  %q: pod-level resources are not supported yet",
			deployment.Name,
		)
	}

	replicas := int32(1)
	replicasDefaulted := deployment.Spec.Replicas == nil

	if deployment.Spec.Replicas != nil {
		replicas = *deployment.Spec.Replicas
	}

	if replicas < 0 {
		return Worload{}, fmt.Errorf(
			"Deployment  %q: replicas must be non-negative",
			deployment.Name,
		)
	}

	containers := make([]Container, 0, len(podsSpec.Containers))

	for _, c := range podsSpec.Containers {
		if c.Name == "" {
			return Worload{}, fmt.Errorf(
				"Deployment  %q: container name is required",
				deployment.Name,
			)
		}

		cpu := requestQuantity(c.Resources.Requests, corev1.ResourceCPU)
		memory := requestQuantity(c.Resources.Requests, corev1.ResourceMemory)

		if cpu != nil && cpu.Sign() < 0 {
			return Worload{}, fmt.Errorf(
				"Deployment  %q: container %q: CPU request must be non-negative",
				deployment.Name,
				c.Name,
			)
		}

		if memory != nil && memory.Sign() < 0 {
			return Worload{}, fmt.Errorf(
				"Deployment  %q: container %q: memory request must be non-negative",
				deployment.Name,
				c.Name,
			)
		}

		containers = append(containers, Container{
			Name:          c.Name,
			CPURequest:    cpu,
			MemoryRequest: memory,
		})
	}

	return Worload{
		APIVersion:        deployment.APIVersion,
		Kind:              deployment.Kind,
		Namespace:         deployment.Namespace,
		Name:              deployment.Name,
		Replicas:          replicas,
		ReplicasDefaulted: replicasDefaulted,
		Containers:        containers,
	}, nil
}

func requestQuantity(
	requests corev1.ResourceList,
	resourceName corev1.ResourceName) *resource.Quantity {
	quantity, exists := requests[resourceName]
	if !exists {
		return nil
	}

	copied := quantity.DeepCopy()
	return &copied
}
