package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yourproject",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application.`,
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
