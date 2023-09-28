package cmd

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/araddon/dateparse"
	"github.com/spf13/cobra"
)

var LogEntries []NginxAccessLog
var numLines int
var Interactive bool
var Start, End, GroupBy string

type NginxAccessLog struct {
	IP              string
	Timestamp       time.Time
	StatusCode      string
	BytesSent       string
	RequestMethod   string
	RequestURL      string
	RequestProtocol string
	Referrer        string
	UserAgent       string
	Checksum        string
}

func (log NginxAccessLog) GetField(field string) string {
	switch strings.ToUpper(field) {
	case "IP":
		return log.IP
	case "TIMESTAMP":
		return log.Timestamp.String()
	case "STATUSCODE":
		return log.StatusCode
	case "BYTESSENT":
		return log.BytesSent
	case "REQUESTMETHOD":
		return log.RequestMethod
	case "REQUESTURL":
		return log.RequestURL
	case "REQUESTPROTOCOL":
		return log.RequestProtocol
	case "REFERRER":
		return log.Referrer
	case "USERAGENT":
		return log.UserAgent
	case "CHECKSUM":
		return log.Checksum
	default:
		return ""
	}
}

func init() {
	RootCmd.PersistentFlags().IntVarP(&numLines, "number-lines", "n", 0, "Number of lines to read from the file or stdin")
	RootCmd.PersistentFlags().BoolVarP(&Interactive, "interactive", "i", false, "Enable interactive mode")
	RootCmd.PersistentFlags().StringVarP(&Start, "start", "s", "", "Limit entries on or after this datetime")
	RootCmd.PersistentFlags().SetAnnotation("start", cobra.BashCompFilenameExt, []string{"date"})
	RootCmd.PersistentFlags().StringVarP(&End, "end", "e", "", "Limit entries on or before this datetime")
	RootCmd.PersistentFlags().SetAnnotation("end", cobra.BashCompFilenameExt, []string{"date"})
	RootCmd.PersistentFlags().StringVarP(&GroupBy, "group-by", "g", "", "Group entries by this field")
}

var RootCmd = &cobra.Command{
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
				if false {
					if line == "" {
						break
					}
				}
				// Parse the line here
				log := ParseNginxLogLine(line)
				LogEntries = append(LogEntries, log)
			}
		} else {
			// Open the file and parse it line by line
			file, err := os.Open(file)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			linesRead := 0
			for scanner.Scan() {
				line := scanner.Text()
				// Parse the line here
				log := ParseNginxLogLine(line)
				LogEntries = append(LogEntries, log)
				linesRead++
				if numLines > 0 && linesRead >= numLines {
					break
				}
			}
		}
		fmt.Printf("Count of entries: %d\n", len(LogEntries))
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
func checksum(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}

func ParseNginxLogLine(line string) NginxAccessLog {
	fields := strings.Fields(line)

	timestampStr := strings.Trim(fields[3], "[]") + " " + strings.Trim(fields[4], "[]")
	// timestamp, err := time.Parse("02/Jan/2006:15:04:05 -0700", timestampStr)
	timestamp, err := dateparse.ParseAny(timestampStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing timestamp: %v\n", err)
	}
	log := NginxAccessLog{
		IP:              fields[0],
		Timestamp:       timestamp,
		StatusCode:      fields[8],
		BytesSent:       strings.Trim(fields[9], "\""),
		RequestMethod:   strings.Trim(fields[5], "\""),
		RequestURL:      fields[6],
		RequestProtocol: strings.Trim(fields[7], "\""),
		Referrer:        strings.Trim(fields[10], "\""),
		UserAgent:       strings.Join(fields[11:len(fields)-1], " "),
		Checksum:        checksum(line),
	}
	return log
}
var exitCmd = &cobra.Command{
	Use:     "exit",
	Aliases: []string{"q", "Q", "bye"},
	Short:   "Exit the application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Goodbye!")
		os.Exit(0)
	},
}

func init() {
	RootCmd.AddCommand(exitCmd)
}
