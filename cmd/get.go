package cmd

import (
	"fmt"
	"os"

	"devtool-cli/internal/utils"

	"github.com/spf13/cobra"
)

// saveFile is the local flag for --save
var saveFile string

var getCmd = &cobra.Command{
	Use:   "get <url>",
	Short: "Fetch API data and display response",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		url := args[0]

		if Verbose {
			utils.PrintDebug(fmt.Sprintf("Target URL: %s", url))
		}

		utils.PrintInfo(fmt.Sprintf("Fetching data from %s...", url))

		data, err := utils.FetchData(url)
		if err != nil {
			utils.PrintError(err.Error())
			return err
		}

		// If --save flag is provided, write to file instead of stdout
		if saveFile != "" {
			if Verbose {
				utils.PrintDebug(fmt.Sprintf("Saving response to file: %s", saveFile))
			}
			if err := os.WriteFile(saveFile, []byte(data), 0644); err != nil {
				return fmt.Errorf("failed to save to file '%s': %w", saveFile, err)
			}
			utils.PrintSuccess(fmt.Sprintf("Response saved to %s", saveFile))
			return nil
		}

		fmt.Println(data)
		utils.PrintSuccess("Request completed successfully")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	// Local flag: only available on the "get" command
	getCmd.Flags().StringVarP(&saveFile, "save", "s", "", "save response to a file (e.g., --save output.json)")
}
