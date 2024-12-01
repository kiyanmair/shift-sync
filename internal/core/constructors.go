package core

import (
	"fmt"

	"github.com/kiyanmair/shift-sync/internal/config"
)

func CreateIntegrations[T Integration](
	configs map[string]config.Integration,
	direction Direction,
) (map[string]T, []error) {
	results := make(map[string]T)
	var errs []error

	for name, cfg := range configs {
		nameAndType := fmt.Sprintf("%s (type %s)", name, cfg.Type)

		constructor, exists := integrationRegistry[cfg.Type]
		if !exists {
			errs = append(errs, fmt.Errorf("integration %s is not supported", nameAndType))
			continue
		}

		integ, err := constructor(cfg)
		if err != nil {
			errs = append(errs, fmt.Errorf("failed to construct %s: %w", nameAndType, err))
			continue
		}

		instance, ok := integ.(T)
		if !ok {
			errs = append(errs, fmt.Errorf("cannot use %s as a %s", nameAndType, direction))
			continue
		}

		if err := integ.Validate(direction); err != nil {
			errs = append(errs, fmt.Errorf("definition for %s is invalid: %v", nameAndType, err))
			continue
		}

		results[name] = instance
	}

	if len(errs) > 0 {
		return nil, errs
	}
	return results, nil
}
