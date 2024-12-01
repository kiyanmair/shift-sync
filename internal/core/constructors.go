package core

import (
	"fmt"

	"github.com/kiyanmair/shift-sync/internal/config"
)

func NewIntegration(config config.Integration) (Integration, error) {
	constructor, exists := integrationRegistry[config.Type]
	if !exists {
		return nil, fmt.Errorf("unsupported integration type: %s", config.Type)
	}
	integ, err := constructor(config)
	if err != nil {
		return nil, fmt.Errorf("failed to construct integration: %w", err)
	}
	return integ, nil
}

func CreateIntegrations[T Integration](
	configs map[string]config.Integration,
	direction IntegrationDirection,
) (map[string]T, []error) {
	results := make(map[string]T)
	var errs []error

	for name, cfg := range configs {
		integ, err := NewIntegration(cfg)
		if err != nil {
			errs = append(errs, fmt.Errorf("failed to create %s: %v", name, err))
			continue
		}

		instance, ok := integ.(T)
		if !ok {
			errs = append(errs, fmt.Errorf("cannot use integration type %s as a %s", cfg.Type, direction))
			continue
		}

		results[name] = instance
	}

	if len(errs) > 0 {
		return nil, errs
	}
	return results, nil
}
