package complementoconcepto

import (
	"encoding/xml"
	obrasantiguedades2 "github.com/SaulEnriqueMR/kore-models/models/complemento/complementoconcepto"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetObrasAntiguedades10ForTest(filename string, t *testing.T) (obrasantiguedades2.ObrasAntiguedades10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed obrasantiguedades2.ObrasAntiguedades10
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	return parsed, errUnmarshal
}

func TestObrasAntiguedades10(t *testing.T) {
	obrasAntiguedades10, _ := GetObrasAntiguedades10ForTest("./obrasantiguedades10.xml", t)
	assert.Equal(t, "1.0", obrasAntiguedades10.Version)
	assert.Equal(t, "01", obrasAntiguedades10.TipoBien)
	assert.NotNil(t, *obrasAntiguedades10.OtrosTipoBien)
	assert.Equal(t, "12", *obrasAntiguedades10.OtrosTipoBien)
	assert.Equal(t, "01", obrasAntiguedades10.TituloAdquirido)
	assert.NotNil(t, *obrasAntiguedades10.OtrosTitulosAdquiridos)
	assert.Equal(t, "21", *obrasAntiguedades10.OtrosTitulosAdquiridos)
	assert.Equal(t, 10000.00, obrasAntiguedades10.Subtotal)
	assert.Equal(t, 16.00, obrasAntiguedades10.IVA)
	assert.Equal(t, "2016-04-03", obrasAntiguedades10.FechaAdquisicionString)
	assert.Equal(t, "02", obrasAntiguedades10.CaracterísticasDeObraoPieza)
}
