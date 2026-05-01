package cmd

import (
	"fmt"
	"os"

	"devtool-cli/internal/utils"

	"github.com/spf13/cobra"
)

// saveFormatted is the local flag for --save on format command
var saveFormatted string

var formatCmd = &cobra.Command{
	Use:   "format <file.json>",
	Short: "Pretty-print JSON files",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		filePath := args[0]

		if Verbose {
			utils.PrintDebug(fmt.Sprintf("Reading JSON file: %s", filePath))
		}

		formattedJSON, err := utils.FormatJSON(filePath)
		if err != nil {
			utils.PrintError(err.Error())
			return err
		}

		// If --save flag is provided, write formatted JSON to that file
		if saveFormatted != "" {
			if Verbose {
				utils.PrintDebug(fmt.Sprintf("Writing formatted output to: %s", saveFormatted))
			}
			if err := os.WriteFile(saveFormatted, []byte(formattedJSON+"\n"), 0644); err != nil {
				return fmt.Errorf("failed to save to file '%s': %w", saveFormatted, err)
			}
			utils.PrintSuccess(fmt.Sprintf("Formatted JSON saved to %s", saveFormatted))
			return nil
		}

		fmt.Println(formattedJSON)
		utils.PrintSuccess("JSON formatted successfully")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(formatCmd)
	formatCmd.Flags().StringVarP(&saveFormatted, "save", "s", "", "save formatted output to a file (e.g., --save pretty.json)")
}
