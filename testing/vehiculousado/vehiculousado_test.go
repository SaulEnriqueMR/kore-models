package vehiculousado

import (
	"encoding/xml"
	"testing"

	vehiculousado1 "github.com/SaulEnriqueMR/kore-models/models/vehiculousado"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetTurisVehiculoUsadoForTest(filename string, t *testing.T) (vehiculousado1.VehiculoUsado10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed vehiculousado1.VehiculoUsado10
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	return parsed, errUnmashal
}

func TestFullVehiculoUsado(t *testing.T) {
	vahiculousado, _ := GetTurisVehiculoUsadoForTest("./vehiculousado10.xml", t)
	assert.NotNil(t, vahiculousado)
	InternalTestBaseVehiculoUsado(t, vahiculousado)
}

func InternalTestBaseVehiculoUsado(t *testing.T, vehiculo vehiculousado1.VehiculoUsado10) {
	assert.Equal(t, "1.0", vehiculo.Version)
	assert.Equal(t, 150000.00, vehiculo.MontoAdquisicion)
	assert.Equal(t, 140000.00, vehiculo.MontoEnajenacion)
	assert.Equal(t, "1234567", vehiculo.ClaveVehicular)
	assert.Equal(t, "Toyota", vehiculo.Marca)
	assert.Equal(t, "Sedan", vehiculo.Tipo)
	assert.Equal(t, "2018", vehiculo.Modelo)
	assert.Equal(t, "1HGBH41JXMN109186", *vehiculo.NumeroMotor)
	assert.Equal(t, "2HGBH41JXMN109186", *vehiculo.NumeroSerie)
	assert.Equal(t, "JTHBJ46G992109186", *vehiculo.NIV)
	assert.Equal(t, 135000.00, vehiculo.Valor)
	InternalTestInformacionAduanera(t, vehiculo.InformacionAduanera)
	InternalTestInformacionAduanera1(t, vehiculo.InformacionAduanera)
}

func InternalTestInformacionAduanera(t *testing.T, infoArray *[]vehiculousado1.InformacionAduanera) {
	assert.NotNil(t, infoArray)
	assert.Equal(t, len(*infoArray), 3)
	info := (*infoArray)[0]
	assert.Equal(t, "12345", info.Numero)
	assert.Equal(t, "2022-05-15", info.FechaString)
	assert.Equal(t, "Aduana del Puerto de Veracruz", *info.Aduana)
}

func InternalTestInformacionAduanera1(t *testing.T, infoArray *[]vehiculousado1.InformacionAduanera) {
	info := (*infoArray)[2]
	assert.Equal(t, "54321", info.Numero)
	assert.Equal(t, "2020-12-10", info.FechaString)
	assert.Equal(t, "Aduana de Lázaro Cárdenas", *info.Aduana)
}
