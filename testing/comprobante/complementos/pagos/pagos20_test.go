package pagos

import (
	"encoding/xml"
	"testing"

	pagos20 "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/pagos"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetPagos20ForTest(filename string, t *testing.T) (pagos20.Pagos20, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed pagos20.Pagos20
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	testing2.GenerateJSONFromStructure("pagos20.json", parsed)
	return parsed, errUnmashal
}

func TestFullPagos20(t *testing.T) {
	pagos, _ := GetPagos20ForTest("./pagos20.xml", t)
	assert.NotNil(t, pagos)
	InternalTestBaseAttributes(t, pagos)
}

func InternalTestBaseAttributes(t *testing.T, pagos pagos20.Pagos20) {
	assert.Equal(t, "2.0", pagos.Version)
	InternalTestTotalesPagos20(t, pagos.Totales)
	InternalTestPagoPagos20(t, pagos.Pago)
}

func InternalTestTotalesPagos20(t *testing.T, totales pagos20.TotalesPagos20) {
	assert.NotNil(t, totales)
	assert.Equal(t, 500.00, *totales.TotalRetencionesIva)
	assert.Equal(t, 300.00, *totales.TotalRetencionesIsr)
	assert.Equal(t, 0.00, *totales.TotalRetencionesIeps)
	assert.Equal(t, 10000.00, *totales.TotalTrasladosBaseIva16)
	assert.Equal(t, 1600.00, *totales.TotalTrasladosImpuestoIva16)
	assert.Equal(t, 0.00, *totales.TotalTrasladosBaseIva8)
	assert.Equal(t, 0.00, *totales.TotalTrasladosImpuestoIva8)
	assert.Equal(t, 0.00, *totales.TotalTrasladosBaseIva0)
	assert.Equal(t, 0.00, *totales.TotalTrasladosImpuestoIva0)
	assert.Equal(t, 0.00, *totales.TotalTrasladosBaseIvaExento)
	assert.Equal(t, 5000.00, totales.MontoTotalPagos)
}

func InternalTestPagoPagos20(t *testing.T, pagos []pagos20.PagoPagos20) {
	assert.NotNil(t, pagos)
	pago := pagos[0]
	assert.NotNil(t, pago)
	assert.Equal(t, "2024-09-15T12:00:00", pago.FechaPagoString)
	assert.Equal(t, "03", pago.FormaPago)
	assert.Equal(t, "USD", pago.Moneda)
	assert.Equal(t, 17.500000, *pago.TipoCambio)
	assert.Equal(t, 5000.00, pago.Monto)
	assert.Equal(t, "1234567890", *pago.NoOperacion)
	assert.Equal(t, "XEXX010101000", *pago.RfcEmisorCuentaOrdenante)
	assert.Equal(t, "Banco Internacional", *pago.NombreBancoOrdenanteExtranjero)
	assert.Equal(t, "1234567890123456", *pago.CuentaOrdenante)
	assert.Equal(t, "ABC123456T12", *pago.RfcEmisorCuentaBeneficiario)
	assert.Equal(t, "6543210987654321", *pago.CuentaBeneficiario)
	assert.Equal(t, "01", *pago.TipoCadenaPago)
	assert.Equal(t, "MIIEBjCCAm4CCQD6P5Iks0ACxjANBgkqhkiG9w0BAQsFADCBhTELMAkGA1UEBhMCVVMxEzARBgNV", *pago.CertificadoPago)
	assert.Equal(t, "00001|ABC123456T12|2024-09-15T12:00:00|5000.00|17.500000", *pago.CadenaPago)
	assert.Equal(t, "MIIEBjCCAm4CCQD6P5Iks0ACxjANBgkqhkiG9w0BAQsFADCBhTELMAkGA1UEBhMCVVMxEzARBgNV", *pago.SelloPago)
	InternalTestDoctoRelacionado(t, pago.DocumentosRelacionados)
	InternalTestRetencionP(t, pago.Impuestos.Retenciones)
	InternalTestTrasladoP(t, pago.Impuestos.Traslados)

}

func InternalTestDoctoRelacionado(t *testing.T, doctosRelacionado []pagos20.DoctoRelacionadoPagos20) {
	assert.NotNil(t, doctosRelacionado)
	assert.Equal(t, len(doctosRelacionado), 1)
	doctoRelacionado := doctosRelacionado[0]
	assert.Equal(t, "123e4567-e89b-12d3-a456-426614174000", doctoRelacionado.IdDocumento)
	assert.Equal(t, "A", *doctoRelacionado.Serie)
	assert.Equal(t, "1001", *doctoRelacionado.Folio)
	assert.Equal(t, "MXN", doctoRelacionado.Moneda)
	assert.Equal(t, 1.0, doctoRelacionado.NoParcialidad)
	assert.Equal(t, 1500.00, doctoRelacionado.ImporteSaldoAnterior)
	assert.Equal(t, 1500.00, doctoRelacionado.ImportePagado)
	assert.Equal(t, 0.00, doctoRelacionado.ImporteSaldoInsoluto)
	assert.Equal(t, "01", doctoRelacionado.ObjetoImpuesto)

	InternalTestRetencionDR(t, doctoRelacionado.Impuestos.Retenciones)
	InternalTestTrasladoDR(t, doctoRelacionado.Impuestos.Traslados)

}

func InternalTestRetencionDR(t *testing.T, retencionesDR *[]pagos20.RetencionDRPagos20) {
	assert.NotNil(t, retencionesDR)
	assert.Equal(t, len(*retencionesDR), 2)
	retencionDR := (*retencionesDR)[0]
	assert.Equal(t, 1000.000001, retencionDR.Base)
	assert.Equal(t, "001", retencionDR.Impuesto)
	assert.Equal(t, "Tasa", retencionDR.TipoFactor)
	assert.Equal(t, 0.160000, retencionDR.TasaOCuota)
	assert.Equal(t, 160.00, retencionDR.Importe)
}

func InternalTestTrasladoDR(t *testing.T, trasladosDR *[]pagos20.TrasladoDRPagos20) {
	assert.NotNil(t, trasladosDR)
	assert.Equal(t, len(*trasladosDR), 2)
	trasladoDR := (*trasladosDR)[0]
	assert.Equal(t, 2000.000001, trasladoDR.Base)
	assert.Equal(t, "002", trasladoDR.Impuesto)
	assert.Equal(t, "Tasa", trasladoDR.TipoFactor)
	assert.Equal(t, 0.160000, *trasladoDR.TasaOCuota)
	assert.Equal(t, 320.00, *trasladoDR.Importe)
}

func InternalTestRetencionP(t *testing.T, retencionesP *[]pagos20.RetencionPagos20) {
	assert.NotNil(t, retencionesP)
	assert.Equal(t, len(*retencionesP), 2)
	retencionP := (*retencionesP)[0]
	assert.Equal(t, "ISR", retencionP.Impuesto)
	assert.Equal(t, float32(1500.00), retencionP.Importe)
}

func InternalTestTrasladoP(t *testing.T, trasladosP *[]pagos20.TrasladoPagos20) {
	assert.NotNil(t, trasladosP)
	assert.Equal(t, len(*trasladosP), 2)
	trasladoP := (*trasladosP)[0]
	assert.Equal(t, 5000.00, trasladoP.Base)
	assert.Equal(t, "IVA", trasladoP.Impuesto)
	assert.Equal(t, "Tasa", trasladoP.TipoFactor)
	assert.Equal(t, 0.16, *trasladoP.TasaOCuota)
	assert.Equal(t, 800.00, *trasladoP.Importe)
}
