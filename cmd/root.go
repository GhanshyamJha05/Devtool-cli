package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Verbose is a global flag accessible by all subcommands for debug logging.
var Verbose bool

var rootCmd = &cobra.Command{
	Use:   "devtool",
	Short: "A production-quality CLI developer tool",
	Long: `devtool-cli is a modular command-line application that automates
common developer tasks like fetching APIs, formatting JSON, and organizing files.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// PersistentFlags are inherited by ALL child commands.
	// This is how you create a truly global flag in Cobra.
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "enable verbose/debug output")
}
