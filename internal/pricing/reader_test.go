package pricing

import (
	"strings"
	"testing"
)

func validCatalogYAML() string {
	return strings.Join([]string{
		"apiVersion: vigilo.io/v1alpha1",
		"kind: PricingCatalog",
		"metadata:",
		strings.Repeat(" ", 2) + "name: aws-eu-west-3-reference",
		"spec:",
		strings.Repeat(" ", 2) + "currency: USD",
		strings.Repeat(" ", 2) + "hoursPerMonth: 730",
		strings.Repeat(" ", 2) + "cpuPerCoreHour: 0.041",
		strings.Repeat(" ", 2) + "memoryPerGiBHour: 0.005",
		strings.Repeat(" ", 2) + "metadata:",
		strings.Repeat(" ", 4) + "provider: aws",
		strings.Repeat(" ", 4) + "region: eu-west-3",
		strings.Repeat(" ", 4) + `effectiveDate: "2026-08-01"`,
		strings.Repeat(" ", 4) + `sourceURL: "Illustrative Vigilo reference rates"`,
		"",
	}, "\n")
}

func replaceCatalogValue(t *testing.T, old, replacement string) string {
	t.Helper()

	input := validCatalogYAML()

	if count := strings.Count(input, old); count != 1 {
		t.Fatalf(
			"invalid test setup: expected one occurrence of %q, got %d",
			old,
			count,
		)
	}

	return strings.Replace(input, old, replacement, 1)
}

func TestReadCatalog(t *testing.T) {
	t.Parallel()

	catalog, err := Read(strings.NewReader(validCatalogYAML()))
	if err != nil {
		t.Fatalf("Read() returned an error: %v", err)
	}

	if catalog == nil {
		t.Fatal("expected a catalog, got nil")
	}

	if catalog.Metadata.Name != "aws-eu-west-3-reference" {
		t.Errorf("unexpected catalog name: %q", catalog.Metadata.Name)
	}

	if catalog.Spec.Currency != "USD" {
		t.Errorf("unexpected currency: %q", catalog.Spec.Currency)
	}

	tests := []struct {
		name string
		got  *float64
		want float64
	}{
		{"hoursPerMonth", catalog.Spec.HoursPerMonth, 730},
		{"cpuPerCoreHour", catalog.Spec.CPUPerCoreHour, 0.041},
		{"memoryPerGiBHour", catalog.Spec.MemoryPerGiBHour, 0.005},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got == nil {
				t.Fatal("expected a value, got nil")
			}

			if *tt.got != tt.want {
				t.Errorf("got %v, want %v", *tt.got, tt.want)
			}
		})
	}

	if catalog.Spec.Metadata.EffectiveDate != "2026-08-01" {
		t.Errorf("unexpected effective date: %q", catalog.Spec.Metadata.EffectiveDate)
	}
}

func TestReadCatalogAllowsZeroRates(t *testing.T) {
	t.Parallel()

	input := replaceCatalogValue(
		t,
		"cpuPerCoreHour: 0.041",
		"cpuPerCoreHour: 0",
	)
	input = strings.Replace(
		input,
		"memoryPerGiBHour: 0.005",
		"memoryPerGiBHour: 0",
		1,
	)

	catalog, err := Read(strings.NewReader(input))
	if err != nil {
		t.Fatalf("Read() returned an error: %v", err)
	}

	if catalog == nil {
		t.Fatal("expected a catalog, got nil")
	}

	if catalog.Spec.CPUPerCoreHour == nil ||
		catalog.Spec.MemoryPerGiBHour == nil {
		t.Fatal("expected explicitly declared rates")
	}

	if *catalog.Spec.CPUPerCoreHour != 0 ||
		*catalog.Spec.MemoryPerGiBHour != 0 {
		t.Fatal("expected zero rate to be preserved")
	}
}

func TestReadCatalogErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   string
		message string
	}{
		{
			name:    "empty input",
			input:   "",
			message: "pricing catalog is empty",
		},
		{
			name:    "null document",
			input:   "null\n",
			message: "pricing catalog is empty",
		},
		{
			name: "unsupported API version",
			input: replaceCatalogValue(
				t, "vigilo.io/v1alpha1", "vigilo.io/v9",
			),
			message: "unsupported pricing apiVersion",
		},
		{
			name: "unsupported kind",
			input: replaceCatalogValue(
				t, "kind: PricingCatalog", "kind: VigiloPolicy",
			),
			message: "unsupported pricing kind",
		},
		{
			name: "missing name",
			input: replaceCatalogValue(
				t, "name: aws-eu-west-3-reference", `name: ""`,
			),
			message: "metadata.name is required",
		},
		{
			name: "invalid currency",
			input: replaceCatalogValue(
				t, "currency: USD", "currency: usd",
			),
			message: "spec.currency",
		},
		{
			name: "missing hours",
			input: replaceCatalogValue(
				t, strings.Repeat(" ", 2)+"hoursPerMonth: 730\n", "",
			),
			message: "spec.hoursPerMonth is required",
		},
		{
			name: "zero hours",
			input: replaceCatalogValue(
				t, "hoursPerMonth: 730", "hoursPerMonth: 0",
			),
			message: "spec.hoursPerMonth must be finite and greater than zero",
		},
		{
			name: "missing CPU rate",
			input: replaceCatalogValue(
				t, strings.Repeat(" ", 2)+"cpuPerCoreHour: 0.041\n", "",
			),
			message: "spec.cpuPerCoreHour is required",
		},
		{
			name: "missing memory rate",
			input: replaceCatalogValue(
				t, strings.Repeat(" ", 2)+"memoryPerGiBHour: 0.005\n", "",
			),
			message: "spec.memoryPerGiBHour is required",
		},
		{
			name: "negative CPU rate",
			input: replaceCatalogValue(
				t, "cpuPerCoreHour: 0.041", "cpuPerCoreHour: -0.041",
			),
			message: "spec.cpuPerCoreHour must be finite and non-negative",
		},
		{
			name: "infinite memory rate",
			input: replaceCatalogValue(
				t, "memoryPerGiBHour: 0.005", "memoryPerGiBHour: .inf",
			),
			message: "spec.memoryPerGiBHour must be finite and non-negative",
		},
		{
			name: "NaN CPU rate",
			input: replaceCatalogValue(
				t, "cpuPerCoreHour: 0.041", "cpuPerCoreHour: .nan",
			),
			message: "spec.cpuPerCoreHour must be finite and non-negative",
		},
		{
			name: "invalid date",
			input: replaceCatalogValue(
				t, "2026-08-01", "2026-02-30",
			),
			message: "spec.metadata.effectiveDate",
		},
		{
			name: "unknown field",
			input: replaceCatalogValue(
				t, "memoryPerGiBHour:", "memoryPerGBHour:",
			),
			message: "decode pricing catalog",
		},
		{
			name: "duplicate field",
			input: replaceCatalogValue(
				t,
				strings.Repeat(" ", 2)+"currency: USD",
				strings.Join([]string{
					strings.Repeat(" ", 2) + "currency: USD",
					strings.Repeat(" ", 2) + "currency: EUR",
				}, "\n"),
			),
			message: "decode pricing catalog",
		},
		{
			name:    "invalid YAML",
			input:   "apiVersion: [",
			message: "decode pricing catalog",
		},
		{
			name: "multiple documents",
			input: validCatalogYAML() + "---\n" +
				validCatalogYAML(),
			message: "exactly one YAML document",
		},
		{
			name:    "additional empty document",
			input:   validCatalogYAML() + "---\n",
			message: "exactly one YAML document",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			catalog, err := Read(strings.NewReader(tt.input))
			if err == nil {
				t.Fatal("expected an error, got nil")
			}

			if !strings.Contains(err.Error(), tt.message) {
				t.Fatalf(
					"expected error containing %q, got %q",
					tt.message,
					err.Error(),
				)
			}

			if catalog != nil {
				if catalog != nil {
					t.Logf(
						"decoded effectiveDate: %q",
						catalog.Spec.Metadata.EffectiveDate,
					)
				}

				t.Fatalf("expected no catalog on error, got nil; input:\n%s", tt.input)
			}
		})
	}
}
