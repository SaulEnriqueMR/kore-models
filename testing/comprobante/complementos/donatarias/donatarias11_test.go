package complementos

import (
	"encoding/xml"
	"testing"

	donatarias2 "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/donatarias"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetDonatarias11ForTest(filename string, t *testing.T) (donatarias2.Donatarias11, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed donatarias2.Donatarias11
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	testing2.GenerateJSONFromStructure("donatarias11.json", parsed)
	return parsed, errUnmarshal
}

func TestDonatarias11(t *testing.T) {
	donatarias11, _ := GetDonatarias11ForTest("./donatarias11.xml", t)
	InternalTestFullAtributesDonatarias11(t, donatarias11)
}

func InternalTestFullAtributesDonatarias11(t *testing.T, donatarias donatarias2.Donatarias11) {
	assert.Equal(t, "1.1", donatarias.Version)
	assert.Equal(t, "SP-DF-12345", donatarias.NoAutorizacion)
	assert.Equal(t, "2014-01-01", donatarias.FechaAutString)
	assert.Equal(t, "Este importe corresponde a una donación", donatarias.Leyenda)
}
