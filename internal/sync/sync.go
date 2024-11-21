package sync

import (
	"fmt"

	"github.com/kiyanmair/shift-sync/config"
	"github.com/kiyanmair/shift-sync/internal/destination"
	"github.com/kiyanmair/shift-sync/internal/source"
)

func Sync(syncCfg config.Sync, sources map[string]source.Source, destinations map[string]destination.Destination) error {
	source, sourceExists := sources[syncCfg.SourceID]
	if !sourceExists {
		return fmt.Errorf("source %s not found", syncCfg.SourceID)
	}

	dest, destExists := destinations[syncCfg.DestinationID]
	if !destExists {
		return fmt.Errorf("destination %s not found", syncCfg.DestinationID)
	}

	users, err := source.FetchUsers()
	if err != nil {
		return fmt.Errorf("failed to fetch users for source %s: %w", syncCfg.SourceID, err)
	}

	if err := dest.UpdateUsers(users); err != nil {
		return fmt.Errorf("failed to update users for destination %s: %w", syncCfg.DestinationID, err)
	}

	return nil
}
