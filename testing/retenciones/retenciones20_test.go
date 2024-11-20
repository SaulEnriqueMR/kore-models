package retenciones

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/SaulEnriqueMR/kore-models/models/retenciones"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func TestFullRetenciones20(t *testing.T) {
	data := testing2.GetFileContentForTest("./retenciones20.xml", t)
	var parsed retenciones.Retenciones20
	errUnmashal := xml.Unmarshal(data, &parsed)

	testing2.GenerateJSONFromStructure("retenciones20.json", parsed)

	assert.NoError(t, errUnmashal)
	InternalTestFullAttributes20(t, parsed)
	InternalTestRetencionRelacionada20(t, parsed.RetencionRelacionada)
	InternalTestEmisor20(t, parsed.Emisor)
	InternalTestReceptor20(t, parsed.Receptor)
	InternalTestPeriodo20(t, parsed.Periodo)
	InternalTestTotales20(t, parsed.Totales)
}

func InternalTestFullAttributes20(t *testing.T, r retenciones.Retenciones20) {
	assert.Equal(t, "2.0", r.Version)
	assert.Equal(t, "bG3OOBPprTCS9TJXLvRl4", r.Sello)
	assert.Equal(t, "30001000000500003416", r.NoCertificado)
	assert.Equal(t, "MIIFsDCCA5igAwIBAgIUMz", r.Certificado)
	assert.Equal(t, "2024-03-04T00:00:00", r.Fecha)
	assert.Equal(t, "45110", r.LugarExpedicion)
	assert.Equal(t, "01", r.ClaveRetencion)

	assert.NotNil(t, r.Folio)
	assert.Equal(t, "9", *r.Folio)

	assert.NotNil(t, r.DescripcionRetencion)
	assert.Equal(t, "DESCRIPCION", *r.DescripcionRetencion)
}

func InternalTestRetencionRelacionada20(t *testing.T, retencionesRelacionadas *retenciones.RetencionRelacionada20) {
	assert.NotNil(t, retencionesRelacionadas)

	assert.Equal(t, "02", retencionesRelacionadas.TipoRelacion)
	assert.Equal(t, "9341a0a9-eade-4c20-9f6e-a58aa5438a6a", retencionesRelacionadas.UUID)
	assert.Equal(t, strings.ToUpper("9341a0a9-eade-4c20-9f6e-a58aa5438a6a"), retencionesRelacionadas.Uuid)
}

func InternalTestEmisor20(t *testing.T, emisor retenciones.Emisor20) {
	assert.Equal(t, "EKU9003173C9", emisor.Rfc)
	assert.Equal(t, "ESCUELA KEMPER URGATE", emisor.Nombre)
	assert.Equal(t, "601", emisor.RegimenFiscal)
}

func InternalTestReceptor20(t *testing.T, receptor20 retenciones.Receptor20) {
	assert.Equal(t, "Nacional", receptor20.Nacionalidad)

	assert.NotNil(t, receptor20.Nacional)
	assert.Equal(t, "URE180429TM6", receptor20.Nacional.Rfc)
	assert.Equal(t, "UNIVERSIDAD ROBOTICA ESPAÑOLA", receptor20.Nacional.Nombre)
	assert.Equal(t, "86991", receptor20.Nacional.DomicilioFiscal)
	assert.NotNil(t, receptor20.Nacional.Curp)
	assert.Equal(t, "UREE180429TMRQPM10", *receptor20.Nacional.Curp)

	assert.NotNil(t, receptor20.Extranjero)
	assert.Equal(t, "UNIVERSIDAD ROBOTICA ESPAÑOLA", receptor20.Extranjero.Nombre)
	assert.NotNil(t, receptor20.Extranjero.NumRegIdTrib)
	assert.Equal(t, "URE180429TM6", *receptor20.Extranjero.NumRegIdTrib)
}

func InternalTestPeriodo20(t *testing.T, periodo20 retenciones.Periodo20) {
	assert.Equal(t, "01", periodo20.MesInicio)
	assert.Equal(t, "03", periodo20.MesFin)
	assert.Equal(t, "2023", periodo20.Ejercicio)
}

func InternalTestTotales20(t *testing.T, totales20 retenciones.Totales20) {
	assert.Equal(t, 2000.0, totales20.MontoTotalOperacion)
	assert.Equal(t, 2000.0, totales20.MontoTotalGravado)
	assert.Equal(t, 0.0, totales20.MontoTotalExento)
	assert.Equal(t, 580.0, totales20.MontoTotalRetenido)
	assert.NotNil(t, totales20.UtilidadBimestral)
	assert.Equal(t, 0.0, *totales20.UtilidadBimestral)
	assert.NotNil(t, totales20.IsrCorrespondiente)
	assert.Equal(t, 0.0, *totales20.IsrCorrespondiente)

	InternalTestImpuestosRetenidos20(t, totales20.ImpuestosRetenidos)
}

func InternalTestImpuestosRetenidos20(t *testing.T, impuestosRetenidos *[]retenciones.ImpuestoRetenido20) {
	assert.NotNil(t, impuestosRetenidos)

	assert.Equal(t, 2, len(*impuestosRetenidos))

	first := (*impuestosRetenidos)[0]

	assert.Equal(t, 580.00, first.Monto)
	assert.Equal(t, "03", first.TipoPago)
	assert.NotNil(t, first.Base)
	assert.Equal(t, 2000.0, *first.Base)
	assert.NotNil(t, first.Impuesto)
	assert.Equal(t, "001", *first.Impuesto)
}
