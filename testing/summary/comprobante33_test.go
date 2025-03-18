package summary

import (
	"bytes"
	"encoding/xml"
	"os"
	"testing"

	summarized "github.com/SaulEnriqueMR/kore-models/models/comprobante"
)

func TestFilename33(t *testing.T) {
	file, err := os.ReadFile("./1362A081-72FD-4631-8573-58B5DB6E2EA9_33.xml")
	if err != nil {
		t.Error("Error reading file:", err)
	}

	filename, err := summarized.MyFilename(xml.NewDecoder(bytes.NewReader(file)))
	if err != nil {
		t.Error("Error reading file:", err)
	}
	if filename != "XAXX010101001/XAXX010101000/2023/4/1362A081-72FD-4631-8573-58B5DB6E2EA9.xml" {
		t.Log(filename)
		t.Error("Filename does not match")
	}
}
