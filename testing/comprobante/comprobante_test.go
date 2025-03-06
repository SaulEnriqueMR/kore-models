package comprobante

import (
	comprobante2 "github.com/SaulEnriqueMR/kore-models/models/comprobante"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func GetFileForTest(filename string, t *testing.T) []byte {
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

func TestComprobante(t *testing.T) {
	bytesComprobante32 := GetFileForTest("./comprobante32.xml", t)
	comprobante32, errComprobante32 := comprobante2.SerializeComprobanteFromXml(bytesComprobante32)
	assert.Nil(t, errComprobante32)
	assert.NotNil(t, comprobante32)

	assert.NotNil(t, comprobante32.Comprobante32)
	assert.Nil(t, comprobante32.Comprobante33)
	assert.Nil(t, comprobante32.Comprobante40)

	bytesComprobante33 := GetFileForTest("./comprobante33.xml", t)
	comprobante33, errComprobante33 := comprobante2.SerializeComprobanteFromXml(bytesComprobante33)
	assert.Nil(t, errComprobante33)
	assert.NotNil(t, comprobante33)

	assert.Nil(t, comprobante33.Comprobante32)
	assert.NotNil(t, comprobante33.Comprobante33)
	assert.Nil(t, comprobante33.Comprobante40)

	bytesComprobante40 := GetFileForTest("./comprobante40.xml", t)
	comprobante40, errComprobante40 := comprobante2.SerializeComprobanteFromXml(bytesComprobante40)
	assert.Nil(t, errComprobante40)
	assert.NotNil(t, comprobante40)

	assert.Nil(t, comprobante40.Comprobante32)
	assert.Nil(t, comprobante40.Comprobante33)
	assert.NotNil(t, comprobante40.Comprobante40)
}
