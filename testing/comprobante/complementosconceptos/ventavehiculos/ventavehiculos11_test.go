package ventavehiculos

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"os"
	"testing"

	ventavehiculos11 "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementosconcepto/ventavehiculos"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetVentaVehiculos11ForTest(filename string, t *testing.T) (ventavehiculos11.VentaVehiculos11, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed ventavehiculos11.VentaVehiculos11
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	GenerateJSONFromXMLVentaVehiculos11("ventavehiculos11.json", parsed)
	return parsed, errUnmashal
}

func TestVentaVehiculos11(t *testing.T) {
	ventaVehiculo11, _ := GetVentaVehiculos11ForTest("./ventavehiculos11.xml", t)
	InternalTestFullAtributes11(t, ventaVehiculo11)
	InternalTestFullParte11(t, ventaVehiculo11.Parte)
	Parte := ventaVehiculo11.Parte
	informacionAduanera := (*Parte)[0].InformacionAduanera
	InternalTestFullInformacionAduanera11(t, informacionAduanera)

}

func InternalTestFullAtributes11(t *testing.T, ventaVehiculo ventavehiculos11.VentaVehiculos11) {
	assert.Equal(t, "1.1", ventaVehiculo.Version)
	assert.Equal(t, "KLS829129990200032712", ventaVehiculo.ClaveVehicular)
	assert.Equal(t, "NIVEH287212", ventaVehiculo.Niv)
}

func InternalTestFullParte11(t *testing.T, parte11 *[]ventavehiculos11.ParteVentVehi11) {
	assert.NotNil(t, parte11)
	assert.Equal(t, 1, len(*parte11))
	first := (*parte11)[0]

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

func InternalTestFullInformacionAduanera11(t *testing.T, informacionAduanera *[]ventavehiculos11.InformacionAduaneraVentVehi11) {
	assert.NotNil(t, informacionAduanera)
	assert.Equal(t, 1, len(*informacionAduanera))
	first := (*informacionAduanera)[0]
	assert.Equal(t, "AAA1212129", first.Numero)
	assert.Equal(t, "2016-04-11", first.FechaString)
	assert.NotNil(t, first.Aduana)
	assert.Equal(t, "Tijuana", *first.Aduana)
}

func GenerateJSONFromXMLVentaVehiculos11(namefile string, data ventavehiculos11.VentaVehiculos11) {
	jsonData, err := json.MarshalIndent(data, "", "	")
	err = os.WriteFile(namefile, jsonData, 0644)
	if err != nil {
		log.Println(err)
	}
}
