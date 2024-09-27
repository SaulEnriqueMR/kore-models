package comprobante

import (
	"encoding/xml"
	comprobante2 "github.com/SaulEnriqueMR/kore-models/models/comprobante"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
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
	comprobante, _ := GetComprobante40ForTest("./comprobante40.xml", t)
	InternalTestFullAttributes40(t, comprobante)
	InternalTestFullInformacionGlobal40(t, comprobante.InformacionGlobal)
	InternalTestFullCfdisRelacionados40(t, comprobante.CfdisRelacionados)
	InternalTestFullEmisor40(t, comprobante.Emisor)
	InternalTestFullReceptor40(t, comprobante.Receptor)
	InternalTestFullConcepto40(t, comprobante.Conceptos)
	InternalTestFullImpuestos40(t, comprobante.Impuestos)
}

func InternalTestFullAttributes40(t *testing.T, comprobante comprobante2.Comprobante40) {
	assert.Equal(t, "4.0", comprobante.Version)
	assert.Equal(t, "2024-04-29T00:00:55", comprobante.Fecha)
	assert.Equal(t, "AMifipYnPS5FuNWEf3ysVCru4FVG0eGp3", comprobante.Sello)
	assert.Equal(t, "MIIFsDCCA5igAwIBAgIUMzAwMDE", comprobante.Certificado)
	assert.Equal(t, "30001000000500003416", comprobante.NoCertificado)
	assert.Equal(t, "MXN", comprobante.Moneda)
	assert.Equal(t, "I", comprobante.TipoComprobante)
	assert.Equal(t, "01", comprobante.Exportacion)
	assert.Equal(t, "20000", comprobante.LugarExpedicion)

	assert.Equal(t, 200.00, comprobante.Subtotal)
	assert.Equal(t, 199.96, comprobante.Total)

	assert.NotNil(t, comprobante.Serie)
	assert.Equal(t, "Serie", *comprobante.Serie)

	assert.NotNil(t, comprobante.Folio)
	assert.Equal(t, "Folio", *comprobante.Folio)

	assert.NotNil(t, comprobante.FormaPago)
	assert.Equal(t, "99", *comprobante.FormaPago)

	assert.NotNil(t, comprobante.CondicionesPago)
	assert.Equal(t, "CondicionesDePago", *comprobante.CondicionesPago)

	assert.NotNil(t, comprobante.Descuento)
	assert.Equal(t, 200.00, *comprobante.Descuento)

	assert.NotNil(t, comprobante.TipoCambio)
	assert.Equal(t, 1.0, *comprobante.TipoCambio)

	assert.NotNil(t, comprobante.MetodoPago)
	assert.Equal(t, "PPD", *comprobante.MetodoPago)

	assert.NotNil(t, comprobante.Confirmacion)
	assert.Equal(t, "UBDWEIUFIRGBYGIEFEWHFEURGHERIFHREU", *comprobante.Confirmacion)
}

func InternalTestFullInformacionGlobal40(t *testing.T, informacionGlobal *comprobante2.InformacionGlobal40) {
	assert.NotNil(t, informacionGlobal)
	assert.Equal(t, "01", informacionGlobal.Periodicidad)
	assert.Equal(t, "08", informacionGlobal.Meses)
	assert.Equal(t, "2024", informacionGlobal.Anio)
}

func InternalTestFullCfdisRelacionados40(t *testing.T, cfdisRelacionados *[]comprobante2.CfdisRelacionados40) {
	assert.NotNil(t, cfdisRelacionados)
	assert.Equal(t, 2, len(*cfdisRelacionados))

	first := (*cfdisRelacionados)[0]
	assert.Equal(t, "01", first.TipoRelacion)
	uuidsFirst := first.UuidRelacionados
	assert.Equal(t, 2, len(uuidsFirst))
	assert.Equal(t, "09755d3a-e6fe-4863-b1f0-ec08a5883324", uuidsFirst[0].UUID)
	assert.Equal(t, "709a6642-acac-4b09-a7ae-1c667a1173f8", uuidsFirst[1].UUID)

	second := (*cfdisRelacionados)[1]
	assert.Equal(t, "02", second.TipoRelacion)
	uuidsSecond := second.UuidRelacionados
	assert.Equal(t, 1, len(uuidsSecond))
	assert.Equal(t, "989427db-2363-4291-8c31-f7b84223406f", uuidsSecond[0].UUID)
}

func InternalTestFullEmisor40(t *testing.T, emisor comprobante2.Emisor40) {
	assert.Equal(t, "EKU9003173C9", emisor.Rfc)
	assert.Equal(t, "ESCUELA KEMPER URGATE", emisor.Nombre)
	assert.Equal(t, "601", emisor.RegimenFiscal)

	assert.NotNil(t, emisor.FactAtrAdquirent)
	assert.Equal(t, "ABCBSCBA", *emisor.FactAtrAdquirent)
}

