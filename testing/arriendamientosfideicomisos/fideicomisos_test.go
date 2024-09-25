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
	InternalTestFullAttributes(t, arriendamiento)
}

func InternalTestFullAttributes(t *testing.T, arriendamiento arriendamientos10.ArriendamientosFideicomisos10) {
	assert.Equal(t, "1.0", arriendamiento.Version)
	assert.Equal(t, 3.0, arriendamiento.Pago)
	assert.Equal(t, 3.0, arriendamiento.Rendimiento)
	assert.Equal(t, 4.0, arriendamiento.Deduccion)
	assert.Equal(t, 5.0, arriendamiento.MontoTotal)
	assert.Equal(t, 6.0, arriendamiento.MontoResultado)
	assert.Equal(t, 7.0, arriendamiento.MontoOtros)
	assert.Equal(t, "ArrendamientoEnFideicomiso", arriendamiento.DescripcionMontos)

}
