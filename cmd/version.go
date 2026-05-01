package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Build-time variables — injected via `go build -ldflags`
var (
	version = "1.0.0"
	commit  = "dev"
	date    = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of devtool-cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("devtool-cli v%s\n", version)
		fmt.Printf("  commit : %s\n", commit)
		fmt.Printf("  built  : %s\n", date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
