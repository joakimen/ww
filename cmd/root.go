/*
Package cmd marks the entry point for the CLI application.

Copyright © 2024 Joakim Lindeng Engeset <joakim.engeset@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ww",
	Short: "Creates tasks from data in various places",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("verbose", "v", false, "Enable verbose output")
}
