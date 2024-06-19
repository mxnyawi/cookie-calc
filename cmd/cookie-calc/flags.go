package main

import (
	"flag"
	"log"
)

// Flags represents the command line flags
type Flags struct {
	Filepath string
	Date     string
}

// ParseFlags parses the command line flags
func ParseFlags() Flags {
	filepath := flag.String("f", "", "CSV file path")
	date := flag.String("d", "", "UTC date")

	flag.Parse()

	if *filepath == "" || *date == "" {
		log.Fatalln("Please provide both -f and -d flags.")
	}

	return Flags{
		Filepath: *filepath,
		Date:     *date,
	}
}
