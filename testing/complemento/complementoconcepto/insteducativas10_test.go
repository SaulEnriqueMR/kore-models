package complementoconcepto

import (
	"encoding/xml"
	"testing"

	insteducativas2 "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementosconcepto/institeducativas"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetInstEducativas10forTest(filename string, t *testing.T) (insteducativas2.InstitucioneEducativas10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed insteducativas2.InstitucioneEducativas10
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	return parsed, errUnmashal
}

func TestInsteducativas10(t *testing.T) {
	institucion, _ := GetInstEducativas10forTest("./insteducativas10.xml", t)
	InternalTestFullAtributes(t, institucion)
}

func InternalTestFullAtributes(t *testing.T, institucion insteducativas2.InstitucioneEducativas10) {
	assert.Equal(t, "1.0", institucion.Version)
	assert.Equal(t, "Francisco Hernandez", institucion.NombreAlumno)
	assert.Equal(t, "BAHF840120HOCTRR04", institucion.Curp)
	assert.Equal(t, "Preescolar", institucion.NivelEducativo)
	assert.Equal(t, "CCA29123210", institucion.AutRVOE)
	assert.NotNil(t, institucion.RfcPago)
	assert.Equal(t, "XAXX010101333", *institucion.RfcPago)
}
