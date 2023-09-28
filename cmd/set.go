package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set [flag] [value]",
	Short: "Set the value of a flag",
	Long: `Set the value of a flag. If no value is provided for a boolean flag, it will toggle the state of the flag.
If no arguments are given, it will print the state/value of all flags.`,
	Args: cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Printf("number-lines: %d\n", numLines)
			fmt.Printf("interactive: %t\n", interactive)
			fmt.Printf("start: %s\n", start)
			fmt.Printf("end: %s\n", end)
			fmt.Printf("group-by: %s\n", groupBy)
		} else if len(args) == 1 {
			switch args[0] {
			case "interactive":
				interactive = !interactive
			default:
				fmt.Printf("Unknown flag: %s\n", args[0])
			}
		} else {
			switch args[0] {
			case "number-lines":
				num, err := strconv.Atoi(args[1])
				if err != nil {
					fmt.Printf("Invalid value for number-lines: %s\n", args[1])
				} else {
					numLines = num
				}
			case "start":
				start = args[1]
			case "end":
				end = args[1]
			case "group-by":
				groupBy = args[1]
			default:
				fmt.Printf("Unknown flag: %s\n", args[0])
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
