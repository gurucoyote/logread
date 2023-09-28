package cmd

import (
	"fmt"
	"sort"
	"strconv"
	"github.com/spf13/cobra"
)

var topCmd = &cobra.Command{
	Use:   "top [field]",
	Short: "Display the top n entries for a specified field",
	Long: `This command will go through the global log entries slice, and count how many entries there are for each unique value of the field specified.
It will then output a sorted top n list from the findings, with count and the field-name as columns.`,
	Args: cobra.MaximumNArgs(2),


Run: func(cmd *cobra.Command, args []string) {
				field, err := ValidField(args[0])
				if err != nil {
					fmt.Println(err)
					return
				}
					fmt.Println("top for: ", field)
	// Create a map to store the count of each unique value of the specified field
	counts := make(map[string]int)

	// Iterate over the global log entries slice
	for _, entry := range LogEntries {
		// For each entry, increment the count of the specified field's value in the map
		counts[entry.GetField(field)]++
	}

	// Convert the map into a slice of pairs (value, count) for sorting
	type pair struct {
		Value string
		Count int
	}
	var pairs []pair
	for value, count := range counts {
		pairs = append(pairs, pair{value, count})
	}

	// Sort the slice in descending order of count
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Count > pairs[j].Count
	})

	// Determine the number of lines to print
	var numLinesToPrint int
	if len(args) > 1 {
		numLinesToPrint, err = strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid number of lines:", args[1])
			return
		}
	} else {
		numLinesToPrint = len(pairs)
	}

	// Print the top n entries from the sorted slice
	for i := 0; i < numLinesToPrint && i < len(pairs); i++ {
		fmt.Printf("%s: %d\n", pairs[i].Value, pairs[i].Count)
	}
},
}

var Quiet bool

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Display metrics about the currently loaded set of log entries",
	Long:  `This command will output metrics about the currently loaded set of log entries. If the -q or --quiet flag is set, it will only output the total count of log entries. Otherwise, it will output the number of unique values for each valid field name.`,
	Run:   statusCmdRun,
}

func statusCmdRun(cmd *cobra.Command, args []string) {
	if Quiet {
		fmt.Printf("Total log entries: %d\n", len(LogEntries))
	} else {
		for _, fieldName := range GetValidFieldNames() {
			uniqueValues := make(map[string]struct{})
			for _, entry := range LogEntries {
				uniqueValues[entry.GetField(fieldName)] = struct{}{}
			}
			fmt.Printf("%s: %d unique values\n", fieldName, len(uniqueValues))
		}
	}
}

func init() {
	RootCmd.AddCommand(topCmd)
	RootCmd.AddCommand(statusCmd)
	statusCmd.Flags().BoolVarP(&Quiet, "quiet", "q", false, "Only output the total count of log entries")
}
