package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/araddon/dateparse"
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
			fmt.Printf("interactive: %t\n", Interactive)
			fmt.Printf("start: %s\n", Start)
			fmt.Printf("end: %s\n", End)
			fmt.Printf("group-by: %s\n", GroupBy)
		} else if len(args) == 1 {
			switch args[0] {
			case "interactive":
				Interactive = !Interactive
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
				if isValidDate(args[1]) {
					Start = args[1]
				} else {
					fmt.Printf("Invalid date format for start: %s\n", args[1])
				}
			case "end":
				if isValidDate(args[1]) {
					End = args[1]
				} else {
					fmt.Printf("Invalid date format for end: %s\n", args[1])
				}
			case "group-by":
				field, err := ValidField(args[1])
				if err != nil {
					fmt.Printf("Invalid field for group-by: %s. Valid fields are: %s\n", args[1], strings.Join(GetValidFieldNames(), ", "))
				} else {
					GroupBy = field
				}
			default:
				fmt.Printf("Unknown flag: %s\n", args[0])
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}
func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

func isValidDate(date string) bool {
	_, err := dateparse.ParseAny(date)
	return err == nil
}
func ValidField(fieldName string) (string, error) {
	validFields := map[string]string{
		"IP": "IP", "TIMESTAMP": "Timestamp", "STATUSCODE": "StatusCode", "BYTESSENT": "BytesSent",
		"REQUESTMETHOD": "RequestMethod", "REQUESTURL": "RequestURL", "REQUESTPROTOCOL": "RequestProtocol",
		"REFERRER": "Referrer", "USERAGENT": "UserAgent", "CHECKSUM": "Checksum",
	}
	field, ok := validFields[strings.ToUpper(fieldName)]
	if !ok {
		return "", fmt.Errorf("Invalid field name: %s", fieldName)
	}
	return field, nil
}

import "reflect"

func GetValidFieldNames() []string {
	e := reflect.ValueOf(&NginxAccessLog{}).Elem()
	var fieldNames []string
	for i := 0; i < e.NumField(); i++ {
		fieldNames = append(fieldNames, e.Type().Field(i).Name)
	}
	return fieldNames
}
