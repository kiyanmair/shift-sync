package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/kiyanmair/shift-sync/config"
	"github.com/kiyanmair/shift-sync/internal/destination"
	"github.com/kiyanmair/shift-sync/internal/source"
	"github.com/kiyanmair/shift-sync/internal/sync"
)

func main() {
	var rootCmd = &cobra.Command{Use: "shiftsync"}

	var configPath string
	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "Execute syncs from sources to destinations",
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.LoadConfig(configPath)
			if err != nil {
				log.Fatalf("Failed to load config: %v", err)
			}

			sources := make(map[string]source.Source)
			for _, srcCfg := range cfg.Sources {
				src, err := source.NewSource(srcCfg)
				if err != nil {
					log.Printf("Failed to create source %s: %v", srcCfg.ID, err)
					continue
				}
				sources[srcCfg.ID] = src
			}

			destinations := make(map[string]destination.Destination)
			for _, destCfg := range cfg.Destinations {
				dest, err := destination.NewDestination(destCfg)
				if err != nil {
					log.Printf("Failed to create destination %s: %v", destCfg.ID, err)
					continue
				}
				destinations[destCfg.ID] = dest
			}

			var errs []error
			for _, syncCfg := range cfg.Syncs {
				err := sync.Sync(syncCfg, sources, destinations)
				if err != nil {
					errs = append(errs, err)
					continue
				}
				log.Printf("Synced source %s to destination %s", syncCfg.SourceID, syncCfg.DestinationID)
			}

			numTotal := len(cfg.Syncs)
			numFailed := len(errs)
			numSuccessful := numTotal - numFailed

			log.Printf("Completed %d/%d syncs", numSuccessful, numTotal)
			if numFailed > 0 {
				log.Fatalf("Encountered errors:\n%v", errors.Join(errs...))
			}
		},
	}

	runCmd.Flags().StringVarP(&configPath, "config-path", "c", "", "Path to configuration file")
	runCmd.MarkFlagRequired("config-path")

	rootCmd.AddCommand(runCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
