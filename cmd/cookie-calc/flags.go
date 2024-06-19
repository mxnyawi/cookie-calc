package main

import (
	"flag"
	"fmt"
	"os"
)

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
		fmt.Println("Please provide both -f and -d flags.")
		os.Exit(1)
	}

	return Flags{
		Filepath: *filepath,
		Date:     *date,
	}
}

func Verify() {
	flags := ParseFlags()
	fmt.Println(flags)
}
