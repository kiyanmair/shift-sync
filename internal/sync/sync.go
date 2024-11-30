package sync

import (
	"errors"
	"fmt"
	"log"

	"github.com/kiyanmair/shift-sync/internal/config"
)

func (s *Syncer) RunSyncs() {
	var errs []error
	for _, syncCfg := range s.cfg.Syncs {
		if err := s.singleSync(syncCfg); err != nil {
			errs = append(errs, err)
		} else {
			log.Printf("Synced source %s to destination %s", syncCfg.Source, syncCfg.Destination)
		}
	}

	numTotal := len(s.cfg.Syncs)
	numFailed := len(errs)
	numSuccessful := numTotal - numFailed

	log.Printf("Completed %d/%d syncs", numSuccessful, numTotal)
	if numFailed > 0 {
		log.Fatalf("Encountered errors:\n%v", errors.Join(errs...))
	}
}

func (s *Syncer) singleSync(syncCfg config.Sync) error {
	source, exists := s.sources[syncCfg.Source]
	if !exists {
		return fmt.Errorf("source %s not found", syncCfg.Source)
	}

	dest, exists := s.destinations[syncCfg.Destination]
	if !exists {
		return fmt.Errorf("destination %s not found", syncCfg.Destination)
	}

	users, err := source.FetchUsers()
	if err != nil {
		return fmt.Errorf("failed to fetch users for source %s: %w", syncCfg.Source, err)
	}

	if err := dest.UpdateUsers(users); err != nil {
		return fmt.Errorf("failed to update users for destination %s: %w", syncCfg.Destination, err)
	}

	return nil
}
