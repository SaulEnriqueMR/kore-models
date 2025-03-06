package helpers

import (
	"github.com/SaulEnriqueMR/kore-models/models/comprobante"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestUnmarshalXMLWithEncoding(t *testing.T) {
	// Open the XML file
	file, errOpen := os.Open("./comprobante33-ISO-8859-1.xml")
	if errOpen != nil {
		t.Error("Error opening file: ", errOpen)
	}
	defer file.Close()

	// Read the contents of the file
	data, errRead := io.ReadAll(file)
	if errRead != nil {
		t.Error("Error reading file:", errRead)
	}

	var comprobante33 comprobante.Comprobante33
	errUnmarshal := helpers.UnmarshalXMLWithEncoding(data, &comprobante33)
	assert.Nil(t, errUnmarshal)
	assert.NotEmpty(t, comprobante33)

	testing2.GenerateJSONFromStructure("./comprobante33-ISO-8859-1.json", comprobante33)
}
