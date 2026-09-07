---
name: Feature request
about: Suggest an improvment, integration, policy rule, or new capability
title: "[Feature]: "
labels: enhancement
assignees: ""
---

## Before submitting

Please confirm the following:

- [ ] I searched existing issues and did not find a similiar request.
- [ ] I read the current project scope and roadmap in README.
- [ ] I understand that Vigilo is in early development.
- [ ] I am describing a user problem, not only a preferred implementation.

## Problem statement

<!-- What Kubernetes FinOps, Platform Engineering, SRE, or GitOps problem are you trying to solve? -->

## Who would benefit?

<!-- Examples: Platform Engineers, SREs, FinOps practitioners, application developers, security teams. -->

## Proposed outcome

<!-- Describe the desired result from a user's perspective. -->

## Example workflow

<!-- Describe how the feature would be used in practice. -->

```text
# Example:
# 1. A developer changes a Deployment resource request.
# 2. A pull request is opened.
# 3. Vigilo evaluates the manifest against the team's policy.
# 4. The CI check reports the estimated monthly cost delta.
# 5. The pull request is blocked only if the configured budget threshold is exceeded.
```

## Example Kubernetes manifest

<!-- Optional but strongly recommended for manifest-related requests. -->

```yaml
# Paste a minimal example manifest here.
```

## Example Vigilo policy

<!-- Optional but strongly recommended for policy-related requests. -->

```yaml
# Paste a proposed Vigilo YAML example here.
```

## Alternative considered

<!-- What other approaches, tools, workflows, or workarounds did you consider? -->
<!-- Explain why they do not fully solve the problem. -->

## Scope and priority

<!-- Select the option that best describes this request. -->

- [ ] CLI validation
- [ ] Policy language
- [ ] Cost estimation or pricing catalogs
- [ ] GitHub Actions or CI/CD integration
- [ ] GitLab CI integration
- [ ] Documentation or examples
- [ ] Kubernetes Operator or runtime behavior
- [ ] Metrics or observability
- [ ] Backstage or Internal Developer Platform integration
- [ ] Other

## Would you like to contribute?

- [ ] I can help clarify the use case.
- [ ] I can provide real-world test cases or anonymized examples
- [ ] I can help write documentation.
- [ ] I can implement this feature.
- [ ] I cannot contribute implementation work, but this feature would be valuable to me.

## Additional context

<!-- Add links, screenshots, public references, or any other relevant context. -->