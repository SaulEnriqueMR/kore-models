package ventavehiculos

import (
	"encoding/xml"
	"testing"

	ventavehiculos10 "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementosconcepto/ventavehiculos"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetVentaVehiculos10ForTest(filename string, t *testing.T) (ventavehiculos10.VentaVehiculos10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed ventavehiculos10.VentaVehiculos10
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	return parsed, errUnmashal
}

func TestVentaVehiculos10(t *testing.T) {
	ventaVehiculo, _ := GetVentaVehiculos10ForTest("./ventavehiculos10.xml", t)
	InternalTestFullAtributes10(t, ventaVehiculo)
	InternalTestFullParte10(t, ventaVehiculo.Parte)
	Parte := ventaVehiculo.Parte
	informacionAduanera := (*Parte)[0].InformacionAduanera
	InternalTestFullInformacionAduanera10(t, informacionAduanera)

}

func InternalTestFullAtributes10(t *testing.T, ventaVehiculo ventavehiculos10.VentaVehiculos10) {
	assert.Equal(t, "1.0", ventaVehiculo.Version)
	assert.Equal(t, "KLS829129990200032712", ventaVehiculo.ClaveVehicular)
}

func InternalTestFullParte10(t *testing.T, parte10 *[]ventavehiculos10.ParteVentVehi10) {
	assert.NotNil(t, parte10)
	assert.Equal(t, 1, len(*parte10))
	first := (*parte10)[0]

	assert.Equal(t, "1", first.Cantidad)
	assert.NotNil(t, first.Unidad)
	assert.Equal(t, "PZA", *first.Unidad)
	assert.NotNil(t, first.NoIdentificacion)
	assert.Equal(t, "123938", *first.NoIdentificacion)
	assert.Equal(t, "CHASIS", first.Descripcion)
	assert.NotNil(t, first.ValorUnitario)
	assert.Equal(t, "10000.00", *first.ValorUnitario)
	assert.NotNil(t, first.Importe)
	assert.Equal(t, "10000.00", *first.Importe)
}

func InternalTestFullInformacionAduanera10(t *testing.T, informacionAduanera *[]ventavehiculos10.InformacionAduaneraVentVehi10) {
	assert.NotNil(t, informacionAduanera)
	assert.Equal(t, 1, len(*informacionAduanera))
	first := (*informacionAduanera)[0]
	assert.Equal(t, "AAA1212129", first.Numero)
	assert.Equal(t, "2016-04-11", first.Fecha)
	assert.NotNil(t, first.Aduana)
	assert.Equal(t, "Tijuana", *first.Aduana)
}
