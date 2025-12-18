package comprobante

import (
	"encoding/xml"
	"testing"

	comprobante2 "github.com/SaulEnriqueMR/kore-models/models/comprobante"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetComprobante30ForTest(filename string, t *testing.T) (comprobante2.Comprobante30, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed comprobante2.Comprobante30
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	testing2.GenerateJSONFromStructure("comprobante30.json", parsed)
	return parsed, errUnmarshal
}

func TestComprobante30(t *testing.T) {
	comprobante30, _ := GetComprobante30ForTest("./comprobante30.xml", t)
	InternalTestFullAtributesComprobante30(t, comprobante30)
	InternalTestFullAtributesEmisorComprobante30(t, comprobante30.Emisor)
	InternalTestFullAtributesReceptor30(t, comprobante30.Receptor)
	InternalTestFullAtributesConceptos30(t, comprobante30.Conceptos)
	InternalTestFullAtributesImpuestos30(t, comprobante30.Impuestos)
}

func InternalTestFullAtributesComprobante30(t *testing.T, c comprobante2.Comprobante30) {
	assert.Equal(t, "3.0", c.Version)
	assert.Equal(t, "A", *c.Serie)
	assert.Equal(t, "5672", *c.Folio)
	assert.Equal(t, "2011-12-09T12:55:04", c.Fecha)
	assert.Equal(t, "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQ...", c.Sello)
	assert.Equal(t, "PAGO EN UNA SOLA EXHIBICION", c.FormaPago)
	assert.Equal(t, "00000000000000000000", c.NoCertificado)
	assert.Equal(t, "MIICqjCCAbKgAwIBAgIJAN6vK1...", c.Certificado)
	assert.Equal(t, "Contado", *c.CondicionesPago)
	assert.Equal(t, 1000.00, c.Subtotal)
	assert.Equal(t, 50.00, *c.Descuento)
	assert.Equal(t, "DESCUENTAZO", *c.MotivoDescuento)
	assert.Equal(t, "13.6569", *c.TipoDeCambio)
	assert.Equal(t, 13.6569, c.TipoCambio)
	assert.Equal(t, "USD", *c.Moneda)
	assert.Equal(t, 1477.84, c.Total)
	assert.Equal(t, "TRANSFERENCIA", *c.MetodoPago)
	assert.Equal(t, "ingreso", c.TipoDeComprobante)
	assert.Equal(t, "I", c.TipoComprobante)
}

func InternalTestFullAtributesEmisorComprobante30(t *testing.T, e comprobante2.Emisor30) {
	assert.Equal(t, "EKU9003173C9", e.Rfc)
	assert.Equal(t, "ESCUELA KEMPER URGATE", e.Nombre)
	InternalTestFullAtributesDomicilioFiscal30(t, *e.DomicilioFiscal)
	InternalTestFullAtributesExpedidoEn30(t, *e.ExpedidoEn)
}

func InternalTestFullAtributesDomicilioFiscal30(t *testing.T, d comprobante2.Domicilio30) {
	assert.NotNil(t, d)
	assert.Equal(t, "CALZADA DE LOS REYES", *d.Calle)
	assert.Equal(t, "1", *d.NoExterior)
	assert.Equal(t, "1", *d.NoInterior)
	assert.Equal(t, "PARQUE INDUSTRIAL FINSA", *d.Colonia)
	assert.Equal(t, "CUAUTITLAN", *d.Localidad)
	assert.Equal(t, "CUAUTITLAN", *d.Referencia)
	assert.Equal(t, "CUAUTITLAN", *d.Municipio)
	assert.Equal(t, "Edo. MEX", *d.Estado)
	assert.Equal(t, "México", *d.Pais)
	assert.Equal(t, "23452", *d.CodigoPostal)
}

func InternalTestFullAtributesExpedidoEn30(t *testing.T, d comprobante2.Domicilio30) {
	assert.NotNil(t, d)
	assert.Equal(t, "CALZADA DE LOS REYES", *d.Calle)
	assert.Equal(t, "1", *d.NoExterior)
	assert.Equal(t, "1", *d.NoInterior)
	assert.Equal(t, "PARQUE INDUSTRIAL FINSA", *d.Colonia)
	assert.Equal(t, "CUAUTITLAN", *d.Localidad)
	assert.Equal(t, "CUAUTITLAN", *d.Referencia)
	assert.Equal(t, "CUAUTITLAN", *d.Municipio)
	assert.Equal(t, "Edo. MEX", *d.Estado)
	assert.Equal(t, "México", *d.Pais)
	assert.Equal(t, "23452", *d.CodigoPostal)
}

func InternalTestFullAtributesReceptor30(t *testing.T, r comprobante2.Receptor30) {
	assert.Equal(t, "XIA190128J61", r.Rfc)
	assert.Equal(t, "XENON INDUSTRIAL ARTICLES", *r.Nombre)
	InternalTestFullAtributesDomicilio30(t, *r.Domicilio)
}

func InternalTestFullAtributesDomicilio30(t *testing.T, d comprobante2.Domicilio30) {
	assert.NotNil(t, d)
	assert.Equal(t, "Corredor Industrial", *d.Calle)
	assert.Equal(t, "1", *d.NoExterior)
	assert.Equal(t, "1", *d.NoInterior)
	assert.Equal(t, "Cd. Mex", *d.Colonia)
	assert.Equal(t, "Localidad", *d.Localidad)
	assert.Equal(t, "Referencia", *d.Referencia)
	assert.Equal(t, "Municipio", *d.Municipio)
	assert.Equal(t, "VER", *d.Estado)
	assert.Equal(t, "MEXICO", *d.Pais)
	assert.Equal(t, "11232", *d.CodigoPostal)
}

func InternalTestFullAtributesConceptos30(t *testing.T, c []comprobante2.Concepto30) {
	assert.Equal(t, 1, len(c))
	assert.Equal(t, float64(1), c[0].Cantidad)
	assert.Equal(t, "PIEZA", *c[0].Unidad)
	assert.Equal(t, "3e432432", *c[0].NoIdentificacion)
	assert.Equal(t, "Consumibles de oficina", c[0].Descripcion)
	assert.Equal(t, 1274.00, c[0].ValorUnitario)
	assert.Equal(t, 1274.00, c[0].Importe)
	InternalTestFullAtributesInformacionAduaneraConcepto30(t, *c[0].InformacionAduanera)
	cuentaPredial30 := c[0].CuentaPredial
	assert.NotNil(t, cuentaPredial30)
	assert.Equal(t, "2132145", cuentaPredial30.Numero)
	InternalTestFullAtributesPartes30(t, *c[0].Parte)
}

func InternalTestFullAtributesInformacionAduaneraConcepto30(t *testing.T, ia []comprobante2.InformacionAduanera30) {
	assert.Equal(t, 1, len(ia))
	assert.Equal(t, "212312321", ia[0].Numero)
	assert.Equal(t, "2025-01-01", ia[0].FechaString)
	assert.Equal(t, "123123", *ia[0].Aduana)
}

func InternalTestFullAtributesPartes30(t *testing.T, p []comprobante2.Parte30) {
	assert.NotNil(t, p)
	assert.Equal(t, 1, len(p))
	assert.Equal(t, float64(2), p[0].Cantidad)
	assert.Equal(t, "PIEZA", *p[0].Unidad)
	assert.Equal(t, "4324234", *p[0].NoIdentificacion)
	assert.Equal(t, "Parte de consumible", p[0].Descripcion)
	assert.Equal(t, 637.00, *p[0].ValorUnitario)
	assert.Equal(t, 1274.00, *p[0].Importe)
}

func InternalTestFullAtributesImpuestos30(t *testing.T, i comprobante2.Impuestos30) {
	assert.Equal(t, 100.00, *i.TotalImpuestosRetenidos)
	assert.Equal(t, 203.84, *i.TotalImpuestosTrasladados)
	InternalTestFullAtributesTraslados30(t, *i.Traslados)
	InternalTestFullAtributesRetenciones30(t, *i.Retenciones)
}

func InternalTestFullAtributesTraslados30(t *testing.T, tr []comprobante2.Traslado30) {
	assert.NotNil(t, tr)
	assert.Equal(t, 1, len(tr))
	assert.Equal(t, "IVA", tr[0].TipoImpuesto)
	assert.Equal(t, "002", tr[0].Impuesto)
	assert.Equal(t, 16.00, tr[0].Tasa)
	assert.Equal(t, 203.84, tr[0].Importe)
}

func InternalTestFullAtributesRetenciones30(t *testing.T, r []comprobante2.Retencion30) {
	assert.NotNil(t, r)
	assert.Equal(t, 1, len(r))
	assert.Equal(t, "IEPS", r[0].TipoImpuesto)
	assert.Equal(t, "003", r[0].Impuesto)
	assert.Equal(t, 10500.00, r[0].Importe)
}
