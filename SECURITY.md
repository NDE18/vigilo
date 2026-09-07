# Security Policy

## Supported versions

Vigilo is currently in early development.

Until the first stable release, security fixes are applied to the latest version on the default branch and, when available, the latest published release.

| Version | Supported |
|---|---|
| `main` | Yes |
| Latest release | Yes, when available |
| Older releases | No |
| Unreleased forks | No |

## Reporting a vulnerability

Please do not report security vulnerabilities through public Github issues, pull requests, discussions, social media, or public chat channels.

If GitHub private vulnerability reporting is enabled for this repository, use the **Report a vulnerability** button in the repository's Security section.

If private vulnerability reporting isn't available, contact the maintainers through the security contact published in the repository profile or organisation profile.

When reporting a vulnerability, please include as much as possible the following informations :

- A clear description of the vulnerability.
- The affected Vigilo version, release, branch, or commit SHA.
- The component or file affected.
- Steps to reproduce the issue.
- A proof of concept, if available.
- The potential impact.
- Any suggested mitigation or fix.
- Your preferred contact method.
- Whether and how you would like to be credited

Please remove credentials, tokens, private repository URLs, customer information, internal infrastructure details, and confidential pricing data from the report.

## What to expect?

The maintainers will make a best-effort attempt to:

1. Acknowledge receipt of a valid report within 7 calendar days.
2. Assess the report and determine its impact.
3. Keep the report confidential while it is investigated.
4. Provide periodic status updates when remediation requires more time.
5. Coordinate a fix, test it, and publish a security advisory when appropriate.
6. Credit the reporter in the advisor, unless the reporter requests anonymity.

Response and remediation timelines may vary because Vigilo is maintained as an open-source project. Acknowledgement doesn't guanrantee that a report will be classified as a security vulnerability or that a fix will be immediately available.

## Scope

Potential security issues may include, but aren't limited to:

- Arbitary code execution through malicious manifests, policies, or pricing fiiles.
- Path traversal or unsafe file handling.
- Unsafe handling of secrets, credentials, or environment variables.
- Dependency vulnerabilities affecting release Vigilo binaries.
- CI/CD integration vulnerabilities.
- Supply-chain risks in release, packaging, signing, or distribution workflows.
- Denial-of-service conditions caused by crafted inputs.
- Incorrect authorization behavior in future Kubernetes runtime components.

## Out of scope

The following are generally out of scope unless they demonstrate a direct, material security impact:

- Requests for new security features.
- Theoretical vulnerabilities without a pratical attack path.
- Vulnerabilities in unsupported versions.
- Issues requiring physical access to a maintainer's computer.
- Vulnerabilities on third-party services outside Vigilo's control.
- Social engineering attacks against maintainers or contributors.
- Missing best-practice headers, informational findings, or non-exploitable configuration suggestions.
- Public disclosure before giving maintainers reasonable time to investigate. 

## Security practices

Vigilo aims to adopt the following security practices as the project matures:

- Dependency updates through automated tooling.
- Continuous integration checks.
- Static analysis and linting.
- Secret scanning and push protection.
- Review requirements for changes to sensitive workflows.
- Reproducible or traceable release processes where pratical.
- Artifact signing for published releases.
- Documented dependency and release provenance.

## Disclosure policy

Please allow maintainers reasonable time to invetigate and address a vulnerability before public disclosure.

The project will coordinate disclosure with the reporter when possible. Public security advisories will describe the issue, affected versions, remediation steps, and acknowledgements where appropriate.

Thank you for helping keep Vigilo and its users secure.
