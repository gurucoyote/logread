package cmd

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

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
					fmt.Printf("Updated 'number-lines' to: %d\n", numLines)
				}
			case "start":
				if t, err := dateparse.ParseAny(args[1]); err == nil {
					Start = t.Format(time.RFC3339)
					fmt.Printf("Updated 'start' to: %s\n", Start)
				} else {
					fmt.Printf("Invalid date format for start: %s\n", args[1])
				}
			case "end":
				if t, err := dateparse.ParseAny(args[1]); err == nil {
					End = t.Format(time.RFC3339)
					fmt.Printf("Updated 'end' to: %s\n", End)
				} else {
					fmt.Printf("Invalid date format for end: %s\n", args[1])
				}
			case "group-by":
				field, err := ValidField(args[1])
				if err != nil {
					fmt.Println(err)
				} else {
					GroupBy = field
					fmt.Printf("Updated 'group-by' to: %s\n", GroupBy)
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
	validFields := GetValidFieldNames()
	matchedField := ""
	for _, validField := range validFields {
		match, _ := regexp.MatchString("(?i)"+fieldName, validField)
		if match {
			matchedField = validField
			break
		}
	}
	if matchedField == "" {
		return "", fmt.Errorf("Invalid field name: %s. Valid fields are: %s", fieldName, strings.Join(GetValidFieldNames(), ", "))
	}

	return matchedField, nil
}

func GetValidFieldNames() []string {
	e := reflect.ValueOf(&NginxAccessLog{}).Elem()
	var fieldNames []string
	for i := 0; i < e.NumField(); i++ {
		fieldNames = append(fieldNames, e.Type().Field(i).Name)
	}
	return fieldNames
}
