package manifest

import (
	"strings"
	"testing"

	"k8s.io/apimachinery/pkg/api/resource"
)

const validDeployment = `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: payment-api
  namespace: payments-prod
spec:
  replicas: 2
  selector:
    matchLabels:
      app: payment-api
  template:
    metadata:
      labels:
        app: payment-api
    spec:
      containers:
        - name: api
          image: example.invalid/payment-api:v1
          resources:
            requests:
              cpu: "500m"
              memory: "512Mi"
`

func TestReadDeployment(t *testing.T) {
	t.Parallel()

	for index, line := range strings.Split(validDeployment, "\n") {
		t.Logf("YAML line %02d: %q", index+1, line)
	}

	workloads, err := Read(strings.NewReader(validDeployment))
	if err != nil {
		t.Fatalf("Read() returned an error: %v", err)
	}

	if len(workloads) != 1 {
		t.Fatalf("Expected 1 workload, got %d", len(workloads))
	}

	workload := workloads[0]

	if workload.Name != "payment-api" {
		t.Errorf(
			"unexpected workload name: got %q, want %q",
			workload.Name,
			"payment-api",
		)
	}

	if workload.Namespace != "payments-prod" {
		t.Errorf(
			"unexpected workload namespace: got %q, want %q",
			workload.Namespace,
			"payments-prod",
		)
	}

	if workload.Replicas != 2 || workload.ReplicasDefaulted {
		t.Fatalf("Unexpected replica information: %+v", workload)
	}

	if len(workload.Containers) != 1 {
		t.Fatalf("Expected 1 container, got %d", len(workload.Containers))
	}

	container := workload.Containers[0]
	wantedCPU := resource.MustParse("500m")
	wantedMemory := resource.MustParse("512Mi")

	if container.CPURequest == nil || container.CPURequest.Cmp(wantedCPU) != 0 {
		t.Fatalf("Unexpected CPU request: got %v, want %v", container.CPURequest, wantedCPU)
	}

	if container.MemoryRequest == nil || container.MemoryRequest.Cmp(wantedMemory) != 0 {
		t.Fatalf("Unexpected memory request: got %v, want %v", container.MemoryRequest, wantedMemory)
	}

	if !workload.HasCompleteRequests() {
		t.Fatal("Expected workload to have complete requests, but it does not")
	}
}

func TestReadReplicaConventions(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     string
		replicas  int32
		defaulted bool
	}{
		{
			name: "Missing replicas default to 1",
			input: strings.Replace(
				validDeployment,
				strings.Repeat(" ", 2)+"replicas: 2\n",
				"",
				1,
			),
			replicas:  1,
			defaulted: true,
		},
		{
			name:      "Explicit zero is preserved",
			input:     strings.Replace(validDeployment, "replicas: 2", "replicas: 0", 1),
			replicas:  0,
			defaulted: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.input == validDeployment {
				t.Fatal("Test setup failed: the manifest was not modified")
			}

			workloads, err := Read(strings.NewReader(tt.input))
			if err != nil {
				t.Fatalf("Read() returned an error: %v", err)
			}

			if len(workloads) != 1 {
				t.Fatalf("Expected 1 workload, got %d", len(workloads))
			}

			workload := workloads[0]

			if workload.Replicas != tt.replicas {
				t.Errorf("Unexpected replicas: got %d, want %d", workload.Replicas, tt.replicas)
			}

			if workload.ReplicasDefaulted != tt.defaulted {
				t.Errorf("Unexpected defaulted status: got %v, want %v", workload.ReplicasDefaulted, tt.defaulted)
			}
		})
	}
}

func TestReadMissingRequests(t *testing.T) {
	t.Parallel()

	input := strings.Replace(validDeployment, `cpu: "500m"`, "", 1)

	workloads, err := Read(strings.NewReader(input))
	if err != nil {
		t.Fatalf("Read() returned an error: %v", err)
	}

	workload := workloads[0]

	if workload.Containers[0].CPURequest != nil {
		t.Fatal("Expected missingCPU request to remain nil")
	}

	if workload.HasCompleteRequests() {
		t.Fatal("Expected workload to not have complete requests, but it does")
	}
}

func TestReadExplicitZeroRequests(t *testing.T) {
	t.Parallel()

	input := strings.Replace(validDeployment, `cpu: "500m"`, `cpu: "0"`, 1)

	workloads, err := Read(strings.NewReader(input))
	if err != nil {
		t.Fatalf("Read() returned an error: %v", err)
	}

	cpu := workloads[0].Containers[0].CPURequest

	if cpu == nil || cpu.Sign() != 0 {
		t.Fatalf("Expected CPU request to be zero, got %v", cpu)
	}
}

