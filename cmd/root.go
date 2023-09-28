package cmd

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var LogEntries []NginxAccessLog
var numLines int

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

func init() {
	rootCmd.PersistentFlags().IntVarP(&numLines, "number-lines", "n", 0, "Number of lines to read from the file or stdin")
}

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
				if line == "" {
					break
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
			for scanner.Scan() {
				line := scanner.Text()
				// Parse the line here
				log := ParseNginxLogLine(line)
				LogEntries = append(LogEntries, log)
			}
			fmt.Printf("Count of entries: %d\n", len(LogEntries))
			jsonData, err := json.Marshal(LogEntries)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(jsonData))
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
func checksum(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}

func ParseNginxLogLine(line string) NginxAccessLog {
	// sample log line:
	// 95.91.212.234 - - [04/Jun/2021:19:40:16 +0000] "POST /unigui_impfung/quicktermin.dll/HandleEvent HTTP/1.1" 200 30 "https://deineanmeldung.de/unigui_impfung/quicktermin.dll/terminbuchung-impfung/2AC0F343-38D8-4DCC-BB03-EE828465D1AE" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4532.2 Safari/537.36" "-"
	// TODO: see if this actually splits the log fields correctly
	fields := strings.Fields(line)

	// TODO: make sure that the actual date/time format from the log is parsed properly here
	timestampStr := strings.Trim(fields[3], "[]")
	timestamp, err := time.Parse("02/Jan/2006:15:04:05 -0700", timestampStr)
	if err != nil {
		// TODO:do not fatal but output the error to stderr
		log.Fatal(err)
	}
	log := NginxAccessLog{
		IP:              fields[0],
		Timestamp:       timestamp,
		StatusCode:      fields[2],
		BytesSent:       fields[3],
		RequestMethod:   fields[4],
		RequestURL:      fields[5],
		RequestProtocol: fields[6],
		Referrer:        fields[7],
		UserAgent:       fields[8],
		Checksum:        checksum(line),
	}
	return log
}
