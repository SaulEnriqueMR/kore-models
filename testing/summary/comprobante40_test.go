package summary

import (
	"bytes"
	"encoding/xml"
	"os"
	"testing"

	summarized "github.com/SaulEnriqueMR/kore-models/models/comprobante"
)

func TestFilename40(t *testing.T) {
	file, err := os.ReadFile("./B3863C54-754E-555C-B282-7E4D47EE2EAA_40.xml")
	if err != nil {
		t.Error("Error reading xile:", err)
	}

	filename, err := summarized.MyFilename(xml.NewDecoder(bytes.NewReader(file)))
	if err != nil {
		t.Error("Error reading file:", err)
	}

	if filename != "XAXX010101001/XAXX010101000/1/1/B3863C54-754E-555C-B282-7E4D47EE2EAA.xml" {
		t.Error("Filename does not match")
	}
}
