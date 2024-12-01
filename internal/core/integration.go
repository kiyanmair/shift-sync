package core

import (
	"fmt"

	"github.com/kiyanmair/shift-sync/internal/config"
)

type Integration interface {
	Valid() (bool, error)
}

type Source interface {
	Integration
	GetUsers() ([]string, error)
}

type Destination interface {
	Integration
	SetUsers(users []string) error
}

func NewIntegration(config config.Integration) (Integration, error) {
	constructor, exists := integrationRegistry[config.Type]
	if !exists {
		return nil, fmt.Errorf("unsupported integration type: %s", config.Type)
	}
	integ, err := constructor(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create example integration: %w", err)
	}
	return integ, nil
}

func IsSource(i Integration) (Source, bool) {
	s, ok := i.(Source)
	return s, ok
}

func IsDestination(i Integration) (Destination, bool) {
	d, ok := i.(Destination)
	return d, ok
}
