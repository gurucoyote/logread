# LogRead - Nginx Logfile Reader

LogRead is a powerful tool written in Go that allows you to read, parse, and analyze Nginx access.log files. It provides useful insights and analytics from your log files, making it easier to understand your server's traffic patterns and troubleshoot issues.

This tool was initially developed in under a day, with the assistance of Aider and GPT4. It was primarily designed to address my personal needs and provide a quick way to gain insights into the activities on a webserver.

While LogRead can be used in a standard mode for one-off actions, its true potential is unlocked when used in interactive mode (-i). This mode allows you to load a set of log entries from a file or stdin and perform various analyses on it, providing a more in-depth understanding of your server's traffic.

## Usage

To use LogRead, you can specify a log file and flags as arguments:

```
logread [file] [flags]
```

You can also use LogRead with a command:

```
logread [command]
```

## Available Commands

- `completion`: Generate the autocompletion script for the specified shell
- `exit`: Exit the application
- `geoip`: Get geolocation information for an IP address
- `help`: Help about any command
- `json`: Output log entries as JSON
- `set`: Set the value of a flag
- `status`: Display metrics about the currently loaded set of log entries
- `top`: Display the top n entries for a specified field

## Flags

- `-e, --end string`: Limit entries on or before this datetime
- `-g, --group-by string`: Group entries by this field
- `-h, --help`: Help for logread
- `-i, --interactive`: Enable interactive mode
- `-n, --number-lines int`: Number of lines to read from the file or stdin
- `-s, --start string`: Limit entries on or after this datetime

For more information about a specific command, use:

```
logread [command] --help
```
