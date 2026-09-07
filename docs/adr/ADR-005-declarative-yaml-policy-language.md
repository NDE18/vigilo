#ADR-005: Use a YAML declarative policy language with Kubernetes-style conventions

- Status: Accepted
- Date: 07-08-2026
- Decisive: Rodrigue NDE
- Technical domain: Public API and policy engine

## Background

Vigilo should allow platform engineers and SREs to express FinOps governance rules which can be:

- Stored and reviewed on Git.
- Applied to local development and CI.
- Read and understand by teams already using Kubernetes and GitOps
- Audited over time.
- Extended without breaking the rules.
- Mapped to future Kubernetes custom resources if an operator is introduced.

The MVP needs a policy format for common rules such as:

- Required resource requests.
- Resource limits required.
- Maximum CPU demand per pod.
- Maximum memory request per pod.
- Maximum number of replicas declared.
- Estimated monthly budget thresholds.
- Namespace targeting.

The format should be accessible to users who are familiar with Kubernetes manifests but may not be political language specialists.

## Decision

Vigilo will define a declarative YAML policy format using Kubernetes-style resource conventions.

The policy documents will contain:

- "apiVersion"
- "nice"
- "metadata"
- "spec"

The first public version of the scheme will be:

```yaml
apiVersion: vigilo.io/v1alpha1
Kind: VigiloPolicy
```

Example:

```yaml
apiVersion: vigilo.io/v1alpha1
Kind: VigiloPolicy
Metadata:
    Name: Payments - Production
:
    Selector:
        Namespaces :
            - Payment production
    princingRef.
        Name: AWS-EU-West-3-Reference

    Rules:
        requireRequests: true
        requireLimits: false

        maxPerPod:
            CPU: "2"
            memory: "4Gi"

        Replies:
            Max: 10

        Budget monthly:
            Amount: 500
            Currency: EUR
```

The MVP policy engine will provide built-in rule types rather than arbitrary expressions that are smashed by the user.

Policy validation will reject:

- Unsupported API versions.
- Unknown required fields.
- Invalid Kubernetes resource amounts.
- Invalid currency or budget values.
- Configuring conflicting rules.
- A policy that cannot be evaluated deterministically.

The public policy schema will be versioned and documented before implementing advanced rules or applying at runtime.

## Consequences

### Positive consequences

- Policies are readable by Kubernetes and GitOps practitioners.
- Policy changes can be viewed via standard pull requests.
- A Kubernetes-like envelope provides a natural path for future migration to CRDs.
- Built-in rules allow for clear error messages and predictable behavior.
- Schema versioning creates an explicit compatibility contract for users.
- YAML policy files can be used consistently by CLI, GitHub Action, GitLab CI, and a future operator.

### Negative consequences

- Built-in rules will initially be less expressive than a general-purpose language such as Rego or CEL.
- Each new rule requires design, implementation, testing, documentation, and version compatibility review.
- YAML can be indentation-sensitive and can become wordy for complex policy sets.
- The early "v1alpha1" pattern may require interrupted changes before reaching "v1".

### Mitigations

- Keep "v1alpha1" deliberately small and mark it as unstable.
- Provide examples, schema reference, and validation errors with accurate file and field locations where possible.
- Use separate policy files for separate teams, namespaces, or environments instead of building an overly complex expression language from the start.
- Design internal evaluation interfaces so that a future CEL, Rego or plugin extension can be added without changing the CLI contract.
- Publish migration notes whenever a schema version changes.
- Do not introduce exceptions, allowlists, approvals or inheritances until the concrete needs of the users have been validated.

## Alternatives considered

### Rego via Open Policy Agent

This option was rejected for the MVP because Rego adds a learning curve and makes cost-based validation logic more difficult to discover and maintain for many platform engineers. The takeover bid remains a possible integration point in the future.

### Only Kyverno policies

This option was rejected because Kyverno is powerful for Kubernetes admission and substantive policy enforcement, but Vigilo needs a domain model dedicated to estimating costs before the merge. The integration of Kyverno can be considered for application to later execution.

### Common Expression Language

This option was rejected for the MVP because the CEL would require defining a stable data model and expression environment before the basic rules of the project had been validated.
It may be appropriate for advanced custom conditions, once the built-in policy catalog is more mature.

### Unversioned custom YAML

This option was rejected because policy files are a public API. Without an explicit version of the API, compatibility, migration, and future support for CRDs would become difficult.

## References

- Vigilo strategic project document.
- Kubernetes API conventions.
- GitOps principles for versioned and declarative configuration.
