package csvreader

import (
	"encoding/csv"
	"fmt"
	"os"
)

// ReadCSV reads a CSV file and returns the data as a slice of slices of strings
func ReadCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("cannot open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("cannot read CSV data: %w", err)
	}

	return data, nil
}
