package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

	DateValidation(*date)

	return Flags{
		Filepath:       *filepath,
		Date:           *date,
		LoggingEnabled: *loggingEnabled,
	}
}

func DateValidation(date string) {
	re := regexp.MustCompile(`^(\d{4})-(\d{2})-(\d{2})$`)
	matches := re.FindStringSubmatch(date)

	if matches == nil {
		log.Fatalln("Please provide a valid date in the form of YYYY-MM-DD.")
	} else {
		_, err1 := strconv.Atoi(matches[1])
		month, err2 := strconv.Atoi(matches[2])
		day, err3 := strconv.Atoi(matches[3])

		if err1 != nil || err2 != nil || err3 != nil || month > 12 || day > 30 {
			log.Fatalln("Please provide a valid date with month <= 12 and day <= 30.")
		}
	}
}
