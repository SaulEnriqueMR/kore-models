package parcialescontruccion

import (
	"encoding/xml"
	"testing"

	parciales1 "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/parcialesconstruccion"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetParcialesConstruccionForTest(filename string, t *testing.T) (parciales1.ParcialesConstruccion10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed parciales1.ParcialesConstruccion10
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	return parsed, errUnmashal
}

func TestFullParcialesConstruccion(t *testing.T) {
	parciales, _ := GetParcialesConstruccionForTest("./parcialesconstruccion10.xml", t)
	assert.NotNil(t, parciales)
	InternalTestBaseParcialesConstruccion10(t, parciales)
}

func InternalTestBaseParcialesConstruccion10(t *testing.T, parcialesConstruccion parciales1.ParcialesConstruccion10) {
	assert.Equal(t, "1.0", parcialesConstruccion.Version)
	assert.Equal(t, "12345-67890", parcialesConstruccion.NumPerLicoAut)
	InternalTestInmueble(t, parcialesConstruccion.Inmueble)
}

func InternalTestInmueble(t *testing.T, inmueble parciales1.Inmueble) {
	assert.Equal(t, "Avenida Reforma", inmueble.Calle)
	assert.Equal(t, "123", *inmueble.NoExterior)
	assert.Equal(t, "A", *inmueble.NoInterior)
	assert.Equal(t, "Centro", *inmueble.Colonia)
	assert.Equal(t, "Ciudad de México", *inmueble.Localidad)
	assert.Equal(t, "Cerca del Monumento a la Independencia", *inmueble.Referencia)
	assert.Equal(t, "Cuauhtémoc", inmueble.Municipio)
	assert.Equal(t, "DF", inmueble.Estado)
	assert.Equal(t, "06000", inmueble.CodigoPostal)
}
