package core

import (
	"github.com/kiyanmair/shift-sync/internal/config"
)

type integrationConstructor func(config.Integration) (Integration, error)

var integrationRegistry = map[string]integrationConstructor{}

func RegisterIntegration(name string, constructor func(config.Integration) (Integration, error)) {
	integrationRegistry[name] = constructor
}
