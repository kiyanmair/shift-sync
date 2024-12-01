package core

import (
	"fmt"

	"github.com/kiyanmair/shift-sync/internal/config"
)

func CreateIntegrations[T Integration](
	configs map[string]config.Integration,
	direction IntegrationDirection,
) (map[string]T, []error) {
	results := make(map[string]T)
	var errs []error

	for name, cfg := range configs {
		constructor, exists := integrationRegistry[cfg.Type]
		if !exists {
			errs = append(errs, fmt.Errorf("unsupported integration type: %s", cfg.Type))
			continue
		}

		integ, err := constructor(cfg)
		if err != nil {
			errs = append(errs, fmt.Errorf("failed to construct integration: %w", err))
			continue
		}

		instance, ok := integ.(T)
		if !ok {
			errs = append(errs, fmt.Errorf("cannot use %s (type %s) as a %s", name, cfg.Type, direction))
			continue
		}

		if err := integ.Validate(direction); err != nil {
			errs = append(errs, fmt.Errorf("definition for %s (type %s) is invalid: %v", name, cfg.Type, err))
			continue
		}

		results[name] = instance
	}

	if len(errs) > 0 {
		return nil, errs
	}
	return results, nil
}
