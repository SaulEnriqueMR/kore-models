package retenciones

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/retenciones"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFullRetenciones10(t *testing.T) {
	data := testing2.GetFileContentForTest("./retenciones10.xml", t)
	var parsed retenciones.Retenciones10
	errUnmashal := xml.Unmarshal(data, &parsed)

	assert.NoError(t, errUnmashal)
	InternalTestFullAttributes10(t, parsed)
	InternalTestEmisor10(t, parsed.Emisor)
	InternalTestReceptor10(t, parsed.Receptor)
	InternalTestPeriodo10(t, parsed.Periodo)
	InternalTestTotales10(t, parsed.Totales)
}

func InternalTestFullAttributes10(t *testing.T, r retenciones.Retenciones10) {
	assert.Equal(t, "1.0", r.Version)
	assert.Equal(t, "RwN+snnlG7o9XzJKvqkAY", r.Sello)
	assert.Equal(t, "20001000000300022755", r.NoCertificado)
	assert.Equal(t, "MIIF7TCCA9WgAwIBAgIUMjAw", r.Certificado)
	assert.Equal(t, "2019-10-20T16:35:28-06:00", r.FechaExp)
	assert.Equal(t, "14", r.ClaveRetencion)

	assert.NotNil(t, r.Folio)
	assert.Equal(t, "54", *r.Folio)

	assert.NotNil(t, r.DescripcionRetencion)
	assert.Equal(t, "DESCRIPCION", *r.DescripcionRetencion)
}

func InternalTestEmisor10(t *testing.T, emisor retenciones.Emisor10) {
	assert.Equal(t, "MAG041126GT8", emisor.Rfc)
	assert.Equal(t, "PREPARATORIA MIGUEL HIDALGO SA", emisor.Nombre)
	assert.Equal(t, "MAG041126GPTWDMA01", emisor.RegimenFiscal)
}

func InternalTestReceptor10(t *testing.T, receptor10 retenciones.Receptor10) {
	assert.Equal(t, "Extranjero", receptor10.Nacionalidad)

	assert.NotNil(t, receptor10.Nacional)
	assert.Equal(t, "REDF465379RT1", receptor10.Nacional.Rfc)
	assert.Equal(t, "Cliente Nacional", receptor10.Nacional.Nombre)
	assert.NotNil(t, receptor10.Nacional.Curp)
	assert.Equal(t, "REDF465379PTRASQ01", *receptor10.Nacional.Curp)

	assert.NotNil(t, receptor10.Extranjero)
	assert.Equal(t, "Cliente Extranjero", receptor10.Extranjero.Nombre)
	assert.NotNil(t, receptor10.Extranjero.NumRegIdTrib)
	assert.Equal(t, "ABCD10231", *receptor10.Extranjero.NumRegIdTrib)
}

func InternalTestPeriodo10(t *testing.T, periodo20 retenciones.Periodo10) {
	assert.Equal(t, "01", periodo20.MesInicio)
	assert.Equal(t, "01", periodo20.MesFin)
	assert.Equal(t, "2018", periodo20.Ejercicio)
}

func InternalTestTotales10(t *testing.T, totales10 retenciones.Totales10) {
	assert.Equal(t, 390.0, totales10.MontoTotalOperacion)
	assert.Equal(t, 90.0, totales10.MontoTotalGravado)
	assert.Equal(t, 10.0, totales10.MontoTotalExento)
	assert.Equal(t, 10.0, totales10.MontoTotalRetenido)

	InternalTestImpuestosRetenidos10(t, totales10.ImpuestosRetenidos)
}

func InternalTestImpuestosRetenidos10(t *testing.T, impuestosRetenidos *[]retenciones.ImpuestoRetenido10) {
	assert.NotNil(t, impuestosRetenidos)

	assert.Equal(t, 2, len(*impuestosRetenidos))

	first := (*impuestosRetenidos)[0]

	assert.Equal(t, 10.0, first.Monto)
	assert.Equal(t, "Pago definitivo", first.TipoPago)
	assert.NotNil(t, first.Base)
	assert.Equal(t, 10.0, *first.Base)
	assert.NotNil(t, first.Impuesto)
	assert.Equal(t, "001", *first.Impuesto)
}
