package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "logread",
	Short: "Read, parse, and analyze (nginx) access.log files",
	Long: `logread is an application that allows you to read, parse, and analyze (nginx) access.log files.
It provides useful insights and analytics from your log files.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("Root command called")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
