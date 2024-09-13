package comprobante

import (
	"encoding/xml"
	comprobante2 "github.com/saulenriquemr/kore-models/models/comprobante"
	testing2 "github.com/saulenriquemr/kore-models/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetComprobante40ForTest(filename string, t *testing.T) (comprobante2.Comprobante40, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed comprobante2.Comprobante40
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	return parsed, errUnmashal
}

func TestFullComprobante40(t *testing.T) {
	parsed, _ := GetComprobante40ForTest("./comprobante40.xml", t)
	InternalTestFullAttributes40(t, parsed)
	InternalTestFullInformacionGlobal40(t, *parsed.InformacionGlobal)
	InternalTestFullCfdisRelacionados40(t, *parsed.CfdisRelacionados)
	InternalTestFullEmisor40(t, parsed.Emisor)
	InternalTestFullReceptor40(t, parsed.Receptor)
	InternalTestFullConcepto40(t, parsed.Conceptos)
}

func InternalTestFullAttributes40(t *testing.T, parsed comprobante2.Comprobante40) {
	// Check that the required attributes are correctly parsed
	assert.Equal(t, "4.0", parsed.Version)
	assert.Equal(t, "2024-04-29T00:00:55", parsed.Fecha)
	assert.Equal(t, "AMifipYnPS5FuNWEf3ysVCru4FVG0eGp3", parsed.Sello)
	assert.Equal(t, "MIIFsDCCA5igAwIBAgIUMzAwMDE", parsed.Certificado)
	assert.Equal(t, "30001000000500003416", parsed.NoCertificado)
	assert.Equal(t, "MXN", parsed.Moneda)
	assert.Equal(t, "I", parsed.TipoComprobante)
	assert.Equal(t, "01", parsed.Exportacion)
	assert.Equal(t, "20000", parsed.LugarExpedicion)

	// Check float attributes
	assert.Equal(t, 200.00, parsed.Subtotal)
	assert.Equal(t, 199.96, parsed.Total)

	// Optional attributes: assert they're not nil and have expected values
	assert.NotNil(t, parsed.Serie)
	assert.Equal(t, "Serie", *parsed.Serie)

	assert.NotNil(t, parsed.Folio)
	assert.Equal(t, "Folio", *parsed.Folio)

	assert.NotNil(t, parsed.FormaPago)
	assert.Equal(t, "99", *parsed.FormaPago)

	assert.NotNil(t, parsed.CondicionesPago)
	assert.Equal(t, "CondicionesDePago", *parsed.CondicionesPago)

	assert.NotNil(t, parsed.Descuento)
	assert.Equal(t, 200.00, *parsed.Descuento)

	assert.NotNil(t, parsed.TipoCambio)
	assert.Equal(t, 1.0, *parsed.TipoCambio)

	assert.NotNil(t, parsed.MetodoPago)
	assert.Equal(t, "PPD", *parsed.MetodoPago)

	assert.NotNil(t, parsed.Confirmacion)
	assert.Equal(t, "UBDWEIUFIRGBYGIEFEWHFEURGHERIFHREU", *parsed.Confirmacion)
}

func InternalTestFullInformacionGlobal40(t *testing.T, informacionGlobal comprobante2.InformacionGlobal40) {
	assert.Equal(t, "01", informacionGlobal.Periodicidad)
	assert.Equal(t, "08", informacionGlobal.Meses)
	assert.Equal(t, "2024", informacionGlobal.Anio)
}

func InternalTestFullCfdisRelacionados40(t *testing.T, parsed []comprobante2.CfdisRelacionados40) {
	assert.NotNil(t, parsed)
	assert.Equal(t, 2, len(parsed))

	first := parsed[0]
	assert.Equal(t, "01", first.TipoRelacion)
	uuidsFirst := first.UuidRelacionados
	assert.Equal(t, 2, len(uuidsFirst))
	assert.Equal(t, "09755d3a-e6fe-4863-b1f0-ec08a5883324", uuidsFirst[0].UUID)
	assert.Equal(t, "709a6642-acac-4b09-a7ae-1c667a1173f8", uuidsFirst[1].UUID)

	second := parsed[1]
	assert.Equal(t, "02", second.TipoRelacion)
	uuidsSecond := second.UuidRelacionados
	assert.Equal(t, 1, len(uuidsSecond))
	assert.Equal(t, "989427db-2363-4291-8c31-f7b84223406f", uuidsSecond[0].UUID)
}

func InternalTestFullEmisor40(t *testing.T, parsed comprobante2.Emisor40) {
	assert.Equal(t, "EKU9003173C9", parsed.Rfc)
	assert.Equal(t, "ESCUELA KEMPER URGATE", parsed.Nombre)
	assert.Equal(t, "601", parsed.RegimenFiscal)

	assert.NotNil(t, parsed.FactAtrAdquirent)
	assert.Equal(t, "ABCBSCBA", *parsed.FactAtrAdquirent)
}

func InternalTestFullReceptor40(t *testing.T, parsed comprobante2.Receptor40) {
	assert.Equal(t, "URE180429TM6", parsed.Rfc)
	assert.Equal(t, "UNIVERSIDAD ROBOTICA ESPAÑOLA", parsed.Nombre)
	assert.Equal(t, "86991", parsed.DomicilioFiscal)
	assert.Equal(t, "601", parsed.RegimenFiscal)
	assert.Equal(t, "G01", parsed.UsoCFDI)

	assert.NotNil(t, parsed.ResidenciaFiscal)
	assert.Equal(t, "MEX", *parsed.ResidenciaFiscal)

	assert.NotNil(t, parsed.NumRegIdTrib)
	assert.Equal(t, "1020321", *parsed.NumRegIdTrib)
}

func InternalTestFullConcepto40(t *testing.T, parsed []comprobante2.Concepto40) {
	assert.Equal(t, 1, len(parsed))

	concepto := parsed[0]
	InternalTestConcepto40Attributes(t, concepto)
}

func InternalTestConcepto40Attributes(t *testing.T, concepto comprobante2.Concepto40) {
	assert.Equal(t, "50211503", concepto.ClaveProdServ)
	assert.Equal(t, 1.0, concepto.Cantidad)
	assert.Equal(t, "H87", concepto.ClaveUnidad)
	assert.Equal(t, "Cigarros", concepto.Descripcion)
	assert.Equal(t, 200.0, concepto.ValorUnitario)
	assert.Equal(t, 200.0, concepto.Importe)
	assert.Equal(t, "02", concepto.ObjetoImp)

	assert.NotNil(t, concepto.NoIdentificacion)
	assert.Equal(t, "ABC", *concepto.NoIdentificacion)

	assert.NotNil(t, concepto.Unidad)
	assert.Equal(t, "Pieza", *concepto.Unidad)

	assert.NotNil(t, concepto.Descuento)
	assert.Equal(t, 100.0, *concepto.Descuento)
}

func InternalTestConceptoImpuestos40(t *testing.T, concepto comprobante2.Concepto40) {

}
