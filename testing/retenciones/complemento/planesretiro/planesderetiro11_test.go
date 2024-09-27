package retenciones

import (
	"encoding/xml"
	"testing"

	planretiro11 "github.com/SaulEnriqueMR/kore-models/models/retenciones/complementos/planesretiro"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetPlanesDeRetiro11ForTest(filename string, t *testing.T) (planretiro11.PlanesDeRetiro11, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed planretiro11.PlanesDeRetiro11
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	testing2.GenerateJSONFromStructure("planesderetiro11.json", parsed)
	return parsed, errUnmarshal
}

func TestPlanesDeRetiro11(t *testing.T) {
	planesderetiro11, _ := GetPlanesDeRetiro11ForTest("./planesderetiro11.xml", t)
	InternalTestFullAtributesPlanesDeRetiro11(t, planesderetiro11)
	InternalTestFullAtributesAportacionesODepositosPlanes11(t, *planesderetiro11.AportacionesODepositos)
}

func InternalTestFullAtributesPlanesDeRetiro11(t *testing.T, retiro11 planretiro11.PlanesDeRetiro11) {
	assert.Equal(t, "1.1", retiro11.Version)

	assert.Equal(t, "SI", retiro11.SistemaFinanciero)
	assert.Equal(t, true, retiro11.ProvienenSistemaFinanciero)

	assert.NotNil(t, retiro11.MontoTotalAportacionesAnioAnterior)
	assert.Equal(t, 1000.00, *retiro11.MontoTotalAportacionesAnioAnterior)
	assert.Equal(t, 100.00, retiro11.MontoInteresesRealesDevengadosAnioAnterior)

	assert.Equal(t, "SI", retiro11.HuboRetirosAnioInmAntPer)
	assert.Equal(t, true, retiro11.HuboRetirosAnioAnteriorAntesPermanencia)

	assert.NotNil(t, retiro11.MontoTotalRetiradoAnioAnteriorPermanencia)
	assert.Equal(t, 10.00, *retiro11.MontoTotalRetiradoAnioAnteriorPermanencia)
	assert.NotNil(t, retiro11.MontoTotalExentoRetiradoAnioAnterior)
	assert.Equal(t, 90.00, *retiro11.MontoTotalExentoRetiradoAnioAnterior)
	assert.NotNil(t, retiro11.MontoTotalExedenteAnioAnterior)
	assert.Equal(t, 900.00, *retiro11.MontoTotalExedenteAnioAnterior)

	assert.Equal(t, "SI", retiro11.HuboRetirosAnioInmAnt)
	assert.Equal(t, true, retiro11.HuboRetirosAnioAnterior)

	assert.NotNil(t, retiro11.MontoTotalRetiradoAnioAnterior)
	assert.Equal(t, 90.00, *retiro11.MontoTotalRetiradoAnioAnterior)
	assert.NotNil(t, retiro11.NoReferencia)
	assert.Equal(t, "123456", *retiro11.NoReferencia)
}

func InternalTestFullAtributesAportacionesODepositosPlanes11(t *testing.T, planes11 []planretiro11.AportacionesODepositosPlanes11) {
	assert.Equal(t, 1, len(planes11))
	assert.Equal(t, "a", planes11[0].Tipo)
	assert.Equal(t, 1750.00, planes11[0].Monto)
	assert.Equal(t, "AAA010101AAA", *planes11[0].RfcFiduciaria)
}
