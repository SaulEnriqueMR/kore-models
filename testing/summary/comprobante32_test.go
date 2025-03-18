package summary

import (
	"bytes"
	"encoding/xml"
	"os"
	"testing"

	summarized "github.com/SaulEnriqueMR/kore-models/models/comprobante"
)

func TestFilename32(t *testing.T) {
	file, err := os.ReadFile("./55ED3268-DB8E-46B6-8BEF-7A1299E76F65_32.xml")
	if err != nil {
		t.Error("Error reading file:", err)
	}
	filename, err := summarized.MyFilename(xml.NewDecoder(bytes.NewReader(file)))
	if err != nil {
		t.Error("Error reading file:", err)
	}

	if filename != "XAXX010101000/XEXX010101000/2017/12/55ED3268-DB8E-46B6-8BEF-7A1299E76F65.xml" {
		t.Error("Filename does not match")
	}
}
