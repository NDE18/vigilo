# Vigilo examples

This directory contains illustrative Vigilo policy, pricing, and Kubernetes manifest examples.

The Vigilo policy language is currently in the `v1alpha1` design phase.
Examples may change before the first stable release.

## Directory structure

```text
examples/
|- manifests/ # Kubernetes workloads evaluated by Vigilo
|- policies/  # Vigilo FinOps policies
|- pricing/   # Explicit request-based pricing catalogs
```

## Example scenario

The examples model a production `payments-prod` namespace.

- `production-guardrails.yaml` defines resources and monthly-cost guardrails.
- `aws-eu-3-reference.yaml` provides illustrative reference rates.
- `payment-api-allowed.yaml` represents a workload that should comply with the policy.
- `payment-api-denied.yaml` represents a workload violate the policy.

The pricing catalog contains illustrative values only. It is not an AWS invoice, a public cloud price source, or a gurantee of actual cloud costs.

Vigilo calculates a request-based monthly governance estimate. It does not predict an exact cloud invoice.