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
	testing2.GenerateJSONFromStructure("serviciosplataformastecnologicas10.json", parsed)
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
	assert.Equal(t, "10", plataformasDigitales10.NoServicio)
	assert.Equal(t, 5000.00, plataformasDigitales10.MontoTotalServicioSinIva)
	assert.Equal(t, 800.00, plataformasDigitales10.TotalIvaTrasladado)
	assert.Equal(t, 200.00, plataformasDigitales10.TotalIvaRetenido)
	assert.Equal(t, 600.00, plataformasDigitales10.TotalIsrRetenido)
	assert.Equal(t, 100.00, plataformasDigitales10.DiferenciaIvaEntregadoPrestadorServicios)
	assert.Equal(t, 300.00, plataformasDigitales10.MontoTotalUsoPlataforma)
	assert.Equal(t, 150.00, *plataformasDigitales10.MontoTotalContribucionGubernamental)
}

func InternalTestFullAtributesDetallesDelServicioPlatTecno10(t *testing.T, tecno10 []platdigital2.DetallesDelServicioPlatTecno10) {
	assert.NotNil(t, tecno10)
	assert.Equal(t, 1, len(tecno10))
	assert.Equal(t, "01", tecno10[0].FormaPagoServicio)
	assert.Equal(t, "A01", tecno10[0].TipoServicio)
	assert.Equal(t, "B02", *tecno10[0].SubtipoServicio)
	assert.Equal(t, "ABC123456789", *tecno10[0].RfcTerceroAutorizado)
	assert.Equal(t, "2023-09-23", tecno10[0].FechaServ)
	assert.Equal(t, 1000.00, tecno10[0].PrecioServicioSinIva)
	InternalTestFullAtributesImpuestosTrasladosDelServicioPlatTecno10(t, tecno10[0].ImpuestosTrasladadosServicio)
	InternalTestFullAtributesContribucionGubernamentalPlatTecno10(t, *tecno10[0].ContribucionGubernamental)
	InternalTestFullAtributesComisionDelServicioPlatTecno10(t, tecno10[0].ComisionServicio)
}

func InternalTestFullAtributesImpuestosTrasladosDelServicioPlatTecno10(t *testing.T, tecno10 platdigital2.ImpuestosTrasladosDelServicioPlatTecno10) {
	assert.NotNil(t, tecno10)
	assert.Equal(t, 1000.00, tecno10.Base)
	assert.Equal(t, "IVA", tecno10.Impuesto)
	assert.Equal(t, "Tasa", tecno10.TipoFactor)
	assert.Equal(t, 0.16, tecno10.TasaOCuota)
	assert.Equal(t, 160.00, tecno10.Importe)
}

func InternalTestFullAtributesContribucionGubernamentalPlatTecno10(t *testing.T, tecno10 platdigital2.ContribucionGubernamentalPlatTecno10) {
	assert.Equal(t, 200.00, tecno10.ImporteContribucion)
	assert.Equal(t, "01", tecno10.EntidadPagoContribucion)
}

func InternalTestFullAtributesComisionDelServicioPlatTecno10(t *testing.T, tecno10 platdigital2.ComisionDelServicioPlatTecno10) {
	assert.NotNil(t, tecno10)
	assert.Equal(t, 150.50, *tecno10.Base)
	assert.NotNil(t, tecno10)
	assert.Equal(t, 0.05, *tecno10.Porcentaje)
	assert.NotNil(t, tecno10)
	assert.Equal(t, 7.53, *tecno10.Importe)
}
