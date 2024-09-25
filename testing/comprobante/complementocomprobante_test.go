package comprobante

import (
	"encoding/xml"
	comprobantes "github.com/SaulEnriqueMR/kore-models/models/comprobante"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetComprobante32ForTestComplementoTest(filename string, t *testing.T) (comprobantes.Comprobante32, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed comprobantes.Comprobante32
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	return parsed, errUnmarshal
}

func GetComprobante33ForTestComplementoTest(filename string, t *testing.T) (comprobantes.Comprobante33, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed comprobantes.Comprobante33
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	return parsed, errUnmarshal
}
func TestComprobanteComplementos(t *testing.T) {
	InternalTestComplementoComprobante32(t)
	InternalTestComplementoComprobante33(t)
}

func InternalTestComplementoComprobante32(t *testing.T) {
	comprobante32, _ := GetComprobante32ForTestComplementoTest("./comprobante32.xml", t)
	InternalTestFullAtributesComplementoComprobantes(t, *comprobante32.Complemento)
}

func InternalTestComplementoComprobante33(t *testing.T) {
	comprobante33, _ := GetComprobante33ForTestComplementoTest("./comprobante33.xml", t)
	InternalTestFullAtributesComplementoComprobantes(t, *comprobante33.Complemento)
}

func InternalTestFullAtributesComplementoComprobantes(t *testing.T, complemento comprobantes.Complemento) {
	assert.NotNil(t, complemento)
	assert.Equal(t, 5, len(*complemento.ImpuestoLocales))
	assert.Equal(t, 7, len(*complemento.ComercioExterior))
	assert.Equal(t, 6, len(*complemento.CartaPorte))
}
