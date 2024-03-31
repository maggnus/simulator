package cli

import (
	shell "github.com/brianstrauch/cobra-shell"
	"github.com/spf13/cobra"
	"simulator/api"
)

var commands api.Command
var configFile string

var rootCmd = &cobra.Command{
	Use:   "client",
	Short: "client is a simple command-line tool to control robot movement.",
}

func init() {
	// Add cli arguments
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "configs/config.yml", "Path to config file")

	// Add subcommands
	rootCmd.AddCommand(instructionCmd)
	rootCmd.AddCommand(webCmd)
	rootCmd.AddCommand(shell.New(rootCmd, nil))
}

func Execute() error {
	return rootCmd.Execute()
}
