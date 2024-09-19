package renovacionsustitucionvehi

import (
	"encoding/xml"
	"testing"

	renovasustivehi1 "github.com/SaulEnriqueMR/kore-models/models/renovacionsustitucionvehi"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetRenovaSustiVehiForTest(filename string, t *testing.T) (renovasustivehi1.RenovacionSustitucion10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed renovasustivehi1.RenovacionSustitucion10
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	return parsed, errUnmashal
}

func TestFullPFIntegrante(t *testing.T) {
	renovacionsustivehi, _ := GetRenovaSustiVehiForTest("./renovacionsustitucionvehi.xml", t)
	assert.NotNil(t, renovacionsustivehi)
	InternalTestBaseRenovaSustiVehi(t, renovacionsustivehi)
}

func InternalTestBaseRenovaSustiVehi(t *testing.T, renovacion renovasustivehi1.RenovacionSustitucion10) {
	assert.Equal(t, "1.0", renovacion.Version)
	assert.Equal(t, "01", renovacion.TipoDecreto)
	InternalTestDecretoRenovacion(t, renovacion.DecretoRenovacion)
	InternalTestDecretoSustitucion(t, renovacion.DecretoSustitucion)
}

func InternalTestDecretoRenovacion(t *testing.T, decreto *renovasustivehi1.DecretoRenovacion) {
	assert.Equal(t, "02", decreto.VehEnaj)
	InternalTestVehiculosUsadosRenovacion(t, decreto.VehiculosUsados)
	InternalTestVehiculosNuevosRenovacion(t, decreto.VehiculosNuevos)
}

func InternalTestDecretoSustitucion(t *testing.T, decreto *renovasustivehi1.DecretoSustitucion) {
	assert.Equal(t, "02", decreto.VehEnaj)
	InternalTestVehiculosNuevosSustitucion(t, decreto.VehiculosNuevos)
	InternalTestVehiculosUsadosSustitucion(t, decreto.VehiculosUsados)
}

func InternalTestVehiculosUsadosRenovacion(t *testing.T, vehiculos []renovasustivehi1.VehiculosUsados) {
	assert.NotNil(t, vehiculos)
	assert.Equal(t, len(vehiculos), 2)
	vehiculo := (vehiculos)[0]
	assert.Equal(t, 250000.00, vehiculo.PrecioVehUsado)
	assert.Equal(t, "AUTOMOVIL", vehiculo.TipoVeh)
	assert.Equal(t, "Toyota", vehiculo.Marca)
	assert.Equal(t, "Sedan", vehiculo.TipooClase)
	assert.Equal(t, 2018, vehiculo.Anio)
	assert.Equal(t, "Corolla", *vehiculo.Modelo)
	assert.Equal(t, "1HGBH41JXMN109186", *vehiculo.NIV)
	assert.Equal(t, "1HGBH41JXMN109186", *vehiculo.NumSerie)
	assert.Equal(t, "XYZ1234", vehiculo.NumPlacas)
	assert.Equal(t, "XJHZ1234567890123", *vehiculo.NumMotor)
	assert.Equal(t, "ABC987654321", vehiculo.NumFolTarjCir)
	assert.Equal(t, "12345678901234567890", *vehiculo.NumPedIm)
	assert.Equal(t, "Aduana de Tijuana", *vehiculo.Aduana)
	assert.Equal(t, "2021-07-15", *vehiculo.FechaRegulVeh)
	assert.Equal(t, "A1B2C3D4E5F6G7H8I9J0", vehiculo.Foliofiscal)
}

func InternalTestVehiculosNuevosRenovacion(t *testing.T, vehiculos []renovasustivehi1.VehiculosNuevos) {
	assert.NotNil(t, vehiculos)
	assert.Equal(t, len(vehiculos), 1)
	vehiculo := (vehiculos)[0]
	assert.Equal(t, 2022, vehiculo.Anio)
	assert.Equal(t, "Hilux", *vehiculo.Modelo)
	assert.Equal(t, "ABC1234", vehiculo.NumPlacas)
	assert.Equal(t, "XAXX010101000", *vehiculo.Rfc)
}

func InternalTestVehiculosUsadosSustitucion(t *testing.T, vehiculos []renovasustivehi1.VehiculosUsados) {
	assert.NotNil(t, vehiculos)
	assert.Equal(t, len(vehiculos), 1)
	vehiculo := (vehiculos)[0]
	assert.Equal(t, 250000.00, vehiculo.PrecioVehUsado)
	assert.Equal(t, "AUTOMOVIL", vehiculo.TipoVeh)
	assert.Equal(t, "Toyota", vehiculo.Marca)
	assert.Equal(t, "Sedan", vehiculo.TipooClase)
	assert.Equal(t, 2018, vehiculo.Anio)
	assert.Equal(t, "Corolla", *vehiculo.Modelo)
	assert.Equal(t, "1HGBH41JXMN109186", *vehiculo.NIV)
	assert.Equal(t, "1HGBH41JXMN109186", *vehiculo.NumSerie)
	assert.Equal(t, "XYZ1234", vehiculo.NumPlacas)
	assert.Equal(t, "XJHZ1234567890123", *vehiculo.NumMotor)
	assert.Equal(t, "ABC987654321", vehiculo.NumFolTarjCir)
	assert.Equal(t, "12345678901234567890", *vehiculo.NumPedIm)
	assert.Equal(t, "Aduana de Tijuana", *vehiculo.Aduana)
	assert.Equal(t, "2021-07-15", *vehiculo.FechaRegulVeh)
	assert.Equal(t, "A1B2C3D4E5F6G7H8I9J0", vehiculo.Foliofiscal)
}

func InternalTestVehiculosNuevosSustitucion(t *testing.T, vehiculos []renovasustivehi1.VehiculosNuevos) {
	assert.NotNil(t, vehiculos)
	assert.Equal(t, len(vehiculos), 3)
	vehiculo := (vehiculos)[2]
	assert.Equal(t, 2022, vehiculo.Anio)
	assert.Equal(t, "Hilux", *vehiculo.Modelo)
	assert.Equal(t, "ABC1234", vehiculo.NumPlacas)
	assert.Equal(t, "XAXX010101000", *vehiculo.Rfc)
}
