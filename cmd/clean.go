package cmd

import (
	"fmt"

	"devtool-cli/internal/utils"

	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean <folder>",
	Short: "Organize files in a directory by file type",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		folderPath := args[0]

		if Verbose {
			utils.PrintDebug(fmt.Sprintf("Target directory: %s", folderPath))
		}

		utils.PrintInfo(fmt.Sprintf("Cleaning up folder: %s...", folderPath))

		if err := utils.OrganizeFolder(folderPath); err != nil {
			utils.PrintError(err.Error())
			return err
		}

		utils.PrintSuccess("Folder successfully organized!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
