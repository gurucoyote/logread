package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var todbCmd = &cobra.Command{
	Use:   "todb",
	Short: "Write the current log entries to the sqlite db",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("todb command called")
		// TODO: Implement the functionality to write log entries to the sqlite db
	},
}

var fromdbCmd = &cobra.Command{
	Use:   "fromdb",
	Short: "Read log entries from an existing sqlite db",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fromdb command called")
		// TODO: Implement the functionality to read log entries from the sqlite db
	},
}

func init() {
	RootCmd.AddCommand(todbCmd)
	RootCmd.AddCommand(fromdbCmd)
}
