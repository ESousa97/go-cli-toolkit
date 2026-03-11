package commands

import (
	"github.com/ESousa97/go-cli-toolkit/internal/config"
	"github.com/spf13/cobra"
)

const (
	toolkitUse   = "toolkit"
	toolkitShort = "A utility CLI toolkit written in Go"
	toolkitLong  = `Toolkit is a command-line interface (CLI) with utilities for network diagnostics 
and data manipulation. Use the available subcommands to interact with the features.`
)

// rootCmd represents the base command when called without subcommands
var rootCmd = &cobra.Command{
	Use:           toolkitUse,
	Short:         toolkitShort,
	Long:          toolkitLong,
	SilenceErrors: true, // main.go handles and logs the error
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(func() {
		if err := config.InitConfig(); err != nil {
			// In a CLI, config errors are reported but don't always stop execution.
			// Depending on criticality, we could use panic or exit.
		}
	})
}