func InternalTestFullReceptor40(t *testing.T, receptor comprobante2.Receptor40) {
	assert.Equal(t, "URE180429TM6", receptor.Rfc)
	assert.Equal(t, "UNIVERSIDAD ROBOTICA ESPAÑOLA", receptor.Nombre)
	assert.Equal(t, "86991", receptor.DomicilioFiscal)
	assert.Equal(t, "601", receptor.RegimenFiscal)
	assert.Equal(t, "G01", receptor.UsoCFDI)

	assert.NotNil(t, receptor.ResidenciaFiscal)
	assert.Equal(t, "MEX", *receptor.ResidenciaFiscal)

	assert.NotNil(t, receptor.NumRegIdTrib)
	assert.Equal(t, "1020321", *receptor.NumRegIdTrib)
}

func InternalTestFullConcepto40(t *testing.T, conceptos []comprobante2.Concepto40) {
	assert.Equal(t, 1, len(conceptos))

	concepto := conceptos[0]
	InternalTestFullConcepto40Attributes(t, concepto)
	InternalTestFullConceptoImpuestos40(t, concepto)
	InternalTestFullACuentaTerceros40(t, concepto.ACuentaTerceros)
	InternalTestFullInformacionAduaneraConcepto40(t, concepto)
	InternalTestFullCuentaPredial40(t, concepto.CuentaPredial)
	InternalTestFullParte40(t, concepto.Parte)
}

func InternalTestFullConcepto40Attributes(t *testing.T, concepto comprobante2.Concepto40) {
	assert.Equal(t, "50211503", concepto.ClaveProdServ)
	assert.Equal(t, 1.0, concepto.Cantidad)
	assert.Equal(t, "H87", concepto.ClaveUnidad)
	assert.Equal(t, "Cigarros", concepto.Descripcion)
	assert.Equal(t, 200.0, concepto.ValorUnitario)
	assert.Equal(t, 200.0, concepto.Importe)
	assert.Equal(t, "02", concepto.ObjetoImpuesto)

	assert.NotNil(t, concepto.NoIdentificacion)
	assert.Equal(t, "ABC", *concepto.NoIdentificacion)

	assert.NotNil(t, concepto.Unidad)
	assert.Equal(t, "Pieza", *concepto.Unidad)

	assert.NotNil(t, concepto.Descuento)
	assert.Equal(t, 100.0, *concepto.Descuento)
}

func InternalTestFullConceptoImpuestos40(t *testing.T, concepto comprobante2.Concepto40) {
	impuestos := concepto.Impuestos
	assert.NotNil(t, impuestos)

	traslados := concepto.Impuestos.Traslados
	InternalTestFullTrasladosConcepto40(t, traslados)

	retenciones := concepto.Impuestos.Retenciones
	InternalTestFullRetencionesConcepto40(t, retenciones)
}

func InternalTestFullTrasladosConcepto40(t *testing.T, traslados *[]comprobante2.TrasladoConcepto40) {
	assert.NotNil(t, traslados)
	assert.Equal(t, len(*traslados), 1)
	first := (*traslados)[0]
	InternalTestFullSingleTrasladoConcepto40(t, first)
}

func InternalTestFullSingleTrasladoConcepto40(t *testing.T, traslado comprobante2.TrasladoConcepto40) {
	assert.Equal(t, 1.0, traslado.Base)
	assert.Equal(t, "002", traslado.Impuesto)
	assert.Equal(t, "Tasa", traslado.TipoFactor)

	assert.NotNil(t, traslado.TasaOCuota)
	assert.Equal(t, 0.16, *traslado.TasaOCuota)

	assert.NotNil(t, traslado.Importe)
	assert.Equal(t, 0.16, *traslado.Importe)
}

func InternalTestFullRetencionesConcepto40(t *testing.T, retenciones *[]comprobante2.RetencionConcepto40) {
	assert.NotNil(t, retenciones)
	assert.Equal(t, len(*retenciones), 2)

	first := (*retenciones)[0]
	InternalTestFullSingleRetencionConcepto40(t, first)
}

func InternalTestFullSingleRetencionConcepto40(t *testing.T, retencion comprobante2.RetencionConcepto40) {
	assert.Equal(t, 1.0, retencion.Base)
	assert.Equal(t, "001", retencion.Impuesto)
	assert.Equal(t, "Tasa", retencion.TipoFactor)

	assert.NotNil(t, retencion.TasaOCuota)
	assert.Equal(t, 0.10, retencion.TasaOCuota)

	assert.NotNil(t, retencion.Importe)
	assert.Equal(t, 0.10, retencion.Importe)
}

