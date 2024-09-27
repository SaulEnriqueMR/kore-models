package intereseshipotecarios

import (
	"encoding/xml"
	"testing"

	intereseshipo1 "github.com/SaulEnriqueMR/kore-models/models/retenciones/complementos/intereseshipotecarios"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetTurisInteresesHipoForTest(filename string, t *testing.T) (intereseshipo1.InteresesHipotecarios10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed intereseshipo1.InteresesHipotecarios10
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	return parsed, errUnmashal
}

func TestFullIntereses(t *testing.T) {
	intereseshipo, _ := GetTurisInteresesHipoForTest("./intereseshipotecario10.xml", t)
	assert.NotNil(t, intereseshipo)
	testing2.GenerateJSONFromStructure("intereseshipotecario10.json", intereseshipo)
	InternalTestInteresesHipotecarios(t, intereseshipo)
}
func InternalTestInteresesHipotecarios(t *testing.T, interesesHipotecarios intereseshipo1.InteresesHipotecarios10) {
	assert.NotNil(t, interesesHipotecarios)
	assert.Equal(t, "1.0", interesesHipotecarios.Version)
	assert.Equal(t, "SI", interesesHipotecarios.CreditoOtorgadoInstitucionFinanciera)
	assert.Equal(t, true, interesesHipotecarios.EsCreditoOtorgadoInstitucionFinanciera)
	assert.Equal(t, 1000000.00, interesesHipotecarios.SaldoInsoluto)
	assert.Equal(t, 0.50, *interesesHipotecarios.ProporcionDeducibleCredito)
	assert.Equal(t, 30000.00, *interesesHipotecarios.MontoTotalInteresesNominalesDevengados)
	assert.Equal(t, 25000.00, *interesesHipotecarios.MontoTotalInteresesNominalesDevengadosYPagados)
	assert.Equal(t, 12000.00, *interesesHipotecarios.MontoTotalInteresesRealesPagsDeducibles)
	assert.Equal(t, "ABC123456", *interesesHipotecarios.NoContrato)
}
