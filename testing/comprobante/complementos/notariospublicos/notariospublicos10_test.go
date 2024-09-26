package complementoconcepto

import (
	"encoding/xml"
	"testing"

	notariosPub2 "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/notariospublicos"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetNotariosPublicos10ForTest(filename string, t *testing.T) (notariosPub2.NotariosPublicos10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed notariosPub2.NotariosPublicos10
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	testing2.GenerateJSONFromStructure("notariospublicos10.json", parsed)
	return parsed, errUnmarshal
}

func TestNotariosPublicos10(t *testing.T) {
	notariosPublicos10, _ := GetNotariosPublicos10ForTest("./notariospublicos10.xml", t)
	InternalTestFullAtributesNotariosPublicos10(t, notariosPublicos10)
	InternalFullTestAtributesDescInmuebleNotariosPub10(t, *notariosPublicos10.DescInmueble)
	InternalTestFullAtributesDatosOperacionNotariosPub10(t, notariosPublicos10.DatosOperacion)
	InternalTestFullAtributesDatosNotarioNortariospub10(t, notariosPublicos10.DatosNotario)
	InternalTestFullAtributesDatosEnajenanteNotariosPub10(t, notariosPublicos10.DatosEnajenante)
	InternalTestFullAtributesDatosAdquirienteNotariosPub10(t, notariosPublicos10.DatosAdquiriente)

}

func InternalTestFullAtributesNotariosPublicos10(t *testing.T, publicos10 notariosPub2.NotariosPublicos10) {
	assert.Equal(t, "1.0", publicos10.Version)
}

func InternalFullTestAtributesDescInmuebleNotariosPub10(t *testing.T, pub10 []notariosPub2.DescInmuebleNotariosPub10) {
	assert.Equal(t, 1, len(pub10))
	assert.Equal(t, "01", pub10[0].TipoInmueble)
	assert.Equal(t, "Conocida", pub10[0].Calle)
	assert.Equal(t, "218", *pub10[0].NoExterior)
	assert.Equal(t, "1", *pub10[0].NoInterior)
	assert.Equal(t, "Colonia1", *pub10[0].Colonia)
	assert.Equal(t, "Localidad1", *pub10[0].Localidad)
	assert.Equal(t, "Referencia1", *pub10[0].Referencia)
	assert.Equal(t, "Oaxaca de Juárez", pub10[0].Municipio)
	assert.Equal(t, "14", pub10[0].Estado)
	assert.Equal(t, "MEX", pub10[0].Pais)
	assert.Equal(t, "68000", pub10[0].CodigoPostal)
}

func InternalTestFullAtributesDatosOperacionNotariosPub10(t *testing.T, pub10 notariosPub2.DatosOperacionNotariosPub10) {
	assert.Equal(t, int64(12345), pub10.NumInstrumentoNotarial)
	assert.Equal(t, "2014-05-05", pub10.FechaInstNotarialString)
	assert.Equal(t, 330000.00, pub10.MontoOperacion)
	assert.Equal(t, 330000.00, pub10.Subtotal)
	assert.Equal(t, 52800.00, pub10.IVA)
}

func InternalTestFullAtributesDatosNotarioNortariospub10(t *testing.T, pub10 notariosPub2.DatosNotarioNotariosPub10) {
	assert.Equal(t, "AAQM010101HCSMNZ00", pub10.Curp)
	assert.Equal(t, int16(3), pub10.NumNotaria)
	assert.Equal(t, "16", pub10.EntidadFederativa)
	assert.Equal(t, "Adscripcion1", *pub10.Adscripcion)
}

func InternalTestFullAtributesDatosEnajenanteNotariosPub10(t *testing.T, pub10 notariosPub2.DatosEnajenanteNotariosPub10) {
	assert.Equal(t, "No", pub10.CoproSocConyugalE)
	InternalTestFullAtributesDatosUnEnajenanteNotariosPub10(t, *pub10.DatosUnEnajenante)
	InternalTestFullAtributesDatosEnajenanteCopSCNotariosPub10(t, *pub10.DatosEnajenanteCopSC)
}

func InternalTestFullAtributesDatosUnEnajenanteNotariosPub10(t *testing.T, pub10 notariosPub2.DatosUnEnajenanteNotariosPub10) {
	assert.Equal(t, "Santiago", pub10.Nombre)
	assert.Equal(t, "Hernandez", pub10.ApellidoPaterno)
	assert.Equal(t, "Martinez", pub10.ApellidoMaterno)
	assert.Equal(t, "XAXX010101000", pub10.Rfc)
	assert.Equal(t, "XAXX010101HOCRRN03", pub10.Curp)
}

func InternalTestFullAtributesDatosEnajenanteCopSCNotariosPub10(t *testing.T, pub10 []notariosPub2.DatosEnajenanteCopSCNotariosPub10) {
	assert.Equal(t, 1, len(pub10))
	assert.Equal(t, "Francisco", pub10[0].Nombre)
	assert.Equal(t, "Vazquez", *pub10[0].ApellidoPaterno)
	assert.Equal(t, "Ramirez", *pub10[0].ApellidoMaterno)
	assert.Equal(t, "XAXX010101000", pub10[0].Rfc)
	assert.Equal(t, "XAXX011001HOCRRN09", *pub10[0].Curp)
	assert.Equal(t, 17.00, pub10[0].Porcentaje)
}

func InternalTestFullAtributesDatosAdquirienteNotariosPub10(t *testing.T, pub10 notariosPub2.DatosAdquirienteNotariosPub10) {
	assert.Equal(t, "Si", pub10.CoproSocConyugalE)
	InternalTestFullAtributesDatosUnAdquirienteNotariosPub10(t, *pub10.DatosUnAdquiriente)
	InternalTestFullAtributesDatosAdquirientesCopSCNotariosPub10(t, *pub10.DatosAdquirientesCopSC)
}

func InternalTestFullAtributesDatosUnAdquirienteNotariosPub10(t *testing.T, pub10 notariosPub2.DatosUnAdquirienteNotariosPub10) {
	assert.Equal(t, "Armando", pub10.Nombre)
	assert.Equal(t, "Contreras", pub10.ApellidoPaterno)
	assert.Equal(t, "Velazquez", pub10.ApellidoMaterno)
	assert.Equal(t, "AACO010101000", pub10.Rfc)
	assert.Equal(t, "XAXX010101HOCRRN03", *pub10.Curp)
}

func InternalTestFullAtributesDatosAdquirientesCopSCNotariosPub10(t *testing.T, pub10 []notariosPub2.DatosAdquirientesCopSCNotariosPub10) {
	assert.NotNil(t, pub10)
	assert.Equal(t, 1, len(pub10))
	assert.Equal(t, "Arturo", pub10[0].Nombre)
	assert.Equal(t, "Montes", *pub10[0].ApellidoPaterno)
	assert.Equal(t, "Herrera", *pub10[0].ApellidoMaterno)
	assert.Equal(t, "AUMO010101000", pub10[0].Rfc)
	assert.Equal(t, "XAXX011001HOCRRN09", *pub10[0].Curp)
	assert.Equal(t, 40.00, pub10[0].Porcentaje)
}
