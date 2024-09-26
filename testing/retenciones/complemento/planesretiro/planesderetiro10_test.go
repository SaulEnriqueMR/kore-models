package retenciones

import (
	"encoding/xml"
	"testing"

	planretiro10 "github.com/SaulEnriqueMR/kore-models/models/retenciones/complementos/planesretiro"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetPlanesDeRetiro10ForTest(filename string, t *testing.T) (planretiro10.PlanesDeRetiro10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed planretiro10.PlanesDeRetiro10
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	return parsed, errUnmarshal
}

func TestPlanesDeRetiro10(t *testing.T) {
	planesderetiro10, _ := GetPlanesDeRetiro10ForTest("./planesderetiro10.xml", t)
	InternalTestFullAtributesPlanesDeRetiro10(t, planesderetiro10)
}

func InternalTestFullAtributesPlanesDeRetiro10(t *testing.T, retiro10 planretiro10.PlanesDeRetiro10) {
	assert.Equal(t, "1.0", retiro10.Version)
	assert.Equal(t, "SI", retiro10.SistemaFinanc)
	assert.NotNil(t, *retiro10.MontTotAportAnioInmAnterior)
	assert.Equal(t, 1000.00, *retiro10.MontTotAportAnioInmAnterior)
	assert.Equal(t, 100.00, retiro10.MontIntRealesDevengAniooInmAnt)
	assert.Equal(t, "SI", retiro10.HuboRetirosAnioInmAntPer)
	assert.NotNil(t, *retiro10.MontTotRetiradoAnioInmAntPer)
	assert.Equal(t, 10.00, *retiro10.MontTotRetiradoAnioInmAntPer)
	assert.NotNil(t, *retiro10.MontTotExentRetiradoAnioInmAnt)
	assert.Equal(t, 90.00, *retiro10.MontTotExentRetiradoAnioInmAnt)
	assert.NotNil(t, *retiro10.MontTotExedenteAnioInmAnt)
	assert.Equal(t, 900.00, *retiro10.MontTotExedenteAnioInmAnt)
	assert.Equal(t, "SI", retiro10.HuboRetirosAnioInmAnt)
	assert.NotNil(t, *retiro10.MontTotRetiradoAnioInmAnt)
	assert.Equal(t, 90.00, *retiro10.MontTotRetiradoAnioInmAnt)
}
