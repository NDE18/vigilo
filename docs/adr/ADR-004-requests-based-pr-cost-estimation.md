#   ADR-004: Estimate pull-request cost from Kubernetes resource rquests

- Status: Accepted
- Date: 2026-08-07
- Decider: Rodrigue NDE
- Technical area: Cost estimation model

## Context

Vigilo operates before a workload is deployed. At pull request time, it can inspect declared Kubernetes manifests but cannot know:

- Actual CPU or memory utilization.
- Future traffic patterns.
- HPA behavior.
- Cluster autoscaler decisions.
- Karpenter or other node-provisioner decisions.
- Node packing efficency.
- Spot availability or effective discounts.
- Runtime cost allocation data.

Vigilo scheduling uses resource requests as an important scheduling input. Requests therefore represent a declared capacity commitment that is actionable before merge and relevant to platform governance.

The project must choose an estimation model is understandable, deterministic, and suitable for pull request validation.

## Decision

For the MVP? vigilo will calculate projected monthly cost from declared Kubernetes resources requests.

The base model is:

```text
monthlyCost =
    replicas x hoursPerMonth x
    (
        requestedCpuCores x cpuPerCoreHour + requestedMemooryGiB x memoryPerGiBHour
    )
```

For a workload with multiple containers, CPU and memory requests will be summed over all regular containers in a pod template.

The initial model will:

- Read `reources.requests.cpu`.
- Read `resources.requests.memory`.
- Read declared replica counts for supported workload types.
- Use a configurable `hoursPerMonth`, defaulting to 730.
- Report CPU and memory cost components separately.
- Make all assumptions visible in the result.

The MVP will initially support `Deployment`. Support for `StatefulSet`, `Job`, and `CronJob` will be added only after their execution and replica semantics are specified and tested.

If a policy requires requests and a container does not define them, validation will report a policy violation rather than silently estimating zero cost.

The MVP will not infer costs from resource limits, historical usage, node instance types, bin-packing models, or runtime allocation dataa.

## Consequences

### Positive consequences

- The estimation model is available before deployment and therefore and supports shift-left FinOps workflows.
- Results are deterministic and easy for developers to understand.
- The model maps directly to values already present in Kubernetes manifests.
- Policy violations can be explained with concrete declared values.
- The calculation is fast and suitable for local execution and CI.
- The model provides a clear baseline form future comparison with runtime allocation.

### Negative consequences

- Estimated costs are not equivalent to actual cloud bills.
- Requests may overestimate costs when workloads are under-utilized or efficiently packed.
- Requests may underestimate costs when overhead, storage, networking, GPU, or other billable resources are material.
- Autoscaling and scheduled jobs need explicit semantics to avoid misleading results.
- A single CPU/memory rate is an intentional simplification of actual node pricing.

### Mitigations

- Label all results as `estimated request-based monthly cost`.
- Include the formula, replica count, resource totals, hours per month, and selected pricing catalog in reports.
- Document exclusions: node overhead, storage, network, GPU, taxes, discounts, usage, and actual allocation.
- Keep the calculator behind a domain interfacte to suport later models, including node-aware, OpenCost-backed, or usage-aware estimation.
- Add test fixtures for millicores, Mi/Gi quantities, multiple containers, missing requests, and multiple replicas.
- Add a future policy option for how to treat  missing replica values and autosclaing-enabled workloads

## Alternatives considered

### Estimate from resource limits

This option was rejected because Kubernetes scheduling is driven primarily by requests. Limits also vary in meaning across workloads and may exaggerate capacity commitment for burstable applications.

### Estimate from observed utilization

this option was rejected for pull request validation because usage data may not exist for a new workload and is not knowable for a proposed code change. Usage-based model are better suited to runtime tooling.

### Estimate from node instance prices and bin packing

This option was rejected for the MVP because it requires assumptions about node families, allocatable capacity, daemonset overhead, scheduling constraints, autoscaling, and placement. It may be valuable in a later advanced estimator.

### Treat missing requests as zero cost

This option was rejected because it would reward incomplete manifests and create a false impression of zero cost. Policies must decide whether missing requests are a warning or a blocking violation.

## References

- Vigilo strategic project document.
- Kubernetes resource request and scheduling concepts.
- OpenCost allocation and resource-cost concepts.