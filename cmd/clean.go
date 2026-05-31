package cmd

import (
	"fmt"

	"github.com/GhanshyamJha05/devtool-cli/internal/utils"

	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean <folder>",
	Short: "Organize files in a directory by file type",
	Long: `Scan a directory and automatically sort files into categorized subfolders
based on their file extension (Images, Documents, Code, Videos, etc.)

Examples:
  devtool clean ./downloads
  devtool clean C:\Users\you\Desktop --verbose`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		folderPath := args[0]

		if Verbose {
			utils.PrintDebug(fmt.Sprintf("Target directory: %s", folderPath))
		}

		utils.PrintInfo(fmt.Sprintf("Scanning folder: %s...", folderPath))

		result, err := utils.OrganizeFolder(folderPath)
		if err != nil {
			utils.PrintError(err.Error())
			return err
		}

		// Handle edge case: nothing to organize
		if result.TotalFiles == 0 {
			utils.PrintWarning("No files found to organize")
			return nil
		}

		// Print a summary of what was moved
		fmt.Println()
		for category, files := range result.Moved {
			utils.PrintSuccess(fmt.Sprintf("%-12s → %d file(s)", category, len(files)))
			if Verbose {
				for _, f := range files {
					utils.PrintDebug(fmt.Sprintf("  moved: %s", f))
				}
			}
		}

		if result.Skipped > 0 {
			utils.PrintWarning(fmt.Sprintf("Skipped %d file(s) with no extension", result.Skipped))
		}

		fmt.Println()
		utils.PrintSuccess(fmt.Sprintf("Done! Organized %d file(s) into %d categories", result.TotalFiles, len(result.Moved)))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
