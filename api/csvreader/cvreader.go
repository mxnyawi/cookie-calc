package csvreader

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/mxnyawi/cache-calc/logger"
)

// ReadCSV reads a CSV file and returns the data as a slice of slices of strings
func ReadCSV(logger logger.LoggerInterface, filePath string) ([][]string, error) {
	logger.Info("Reading CSV file...")
	file, err := os.Open(filePath)
	if err != nil {
		logger.Error("cannot open file: %v", err)
		return nil, fmt.Errorf("cannot open file: %w", err)
	}
	defer file.Close()

	logger.Info("Parsing CSV data...")
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		logger.Error("cannot read CSV data: %v", err)
		return nil, fmt.Errorf("cannot read CSV data: %w", err)
	}

	logger.Info("CSV data read successfully")
	return data, nil
}
