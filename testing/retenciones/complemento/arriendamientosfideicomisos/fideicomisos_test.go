package fideicomisos

import (
	"encoding/xml"
	"testing"

	arriendamientos10 "github.com/SaulEnriqueMR/kore-models/models/retenciones/complementos/arriendamientosfideicomisos"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetArriendamientos10ForTest(filename string, t *testing.T) (arriendamientos10.ArriendamientosFideicomisos10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed arriendamientos10.ArriendamientosFideicomisos10
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	return parsed, errUnmashal
}

func TestFullComprobante40(t *testing.T) {
	arriendamiento, _ := GetArriendamientos10ForTest("./arriendamientos10.xml", t)
	testing2.GenerateJSONFromStructure("arriendamientos10.json", arriendamiento)
	InternalTestFullAttributes(t, arriendamiento)
}

func InternalTestFullAttributes(t *testing.T, arriendamiento arriendamientos10.ArriendamientosFideicomisos10) {
	assert.Equal(t, "1.0", arriendamiento.Version)
	assert.Equal(t, 3.0, arriendamiento.PagoEfectuadoPorFiduciario)
	assert.Equal(t, 3.0, arriendamiento.RendimientoFideicomiso)
	assert.Equal(t, 4.0, arriendamiento.DeduccionCorrespondiente)
	assert.Equal(t, 5.0, arriendamiento.MontoTotalRetenido)
	assert.Equal(t, 6.0, arriendamiento.MontoResultadoFiscalDistribuidoPorFibras)
	assert.Equal(t, 7.0, arriendamiento.MontoOtrosConceptosDistribuidos)
	assert.Equal(t, "ArrendamientoEnFideicomiso", arriendamiento.DescripcionMontos)

}
