package comprobante

import (
	"encoding/xml"
	"testing"

	comprobante3 "github.com/SaulEnriqueMR/kore-models/models/comprobante"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetComprobante33ForTest(filename string, t *testing.T) (comprobante3.Comprobante33, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed comprobante3.Comprobante33
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	testing2.GenerateJSONFromStructure("comprobante33.json", parsed)
	return parsed, errUnmarshal
}

func TestComprobante33(t *testing.T) {
	comprobante33, _ := GetComprobante33ForTest("./comprobante33.xml", t)
	InternalTestFullAtributesComprobante33(t, comprobante33)
	InternalTestFullAtributesCfdiRelacionados(t, *comprobante33.CfdisRelacionados)
	InternalTestFullAtributesEmisorComprobante33(t, comprobante33.Emisor)
	InternalTestFullAtributesReceptorComprobante33(t, comprobante33.Receptor)
	InternalTestFullAtributesImpuestosComprobante33(t, *comprobante33.Impuestos)
	InternalTestFullAtributesConceptos33(t, comprobante33.Conceptos)
	InternalTestFullAtributesComplemento33(t, *comprobante33.Complemento)
}

func InternalTestFullAtributesComprobante33(t *testing.T, comprobante33 comprobante3.Comprobante33) {
	assert.Equal(t, "3.3", comprobante33.Version)
	assert.Equal(t, "A", *comprobante33.Serie)
	assert.Equal(t, "123456", *comprobante33.Folio)
	assert.Equal(t, "2023-09-24T14:30:00", comprobante33.Fecha)
	assert.Equal(t, "Base64SelloDigital", comprobante33.Sello)
	assert.Equal(t, "01", *comprobante33.FormaPago)
	assert.Equal(t, "12345678901234567890", comprobante33.NoCertificado)
	assert.Equal(t, "Base64Certificado", comprobante33.Certificado)
	assert.Equal(t, "Contado", *comprobante33.CondicionesPago)
	assert.Equal(t, 1000.00, comprobante33.Subtotal)
	assert.Equal(t, 50.00, *comprobante33.Descuento)
	assert.Equal(t, "MXN", comprobante33.Moneda)
	assert.Equal(t, 1.000000, *comprobante33.TipoCambio)
	assert.Equal(t, 950.00, comprobante33.Total)
	assert.Equal(t, "I", comprobante33.TipoComprobante)
	assert.Equal(t, "PUE", *comprobante33.MetodoPago)
	assert.Equal(t, "01234", comprobante33.LugarExpedicion)
	assert.Equal(t, "ABCDE", *comprobante33.Confirmacion)
}

func InternalTestFullAtributesCfdiRelacionados(t *testing.T, relacionados33 comprobante3.CfdiRelacionados33) {
	assert.NotNil(t, relacionados33)
	assert.Equal(t, "01", relacionados33.TipoRelacion)
	//Prueba del atributo UUID del nodo CfdiRelacionado
	assert.Equal(t, "123e4567-e89b-12d3-a456-426614174000", relacionados33.UuidRelacionados[0].UUID)
	assert.Equal(t, "123E4567-E89B-12D3-A456-426614174000", relacionados33.UuidRelacionados[0].Uuid)
}

func InternalTestFullAtributesEmisorComprobante33(t *testing.T, emisor33 comprobante3.Emisor33) {
	assert.Equal(t, "XAXX010101000", emisor33.Rfc)
	assert.Equal(t, "Empresa Ejemplo S.A. de C.V.", *emisor33.Nombre)
	assert.Equal(t, "601", emisor33.RegimenFiscal)
}

func InternalTestFullAtributesReceptorComprobante33(t *testing.T, receptor33 comprobante3.Receptor33) {
	assert.Equal(t, "XAXX010101000", receptor33.Rfc)
	assert.Equal(t, "Juan Pérez", *receptor33.Nombre)
	assert.Equal(t, "MEX", *receptor33.ResidenciaFiscal)
	assert.Equal(t, "1234567890", *receptor33.NumRegIdTrib)
	assert.Equal(t, "G03", receptor33.UsoCFDI)
}

func InternalTestFullAtributesConceptos33(t *testing.T, concepto33 []comprobante3.Concepto33) {
	assert.Equal(t, 1, len(concepto33))
	assert.Equal(t, "01010101", concepto33[0].ClaveProdServ)
	assert.Equal(t, "SKU12345", *concepto33[0].NoIdentificacion)
	assert.Equal(t, 10.000000, concepto33[0].Cantidad)
	assert.Equal(t, "H87", concepto33[0].ClaveUnidad)
	assert.Equal(t, "Piezas", *concepto33[0].Unidad)
	assert.Equal(t, "Descripción del producto o servicio", concepto33[0].Descripcion)
	assert.Equal(t, 100.00, concepto33[0].ValorUnitario)
	assert.Equal(t, 1000.00, concepto33[0].Importe)
	assert.Equal(t, 50.00, *concepto33[0].Descuento)
	InternalTestFullAtributesImpuestosConcepto33(t, *concepto33[0].Impuestos)
	assert.Equal(t, "230112340000123", (*concepto33[0].InformacionAduanera)[0].NumeroPedimento)
	assert.Equal(t, "12345678901234567890", concepto33[0].CuentaPredial.Numero)
	InternalTestFullAtributesParteConcepto33(t, *concepto33[0].Parte)
}

func InternalTestFullAtributesImpuestosConcepto33(t *testing.T, concepto33 comprobante3.ImpuestosConcepto33) {
	assert.NotNil(t, concepto33)
	InternalTestFullAtributesTrasladosConcepto33(t, *concepto33.Traslados)
	InternalTestFullAtributesRetencionesConcepto33(t, *concepto33.Retenciones)
}

func InternalTestFullAtributesTrasladosConcepto33(t *testing.T, concepto33 []comprobante3.TrasladoConcepto33) {
	assert.NotNil(t, concepto33)
	assert.Equal(t, 1, len(concepto33))
	assert.Equal(t, 1000.00, concepto33[0].Base)
	assert.Equal(t, "002", concepto33[0].Impuesto)
	assert.Equal(t, "Tasa", concepto33[0].TipoFactor)
	assert.Equal(t, 0.16, *concepto33[0].TasaOCuota)
	assert.Equal(t, 160.00, *concepto33[0].Importe)
}

func InternalTestFullAtributesRetencionesConcepto33(t *testing.T, concepto33 []comprobante3.RetencionConcepto33) {
	assert.NotNil(t, concepto33)
	assert.Equal(t, 1, len(concepto33))
	assert.Equal(t, 1000.00, concepto33[0].Base)
	assert.Equal(t, "003", concepto33[0].Impuesto)
	assert.Equal(t, "Tasa", concepto33[0].TipoFactor)
	assert.Equal(t, 0.10, concepto33[0].TasaOCuota)
	assert.Equal(t, 100.00, concepto33[0].Importe)
}

func InternalTestFullAtributesParteConcepto33(t *testing.T, parte33 []comprobante3.Parte33) {
	assert.NotNil(t, parte33)
	assert.Equal(t, 1, len(parte33))
	assert.Equal(t, "01010101", parte33[0].ClaveProdServ)
	assert.Equal(t, "PN123456", *parte33[0].NoIdentificacion)
	assert.Equal(t, 5.000000, parte33[0].Cantidad)
	assert.Equal(t, "Piezas", *parte33[0].Unidad)
	assert.Equal(t, "Parte de repuesto para maquinaria", parte33[0].Descripcion)
	assert.Equal(t, 150.00, *parte33[0].ValorUnitario)
	assert.Equal(t, 750.00, *parte33[0].Importe)
}

func InternalTestFullAtributesImpuestosComprobante33(t *testing.T, impuestos33 comprobante3.Impuestos33) {
	assert.NotNil(t, impuestos33)
	assert.Equal(t, 100.00, *impuestos33.TotalImpuestosRetenidos)
	assert.Equal(t, 200.00, *impuestos33.TotalImpuestosTrasladados)
	InternalTestFullAtributesRetenciones33(t, *impuestos33.Retenciones)
	InternalTestFullAtributesTraslados(t, *impuestos33.Traslados)
}

func InternalTestFullAtributesRetenciones33(t *testing.T, impuestos33 []comprobante3.RetencionImpuestos33) {
	assert.NotNil(t, impuestos33)
	assert.Equal(t, 1, len(impuestos33))
	assert.Equal(t, "001", impuestos33[0].Impuesto)
	assert.Equal(t, 100.00, impuestos33[0].Importe)
}

func InternalTestFullAtributesTraslados(t *testing.T, impuestos33 []comprobante3.TrasladoImpuestos33) {
	assert.NotNil(t, impuestos33)
	assert.Equal(t, 1, len(impuestos33))
	assert.Equal(t, "002", impuestos33[0].Impuesto)
	assert.Equal(t, "Tasa", impuestos33[0].TipoFactor)
	assert.Equal(t, 0.160000, impuestos33[0].TasaOCuota)
	assert.Equal(t, 200.00, impuestos33[0].Importe)
}

func InternalTestFullAtributesComplemento33(t *testing.T, complemento33 comprobante3.Complemento) {
	assert.NotNil(t, complemento33)
	assert.NotNil(t, complemento33.ImpuestoLocales)
	assert.Equal(t, 5, len(*complemento33.ImpuestoLocales.ImpuestoLocales10))
	impuestosLocales := *complemento33.ImpuestoLocales
	impuestosLocales1 := *impuestosLocales.ImpuestoLocales10
	assert.Equal(t, "1.0", impuestosLocales1[0].Version)
	assert.Equal(t, 1500.00, impuestosLocales1[0].TotalRetenciones)
	assert.Equal(t, 3000.00, impuestosLocales1[0].TotalTraslados)

	retencionesLocales := *impuestosLocales1[0].Retenciones
	trasladosLocales := *impuestosLocales1[0].Traslados
	assert.Equal(t, "ISR", retencionesLocales[0].Impuesto)
	assert.Equal(t, 15.00, retencionesLocales[0].Tasa)
	assert.Equal(t, 1500.00, retencionesLocales[0].Importe)

	assert.Equal(t, "IVA", trasladosLocales[0].Impuesto)
	assert.Equal(t, 16.00, trasladosLocales[0].Tasa)
	assert.Equal(t, 3000.00, trasladosLocales[0].Importe)
}

func TestEmitidaRecibida33(t *testing.T) {
	comprobante33, _ := GetComprobante33ForTest("./comprobante33_emitido-recibido.xml", t)

	rfcEmitido := "ABC123456789"
	comprobante33.DefineTransaccion(rfcEmitido)
	assert.Equal(t, "EMITIDO", comprobante33.Transaccion)

	rfcRecibido := "DEF123456789"
	comprobante33.DefineTransaccion(rfcRecibido)
	assert.Equal(t, "RECIBIDO", comprobante33.Transaccion)
}

func TestBasePath33(t *testing.T) {
	comprobante33, _ := GetComprobante33ForTest("./comprobante33.xml", t)

	assert.Equal(t, comprobante33.GetFileName(), comprobante33.GetBasePath()+".xml")
}
