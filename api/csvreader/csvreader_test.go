package csvreader

import (
	"encoding/csv"
	"os"
	"testing"

	"github.com/mxnyawi/cache-calc/logger/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestReadCSV(t *testing.T) {
	mockLogger := new(mocks.MockLoggerInterface)
	mockLogger.On("Info", mock.Anything, mock.Anything).Maybe()
	mockLogger.On("Error", mock.Anything, mock.Anything).Maybe()
	tempFile, err := os.CreateTemp(os.TempDir(), "prefix")
	if err != nil {
		t.Fatalf("Cannot create temporary file: %s", err)
	}

	defer os.Remove(tempFile.Name())

	writer := csv.NewWriter(tempFile)
	writer.Write([]string{"header1", "header2"})
	writer.Write([]string{"data1", "data2"})
	writer.Flush()

	data, err := ReadCSV(mockLogger, tempFile.Name())
	mockLogger.AssertExpectations(t)
	if err != nil {
		t.Fatalf("ReadCSV() error = %v", err)
	}

	require.Len(t, data, 2)
	require.Equal(t, []string{"header1", "header2"}, data[0])
	require.Equal(t, []string{"data1", "data2"}, data[1])
}
