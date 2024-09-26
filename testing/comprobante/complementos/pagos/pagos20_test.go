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
	assert.Equal(t, 500.00, *totales.TotalRetencionesIVA)
	assert.Equal(t, 300.00, *totales.TotalRetencionesISR)
	assert.Equal(t, 0.00, *totales.TotalRetencionesIEPS)
	assert.Equal(t, 10000.00, *totales.TotalTrasladosBaseIVA16)
	assert.Equal(t, 1600.00, *totales.TotalTrasladosImpuestoIVA16)
	assert.Equal(t, 0.00, *totales.TotalTrasladosBaseIVA8)
	assert.Equal(t, 0.00, *totales.TotalTrasladosImpuestoIVA8)
	assert.Equal(t, 0.00, *totales.TotalTrasladosBaseIVA0)
	assert.Equal(t, 0.00, *totales.TotalTrasladosImpuestoIVA0)
	assert.Equal(t, 0.00, *totales.TotalTrasladosBaseIVAExento)
	assert.Equal(t, 5000.00, totales.MontoTotalPagos)
}

func InternalTestPagoPagos20(t *testing.T, pagos []pagos20.PagoPagos20) {
	assert.NotNil(t, pagos)
	pago := pagos[0]
	assert.NotNil(t, pago)
	assert.Equal(t, "2024-09-15T12:00:00", pago.FechaPago)
	assert.Equal(t, "03", pago.FormaPagoP)
	assert.Equal(t, "USD", pago.MonedaP)
	assert.Equal(t, 17.500000, *pago.TipoCambioP)
	assert.Equal(t, 5000.00, pago.Monto)
	assert.Equal(t, "1234567890", *pago.NumOperacion)
	assert.Equal(t, "XEXX010101000", *pago.RfcEmisorCtaOrd)
	assert.Equal(t, "Banco Internacional", *pago.NomBancoOrdExt)
	assert.Equal(t, "1234567890123456", *pago.CtaOrdenante)
	assert.Equal(t, "ABC123456T12", *pago.RfcEmisorCtaBen)
	assert.Equal(t, "6543210987654321", *pago.CtaBeneficiario)
	assert.Equal(t, "01", *pago.TipoCadPago)
	assert.Equal(t, "MIIEBjCCAm4CCQD6P5Iks0ACxjANBgkqhkiG9w0BAQsFADCBhTELMAkGA1UEBhMCVVMxEzARBgNV", *pago.CertPago)
	assert.Equal(t, "00001|ABC123456T12|2024-09-15T12:00:00|5000.00|17.500000", *pago.CadPago)
	assert.Equal(t, "MIIEBjCCAm4CCQD6P5Iks0ACxjANBgkqhkiG9w0BAQsFADCBhTELMAkGA1UEBhMCVVMxEzARBgNV", *pago.SelloPago)
	InternalTestDoctoRelacionado(t, pago.DoctoRelacionado)
	InternalTestRetencionP(t, pago.ImpuestosP.RetencionesP)
	InternalTestTrasladoP(t, pago.ImpuestosP.TrasladosP)

}

func InternalTestDoctoRelacionado(t *testing.T, doctosRelacionado []pagos20.DoctoRelacionadoPagos20) {
	assert.NotNil(t, doctosRelacionado)
	assert.Equal(t, len(doctosRelacionado), 1)
	doctoRelacionado := doctosRelacionado[0]
	assert.Equal(t, "123e4567-e89b-12d3-a456-426614174000", doctoRelacionado.IdDocumento)
	assert.Equal(t, "A", *doctoRelacionado.Serie)
	assert.Equal(t, "1001", *doctoRelacionado.Folio)
	assert.Equal(t, "MXN", doctoRelacionado.MonedaDR)
	assert.Equal(t, 1.0, doctoRelacionado.NumParcialidad)
	assert.Equal(t, 1500.00, doctoRelacionado.ImpSaldoAnt)
	assert.Equal(t, 1500.00, doctoRelacionado.ImpPagado)
	assert.Equal(t, 0.00, doctoRelacionado.ImpSaldoInsoluto)
	assert.Equal(t, "01", doctoRelacionado.ObjetoImpDR)

	InternalTestRetencionDR(t, doctoRelacionado.ImpuestosDR.RetencionesDR)
	InternalTestTrasladoDR(t, doctoRelacionado.ImpuestosDR.TrasladosDR)

}

func InternalTestRetencionDR(t *testing.T, retencionesDR *[]pagos20.RetencionDRPagos20) {
	assert.NotNil(t, retencionesDR)
	assert.Equal(t, len(*retencionesDR), 2)
	retencionDR := (*retencionesDR)[0]
	assert.Equal(t, 1000.000001, retencionDR.BaseDR)
	assert.Equal(t, "001", retencionDR.ImpuestoDR)
	assert.Equal(t, "Tasa", retencionDR.TipoFactorDR)
	assert.Equal(t, 0.160000, retencionDR.TasaOCuotaDR)
	assert.Equal(t, 160.00, retencionDR.ImporteDR)
}

func InternalTestTrasladoDR(t *testing.T, trasladosDR *[]pagos20.TrasladoDRPagos20) {
	assert.NotNil(t, trasladosDR)
	assert.Equal(t, len(*trasladosDR), 2)
	trasladoDR := (*trasladosDR)[0]
	assert.Equal(t, 2000.000001, trasladoDR.BaseDR)
	assert.Equal(t, "002", trasladoDR.ImpuestoDR)
	assert.Equal(t, "Tasa", trasladoDR.TipoFactorDR)
	assert.Equal(t, 0.160000, *trasladoDR.TasaOCuotaDR)
	assert.Equal(t, 320.00, *trasladoDR.ImporteDR)
}

func InternalTestRetencionP(t *testing.T, retencionesP *[]pagos20.RetencionPPagos20) {
	assert.NotNil(t, retencionesP)
	assert.Equal(t, len(*retencionesP), 2)
	retencionP := (*retencionesP)[0]
	assert.Equal(t, "ISR", retencionP.ImpuestoP)
	assert.Equal(t, float32(1500.00), retencionP.ImporteP)
}

func InternalTestTrasladoP(t *testing.T, trasladosP *[]pagos20.TrasladoPPagos20) {
	assert.NotNil(t, trasladosP)
	assert.Equal(t, len(*trasladosP), 2)
	trasladoP := (*trasladosP)[0]
	assert.Equal(t, 5000.00, trasladoP.BaseP)
	assert.Equal(t, "IVA", trasladoP.ImpuestoP)
	assert.Equal(t, "Tasa", trasladoP.TipoFactorP)
	assert.Equal(t, 0.16, *trasladoP.TasaOCuotaP)
	assert.Equal(t, 800.00, *trasladoP.ImporteP)
}
