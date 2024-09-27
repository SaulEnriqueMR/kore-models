package acreditamientoieps

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementosconcepto/acreditamientoieps"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetAcreditamientoIeps(filename string, t *testing.T) (acreditamientoieps.AcreditamientoIeps10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed acreditamientoieps.AcreditamientoIeps10
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	testing2.GenerateJSONFromStructure("acreditamientoieps.json", parsed)
	return parsed, errUnmashal
}

func TestAcreditamientoIeps10Full(t *testing.T) {

}
