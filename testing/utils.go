package testing

import (
	"io"
	"os"
	"testing"
)

func GetFileContentForTest(filename string, t *testing.T) []byte {
	// Open the XML file
	file, errOpen := os.Open(filename)
	if errOpen != nil {
		t.Error("Error opening file: ", errOpen)
	}
	defer file.Close()

	// Read the contents of the file
	data, errRead := io.ReadAll(file)
	if errRead != nil {
		t.Error("Error reading file:", errRead)
	}
	return data
}
