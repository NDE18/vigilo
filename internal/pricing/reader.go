package pricing

import (
	"errors"
	"fmt"
	"io"

	"go.yaml.in/yaml/v3"
)

// Read loads exactly one pricing catalog and validates its contents.
//
// It returns no catalog when decoding or validation fails.
func Read(reader io.Reader) (*Catalog, error) {
	decoder := yaml.NewDecoder(reader)
	decoder.KnownFields(true)

	var catalog *Catalog

	if err := decoder.Decode(&catalog); err != nil {
		if errors.Is(err, io.EOF) {
			return nil, errors.New("pricing catalog is empty")
		}

		return nil, fmt.Errorf("decode pricing catalog: %w", err)
	}

	if catalog == nil {
		return nil, errors.New("pricing catalog is empty")
	}

	var extra yaml.Node

	err := decoder.Decode(&extra)
	if err == nil {
		return nil, errors.New("pricing input must contain exactly one YAML document")
	}

	if !errors.Is(err, io.EOF) {
		return nil, fmt.Errorf("decode additional pricing document: %w", err)
	}

	if err := catalog.Validate(); err != nil {
		return nil, fmt.Errorf("invalid pricing catalog: %w", err)
	}

	return catalog, nil
}
