package intereseshipotecarios

import (
	"encoding/xml"
	"testing"

	intereseshipo1 "github.com/SaulEnriqueMR/kore-models/models/intereseshipotecarios"
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
	InternalTestInteresesHipotecarios(t, intereseshipo)
}
func InternalTestInteresesHipotecarios(t *testing.T, interesesHipotecarios intereseshipo1.InteresesHipotecarios10) {
	assert.NotNil(t, interesesHipotecarios)
	assert.Equal(t, "1.0", interesesHipotecarios.Version)
	assert.Equal(t, "SI", interesesHipotecarios.CreditoDeInstFinanc)
	assert.Equal(t, 1000000.00, interesesHipotecarios.SaldoInsoluto)
	assert.Equal(t, 0.50, *interesesHipotecarios.PropDeducDelCredit)
	assert.Equal(t, 30000.00, *interesesHipotecarios.MontTotIntNominalesDev)
	assert.Equal(t, 25000.00, *interesesHipotecarios.MontTotIntNominalesDevYPag)
	assert.Equal(t, 12000.00, *interesesHipotecarios.MontTotIntRealPagDeduc)
	assert.Equal(t, "ABC123456", *interesesHipotecarios.NumContrato)
}