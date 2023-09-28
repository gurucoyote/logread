package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "Output log entries as JSON",
	Long:  `This command outputs the log entries as a JSON formatted string.`,
	Run: func(cmd *cobra.Command, args []string) {
		var entriesToOutput []NginxAccessLog
		if numLines > 0 && numLines < len(LogEntries) {
			entriesToOutput = LogEntries[:numLines]
		} else {
			entriesToOutput = LogEntries
		}
		jsonData, err := json.MarshalIndent(entriesToOutput, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(jsonData))
	},
}

func init() {
	RootCmd.AddCommand(jsonCmd)
}
