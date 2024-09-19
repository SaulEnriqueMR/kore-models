package fideicomisonoempresarial

import (
	"encoding/xml"
	"testing"

	fideicomiso1 "github.com/SaulEnriqueMR/kore-models/models/fideicomisonoempresarial"
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
	fideicomiso, _ := GetTurisFideicomisoForTest("./fideicomisonoempre.xml", t)
	assert.NotNil(t, fideicomiso)
	InternalTestBaseFideicomisoEmpresarial(t, fideicomiso)
}
func InternalTestBaseFideicomisoEmpresarial(t *testing.T, fideicomiso fideicomiso1.FideicomisoEmpresarial10) {
	assert.Equal(t, "1.0", fideicomiso.Version)

	assert.Equal(t, 300.00, fideicomiso.IngresosOEntradas.MontTotEntradasPeriodo)
	assert.Equal(t, 300.00, fideicomiso.IngresosOEntradas.PartPropAcumDelFideicom)
	assert.Equal(t, 800.000000, fideicomiso.IngresosOEntradas.PropDelMontTot)
	assert.Equal(t, "CONCEPTO2", fideicomiso.IngresosOEntradas.IntegracIngresos.Concepto)

	assert.Equal(t, 50.00, fideicomiso.DeduccOSalidas.MontTotEgresPeriodo)
	assert.Equal(t, 0.00, fideicomiso.DeduccOSalidas.PartPropDelFideicom)
	assert.Equal(t, 800.000000, fideicomiso.DeduccOSalidas.PropDelMontTot)

	assert.Equal(t, 750.45, fideicomiso.Retenciones.MontRetRelPagFideic)
	assert.Equal(t, "DESCRIPCION", fideicomiso.Retenciones.DescRetRelPagFideic)
}
