package sync

import (
	"log"

	"github.com/kiyanmair/shift-sync/internal/config"
	"github.com/kiyanmair/shift-sync/internal/core"
)

type Syncer struct {
	cfg          *config.Config
	sources      map[string]core.Source
	destinations map[string]core.Destination
}

func NewSyncer(configPath string) *Syncer {
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	sources := make(map[string]core.Source)
	for srcName, srcCfg := range cfg.Sources {
		integ, err := core.NewIntegration(srcCfg)
		if err != nil {
			log.Printf("Failed to create integration %s: %v", srcName, err)
			continue
		}

		src, ok := integ.(core.Source)
		if !ok {
			log.Printf("Integration type %s cannot be used as a source", srcCfg.Type)
			continue
		}

		sources[srcName] = src
	}

	destinations := make(map[string]core.Destination)
	for destName, destCfg := range cfg.Destinations {
		integ, err := core.NewIntegration(destCfg)
		if err != nil {
			log.Printf("Failed to create integration %s: %v", destName, err)
			continue
		}

		dest, ok := integ.(core.Destination)
		if !ok {
			log.Printf("Integration type %s cannot be used as a destination", destCfg.Type)
			continue
		}

		destinations[destName] = dest
	}

	syncer := Syncer{
		cfg:          cfg,
		sources:      sources,
		destinations: destinations,
	}

	return &syncer
}
