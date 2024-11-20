package renovacionsustitucionvehi

import (
	"encoding/xml"
	"testing"

	renovasustivehi1 "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/renovacionsustitucionvehi"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetRenovaSustiVehiForTest(filename string, t *testing.T) (renovasustivehi1.RenovacionSustitucion10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed renovasustivehi1.RenovacionSustitucion10
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	testing2.GenerateJSONFromStructure("renovacion10.json", parsed)
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
	assert.Equal(t, "02", decreto.VehiculoEnajenado)
	InternalTestVehiculosUsadosRenovacion(t, decreto.VehiculosUsados)
	InternalTestVehiculosNuevosRenovacion(t, decreto.VehiculosNuevos)
}

func InternalTestDecretoSustitucion(t *testing.T, decreto *renovasustivehi1.DecretoSustitucion) {
	assert.Equal(t, "02", decreto.VehiculoEnajenado)
	InternalTestVehiculosNuevosSustitucion(t, decreto.VehiculosNuevos)
	InternalTestVehiculosUsadosSustitucion(t, decreto.VehiculosUsados)
}

func InternalTestVehiculosUsadosRenovacion(t *testing.T, vehiculos []renovasustivehi1.VehiculosUsados) {
	assert.NotNil(t, vehiculos)
	assert.Equal(t, len(vehiculos), 2)
	vehiculo := (vehiculos)[0]
	vehiculo1 := (vehiculos)[1]
	assert.Equal(t, 250000.00, vehiculo.PrecioVehUsado)
	assert.Equal(t, "AUTOMOVIL", vehiculo.Tipo)
	assert.Equal(t, "Toyota", vehiculo.Marca)
	assert.Equal(t, "Sedan", vehiculo.TipoOClase)
	assert.Equal(t, "2018", vehiculo.Anio)
	assert.Equal(t, "Corolla", *vehiculo.Modelo)
	assert.Equal(t, "1HGBH41JXMN109186", *vehiculo.NoIdentificacionVehicular)
	assert.Equal(t, "1HGBH41JXMN109186", *vehiculo.NoSerie)
	assert.Equal(t, "XYZ1234", vehiculo.NoPlacas)
	assert.Equal(t, "XJHZ1234567890123", *vehiculo.NoMotor)
	assert.Equal(t, "ABC987654321", vehiculo.NoFolioTarjetaCirculacion)
	assert.Equal(t, "12345678901234567890", *vehiculo.NumeroPedimentoImportacion)
	assert.Equal(t, "Aduana de Tijuana", *vehiculo.Aduana)
	assert.Equal(t, "2021-07-15", *vehiculo.FechaRegulVeh)
	assert.Equal(t, "A1B2C3D4E5F6G7H8I9J0", vehiculo.FolioFiscal)
	assert.Equal(t, "2022-07-15", *vehiculo1.FechaRegulVeh)
}

func InternalTestVehiculosNuevosRenovacion(t *testing.T, vehiculos []renovasustivehi1.VehiculosNuevos) {
	assert.NotNil(t, vehiculos)
	assert.Equal(t, len(vehiculos), 1)
	vehiculo := (vehiculos)[0]
	assert.Equal(t, "2022", vehiculo.Anio)
	assert.Equal(t, "Hilux", *vehiculo.Modelo)
	assert.Equal(t, "ABC1234", vehiculo.NoPlacas)
	assert.Equal(t, "XAXX010101000", *vehiculo.Rfc)
}

func InternalTestVehiculosUsadosSustitucion(t *testing.T, vehiculos []renovasustivehi1.VehiculosUsados) {
	assert.NotNil(t, vehiculos)
	assert.Equal(t, len(vehiculos), 1)
	vehiculo := (vehiculos)[0]

	assert.Equal(t, 250000.00, vehiculo.PrecioVehUsado)
	assert.Equal(t, "AUTOMOVIL", vehiculo.Tipo)
	assert.Equal(t, "Toyota", vehiculo.Marca)
	assert.Equal(t, "Sedan", vehiculo.TipoOClase)
	assert.Equal(t, "2018", vehiculo.Anio)
	assert.Equal(t, "Corolla", *vehiculo.Modelo)
	assert.Equal(t, "1HGBH41JXMN109186", *vehiculo.NoIdentificacionVehicular)
	assert.Equal(t, "1HGBH41JXMN109186", *vehiculo.NoSerie)
	assert.Equal(t, "XYZ1234", vehiculo.NoPlacas)
	assert.Equal(t, "XJHZ1234567890123", *vehiculo.NoMotor)
	assert.Equal(t, "ABC987654321", vehiculo.NoFolioTarjetaCirculacion)
	assert.Equal(t, "12345678901234567890", *vehiculo.NumeroPedimentoImportacion)
	assert.Equal(t, "Aduana de Tijuana", *vehiculo.Aduana)
	assert.Equal(t, "2023-07-15", *vehiculo.FechaRegulVeh)
	assert.Equal(t, "A1B2C3D4E5F6G7H8I9J0", vehiculo.FolioFiscal)
}

func InternalTestVehiculosNuevosSustitucion(t *testing.T, vehiculos []renovasustivehi1.VehiculosNuevos) {
	assert.NotNil(t, vehiculos)
	assert.Equal(t, len(vehiculos), 3)
	vehiculo := (vehiculos)[2]
	assert.Equal(t, "2022", vehiculo.Anio)
	assert.Equal(t, "Hilux", *vehiculo.Modelo)
	assert.Equal(t, "ABC1234", vehiculo.NoPlacas)
	assert.Equal(t, "XAXX010101000", *vehiculo.Rfc)
}
