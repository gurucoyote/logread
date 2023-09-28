package cmd

import (
	"fmt"
	"strings"
	"github.com/spf13/cobra"
)

var topCmd = &cobra.Command{
	Use:   "top [field]",
	Short: "Display the top n entries for a specified field",
	Long: `This command will go through the global log entries slice, and count how many entries there are for each unique value of the field specified.
It will then output a sorted top n list from the findings, with count and the field-name as columns.`,
	Args: cobra.ExactArgs(1),


Run: func(cmd *cobra.Command, args []string) {
	var validFields = map[string]string{
		"IP": "IP", "TIMESTAMP": "Timestamp", "STATUSCODE": "StatusCode", "BYTESSENT": "BytesSent",
		"REQUESTMETHOD": "RequestMethod", "REQUESTURL": "RequestURL", "REQUESTPROTOCOL": "RequestProtocol",
		"REFERRER": "Referrer", "USERAGENT": "UserAgent", "CHECKSUM": "Checksum",
	}
	upperArg := strings.ToUpper(args[0])
	if _, ok := validFields[upperArg]; !ok {
		keys := make([]string, 0, len(validFields))
		for k := range validFields {
			keys = append(keys, k)
		}
		fmt.Printf("Invalid field: %s. Valid fields are: %s\n", args[0], strings.Join(keys, ", "))
		return
	}
	field := validFields[upperArg]

	// Create a map to store the count of each unique value of the specified field
	counts := make(map[string]int)

	// Iterate over the global log entries slice
	for _, entry := range LogEntries {
		// For each entry, increment the count of the specified field's value in the map
		counts[entry[field]]++
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

	// Print the top n entries from the sorted slice
	for i := 0; i < numLines && i < len(pairs); i++ {
		fmt.Printf("%s: %d\n", pairs[i].Value, pairs[i].Count)
	}
},
}

func init() {
	RootCmd.AddCommand(topCmd)
}
