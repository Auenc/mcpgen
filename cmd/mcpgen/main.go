package main

import (
	"fmt"
	"os"

	"github.com/auenc/mcpgen/internal/config"
	"github.com/auenc/mcpgen/internal/generator"
	"github.com/spf13/cobra"
)

var (
	configPath string
	isDryRun   bool

	rootCmd = &cobra.Command{
		Use:   "mcpgen",
		Short: "Go MCP server generator",
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.LoadConfig(configPath)
			if err != nil {
				fmt.Printf("Error loading config: %v\n", err)
				os.Exit(1)
			}

			// Adding to config to avoid having to pass a lot of params
			cfg.DryRun = isDryRun

			generator.Generate(*cfg)
		},
	}
)

func main() {
	rootCmd.Flags().StringVarP(&configPath, "config", "c", "mcpgen.yaml", "Path to the YAML mcpgen configuration file")
	rootCmd.Flags().BoolVar(&isDryRun, "dry", false, "Show what would be done, but do not create files")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
