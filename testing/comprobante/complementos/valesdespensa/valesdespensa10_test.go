package valesdespensa

import (
	"encoding/xml"
	"testing"

	valesdespensa1 "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/valesdespensa"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetTurisValesDespensaForTest(filename string, t *testing.T) (valesdespensa1.ValesDespensa10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed valesdespensa1.ValesDespensa10
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	return parsed, errUnmashal
}

func TestFullValesDespensa(t *testing.T) {
	valesdespensa, _ := GetTurisValesDespensaForTest("./valesdespensa10.xml", t)
	assert.NotNil(t, valesdespensa)
	InternalTestBase(t, valesdespensa)
}

func InternalTestBase(t *testing.T, vales valesdespensa1.ValesDespensa10) {
	assert.Equal(t, "1.0", vales.Version)
	assert.Equal(t, "monedero electrónico", vales.TipoOperacion)
	assert.Equal(t, "12345678901234567890", vales.NumeroCuenta)
	assert.Equal(t, "RP1234567890123456", *vales.RegistroPatronal)
	assert.Equal(t, 5000.00, vales.Total)
	InternalTestConceptos1(t, vales.Conceptos)
	InternalTestConceptos2(t, vales.Conceptos)
}

func InternalTestConceptos1(t *testing.T, conceptos []valesdespensa1.Conceptos) {
	assert.NotNil(t, conceptos)
	assert.Equal(t, len(conceptos), 3)
	concepto := conceptos[0]
	assert.Equal(t, "ABC12345678901234567", concepto.Identificador)
	assert.Equal(t, "2024-09-19T14:30:00", concepto.FechaString)
	assert.Equal(t, "ABC123456T56", concepto.Rfc)
	assert.Equal(t, "ABCD890123HMNLLL05", concepto.Curp)
	assert.Equal(t, "Juan Perez", concepto.Nombre)
	assert.Equal(t, "123456789012345", concepto.NumSeguridadSocial)
	assert.Equal(t, 1500.00, concepto.Importe)
}

func InternalTestConceptos2(t *testing.T, conceptos []valesdespensa1.Conceptos) {
	concepto := conceptos[2]
	assert.Equal(t, "XYZ45678901234567890", concepto.Identificador)
	assert.Equal(t, "2024-09-19T15:00:00", concepto.FechaString)
	assert.Equal(t, "XYZ456789G21", concepto.Rfc)
	assert.Equal(t, "XYZD890123HMNLLL07", concepto.Curp)
	assert.Equal(t, "Carlos Sanchez", concepto.Nombre)
	assert.Equal(t, "987654321098766", concepto.NumSeguridadSocial)
	assert.Equal(t, 1500.00, concepto.Importe)
}
