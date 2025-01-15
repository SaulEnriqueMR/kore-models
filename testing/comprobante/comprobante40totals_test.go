package comprobante

import (
	"encoding/xml"
	comprobante2 "github.com/SaulEnriqueMR/kore-models/models/comprobante"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetComprobante40ForTestTotals(filename string, t *testing.T) (comprobante2.Comprobante40, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed comprobante2.Comprobante40
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	testing2.GenerateJSONFromStructure("comprobante40totals.json", parsed)
	return parsed, errUnmashal
}

func TestFullComprobante40Totals(t *testing.T) {
	c40, _ := GetComprobante40ForTestTotals("./comprobante40totals.xml", t)
	assert.Equal(t, 4000.00, c40.TotalesMonedaLocal.Total)
	assert.Equal(t, 4000.00, c40.TotalesMonedaLocal.Subtotal)

	assert.Nil(t, c40.TotalesMonedaLocal.Descuento)

	assert.Equal(t, 2000.00, *c40.TotalesMonedaLocal.TotalImpuestosTrasladados)
	assert.Equal(t, 2000.00, *c40.TotalesMonedaLocal.TotalImpuestosRetenidos)
}