func TestReadMultipleContainers(t *testing.T) {
	t.Parallel()

	sidecarLines := []string{
		strings.Repeat(" ", 8) + "- name: sidecar",
		strings.Repeat(" ", 10) + "image: example.invalid/sidecar:1.0.0",
		strings.Repeat(" ", 10) + "resources:",
		strings.Repeat(" ", 12) + "requests:",
		strings.Repeat(" ", 14) + `cpu: "100m"`,
		strings.Repeat(" ", 14) + `memory: "64Mi"`,
	}

	input := strings.TrimRight(validDeployment, "\r\n") +
		"\n" +
		strings.Join(sidecarLines, "\n") +
		"\n"

	workloads, err := Read(strings.NewReader(input))
	if err != nil {
		for index, line := range strings.Split(input, "\n") {
			t.Logf("YAML line %02d: %q", index+1, line)
		}

		t.Fatalf("Read() returned an error: %v", err)
	}

	if len(workloads) != 1 {
		t.Fatalf("expected 1 workload, got %d", len(workloads))
	}

	containers := workloads[0].Containers

	if len(containers) != 2 {
		t.Fatalf("expected 2 containers, got %d", len(containers))
	}

	sidecar := containers[1]

	if sidecar.Name != "sidecar" {
		t.Errorf("unexpected sidecar name: got %q, want %q", sidecar.Name, "sidecar")
	}

	wantCPU := resource.MustParse("100m")
	wantMemory := resource.MustParse("64Mi")

	if sidecar.CPURequest == nil || sidecar.CPURequest.Cmp(wantCPU) != 0 {
		t.Errorf(
			"unexpected sidecar CPU request: got %v, want %v",
			sidecar.CPURequest,
			wantCPU,
		)
	}

	if sidecar.MemoryRequest == nil || sidecar.MemoryRequest.Cmp(wantMemory) != 0 {
		t.Errorf(
			"unexpected sidecar memory request: got %v, want %v",
			sidecar.MemoryRequest,
			wantMemory,
		)
	}
}

func TestReadMultipleDocuments(t *testing.T) {
	t.Parallel()

	input := validDeployment + "\n---\n" + strings.Replace(
		validDeployment,
		"name: payment-api",
		"name: another-api",
		1,
	)

	workloads, err := Read(strings.NewReader(input))
	if err != nil {
		t.Fatalf("Read() returned an error: %v", err)
	}

	if len(workloads) != 2 {
		t.Fatalf("Expected 2 workloads, got %d", len(workloads))
	}
}

func TestReadErrors(t *testing.T) {
	t.Parallel()

	replaceOnce := func(old, replacement string) string {
		t.Helper()

		if count := strings.Count(validDeployment, old); count != 1 {
			t.Fatalf(
				"Invalid test setup: expected exactly one occurrence of %q, got %d",
				old,
				count,
			)
		}

		return strings.Replace(validDeployment, old, replacement, 1)
	}

	containersLine := strings.Repeat(" ", 6) + "containers:"

	withInitContainers := replaceOnce(
		containersLine,
		strings.Join([]string{
			strings.Repeat(" ", 6) + "initContainers:",
			strings.Repeat(" ", 8) + "- name: init",
			strings.Repeat(" ", 10) + "image: example.invalid/init:v1",
			containersLine,
		}, "\n"),
	)

	withPodResources := replaceOnce(
		containersLine,
		strings.Join([]string{
			strings.Repeat(" ", 6) + "resources:",
			strings.Repeat(" ", 8) + "requests:",
			strings.Repeat(" ", 10) + `cpu: "1"`,
			containersLine,
		}, "\n"),
	)

	tests := []struct {
		name    string
		input   string
		message string
	}{
		{
			name:    "Empty input",
			input:   "",
			message: "no workload manifests found",
		},
		{
			name: "Unsupported resource kind",
			input: replaceOnce(
				"kind: Deployment",
				"kind: StatefulSet",
			),
			message: "unsupported resource",
		},
		{
			name: "Unsupported apiVersion",
			input: replaceOnce(
				"apiVersion: apps/v1",
				"apiVersion: v1",
			),
			message: "unsupported resource",
		},
		{
			name: "Negative replicas",
			input: replaceOnce(
				"replicas: 2",
				"replicas: -1",
			),
			message: "replicas must be non-negative",
		},
		{
			name: "Missing name",
			input: replaceOnce(
				strings.Repeat(" ", 2)+"name: payment-api\n",
				"",
			),
			message: "metadata.name is required",
		},
		{
			name: "Negative CPU request",
			input: replaceOnce(
				`cpu: "500m"`,
				`cpu: "-500m"`,
			),
			message: "CPU request must be non-negative",
		},
		{
			name: "Negative memory request",
			input: replaceOnce(
				`memory: "512Mi"`,
				`memory: "-512Mi"`,
			),
			message: "memory request must be non-negative",
		},
		{
			name: "Invalid memory quantity",
			input: replaceOnce(
				`memory: "512Mi"`,
				`memory: "512XYZ"`,
			),
			message: "decode Deployment",
		},
		{
			name:    "Invalid YAML",
			input:   "apiVersion: [",
			message: "decode manifest",
		},
		{
			name:    "Init containers not supported",
			input:   withInitContainers,
			message: "init containers are not supported yet",
		},
		{
			name:    "Unsupported pod-level resources",
			input:   withPodResources,
			message: "pod-level resources are not supported yet",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			workloads, err := Read(strings.NewReader(tt.input))
			if err == nil {
				for index, line := range strings.Split(tt.input, "\n") {
					t.Logf("YAML line %02d: %q", index+1, line)
				}

				t.Fatalf("expected an error, got nil")
			}

			if !strings.Contains(err.Error(), tt.message) {
				for index, line := range strings.Split(tt.input, "\n") {
					t.Logf("YAML line %02d: %q", index+1, line)
				}

				t.Fatalf(
					"unexpected error containing %q, got %q",
					tt.message,
					err.Error(),
				)
			}

			if workloads != nil {
				t.Fatalf("expected no partial results on error")
			}
		})
	}
}
