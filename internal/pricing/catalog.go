package pricing

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strings"
	"time"
)

var currencyPattern = regexp.MustCompile(`^[A-Z]{3}$`)

// Catalog defines reference rates for request-based cost estimation.
type Catalog struct {
	APIVersion string          `yaml:"apiVersion"`
	Kind       string          `yaml:"kind"`
	Metadata   CatalogMetadata `yaml:"metadata"`
	Spec       CatalogSpec     `yaml:"spec"`
}

// CatalogMetadata identifies a pricing catalog.
type CatalogMetadata struct {
	Name string `yaml:"name"`
}

// CatalogSpec defines the currency, duration, rates, and provenance.
type CatalogSpec struct {
	Currency         string       `yaml:"currency"`
	HoursPerMonth    *float64     `yaml:"hoursPerMonth"`
	CPUPerCoreHour   *float64     `yaml:"cpuPerCoreHour"`
	MemoryPerGiBHour *float64     `yaml:"memoryPerGiBHour"`
	Metadata         RateMetadata `yaml:"metadata"`
}

// RateMetadata documents the origin and applicability of reference rates.
type RateMetadata struct {
	Provider      string `yaml:"provider"`
	Region        string `yaml:"region"`
	EffectiveDate string `yaml:"effectiveDate"`
	Source        string `yaml:"sourceURL"`
}

// Validate checks that the catalog is well-formed and has valid rates.
func (c Catalog) Validate() error {
	if c.APIVersion != "vigilo.io/v1alpha1" {
		return fmt.Errorf("unsupported pricing apiVersion %q", c.APIVersion)
	}

	if c.Kind != "PricingCatalog" {
		return fmt.Errorf("unsupported pricing kind %q", c.Kind)
	}

	if strings.TrimSpace(c.Metadata.Name) == "" {
		return errors.New("metadata.name is required")
	}

	if !currencyPattern.MatchString(c.Spec.Currency) {
		return errors.New("spec.currency must contain exactly three uppercase letters, e.g. \"USD\" or \"EUR\"")
	}

	hours := c.Spec.HoursPerMonth
	if hours == nil {
		return errors.New("spec.hoursPerMonth is required")
	}

	if math.IsNaN(*hours) || math.IsInf(*hours, 0) || *hours <= 0 {
		return errors.New("spec.hoursPerMonth must be finite and greater than zero")
	}

	if err := validateRate("spec.cpuPerCoreHour", c.Spec.CPUPerCoreHour); err != nil {
		return err
	}

	if err := validateRate(
		"spec.memoryPerGiBHour",
		c.Spec.MemoryPerGiBHour); err != nil {
		return err
	}

	if date := c.Spec.Metadata.EffectiveDate; date != "" {
		if _, err := time.Parse(time.DateOnly, date); err != nil {
			return fmt.Errorf("spec.metadata.effectiveDate must be in YYYY-MM-DD format: %w", err)
		}
	}

	return nil
}

func validateRate(field string, value *float64) error {
	if value == nil {
		return fmt.Errorf("%s is required", field)
	}

	if math.IsNaN(*value) || math.IsInf(*value, 0) || *value < 0 {
		return fmt.Errorf("%s must be finite and non-negative", field)
	}

	return nil
}
