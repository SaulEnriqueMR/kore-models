package fideicomisonoempresarial

import (
	"encoding/xml"
	"testing"

	fideicomiso1 "github.com/SaulEnriqueMR/kore-models/models/retenciones/complementos/fideicomisonoempresarial"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetTurisFideicomisoForTest(filename string, t *testing.T) (fideicomiso1.FideicomisoEmpresarial10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed fideicomiso1.FideicomisoEmpresarial10
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	return parsed, errUnmashal
}

func TestFullFideicomiso(t *testing.T) {
	fideicomiso, _ := GetTurisFideicomisoForTest("./fideicomisonoempresarial10.xml", t)
	assert.NotNil(t, fideicomiso)
	testing2.GenerateJSONFromStructure("fideicomisonoempresarial10.json", fideicomiso)
	InternalTestBaseFideicomisoEmpresarial(t, fideicomiso)
}
func InternalTestBaseFideicomisoEmpresarial(t *testing.T, fideicomiso fideicomiso1.FideicomisoEmpresarial10) {
	assert.Equal(t, "1.0", fideicomiso.Version)

	assert.Equal(t, 300.00, fideicomiso.IngresosOEntradas.MontoTotalEntradasPeriodo)
	assert.Equal(t, 300.00, fideicomiso.IngresosOEntradas.ParteProporcionalAcumulableFideicomiso)
	assert.Equal(t, 800.000000, fideicomiso.IngresosOEntradas.ProporcionMontoTotal)
	assert.Equal(t, "CONCEPTO2", fideicomiso.IngresosOEntradas.IntegracionIngresos.Concepto)

	assert.Equal(t, 50.00, fideicomiso.DeduccionesOSalidas.MontoTotalEgresosPeriodo)
	assert.Equal(t, 0.00, fideicomiso.DeduccionesOSalidas.ParteProporcionalFideicomiso)
	assert.Equal(t, 800.000000, fideicomiso.DeduccionesOSalidas.ProporcionMontoTotal)
	assert.Equal(t, "CONCEPTO2", fideicomiso.DeduccionesOSalidas.IntegracionEgresos.Concepto)

	assert.Equal(t, 750.45, fideicomiso.RetencionesEfectuadasFideicomiso.MontoRetencionRelacionadaFideicomiso)
	assert.Equal(t, "DESCRIPCION", fideicomiso.RetencionesEfectuadasFideicomiso.DescripcionRetencionRelacionadaFideicomiso)
}
