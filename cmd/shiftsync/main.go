package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/kiyanmair/shift-sync/config"
)

func main() {
	var rootCmd = &cobra.Command{Use: "shiftsync"}

	var configPath string
	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "Execute syncs from sources to destinations",
		Run: func(cmd *cobra.Command, args []string) {
			_, err := config.LoadConfig(configPath)
			if err != nil {
				log.Fatalf("Failed to load config: %v", err)
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
