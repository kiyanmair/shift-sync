package core

import "github.com/kiyanmair/shift-sync/config"

var sourceRegistry = map[string]func(config.Source) (Source, error){}

var destinationRegistry = map[string]func(config.Destination) (Destination, error){}

func RegisterSource(name string, constructor func(config.Source) (Source, error)) {
	sourceRegistry[name] = constructor
}

func RegisterDestination(name string, constructor func(config.Destination) (Destination, error)) {
	destinationRegistry[name] = constructor
}
