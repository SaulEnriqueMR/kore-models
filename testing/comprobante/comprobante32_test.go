package comprobante

import (
	"encoding/xml"

	comprobante2 "github.com/SaulEnriqueMR/kore-models/models/comprobante"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"

	"testing"
)

func GestComprobante32ForTest(filename string, t *testing.T) (comprobante2.Comprobante32, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed comprobante2.Comprobante32
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	testing2.GenerateJSONFromStructure("comprobante32.json", parsed)
	return parsed, errUnmarshal
}

func TestComprobante32(t *testing.T) {
	comprobante32, _ := GestComprobante32ForTest("./comprobante32.xml", t)
	InternalTestFullAtributesComprobante32(t, comprobante32)
	InternalTestFullAtributesEmisorComprobante32(t, comprobante32.Emisor)
	InternalTestFullAtributesReceptor32(t, comprobante32.Receptor)
	InternalTestFullAtributesConceptos32(t, comprobante32.Conceptos)
	InternalTestFullAtributesImpuestos32(t, comprobante32.Impuestos)
}

func InternalTestFullAtributesComprobante32(t *testing.T, comprobante32 comprobante2.Comprobante32) {
	assert.Equal(t, "3.2", comprobante32.Version)
	assert.Equal(t, "A", *comprobante32.Serie)
	assert.Equal(t, "123456", *comprobante32.Folio)
	assert.Equal(t, "2023-09-23T12:00:00", comprobante32.Fecha)
	assert.Equal(t, "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQ...", comprobante32.Sello)
	assert.Equal(t, "01", comprobante32.FormaPago)
	assert.Equal(t, "00000000000000000000", comprobante32.NoCertificado)
	assert.Equal(t, "MIICqjCCAbKgAwIBAgIJAN6vK1...", comprobante32.Certificado)
	assert.Equal(t, "Contado", *comprobante32.CondicionesPago)
	assert.Equal(t, 1000.00, comprobante32.Subtotal)
	assert.Equal(t, 50.00, *comprobante32.Descuento)
	assert.Equal(t, "Promoción especial", *comprobante32.MotivoDescuento)
	assert.Equal(t, "20.5", *comprobante32.TipoCambio)
	assert.Equal(t, "MXN", *comprobante32.Moneda)
	assert.Equal(t, 1000.00, comprobante32.Total)
	assert.Equal(t, "ingreso", comprobante32.TipoComprobante)
	assert.Equal(t, "Transferencia bancaria", comprobante32.MetodoPago)
	assert.Equal(t, "Ciudad de México", comprobante32.LugarExpedicion)
	assert.Equal(t, "1234", *comprobante32.NumeroCuentaPago)
	assert.Equal(t, "ABC123456", *comprobante32.FolioFiscalOriginal)
	assert.Equal(t, "A1", *comprobante32.SerieFolioFiscalOriginal)
	assert.Equal(t, "2023-09-23T12:00:00", *comprobante32.FechaFolioFiscalOriginal)
	assert.Equal(t, "1000.00", *comprobante32.MontoFolioFiscalOriginal)
}

func InternalTestFullAtributesEmisorComprobante32(t *testing.T, emisor32 comprobante2.Emisor32) {
	assert.Equal(t, "ABC123456789", emisor32.Rfc)
	assert.Equal(t, "Juan Pérez S.A. de C.V.", *emisor32.Nombre)
	InternalTestFullAtributesRegimenFiscal32(t, emisor32.RegimenesFiscales)
	InternalTestFullAtributesDomicilioFiscal(t, *emisor32.DomicilioFiscal)
	InternalTestFullAtributesExpedidoEn32(t, *emisor32.ExpedidoEn)
}

func InternalTestFullAtributesRegimenFiscal32(t *testing.T, fiscal32 []comprobante2.RegimenFiscal32) {
	assert.Equal(t, 3, len(fiscal32))
	assert.Equal(t, "Persona Física con Actividades Empresariales", fiscal32[0].RegimenFiscal)
	assert.Equal(t, "Régimen de Incorporación Fiscal", fiscal32[1].RegimenFiscal)
	assert.Equal(t, "Persona Moral", fiscal32[2].RegimenFiscal)
}

func InternalTestFullAtributesDomicilioFiscal(t *testing.T, domicilio32 comprobante2.Domicilio32) {
	assert.NotNil(t, domicilio32)
	assert.Equal(t, "Avenida Insurgentes", domicilio32.Calle)
	assert.Equal(t, "456", *domicilio32.NoExterior)
	assert.Equal(t, "A", *domicilio32.NoInterior)
	assert.Equal(t, "Roma", *domicilio32.Colonia)
	assert.Equal(t, "Ciudad de México", *domicilio32.Localidad)
	assert.Equal(t, "Cerca del Parque", *domicilio32.Referencia)
	assert.Equal(t, "Cuauhtémoc", domicilio32.Municipio)
	assert.Equal(t, "CDMX", domicilio32.Estado)
	assert.Equal(t, "México", domicilio32.Pais)
	assert.Equal(t, "06700", domicilio32.CodigoPostal)
}

