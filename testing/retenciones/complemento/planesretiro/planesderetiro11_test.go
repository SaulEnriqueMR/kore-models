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
	assert.Equal(t, "SI", retiro11.SistemaFinanc)
	assert.NotNil(t, *retiro11.MontTotAportAnioInmAnterior)
	assert.Equal(t, 1000.00, *retiro11.MontTotAportAnioInmAnterior)
	assert.Equal(t, 100.00, retiro11.MontIntRealesDevengAniooInmAnt)
	assert.Equal(t, "SI", retiro11.HuboRetirosAnioInmAntPer)
	assert.NotNil(t, *retiro11.MontTotRetiradoAnioInmAntPer)
	assert.Equal(t, 10.00, *retiro11.MontTotRetiradoAnioInmAntPer)
	assert.NotNil(t, *retiro11.MontTotExentRetiradoAnioInmAnt)
	assert.Equal(t, 90.00, *retiro11.MontTotExentRetiradoAnioInmAnt)
	assert.NotNil(t, *retiro11.MontTotExedenteAnioInmAnt)
	assert.Equal(t, 900.00, *retiro11.MontTotExedenteAnioInmAnt)
	assert.Equal(t, "SI", retiro11.HuboRetirosAnioInmAnt)
	assert.NotNil(t, *retiro11.MontTotRetiradoAnioInmAnt)
	assert.Equal(t, 90.00, *retiro11.MontTotRetiradoAnioInmAnt)
	assert.NotNil(t, *retiro11.NumReferencia)
	assert.Equal(t, "123456", *retiro11.NumReferencia)
}

func InternalTestFullAtributesAportacionesODepositosPlanes11(t *testing.T, planes11 []planretiro11.AportacionesODepositosPlanes11) {
	assert.Equal(t, 1, len(planes11))
	assert.Equal(t, "a", planes11[0].TipoAportacionODeposito)
	assert.Equal(t, 1750.00, planes11[0].MontAportODep)
	assert.Equal(t, "AAA010101AAA", *planes11[0].RFCFiduciaria)
}
