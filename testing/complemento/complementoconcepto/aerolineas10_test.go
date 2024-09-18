package complementoconcepto

import (
	"encoding/xml"
	aerolineas10 "github.com/SaulEnriqueMR/kore-models/models/complemento/complementoconcepto"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetAerolineas10ForTest(filename string, t *testing.T) (aerolineas10.Aerolineas10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed aerolineas10.Aerolineas10
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	return parsed, errUnmashal
}

func TestAerolineas10(t *testing.T) {
	aerolinea10, _ := GetAerolineas10ForTest("./aerolineas10.xml", t)
	InternalTestFullAtributesAerolineas10(t, aerolinea10)
	otrosCargos := aerolinea10.Otroscargos
	InternalTestFullOtrosCargosAero11(t, otrosCargos)
	cargo := (*otrosCargos)[0].Cargo
	InternalTestFullCargoAero11(t, cargo)
}

func InternalTestFullAtributesAerolineas10(t *testing.T, aerolinea10 aerolineas10.Aerolineas10) {
	assert.Equal(t, "1.0", aerolinea10.Version)
	assert.Equal(t, "1000.00", aerolinea10.Tua)
}

func InternalTestFullOtrosCargosAero11(t *testing.T, otrosCargos *[]aerolineas10.OtrosCargosAero10) {
	assert.NotNil(t, otrosCargos)
	assert.Equal(t, 1, len(*otrosCargos))
	firstOtrosCargo := (*otrosCargos)[0]

	assert.Equal(t, "60.00", firstOtrosCargo.TotalCargos)
}

func InternalTestFullCargoAero11(t *testing.T, cargo []aerolineas10.CargoAero10) {
	assert.NotNil(t, cargo)
	assert.Equal(t, 1, len(cargo))
	firstCargo := (cargo)[0]

	assert.Equal(t, "AA", firstCargo.CodigoCargo)
	assert.Equal(t, "60.00", firstCargo.Importe)
}
