package implocales

import (
	"encoding/xml"
	"testing"

	implocales10 "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/impuestoslocales"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetImpuestosLocales10ForTest(filename string, t *testing.T) (implocales10.ImpuestoLocales10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed implocales10.ImpuestoLocales10
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	testing2.GenerateJSONFromStructure("implocales10.json", parsed)
	return parsed, errUnmashal
}

func TestFullImpLocales10(t *testing.T) {
	implocales, _ := GetImpuestosLocales10ForTest("./implocales.xml", t)
	InternalTestFullAttributes(t, implocales)
	InternalTestFullRetencionesLocales(t, implocales.Retenciones)
	InternalTestFullTrasladosLocales(t, implocales.Traslados)
}

func InternalTestFullAttributes(t *testing.T, implocales implocales10.ImpuestoLocales10) {
	assert.Equal(t, "1.0", implocales.Version)
	assert.Equal(t, 1500.00, implocales.TotalRetenciones)
	assert.Equal(t, 3000.00, implocales.TotalTraslados)

}

func InternalTestFullRetencionesLocales(t *testing.T, implocales *[]implocales10.RetencionesLocalesImpLoc10) {
	assert.NotNil(t, implocales)
	first := (*implocales)[0]
	assert.Equal(t, "ISR", first.Impuesto)
	assert.Equal(t, 1500.00, first.Importe)
	assert.Equal(t, 15.00, first.Tasa)
}

func InternalTestFullTrasladosLocales(t *testing.T, implocales *[]implocales10.TrasladosLocalesImpLoc10) {
	assert.NotNil(t, implocales)
	first := (*implocales)[0]
	assert.Equal(t, "IVA", first.Impuesto)
	assert.Equal(t, 3000.00, first.Importe)
	assert.Equal(t, 16.00, first.Tasa)

}
