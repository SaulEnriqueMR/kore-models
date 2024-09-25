package retenciones

import (
	"encoding/xml"
	"testing"

	platdigital2 "github.com/SaulEnriqueMR/kore-models/models/retenciones/complementos/serviciosplataformastecnologicas"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetServicioPlataformasDigitalesForTest(filename string, t *testing.T) (platdigital2.ServiciosPlataformasTecnologicas10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed platdigital2.ServiciosPlataformasTecnologicas10
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	return parsed, errUnmarshal
}

func TestServiciosPlataformasDigitales10ForTest(t *testing.T) {
	plataformasDigitales10, _ := GetServicioPlataformasDigitalesForTest("./serviciosplataformastecnologicas10.xml", t)
	InternalTestFullAtributesServiciosPlataformasDigitales(t, plataformasDigitales10)
	InternalTestFullAtributesDetallesDelServicioPlatTecno10(t, *plataformasDigitales10.DetallesDelServicio)
}

func InternalTestFullAtributesServiciosPlataformasDigitales(t *testing.T, plataformasDigitales10 platdigital2.ServiciosPlataformasTecnologicas10) {
	assert.Equal(t, "1.0", plataformasDigitales10.Version)
	assert.Equal(t, "Mensual", plataformasDigitales10.Periodicidad)
	assert.Equal(t, "10", plataformasDigitales10.NumServ)
	assert.Equal(t, 5000.00, plataformasDigitales10.MonTotServSIVA)
	assert.Equal(t, 800.00, plataformasDigitales10.TotalIVATrasladado)
	assert.Equal(t, 200.00, plataformasDigitales10.TotalIVARetenido)
	assert.Equal(t, 600.00, plataformasDigitales10.TotalISRRetenido)
	assert.Equal(t, 100.00, plataformasDigitales10.DifIVAEntregadoPrestServ)
	assert.Equal(t, 300.00, plataformasDigitales10.MonTotalporUsoPlataforma)
	assert.Equal(t, 150.00, *plataformasDigitales10.MonTotalContribucionGubernamental)
}

func InternalTestFullAtributesDetallesDelServicioPlatTecno10(t *testing.T, tecno10 []platdigital2.DetallesDelServicioPlatTecno10) {
	assert.NotNil(t, tecno10)
	assert.Equal(t, 1, len(tecno10))
	assert.Equal(t, "01", tecno10[0].FormaPagoServ)
	assert.Equal(t, "A01", tecno10[0].TipoDeServ)
	assert.Equal(t, "B02", *tecno10[0].SubTipServ)
	assert.Equal(t, "ABC123456789", *tecno10[0].RFCTerceroAutorizado)
	assert.Equal(t, "2023-09-23", tecno10[0].FechaServString)
	assert.Equal(t, 1000.00, tecno10[0].PrecioServSinIVA)
	InternalTestFullAtributesImpuestosTrasladosDelServicioPlatTecno10(t, tecno10[0].ImpuestosTrasladadosDelServicio)
	InternalTestFullAtributesContribucionGubernamentalPlatTecno10(t, *tecno10[0].ContribucionGubernamental)
	InternalTestFullAtributesComisionDelServicioPlatTecno10(t, tecno10[0].ComisionDelServicio)
}

func InternalTestFullAtributesImpuestosTrasladosDelServicioPlatTecno10(t *testing.T, tecno10 platdigital2.ImpuestosTrasladosDelServicioPlatTecno10) {
	assert.NotNil(t, tecno10)
	assert.Equal(t, 1000.00, tecno10.Base)
	assert.Equal(t, "IVA", tecno10.Impuesto)
	assert.Equal(t, "Tasa", tecno10.TipoFactor)
	assert.Equal(t, 0.16, tecno10.TasaCuota)
	assert.Equal(t, 160.00, tecno10.Importe)
}

func InternalTestFullAtributesContribucionGubernamentalPlatTecno10(t *testing.T, tecno10 platdigital2.ContribucionGubernamentalPlatTecno10) {
	assert.Equal(t, 200.00, tecno10.ImpContrib)
	assert.Equal(t, "01", tecno10.EntidadDondePagaLaContribucion)
}

func InternalTestFullAtributesComisionDelServicioPlatTecno10(t *testing.T, tecno10 platdigital2.ComisionDelServicioPlatTecno10) {
	assert.NotNil(t, tecno10)
	assert.Equal(t, 150.50, *tecno10.Base)
	assert.NotNil(t, tecno10)
	assert.Equal(t, 0.05, *tecno10.Porcentaje)
	assert.NotNil(t, tecno10)
	assert.Equal(t, 7.53, *tecno10.Importe)
}