func InternalTestFullACuentaTerceros40(t *testing.T, aCuentaTerceros *comprobante2.ACuentaTerceros40) {
	assert.NotNil(t, aCuentaTerceros)

	assert.Equal(t, "CACX7605101P8", aCuentaTerceros.Rfc)
	assert.Equal(t, "XOCHILT CASAS CHAVEZ", aCuentaTerceros.Nombre)
	assert.Equal(t, "601", aCuentaTerceros.RegimenFiscal)
	assert.Equal(t, "36257", aCuentaTerceros.DomicilioFiscal)
}

func InternalTestFullInformacionAduaneraConcepto40(t *testing.T, concepto comprobante2.Concepto40) {
	assert.NotNil(t, concepto.InformacionAduanera)

	assert.Equal(t, len(*concepto.InformacionAduanera), 2)

	first := (*concepto.InformacionAduanera)[0]
	assert.Equal(t, "123456", first.NumeroPedimento)
	second := (*concepto.InformacionAduanera)[1]
	assert.Equal(t, "789941", second.NumeroPedimento)
}

func InternalTestFullCuentaPredial40(t *testing.T, cuentaPredial *[]comprobante2.CuentaPredial40) {
	assert.NotNil(t, cuentaPredial)

	assert.Equal(t, len(*cuentaPredial), 2)

	first := (*cuentaPredial)[0]
	assert.Equal(t, "aB3cD4eF5gH6iJ7kL8mN9oP0", first.Numero)
	second := (*cuentaPredial)[1]
	assert.Equal(t, "123213FSDFSVSDVSB8mN9oP0", second.Numero)
}

func InternalTestFullParte40(t *testing.T, parte40 *[]comprobante2.Parte40) {
	assert.NotNil(t, parte40)

	assert.Equal(t, len(*parte40), 1)
	first := (*parte40)[0]
	InternalTestFullSingleParte40(t, first)
}

func InternalTestFullSingleParte40(t *testing.T, parte comprobante2.Parte40) {
	assert.Equal(t, "51241200", parte.ClaveProdServ)
	assert.Equal(t, 1.0, parte.Cantidad)
	assert.Equal(t, "ACEITE AJONJOLI", parte.Descripcion)

	assert.NotNil(t, parte.NoIdentificacion)
	assert.Equal(t, "IM0071", *parte.NoIdentificacion)

	assert.NotNil(t, parte.Unidad)
	assert.Equal(t, "Pieza", *parte.Unidad)

	assert.NotNil(t, parte.ValorUnitario)
	assert.Equal(t, 736.89, *parte.ValorUnitario)

	assert.NotNil(t, parte.Importe)
	assert.Equal(t, 736.89, *parte.Importe)

	InternalTestFullInformacionAduaneraParte40(t, parte.InformacionAduanera)
}

func InternalTestFullInformacionAduaneraParte40(t *testing.T, informacionAduanera *[]comprobante2.InformacionAduanera40) {
	assert.NotNil(t, informacionAduanera)

	assert.Equal(t, len(*informacionAduanera), 1)

	first := (*informacionAduanera)[0]
	assert.Equal(t, "345678", first.NumeroPedimento)
}

func InternalTestFullImpuestos40(t *testing.T, impuestos *comprobante2.Impuestos40) {
	assert.NotNil(t, impuestos)

	assert.NotNil(t, impuestos.TotalImpuestosRetenidos)
	assert.Equal(t, 0.20, *impuestos.TotalImpuestosRetenidos)

	InternalTestFullTraslados(t, impuestos.Traslados)

	assert.NotNil(t, impuestos.TotalImpuestosTrasladados)
	assert.Equal(t, 0.16, *impuestos.TotalImpuestosTrasladados)

	InternalTestFullRetenciones(t, impuestos.Retenciones)
}

func InternalTestFullRetenciones(t *testing.T, retenciones *[]comprobante2.RetencionImpuestos40) {
	assert.NotNil(t, retenciones)

	assert.Equal(t, len(*retenciones), 2)

	first := (*retenciones)[0]
	assert.Equal(t, "001", first.Impuesto)
	assert.Equal(t, 0.10, first.Importe)

	second := (*retenciones)[1]
	assert.Equal(t, "002", second.Impuesto)
	assert.Equal(t, 0.10, second.Importe)
}

func InternalTestFullTraslados(t *testing.T, traslados *[]comprobante2.TrasladoImpuestos40) {
	assert.NotNil(t, traslados)

	assert.Equal(t, len(*traslados), 1)

	first := (*traslados)[0]
	assert.Equal(t, 1.0, first.Base)
	assert.Equal(t, "002", first.Impuesto)
	assert.Equal(t, "Tasa", first.TipoFactor)

	assert.NotNil(t, first.Importe)
	assert.Equal(t, 0.16, *first.Importe)

	assert.NotNil(t, first.TasaOCuota)
	assert.Equal(t, 0.16, *first.TasaOCuota)
}
