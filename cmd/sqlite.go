package cmd

import (
	"fmt"
	"os"
	"bufio"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

var dbFileName string

var sqliteCmd = &cobra.Command{
	Use:   "sqlite",
	Short: "Print the current filename of the sqlite db",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(dbFileName); os.IsNotExist(err) {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Database file does not exist. Do you want to create it? (y/n): ")
			text, _ := reader.ReadString('\n')
			if text == "y\n" {
				file, err := os.Create(dbFileName)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
				file.Close()
				fmt.Println("Database file created.")

				db, err := sql.Open("sqlite3", dbFileName)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					return
				}
				defer db.Close()

				sqlStmt := `
				CREATE TABLE accesslog (
					IP TEXT,
					Timestamp TEXT,
					StatusCode TEXT,
					BytesSent TEXT,
					RequestMethod TEXT,
					RequestURL TEXT,
					RequestProtocol TEXT,
					Referrer TEXT,
					UserAgent TEXT,
					Checksum TEXT
				);
				`
				_, err = db.Exec(sqlStmt)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					return
				}
				fmt.Println("Accesslog table created.")
			} else {
				fmt.Println("Database file not created.")
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Current SQLite DB filename:", dbFileName)
	},
}

var todbCmd = &cobra.Command{
	Use:   "todb",
	Short: "Write the current log entries to the sqlite db",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := sql.Open("sqlite3", dbFileName)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		defer db.Close()

		stmt, err := db.Prepare("INSERT INTO accesslog(IP, Timestamp, StatusCode, BytesSent, RequestMethod, RequestURL, RequestProtocol, Referrer, UserAgent, Checksum) values(?,?,?,?,?,?,?,?,?,?)")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		defer stmt.Close()

		for _, log := range LogEntries {
			_, err = stmt.Exec(log.IP, log.Timestamp, log.StatusCode, log.BytesSent, log.RequestMethod, log.RequestURL, log.RequestProtocol, log.Referrer, log.UserAgent, log.Checksum)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
		}
	},
}

var fromdbCmd = &cobra.Command{
	Use:   "fromdb",
	Short: "Read log entries from an existing sqlite db",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fromdb called")
		db, err := sql.Open("sqlite3", dbFileName)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		defer db.Close()

		rows, err := db.Query("SELECT * FROM accesslog")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var log NginxAccessLog
			var timestampStr string
			err = rows.Scan(&log.IP, &timestampStr, &log.StatusCode, &log.BytesSent, &log.RequestMethod, &log.RequestURL, &log.RequestProtocol, &log.Referrer, &log.UserAgent, &log.Checksum)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			log.Timestamp, err = time.Parse(time.RFC3339, timestampStr)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			LogEntries = append(LogEntries, log)
			fmt.Printf("Added log entry: %+v\n", log) // Debug print
		}
		err = rows.Err()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		fmt.Printf("Total log entries after fromdb: %d\n", len(LogEntries)) // Debug print
	},
}

func init() {
	sqliteCmd.PersistentFlags().StringVarP(&dbFileName, "database", "d", "./access.db", "SQLite database file name")
	RootCmd.AddCommand(sqliteCmd)
	sqliteCmd.AddCommand(todbCmd)
	sqliteCmd.AddCommand(fromdbCmd)
}
