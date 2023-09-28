package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "logread [file]",
	Short: "Read, parse, and analyze (nginx) access.log files",
	Long: `logread is an application that allows you to read, parse, and analyze (nginx) access.log files.
It provides useful insights and analytics from your log files.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		if file == "-" {
			// Read from stdin
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				line := scanner.Text()
				// Parse the line here
				fmt.Println(line)
			}
		} else {
			// Open the file and parse it line by line
			file, err := os.Open(file)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()
				// Parse the line here
				fmt.Println(line)
			}
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
