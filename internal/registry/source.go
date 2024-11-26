package registry

import (
	"fmt"

	"github.com/kiyanmair/shift-sync/config"
	"github.com/kiyanmair/shift-sync/internal/source"
)

var sourceRegistry = map[string]func(config.Source) (source.Source, error){}

func RegisterSource(name string, constructor func(config.Source) (source.Source, error)) {
	sourceRegistry[name] = constructor
}

func NewSource(config config.Source) (source.Source, error) {
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