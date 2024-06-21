package main

import (
	"github.com/mxnyawi/cache-calc/api/csvreader"
	"github.com/mxnyawi/cache-calc/logger"
	"github.com/mxnyawi/cache-calc/pkg/calculator"
)

func main() {
	stdOut := logger.StdOut()
	flags := ParseFlags()

	var Logger logger.LoggerInterface
	if flags.LoggingEnabled {
		Logger = logger.SetupLogger()
		defer Logger.Close()
		Logger.Info("Application started")
	} else {
		Logger = logger.NewNoOpLogger()
	}

	data, err := csvreader.ReadCSV(Logger, flags.Filepath)
	if err != nil {
		stdOut.Fatalf("Error reading CSV file: %v", err)
		return
	}

	result, err := calculator.Calculate(Logger, data, flags.Date)
	switch {
	case err == calculator.ErrInvalidData:
		stdOut.Fatalf("Invalid data found in the CSV file")
	case err == calculator.ErrNoCookies:
		stdOut.Fatalf("No cookies found for %v", flags.Date)
	case err != nil:
		stdOut.Fatalf("Error calculating most frequent cookies: %v", err)
	default:
		for _, cookie := range result {
			stdOut.Println(cookie)
		}
	}
}
