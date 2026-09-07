# Contributing to Vigilo

Thank you for your interest in contributing to Vigilo.

Vigilo is an open-source, GitOps-native FinOps policy engine for Kubernetes.
Its goal is to help Platform Engineers and SREs define versioned cost guardrails and validate Kubernetes manifests before they are merged and deployed.

The project is in early development. Feedback on the problem, policy language, cost-estimation model, CLI experience, documentation, and architecture is especially valuable at this stage.

By participating in this project, you agree to follow the [Code of Conduct](CODE_OF_CONDUCT.md).

## Ways to contribute

You do not need to write code to contribute. Useful contributions include:

- Reporting bugs or unclear behavior.
- Proposing improvements to the policy language.
- Sharing Kubernetes FinOps use cases.
- Reviewing documentation and examples.
- Adding test fixtures for Kubernetes manifests.
- Improving error messages and CLI output.
- Suggesting integrations with CI/CD or GitOps tools.
- Contributing code, tests, documentation, or examples.

Before starting substancial work, please open an issue or start a discussion to confirm that the proposal aligns with the project roadmap.

## Before you start

Please read the following documents:

- [README.md](README.md) for the project vision and current scope.
- [Architecture Decision Records](docs/adr/) for major technical decisons.
- [SECURITY.md](SECURITY.md) for security vulnerability reporting.
- [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) for community expectations.

Do not use public GitHub issues to report security vulnerabilities. Follow the process described in [SECURITY.md](SECURITY.md) instead.

## Current scope

The first vigilo MVP focuses on a standalone Go CLI that:

- Reads Kubernetes manifests from files or directories
- Reads declarative Vigilo YAML policies.
- Uses explicit, versioned pricing catalogs.
- Produces request-based monthly cost estimates.
- Validates supported policy rules.
- Returns a non-zero exit code for blocking violations.
- Runs locally and in CI environments.

The following items are intentionally out of scope for the first MVP:

- Kubernetes Operators and Custom Resource Definitions.
- Admission webhooks.
- Runtime enforcement and alerting.
- Dynamic cloud pricing API calls.
- Multi-cloud cost modeling.
- Node-level bin-packing simulation.
- Backstage plugins.

Please avoid submitting large features outside this initial scope without first opening an issue and getting maintainer agreement.

## Development setup

### Prerequisites

The initial development environment requires:

- Go 1.22 or later.
- Git.
- GNU Make, if the repository provides a `Makefile`
- Docker and a local Kubernetes environment only when working on future Kubernetes runtime components.

Verify your Go installation:

```bash
go version
```

Clone the repository:

```bash
git clone https://github.com/nde18/vigilo.git
cd vigilo
```

Install dependencies:

```bash
go mod download
```

Run the tests:

```bash
go test ./...
```

Run static checks:

```bash
go vet ./...
```

When configured, also run:

```bash
golangci-lint run
````

## Branching and pull requests

### Branch names

Use short, descriptive branch names:

```text
feat/add-monthly-budget-rule
fix/parse-memory-quantity
docs/improve-policy-example
test/add-deployment-fixtures
chore/update-go-dependencies
```

### Commit messages

Use clear, imperative commit messages:

```text
Add monthly budget policy validation
Fix CPU quantity parsing for millicores
Document request-based cost assumptions
Add tests for missing resource requests
```

Avoid vague messages such as:

```text
Fix stuff
Update code
Changes
WIP
```

### Pull request requirements

Before opening a pull request, ensure that:

- Your branch is based on the latest default branch.
- The code is formatted with `gofmt`.
- Existing and new tests pass locally.
- New bej=havior includes relevant tests.
- Documentaion is updatedt when behavior, configuration, or the public policy schema changes.
- The pull request has a clear title and description.
- The pull request is focused on one logical change.

A pull request description shoukd explain:

- What problem does this change solve?
- What is the proposed solution?
- What behavior changed?
- How was the change tested?
- Are there policy-schema, compatibility, or documentation implications?
- Are there follow-up tasks that intentionally out of scope?

## Go code guidelines

Vigilo follows standards Go conventions.

- Run `gofmt` on all Go source files.
- Prefer clear, small functions over deeply nested logic.
- Return actionable errors with sufficient context.
- Avoid global mutable state.
- Pass `context.Context` where cancellation or external operations may apply.
- Keep domain logic separate from CLI, file system, and output concerns.
- Avoid adding dependencies unless thes provide a clear and documented benefit.
- Do not make network calls, cloud API calls, or shell execution implicit in the policy evaluation

Use Kubernetes resource quantity types when parsing CPU and memory values. Do not implement custom parsing for standard Kubernetes quantities unless there is a documented reason.

## Testing guidelines

All functional changes should include tests.

At a minimum, test:

- Valid and invalid Kubernetes manifests.
- CPU quantities, including cores and millicores.
- Memory quantities, including Mi and Gi values.
- Missing resource requests or limits.
- Multiple containers in Pod template.
- Replica-aware cost estimation.
- Policy violations and expected messages.
- Invalid policy files and invalid pricing catalogs.

Tests should be deterministic and must not depend on live cloud pricing APIs, external services, local clusters, or the current date time.

Use fixtures in `testdata/` for Kubernetes manifests, policies, and pricing catalogs when appropriate.

## Policy language changes

The Vigilo policy format is a public API.

Before changing `VigiloPolicy`, `PricingCatalog`, or other public YAML structures:

1. Open an issue describing the use case and proposed schema.
2. Explain compatibility implications.
3. Add or update examples.
4. Add schema validation and tests.
5. Update the policy reference documentation.
6. Consider whether an Architecture Decision Record is needed.

The initial schema uses `v1alpha1`. Breaking changes may occur while the API is alpha, but they still be documented in release notes and migration guidance.

## Documentation guidelines

Documentation is part of the product.

Please:

- Write documentation in English.
- Use concose and direct language.
- Keep examples valid and copy-pasteable.
- State assumptions and limitations clearly.
- Distinguish current behavior from planned behavior.
- Do not present request-based estimates as cloud invoice forecasts.
- Update `README.md`, `QUICKSTART.md`, or the documentation site when users need new instructions to use a change.

## Reporting bugs

Before opening a bug report:

1. Search existing issues to avoid duplicates
2. Confirm that the issue is not a security vulnerability
3. Provide a minimal reproducible example.

A helpful bug report includes:

- Vigilo version or commit SHA.
- Operating system and architecture
- Go version, when building from source.
- Command excuted.
- Kubernetes manifest, policy, and pricing catalog excerpts.
- Expected behavior.
- Actual behavior.
- Full error output, with secrets removed.

Never include credentials, tokens, internal endpoints, customer data, or confidential pricing information in a public issue.

## Feature requests

Feature requests should describe the user problem before proposing a solution.

Please include:

- The Kubernetes FinOps or Platform Engineering use case.
- The affected users.
- Why current Vigilo behavior is insufficient.
- A concrete example manifest and policy, if possible.
- How the feature fits the projects scope.
- Potential alternatives or trade-offs.

Large features may be discussed in issues before implementation begins.

## License and Contributor Agreement

By submitting a contribution, you certify that:

- You have the right to submit the contribution.
- The contribution is your original work or you have permission to submit it.
- You grant the project the right to use, modify, and distribute your contribution under the [Apache License 2.0](LICENSE).
- You understand that contributions may relicensed only as permitted by the project license and applicable law.

## Questions and discussions

If you have a question, an early design idea, or a use case that is not yet ready for an issue, please use GitHub Discussions when available.

Thank you for helping make Kubernetes cost governance more transparent, collaborative, and actionable.