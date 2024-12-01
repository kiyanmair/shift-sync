package core

import (
	"fmt"

	"github.com/kiyanmair/shift-sync/internal/config"
)

func NewDestination(config config.Integration) (Destination, error) {
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
