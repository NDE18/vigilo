## Summary

<!-- Clearly describe what this pull request changes and why. -->

## Related issue

<!-- Link the issue this pull request addresses. -->
<!-- Use "Closes #123" when merging this PR should automatically close the issue. -->

Closes #

## Type of change

<!-- Select all applicable options. -->

- [ ] Bug fix
- [ ] New feature 
- [ ] Policy language change
- [ ] Cost estimation or pricing change
- [ ] CLI behavior change
- [ ] CI/CD or GitHub Actions change
- [ ] Documentation change
- [ ] Test-only change
- [ ] Refactoring
- [ ] Build, dependency or tooling change
- [ ] Breaking change

## What changed?

<!-- Describe the main implementation choices. -->

- 
- 
- 

## User-facing behavior

<!-- Explain how a user experiences this change. -->
<!-- Include example commands, output, manifests, policies, or migration steps when relevant. -->

```text
# Example command or output, if applicable.
```

## Policy schema impact

<!-- Complete this section if the PR changes public YAML policies, pricing catalogs, or output formats. -->

- [ ] No public schema or output contract changes
- [ ] Adds an optional field
- [ ] Adds a new rule or policy capability 
- [ ] Changes existing behavior
- [ ] Introduces a breaking change

<!-- If applicabe, describe the API version, compatibility impact, and migration path. -->

## Testing

<!-- Describe the tests you ran locally -->

- [ ] `go test ./...`
- [ ] `go test -race ./...`
- [ ] `go vet ./...`
- [ ] `golangci-lint run`
- [ ] `go build -trimpath -o bin/vigilo ./cmd/vigilo`
- [ ] Manual CLI validation
- [ ] Added or updated unit tests
- [ ] Added or updated test fixtures
- [ ] Not applicabe

<-- Add relevant commands, results, or edge cases tested. -->

```text
# Paste test commands and relevant output here.
```

### Documentation

- [ ] No documentation changes are required
- [ ] Updated README
- [ ] Updated QUICKSTART
- [ ] Updated policy reference
- [ ] Updated Architecture Decision Record
- [ ] Added release or migration notes

## Checklist

- [ ] My change is focused and does not include unrelated refactoring.
- [ ] I followed the project's code and contribution guideline.
- [ ] I formatted Go code with `gofmt`.
- [ ] I added or updated test for changed behavior.
- [ ] I updated documentation where needed.
- [ ] I removed debug code, temporary files, and unrelated changes.
- [ ] I did not add secrets, credentials, private URLs, or confidential data.
- [ ] I did not introduce unreviewed dependencies.
- [ ] I considered backward compatibility.
- [ ] I confirm that this contribution may be distributed under the Apache License 2.0.

## Screenshots or output

<!-- Optional: Add screeshots, CLI output, or before-and-after behavior. -->

## Additional context

<!-- Add reviewer notes, trade-offs, follow-up work, or known limitations. -->