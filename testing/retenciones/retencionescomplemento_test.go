package retenciones

import (
	"encoding/xml"
	retenciones2 "github.com/SaulEnriqueMR/kore-models/models/retenciones"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetRetenciones10ForTestComplementoTest(filename string, t *testing.T) (retenciones2.Retenciones10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed retenciones2.Retenciones10
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	return parsed, errUnmarshal
}

func GetRetenciones20ForTestComplementoTest(filename string, t *testing.T) (retenciones2.Retenciones20, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed retenciones2.Retenciones20
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	return parsed, errUnmarshal
}

func TestRetencionesComplemento(t *testing.T) {
	InternalTestComplementoRetenciones10(t)
	InternalTestComplementoRetenciones20(t)
}

func InternalTestComplementoRetenciones10(t *testing.T) {
	retenciones10, _ := GetRetenciones10ForTestComplementoTest("./retenciones10.xml", t)
	InternalTestFullAtributesComplementoRetenciones(t, *retenciones10.Complemento)
}

func InternalTestComplementoRetenciones20(t *testing.T) {
	retenciones20, _ := GetRetenciones20ForTestComplementoTest("./retenciones20.xml", t)
	InternalTestFullAtributesComplementoRetenciones(t, *retenciones20.Complemento)
}

func InternalTestFullAtributesComplementoRetenciones(t *testing.T, complemento retenciones2.Complemento) {
	assert.NotNil(t, complemento)
	assert.Equal(t, 3, len(*complemento.ArriendamientosFideicomisos.ArriendamientosFideicomisos10))
	assert.Equal(t, 3, len(*complemento.Dividendos.Dividendos10))
	assert.Equal(t, 3, len(*complemento.EnajenacionDeAcciones.EnajenacionDeAcciones10))
	assert.Equal(t, 3, len(*complemento.FideicomisoEmpresarial.FideicomisoEmpresarial10))
	assert.Equal(t, 3, len(*complemento.Intereses.Intereses10))
	assert.Equal(t, 3, len(*complemento.InteresesHipotecarios.InteresesHipotecarios10))
	assert.Equal(t, 3, len(*complemento.OperacionesDerivados.OperacionesDerivados10))
	assert.Equal(t, 2, len(*complemento.PlanesDeRetiro.PlanesDeRetiro10))
	assert.Equal(t, 2, len(*complemento.PlanesDeRetiro.PlanesDeRetiro11))
	assert.Equal(t, 2, len(*complemento.Premios.Premios10))
	assert.Equal(t, 3, len(*complemento.SectorFinanciero.SectorFinanciero10))
	assert.Equal(t, 3, len(*complemento.ServiciosPlataformasTecnologicas.ServiciosPlataformasTecnologicas10))
	assert.Equal(t, 2, len(*complemento.PagosAExtranjeros.PagosAExtranjeros10))
	assert.NotNil(t, complemento.TimbreFiscalDigital.TimbreFiscalDigital10)
	assert.NotNil(t, complemento.TimbreFiscalDigital.TimbreFiscalDigital11)
}
