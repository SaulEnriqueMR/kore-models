package comprobante

import (
	"encoding/xml"
	comprobantes "github.com/SaulEnriqueMR/kore-models/models/comprobante"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetComprobante32FortTestComplementoConceptoTest(filename string, t *testing.T) (comprobantes.Comprobante32, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed comprobantes.Comprobante32
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	return parsed, errUnmarshal
}

func GetComprobante33ForTestComplementoConceptoTest(filename string, t *testing.T) (comprobantes.Comprobante33, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed comprobantes.Comprobante33
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	return parsed, errUnmarshal
}

func TestComplementoConcepto(t *testing.T) {
	InternalTestComplementoConceptoComprobante32(t)
	InternalTestComplementoConceptoComprobante33(t)
}

func InternalTestComplementoConceptoComprobante32(t *testing.T) {
	comprobante32, _ := GetComprobante32FortTestComplementoConceptoTest("./comprobante32.xml", t)
	InternalTestFullAtributesComplementoConcepto(t, *comprobante32.Conceptos[0].ComplementoConcepto)
}

func InternalTestComplementoConceptoComprobante33(t *testing.T) {
	comprobante33, _ := GetComprobante33ForTestComplementoConceptoTest("./comprobante33.xml", t)
	InternalTestFullAtributesComplementoConcepto(t, *comprobante33.Conceptos[0].ComplementoConcepto)
}

func InternalTestFullAtributesComplementoConcepto(t *testing.T, complementoConcepto comprobantes.ComplementoConcepto) {
	assert.NotNil(t, complementoConcepto)
	assert.Equal(t, 3, len(*complementoConcepto.AcreditamientoIeps.AcreditamientoIeps10))
	assert.Equal(t, 3, len(*complementoConcepto.InstitucionesEducativas.InstitucioneEducativas10))
	assert.Equal(t, 3, len(*complementoConcepto.ACuentaTerceros.ACuentaTerceros11))
	assert.Equal(t, 2, len(*complementoConcepto.VentaVehiculos.VentaVehiculos11))
	assert.Equal(t, 2, len(*complementoConcepto.VentaVehiculos.VentaVehiculos10))
}
