package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	_ "github.com/kiyanmair/shift-sync/internal/integration/destination"
	_ "github.com/kiyanmair/shift-sync/internal/integration/source"
	"github.com/kiyanmair/shift-sync/internal/sync"
)

func main() {
	var rootCmd = &cobra.Command{Use: "shiftsync"}

	var configPath string
	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "Execute syncs from sources to destinations",
		Run: func(cmd *cobra.Command, args []string) {
			syncer := sync.NewSyncer(configPath)
			syncer.RunSyncs()
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
