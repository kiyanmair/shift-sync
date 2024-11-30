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
		src, err := core.NewSource(srcCfg)
		if err != nil {
			log.Printf("Failed to create source %s: %v", srcName, err)
			continue
		}
		sources[srcName] = src
	}

	destinations := make(map[string]core.Destination)
	for destName, destCfg := range cfg.Destinations {
		dest, err := core.NewDestination(destCfg)
		if err != nil {
			log.Printf("Failed to create destination %s: %v", destName, err)
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
