#   ADR-003: Use static Git-versioned pricing data for the MVP


- Status: Accepted
- Date: 2026-08-06
- Decider: Rodrigue NDE
- Technical area: Cost estimation

##  Context

Vigilo needs to estimate the monthly cost impact of Kubernetes resource requests during pull request validation.

Cloud-provider pricing is complex and may depend on:

- Cloud provider and account agreements.
- Region and availability zone.
- Compute instance family.
- Operating system and tenancy.
- On-demand, Savings Plans, Reserved Instances, or Spot commitments.
- Cluster autoscaler and node-provisioning behavior.
- Node overhead and allocatable capacity.
- Shared cluster costs and allocation methodology.

Calling a cloud pricing API directly from the first CLI version would introduce additional complexity:

- Authentication and credentials management.
- Network access from local development and CI.
- API rate limits and caching.
- Regional catalog normalization.
- Non-reproducible pull request results when cloud pricing changes.
- Unclear responsibility for selecting effective versus list prices.

Vigilo must provide useful, auditable pre-merge estimates without claiming to predict a cloud invoice precisely.

## Decision

The Vigilo MVP will use static, Git-versioned pricing files.

Pricing will be supplied explicitly through a YAML file selected by the user or CI configuration. The pricing file will include:

- Currency.
- Cost per requested CPU core-hour.
- Cost per requested GiB-hour of memory.
- Default hours per month.
- Optional metadata such as cloud provider, region, effective date, and source.

Example:

```yaml
apiVersion: kubebudget.io/v1alpha1
kind: PricingCatalog
metadata:
    name: aws-eu-west-3-reference
spec:
    currency: EUR
    hoursPerMonth: 730
    cpuPerCoreHour: 0.042
    memoryPerGiBHour: 0.005
    metadata:
        provider: aws
        region: eu-west-3
        effectiveDate: "2026-08-01"
        source: "Internal reference rates"
```

The MVP will use the pricing catalog as an explicit policy input. It will not call AWS, Azure, Google Cloud, OpenCost, or third-party pricing APIs by default.

Support for dynamic pricing providers may be added later behind explicit interfaces, caching, documentation, and reproductibility controls.

## Consequences

### Positive consequences

- Pull request results are deterministic for a given commit, manifest, policy, and pricing catalog.
- Pricing assumptions are visible, reviewable, and auditable in Git.
- The CLI can run offline and without cloud credentials.
- CI intergration remains simple and secure.
- Users can model their own effective rates rather than being limited to public list pricing.
- The core policy engine is decoupled from cloud-provider APIs

### Negative consequences

- Users must maintain and update reference prices.
- Initial estimates may diverge from actual billed costs.
- Pricing files based only on CPU and memory cannot represent every cloud billing model.
- Static rates do not automatically account for discounts, Spot interruptions, node-family constraints, node-level overhead.

### Mitigations

- Require pricing catalog metadata, including currency, effective date, region, and region.
- Display the pricing catalog name and rate assumptions in the CLI report.
- Document that Vigilo produces a governance estimate, not an invoice forecast.
- Provide example pricing catalogs only as illustrative templates, not authoritative  cloud price sources.
- Define a `PricingProvider` interface so future dynamic and OpenCost-backed implementations do not change the policy engine API.
- Add pricing-file validation and fail clearly when required values are absent.

## Alternatives considered

### Query the AWS Pricing API during validation

this option was rejected for the MVP because it requires credentials, network access, caching, product catalog mapping, and rate-selection rules. It also makes pull request results harder to reproduce.

### Use an external SaaS pricing service

This option was rejected because it adds a vendor dependency, may require secrets or paid accounts, and conflicts with Vigilo's open-source, self-contained positioning.

### Depend on OpenCost from the first release

THe option was rejected because OpenCost is optimized for runtime allocation and measurement. Vigilo's MVP operates before deployment, where future resource usage and realized allocation are unknown. OpenCost remains a potential integration for later runtime-aware capabilities.

## References

- Vigilo strategic project document.
- OpenCost cost allocation concepts and specification.