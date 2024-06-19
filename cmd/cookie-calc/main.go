package main

import (
	"log"

	"github.com/mxnyawi/cache-calc/api/csvreader"
	"github.com/mxnyawi/cache-calc/pkg/calculator"
)

func main() {
	flags := ParseFlags()
	log.Printf("Calculating most frequent cookies for %v on date %v\n", flags.Filepath, flags.Date)

	data, err := csvreader.ReadCSV(flags.Filepath)
	if err != nil {
		log.Fatalf("Error reading CSV file: %v", err)
		return
	}

	result, err := calculator.Calculate(data, flags.Date)
	switch {
	case err == calculator.ErrInvalidData:
		log.Println("Invalid data found in the CSV file")
	case err == calculator.ErrNoCookies:
		log.Printf("No cookies found for %v", flags.Date)
	case err != nil:
		log.Fatalf("Error calculating most frequent cookies: %v", err)
	default:
		for _, cookie := range result {
			log.Println(cookie)
		}
	}
}
