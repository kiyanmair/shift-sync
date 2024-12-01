package sync

import (
	"errors"
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

	sources, srcErrs := core.CreateIntegrations[core.Source](
		cfg.Sources,
		core.SourceDirection,
	)

	destinations, destErrs := core.CreateIntegrations[core.Destination](
		cfg.Destinations,
		core.DestinationDirection,
	)

	errs := append(srcErrs, destErrs...)
	if len(errs) > 0 {
		log.Fatalf("Encountered errors while initialising:\n%v", errors.Join(errs...))
	}

	syncer := Syncer{
		cfg:          cfg,
		sources:      sources,
		destinations: destinations,
	}

	return &syncer
}
