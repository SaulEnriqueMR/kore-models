package comprobante

import (
	"encoding/xml"
	comprobante2 "github.com/SaulEnriqueMR/kore-models/models/comprobante"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetComprobante33ForTestTotals(filename string, t *testing.T) (comprobante2.Comprobante33, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed comprobante2.Comprobante33
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	testing2.GenerateJSONFromStructure("comprobante33totals.json", parsed)
	return parsed, errUnmashal
}

func TestFullComprobante33Totals(t *testing.T) {
	c33, _ := GetComprobante33ForTestTotals("./comprobante33totals.xml", t)
	assert.Equal(t, 4000.00, c33.TotalesMonedaLocal.Total)
	assert.Equal(t, 4000.00, c33.TotalesMonedaLocal.Subtotal)

	assert.Equal(t, 2000.00, *c33.TotalesMonedaLocal.Descuento)

	assert.Equal(t, 2000.00, *c33.TotalesMonedaLocal.TotalImpuestosTrasladados)
	assert.Nil(t, c33.TotalesMonedaLocal.TotalImpuestosRetenidos)
}
