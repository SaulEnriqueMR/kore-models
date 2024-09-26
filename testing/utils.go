package testing

import (
	"encoding/json"
	"io"
	"log"
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

func GenerateJSONFromStructure(namefile string, data any) {
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Println(err)
	}
	err1 := os.WriteFile(namefile, jsonData, 0644)
	if err1 != nil {
		log.Println(err1)
	}
}
