---
name: Bug report
about: Report unexpected or incorrect Vigilo behavior
title: "[Bug]: "
labels: bug
assigness: ""
---

## Before submitting

Pleasure confirm the following before submitting a bug report:

- [ ] I searched existing issues issues and did not find a duplicate.
- [ ] I am using a supported Vigilo version or the latest `main` branch.
- [ ] This is not a security vulnerability.
- [ ] I removed secrets, credentials, private URLs, customer data, and confidential pricing information.

> Do not report security vulnerabilities in public issues.
> Please follow the process described in [SECURITY.md](../..SECURITY.md).

## Summary

<!-- Provide a short and clear description of the problem. -->

## Expected behavior

<!-- What did you expect Vigilo to do? -->

## Actual behavior

<!-- What happened instead? Include the complete error output when possible. -->

## Steps to reproduce

<!-- Provide the smallest reproduible sequence of steps. -->

1. 
2. 
3. 

## Command executed

```text
# Example:
# vigilo validate \
#   --policy examples/policies/production.yaml \
#   --manifest examples/manifests/payment-api.yaml \
#   --pricing examples/pricing/aws-eu-west-3.yaml
```

## Kubernetes manifest

<!-- Include the smallest manifest that reproduces the issue. -->
<!-- Remove images, internal registry URLs, secrets, credentials, and confidential metadata. -->

```yaml
# Paste a minimal Kubernetes manifest here.
```

## Vigilo policy

<!-- Include the policy used to reproduce the issue. -->

```yaml
# Paste the VigiloPolicy YAML here.
```

## Pricing catalog

<!-- Include the pricing catalog only if it is relevant to the issue. -->
<!-- Do not include confidential commercial rates. Use anonymized values if needed. -->

```yaml
# Paste the PricingCatalog YAML here.
```

## Environment

| Item | Value |
|---|---|
| Vigilo version or commit | <!-- e.g. v0.1.0, dev, or commit SHA --> |
| Installation method | <!-- e.g. GitHub release, go install, go run, source build --> |
| Operating system | <!-- e.g. Windows 11, macOS 15, Ubuntu 24.04 --> |
| CPU architecture | <!-- e.g. amd64, arm64 --> |
| Go version | <!-- output of `go version`, if built from source --> |
| Kubernetes version | <!-- if relevant --> |
| CI provider | <!-- e.g. GitHub Actions, GitLab CI, local only --> |

## Logs and additional context

<!-- Add logs, screenchots, links to public reproductions, or any additional context. -->

```text
# Paste relevant logs here.
```

## Possible fix

<!-- Optional: If you have an idea for a fix, describe it here. -->
