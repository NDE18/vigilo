# ADR-001: Use Go as the implementation language

- Status: Accepted
- Date: 2026-08-06
- Decider: Rodrigue NDE
- Technical area: Core architecture

## Context

Vigilo is a GitOps-native policy engine for Kubernetes platforms.

Its first delivrable is a command-line interface that validates Kubernetes manifests against declarative cost policies during local development and continuous integration.
Future roadmap items may include a GitHub Action, a Kubernetes Operator, Custom Resource Definitions (CRDs), Prometheus metrics, and integrations with GitOps tools.

The projet needs a language that supports:

- Native interoperability with Kubernetes API types and resource quantities.
- Reliable cross-plaform CLI distribution.
- A strong ecosystem for Kubernetes controllers and tooling.
- Readable, maintainable code for open-source contributors.
- Efficient execution in CI pipelines.
- A future path to implementing a Kubernets Operator

## Decision

Vigilo will use Go as its primary implementation language.

The project will use the Go standard library whenever possible. Dependencies will be introduced only when they provide clear value for Kubernetes manifest parsing, CLI ergonomics, schema validation, or testing.

The initial CLI will be implemented independently from Kubernetes controller frameworks. Controller-related dependencies such as `controller-runtime` will be introduced only when runtime Operator phase begins.

## Consequences

### Positive consequences

- Vigilo can reuse Kubernetes-native types, including resource quantity parsing for CPU and memory requests.
- The project aligns with the language and conventions used throughout the Kubernetes ecosystem.
- A single static binary can be produced for Linux, macOS, and Windows.
- Go is well suited to CI usage because binaries start quickly and do not require a language runtime on the target system.
- Go makes it easier for Kubernetes-oriented contributors to understand and review the project.

### Negative consequences

- Contributors without Go experience face an onboarding cost.
- Go's error-handling style and type system require discipline to keep domain code expresive.
- Kubernetes Go module can increase dependency size and version-management effort.
- The team must establish Go formatting, linting, testing, and dependency-management standarts from the begining.

### Mitigations

- Keep the MVP architecture small, modular, and documented.
- Use `gofmt`, `go vet`, unit tests, and `golangci-lint` in CI.
- Document local development in `CONTIBUTING.md`.
- Limit direct Kubernetes dependencies in the CLI to the APIs required for safe manifest and resource quantity parsing.
- Maintain Architecture Decision Records for major dependency and architecture changes.

## Alternatives considered

### Python

Python offers rapid prototyping and accessible YAML tooling. It was not selected because distributing a dependable CLI is more complex, Kubernetes API integration is less native, and the future Operator would likely require a separate implementation language or framework.

### Typescript / Node.js

Typescript would be convenient for a Backstage plugin and GitHub Actions. It was not selected as the core language because Kubernetes controller and API tooling are stronger and more idiomatic in Go. Typescript may still be used later for a dedicated Backstage plugin if required.

### Rust

Rust offers strong safety guarantees and efficient binaries. It was not selected because the Kubernetes controller ecosystem, contributor familiarity, and reference implementations are currently stronger in Go.

## References

- Kubenetes Go client libraries and API conventions.
