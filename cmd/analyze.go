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
	// TODO: Implement the logic to count the occurrences of each unique value of the specified field
	fmt.Printf("Field: %s\n", field)
},
}

func init() {
	RootCmd.AddCommand(topCmd)
}
