package sync

import (
	"log"

	"github.com/kiyanmair/shift-sync/config"
	"github.com/kiyanmair/shift-sync/internal/destination"
	"github.com/kiyanmair/shift-sync/internal/registry"
	"github.com/kiyanmair/shift-sync/internal/source"
)

type Syncer struct {
	cfg          *config.Config
	sources      map[string]source.Source
	destinations map[string]destination.Destination
}

func NewSyncer(configPath string) *Syncer {
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	sources := make(map[string]source.Source)
	for _, srcCfg := range cfg.Sources {
		src, err := registry.NewSource(srcCfg)
		if err != nil {
			log.Printf("Failed to create source %s: %v", srcCfg.ID, err)
			continue
		}
		sources[srcCfg.ID] = src
	}

	destinations := make(map[string]destination.Destination)
	for _, destCfg := range cfg.Destinations {
		dest, err := registry.NewDestination(destCfg)
		if err != nil {
			log.Printf("Failed to create destination %s: %v", destCfg.ID, err)
			continue
		}
		destinations[destCfg.ID] = dest
	}

	syncer := Syncer{
		cfg:          cfg,
		sources:      sources,
		destinations: destinations,
	}

	return &syncer
}
