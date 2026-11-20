package summary

import (
	"bytes"
	"encoding/xml"
	"log"
	"os"
	"testing"

	"github.com/SaulEnriqueMR/kore-models/models/retenciones"
)

func TestFilename20(t *testing.T) {
	file, err := os.ReadFile("./fa170e82-b398-4775-a00d-825b72454e3b.xml")
	// file, err := os.ReadFile("./B91669C8-A7BB-46C6-ABF6-CA3D8E90C2E9.xml")
	if err != nil {
		t.Error("Error reading file:", err)
	}
	filename, err := retenciones.MyFilename(xml.NewDecoder(bytes.NewReader(file)))
	if err != nil {
		t.Error("Error reading file:", err)
	}

	log.Println(filename)
	// if filename != "XAXX010101000/XEXX010101000/2017/12/B91669C8-A7BB-46C6-ABF6-CA3D8E90C2E9.xml" {
	// 	t.Error("Filename does not match")
	// }
}
