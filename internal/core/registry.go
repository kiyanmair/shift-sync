package core

import "github.com/kiyanmair/shift-sync/config"

type sourceConstructor func(config.Source) (Source, error)
type destinationConstructor func(config.Destination) (Destination, error)

var sourceRegistry = map[string]sourceConstructor{}
var destinationRegistry = map[string]destinationConstructor{}

func RegisterSource(name string, constructor func(config.Source) (Source, error)) {
	sourceRegistry[name] = constructor
}

func RegisterDestination(name string, constructor func(config.Destination) (Destination, error)) {
	destinationRegistry[name] = constructor
}
