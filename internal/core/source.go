package core

import (
	"fmt"

	"github.com/kiyanmair/shift-sync/internal/config"
)

func NewSource(config config.Source) (Source, error) {
	constructor, exists := sourceRegistry[config.Type]
	if !exists {
		return nil, fmt.Errorf("unsupported source type: %s", config.Type)
	}
	source, err := constructor(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create example source: %w", err)
	}
	return source, nil
}
