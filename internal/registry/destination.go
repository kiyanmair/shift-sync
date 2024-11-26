package registry

import (
	"fmt"

	"github.com/kiyanmair/shift-sync/config"
	"github.com/kiyanmair/shift-sync/internal/destination"
)

var destinationRegistry = map[string]func(config.Destination) (destination.Destination, error){}

func RegisterDestination(name string, constructor func(config.Destination) (destination.Destination, error)) {
	destinationRegistry[name] = constructor
}

func NewDestination(config config.Destination) (destination.Destination, error) {
	constructor, exists := destinationRegistry[config.Type]
	if !exists {
		return nil, fmt.Errorf("unsupported destination type: %s", config.Type)
	}
	destination, err := constructor(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create example destination: %w", err)
	}
	return destination, nil
}
