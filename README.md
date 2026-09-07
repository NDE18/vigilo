<p align="center">
  <img src="assets/crecerelle.png" alt="Crécerelle" width="400">
</p>

# Vigilo

**GitOps-Native FinOps Policy Engine for Kubernetes**

> FinOps Shift Left. Enforce Cost Policies in Git.
> Vigilo watches your Kubernetes costs like a kestrel watches the ground

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)
[![Status](https://img.shields.io/badge/Status-Early%20Development-orange)]()

Vigilo is an open-source policy engine that brings **cost awarenes** into the Kubernetes development workflow.

Vigilo's first cost model produces a request-based monthly estimate from declared CPU, memory, replica count, and an explicit pricing catalog. It is a governance estimate, not a cloud invoice forecast.

It lets platform engineers and SREs define declarative FinOps policies (budgets, resource limits, cost guardrails) as YAML, version them in git and enforce them:

- **At Pull Request time** (shift-left): block expensive configurations before they are merged
- **At runtime**: alert or enforce when bugets are exceeded (Planned for a later Operator phase)

## Why Vigilo?

Kubernetes FinOps tools focus on **visibility** (who is spending what) or **runtime optimization** (rightsizing). These capabilities are essential, but they often identify costly configuration choices only after deployment.

With Vigilo, teams can keep FinOps guardrails in Git alongside application and platform configuration, review them through pull requests, and use CI to report or block policy violations before merge.

It complements visibility and runtime optimization tools by making FinOps policies reviewable and enforceable in the Git workflow.

## Project Status

Vigilo is in **Phase 0: Preparation**

Current work includes:

- Defining the public policy language.
- Establishing architectural decisions.
- Preparing repository governance and contributor workflows.
- Building the technical foundation for the CLI MVP.

The first usable CLI MVP is planned for phase 1.

## Roadmap

| Component             | Purpose                                       | Target Phase  |
|-----------------------|-----------------------------------------------|---------------|
| CLI (`vigilo`)        | Validate manifests against policies           | Phase 1       |
| GitHub Action         | PR comments + optional merge blocking         | Phase 2       |
| Kubernetes Operator   | Runtime enforcement + Prometheus metrics      | Phase 3       |
| Backstage Plugin      | Optional integration with IDPs                | Later         |

## Policy example

> The policy language is currently a draft and may change before the first stable release.

```yaml
apiVersion: vigilo.io/v1alpha1
kind: BudgetPolicy
metadata:
    name: production-guardrails
spec:
    selector:
        namespace: 
            -   production
            -   prod-*

    pricingRef:
        name: aws-eu-west-3-reference

    rules:
        requireRequests: true
        requireLimits: false

        maxPerPod:
            cpu: "4"
            memory: "8Gi"

        replicas:
            max: 10

        monthlyBudget:
            amount: 5000
            currency: USD
```

## Intended Workflow

```text
Developer updates a Kubernetes manifest
                |
                v
Pull request is opened  
                |
                v
Vigilo validates resources and policy compliance
                |
                +--> Advisory mode: report estimated cost and violations
                |
                +--> Strict mode: fail the CI check and block merge                 
```

## Not in the first MVP

- Kubernetes admission webhook
- Kubernetes Operator or CRDs
- Runtime enforcement
- Dynamic cloud pricing API lookups
- Multi-cloud pricing
- Node-level bin-packing simulation
- Backstage plugin

Thes capabilities may be added after the CLI and policy model validated with early users.

## Contributing

Vigilo is in early development. Feedback on the problem statement,policy language, pricing model, and CLI workflow is especially welcome.

Please read [CONTRIBUTING.md](CONTRIBUTING.md) before opening an issue or pull request.

## Security

To report a security vulnerability, please follow [SECURITY.md](SECURITY.md).

## License

Vigilo is licensed under the [Apache License 2.0](License).