func InternalTestFullAtributesExpedidoEn32(t *testing.T, domicilio32 comprobante2.Domicilio32) {
	assert.NotNil(t, domicilio32)
	assert.Equal(t, "Avenida Reforma", domicilio32.Calle)
	assert.Equal(t, "123", *domicilio32.NoExterior)
	assert.Equal(t, "1", *domicilio32.NoInterior)
	assert.Equal(t, "Centro", *domicilio32.Colonia)
	assert.Equal(t, "Ciudad de México", *domicilio32.Localidad)
	assert.Equal(t, "Referencia 1", *domicilio32.Referencia)
	assert.Equal(t, "Cuauhtémoc", domicilio32.Municipio)
	assert.Equal(t, "CDMX", domicilio32.Estado)
	assert.Equal(t, "México", domicilio32.Pais)
	assert.Equal(t, "06000", domicilio32.CodigoPostal)
}

func InternalTestFullAtributesReceptor32(t *testing.T, receptor32 comprobante2.Receptor32) {
	assert.Equal(t, "ABC123456789", receptor32.Rfc)
	assert.Equal(t, "Juan Pérez López", *receptor32.Nombre)
	InternalTestFullAtributesDomicilioReceptor32(t, *receptor32.Domicilio)
}

func InternalTestFullAtributesDomicilioReceptor32(t *testing.T, domicilio32 comprobante2.Domicilio32) {
	assert.NotNil(t, domicilio32)
	assert.Equal(t, "Avenida Reforma", domicilio32.Calle)
	assert.Equal(t, "123", *domicilio32.NoExterior)
	assert.Equal(t, "1", *domicilio32.NoInterior)
	assert.Equal(t, "Centro", *domicilio32.Colonia)
	assert.Equal(t, "Ciudad de México", *domicilio32.Localidad)
	assert.Equal(t, "Referencia 1", *domicilio32.Referencia)
	assert.Equal(t, "Cuauhtémoc", domicilio32.Municipio)
	assert.Equal(t, "CDMX", domicilio32.Estado)
	assert.Equal(t, "México", domicilio32.Pais)
	assert.Equal(t, "06000", domicilio32.CodigoPostal)
}

func InternalTestFullAtributesConceptos32(t *testing.T, concepto32 []comprobante2.Concepto32) {
	assert.Equal(t, 1, len(concepto32))
	assert.Equal(t, float64(5), concepto32[0].Cantidad)
	assert.Equal(t, "pieza", concepto32[0].Unidad)
	assert.Equal(t, "ABC123", *concepto32[0].NoIdentificacion)
	assert.Equal(t, "Producto Ejemplo", concepto32[0].Descripcion)
	assert.Equal(t, 100.00, concepto32[0].ValorUnitario)
	assert.Equal(t, 500.00, concepto32[0].Importe)
	InternalTestFullAtributesInformacionAduaneraConcepto32(t, *concepto32[0].InformacionAduanera)
	cuentaPredial32 := concepto32[0].CuentaPredial
	assert.NotNil(t, cuentaPredial32)
	assert.Equal(t, "123456789", cuentaPredial32.Numero)
	InternalTestFullAtributesPartes32(t, *concepto32[0].Parte)
}

func InternalTestFullAtributesInformacionAduaneraConcepto32(t *testing.T, aduanera32 []comprobante2.InformacionAduanera32) {
	assert.Equal(t, 1, len(aduanera32))
	assert.Equal(t, "A123456", aduanera32[0].Numero)
	assert.Equal(t, "2023-09-15", aduanera32[0].Fecha)
	assert.Equal(t, "Aduana de Ciudad Juárez", *aduanera32[0].Aduana)
}

func InternalTestFullAtributesPartes32(t *testing.T, parte32 []comprobante2.Parte32) {
	assert.NotNil(t, parte32)
	assert.Equal(t, 1, len(parte32))
	assert.Equal(t, float64(2), parte32[0].Cantidad)
	assert.Equal(t, "unidad", *parte32[0].Unidad)
	assert.Equal(t, "ABC123", *parte32[0].NoIdentificacion)
	assert.Equal(t, "Componente Ejemplo", parte32[0].Descripcion)
	assert.Equal(t, 50.00, *parte32[0].ValorUnitario)
	assert.Equal(t, 100.00, *parte32[0].Importe)
}

func InternalTestFullAtributesImpuestos32(t *testing.T, impuestos32 comprobante2.Impuestos32) {
	assert.Equal(t, 11000.00, *impuestos32.TotalImpuestosRetenidos)
	assert.Equal(t, 11000.00, *impuestos32.TotalImpuestosTrasladados)
	InternalTestFullAtributesTraslados32(t, *impuestos32.Traslados)
	InternalTestFullAtributesRetenciones32(t, *impuestos32.Retenciones)
}

func InternalTestFullAtributesTraslados32(t *testing.T, traslado32 []comprobante2.Traslado32) {
	assert.NotNil(t, traslado32)
	assert.Equal(t, 1, len(traslado32))
	assert.Equal(t, "IEPS", traslado32[0].Impuesto)
	assert.Equal(t, 16.00, traslado32[0].Tasa)
	assert.Equal(t, 100.00, traslado32[0].Importe)
}

func InternalTestFullAtributesRetenciones32(t *testing.T, retenciones32 []comprobante2.Retencion32) {
	assert.NotNil(t, retenciones32)
	assert.Equal(t, 1, len(retenciones32))
	assert.Equal(t, "IEPS", retenciones32[0].Impuesto)
	assert.Equal(t, 10500.00, retenciones32[0].Importe)
}
