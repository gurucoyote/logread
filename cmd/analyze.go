package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var topCmd = &cobra.Command{
	Use:   "top [field]",
	Short: "Display the top n entries for a specified field",
	Long: `This command will go through the global log entries slice, and count how many entries there are for each unique value of the field specified.
It will then output a sorted top n list from the findings, with count and the field-name as columns.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		field := args[0]
		// TODO: Implement the logic to count the occurrences of each unique value of the specified field
		fmt.Printf("Field: %s\n", field)
	},
}

func init() {
	analyzeCmd.AddCommand(topCmd)
}
