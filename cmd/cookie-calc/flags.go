package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// Flags represents the command line flags
type Flags struct {
	Filepath       string
	Date           string
	LoggingEnabled bool
}

// ParseFlags parses the command line flags
func ParseFlags() Flags {
	filepath := flag.String("f", "", "CSV file path")
	date := flag.String("d", "", "UTC date")
	loggingEnabled := flag.Bool("log", false, "Enable logging")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "  -f string\n\tCSV file path\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  -d string\n\tUTC date\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  -log\n\tEnable logging\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  -h\n\tShow help message\n")
	}

	flag.Parse()

	if flag.Lookup("h") != nil {
		flag.Usage()
		os.Exit(0)
	}

	if *filepath == "" || *date == "" {
		log.Fatalln("Please provide both -f and -d flags.")
	}

	return Flags{
		Filepath:       *filepath,
		Date:           *date,
		LoggingEnabled: *loggingEnabled,
	}
}
