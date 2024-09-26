package pagos

import (
	"encoding/xml"
	"testing"

	pagos10 "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/pagos"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetPagos10ForTest(filename string, t *testing.T) (pagos10.Pagos10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed pagos10.Pagos10
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	testing2.GenerateJSONFromStructure("pagos10.json", parsed)
	return parsed, errUnmashal
}

func TestFullPagos10(t *testing.T) {
	pagos, _ := GetPagos10ForTest("./pagos10.xml", t)
	assert.NotNil(t, pagos)
	InternalTestBasePagos10Attributes(t, pagos)
}

func InternalTestBasePagos10Attributes(t *testing.T, pagos pagos10.Pagos10) {
	assert.Equal(t, "1.0", pagos.Version)
	InternalTestPagos10(t, pagos.Pago)

}
func InternalTestPagos10(t *testing.T, pagos []pagos10.PagoPagos10) {
	assert.NotNil(t, pagos)
	assert.Equal(t, len(pagos), 1)
	pago := pagos[0]
	assert.Equal(t, "2024-09-24T12:00:00", pago.FechaPago)
	assert.Equal(t, "01", pago.FormaDePagoP)
	assert.Equal(t, "MXN", pago.MonedaP)
	assert.Equal(t, 1.0, *pago.TipoCambioP)
	assert.Equal(t, 1000.00, pago.Monto)
	assert.Equal(t, "123456789", *pago.NumOperacion)
	assert.Equal(t, "XEXX010101000", *pago.RfcEmisorCtaOrd)
	assert.Equal(t, "Banco Ejemplo", *pago.NomBancoOrdExt)
	assert.Equal(t, "0123456789", *pago.CtaOrdenante)
	assert.Equal(t, "A123456789", *pago.RfcEmisorCtaBen)
	assert.Equal(t, "9876543210", *pago.CtaBeneficiario)
	assert.Equal(t, "01", *pago.TipoCadPago)
	assert.Equal(t, "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvD1X...", *pago.CertPago)
	assert.Equal(t, "Cadena original de comprobante de pago", *pago.CadPago)
	assert.Equal(t, "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvD1X...", *pago.SelloPago)
	InternalTestDoctoRelacionadoPagos10(t, pago.DoctoRelacionado)
	InternalTestImpuestosPagos10(t, pago.Impuestos)
}
func InternalTestDoctoRelacionadoPagos10(t *testing.T, documentos *[]pagos10.DoctoRelacionadoPagos10) {
	assert.NotNil(t, documentos)
	assert.Equal(t, len(*documentos), 3)
	documento := (*documentos)[1]
	assert.Equal(t, "456e1234-e89b-12d3-a456-426614174001", documento.IdDocumento)
	assert.Equal(t, "B", *documento.Serie)
	assert.Equal(t, "0002", *documento.Folio)
	assert.Equal(t, "USD", documento.MonedaDR)
	assert.Equal(t, 1.20, *documento.TipoCambioDR)
	assert.Equal(t, "02", documento.MetodoDePagoDR)
	assert.Equal(t, 1.0, *documento.NumParcialidad)
	assert.Equal(t, 2500.00, *documento.ImpSaldoAnt)
	assert.Equal(t, 1500.00, *documento.ImpPagado)
	assert.Equal(t, 1000.00, *documento.ImpSaldoInsoluto)
}

func InternalTestImpuestosPagos10(t *testing.T, impuestos *[]pagos10.ImpuestosPagos10) {
	assert.NotNil(t, impuestos)
	assert.Equal(t, len(*impuestos), 2)
	impuesto := (*impuestos)[0]
	assert.Equal(t, 150.00, *impuesto.TotalImpuestosRetenidos)
	assert.Equal(t, 50.00, *impuesto.TotalImpuestosTrasladados)
	InternalTestRetencionesPagos10(t, impuesto.Retenciones)
	InternalTestTrasladosPagos10(t, impuesto.Traslados)
}

func InternalTestRetencionesPagos10(t *testing.T, retenciones *[]pagos10.RetencionPagos10) {
	assert.NotNil(t, retenciones)
	assert.Equal(t, len(*retenciones), 2)
	retencion1 := (*retenciones)[0]
	assert.Equal(t, "001", retencion1.Impuesto)
	assert.Equal(t, 100.00, retencion1.Importe)
	retencion2 := (*retenciones)[1]
	assert.Equal(t, "002", retencion2.Impuesto)
	assert.Equal(t, 50.00, retencion2.Importe)
}

func InternalTestTrasladosPagos10(t *testing.T, traslados *[]pagos10.TrasladoPagos10) {
	assert.NotNil(t, traslados)
	assert.Equal(t, len(*traslados), 1)
	traslado := (*traslados)[0]
	assert.Equal(t, "003", traslado.Impuesto)
	assert.Equal(t, "Tasa", traslado.TipoFactor)
	assert.Equal(t, 0.16, traslado.TasaOCuota)
	assert.Equal(t, 50.00, traslado.Importe)
}
