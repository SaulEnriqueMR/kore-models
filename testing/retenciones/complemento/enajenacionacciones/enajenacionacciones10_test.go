package enajenacionacciones

import (
	"encoding/xml"
	"testing"

	enajacciones1 "github.com/SaulEnriqueMR/kore-models/models/retenciones/complementos/enajenaciondeacciones"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetTurisEnajenacionForTest(filename string, t *testing.T) (enajacciones1.EnajenacionDeAcciones10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed enajacciones1.EnajenacionDeAcciones10
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	return parsed, errUnmashal
}

func TestFullEnajenacionAcciones(t *testing.T) {
	enajenacion, _ := GetTurisEnajenacionForTest("./enajenacionacciones10.xml", t)
	assert.NotNil(t, enajenacion)
	testing2.GenerateJSONFromStructure("enajenacionacciones10.json", enajenacion)

	InternalTestEnajenacionDeAcciones(t, enajenacion)
}

func InternalTestEnajenacionDeAcciones(t *testing.T, enajenacion enajacciones1.EnajenacionDeAcciones10) {
	assert.Equal(t, "1.0", enajenacion.Version)
	assert.Equal(t, "Contrato de intermediación de acciones de la empresa XYZ", enajenacion.ContratoIntermediacion)
	assert.Equal(t, 5000.00, enajenacion.Ganancia)
	assert.Equal(t, 1500.50, enajenacion.Perdida)
}
