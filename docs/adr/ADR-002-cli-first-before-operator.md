# ADR-002: Deliver a CLI-first MVP before building a Kubernetes Operator

- Status: Accepted
- Date: 2026-08-05
- Decider: Rodrigue NDE
- Technical area: Product scope and architecture

## Context

Vigilo aims to help Platform Engineers and SREs prevent expensive Kubernetes configurations before they are merged and deployed.

The long-term product vision includes:

- Pull request validation.
- Declarative FinOps policiesd stored in Git.
- Runtime policy evaluation.
- Kubernetes Custom Resources.
- Prometheus metrics and alerts.
- GitOps synchronization through tools such as Argo CD or Flux.
- Optimal Internal Developer Platform integrations.

However, a Kubernetes Operator introduces substantial additional scope:

- CRD API design and versioning.
- Reconciliation logic and indempotency.
- RBAC permissions.
- CLuster installation and upgrade concerns.
- Helm packaging.
- Security reviews.
- Observability and operational support.
- Compatibility testing across Kubernetes versions.

The primary user value must be validated before the project accepts this operational and maintenance burden.

## Decision

Vigilo will follow a CLI-first product strategy.

The first public MVP will be a standalone CLI that:

1. Reads kubernetes manifests from files or directories.
2. Reads a versioned Vigilo policy file.
3. Estimates request-based monthly costs using and explicit pricing configuration.
4. Evaluates policy violations.
5. Produces human-readable and machine-readable results.
6. Retruns a non-zero exit code when blocking violations are found.

The CLI will be designed to run:

- On a developer workstation.
- In Github Actions.
- In GitLab CI.
- In generic CI/CD ssytems.
- As a pre-commit or repositoory validation command.

Tkhe Kubernetes Operator, CRDs, Helm chart, runtime scanning, and Promrtheus metrics are explicitly deferred until the CLI MVP has demonstrated value and stable policy semantics.

## Consequences

### Positive consequences

- The MVP can be installed and evaluated without cluster access, RBAC privileges, or Helm.
- Early adopters can integrate Vigilo into existing CI workflows quickly.
- The core policy model and cost-estimation logic can be tested independently from Kubernetes runtime concerns.
- The project can validate user demand before committing to controller maintenance.
- Security and operational risk are susbtantially lower in the first release.
- GitHub Actions and GitLab CI integrations can reuse the same CLI binary.

### Negatives consequences

- Vigilo will not initially detect drift in workloads already deployed to cluster.
- The first releases will not provide runtime alerts or Prometheus metrics.
- Policies will not initially be represented as Kubernetes-native CRDs.
- Users expecting admission-time blocking inside the cluster will need to use CI enforcement or wait for a later runtime phase.

### Mitigations

- Keep policy files compatibme with future Kubernetes resource conventions:
  `apiVersion`, `kind`, `metadata`, and `spec`.
- Design the domain model independently from Cobra commands and file-system concerns.
- Publish a clear roadmap distinguishing CLI validation from future runtime enforcement.
- Add structured output early so CI providers and future controllers can consume the same evaluation results.
- Gather user feedback before defining CRDs and admission or contoller behavior.

## Alternatives considered

### Build the Operator first

This approach was rejected because it increases timme-to-first-value, requires a cluster for evaluation, and forces early decisions on RBAC, CRD versioning lifecycle management, and runtime enforcement semantics.

### Build an admission webhook first

THis approach was rejected because it requires in-cluster deployment, certificate management, high-availability considerations, failure policy decisions, and production-grade operational guarantees before the policy language is proven.

### Build a Backstage plugin first

This approach was rejected because a plugin would visualize or invoke policy results but would not estalish the core cost-estimation and enforcement engine. The CLI is the reusable foundation for future integrations.

## Reference

- GitOps and pull-request-based policy validation practices