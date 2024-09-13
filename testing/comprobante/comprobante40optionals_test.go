package comprobante

import (
	comprobante2 "github.com/saulenriquemr/koremodels/models/comprobante"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionalComprobante40(t *testing.T) {
	parsed, _ := GetComprobante40ForTest("./comprobante40_requeridos.xml", t)
	InternalTestOptionalAttributes40(t, parsed)
	InternalTestInformacionGlobal40ptional(t, parsed.InformacionGlobal)
	InternalTestCfdisRelacionados4Optional(t, parsed.CfdisRelacionados)
	InternalTestEmisor40Optional(t, parsed.Emisor)
	InternalTestReceptor40Optional(t, parsed.Receptor)
	InternalTestConcepto40Optional(t, parsed.Conceptos)
}

func InternalTestOptionalAttributes40(t *testing.T, parsed comprobante2.Comprobante40) {
	assert.Nil(t, parsed.Serie)
	assert.Nil(t, parsed.Folio)
	assert.Nil(t, parsed.FormaPago)
	assert.Nil(t, parsed.CondicionesPago)
	assert.Nil(t, parsed.Descuento)
	assert.Nil(t, parsed.TipoCambio)
	assert.Nil(t, parsed.MetodoPago)
	assert.Nil(t, parsed.Confirmacion)
}

func InternalTestInformacionGlobal40ptional(t *testing.T, parsed *comprobante2.InformacionGlobal40) {
	assert.Nil(t, parsed)
}

func InternalTestCfdisRelacionados4Optional(t *testing.T, parsed *[]comprobante2.CfdisRelacionados40) {
	assert.Nil(t, parsed)
}

func InternalTestEmisor40Optional(t *testing.T, parsed comprobante2.Emisor40) {
	assert.Nil(t, parsed.FactAtrAdquirent)
}

func InternalTestReceptor40Optional(t *testing.T, parsed comprobante2.Receptor40) {
	assert.Nil(t, parsed.ResidenciaFiscal)
	assert.Nil(t, parsed.NumRegIdTrib)
}

func InternalTestConcepto40Optional(t *testing.T, parsed []comprobante2.Concepto40) {
	concepto := parsed[0]
	InternalTestConcepto40OptionalAttributes(t, concepto)
}

func InternalTestConcepto40OptionalAttributes(t *testing.T, concepto comprobante2.Concepto40) {
	assert.Nil(t, concepto.NoIdentificacion)
	assert.Nil(t, concepto.Unidad)
	assert.Nil(t, concepto.Descuento)
}
