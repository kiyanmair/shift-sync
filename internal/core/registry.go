package core

import "github.com/kiyanmair/shift-sync/internal/config"

type sourceConstructor func(config.Integration) (Source, error)
type destinationConstructor func(config.Integration) (Destination, error)

var sourceRegistry = map[string]sourceConstructor{}
var destinationRegistry = map[string]destinationConstructor{}

func RegisterSource(name string, constructor func(config.Integration) (Source, error)) {
	sourceRegistry[name] = constructor
}

func RegisterDestination(name string, constructor func(config.Integration) (Destination, error)) {
	destinationRegistry[name] = constructor
}
