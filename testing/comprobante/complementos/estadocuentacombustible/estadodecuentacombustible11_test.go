package complementos

import (
	"encoding/xml"
	edocuentacombus11 "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/estadodecuentacombustible"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetEstadoDeCuentaCombustible11ForTest(filename string, t *testing.T) (edocuentacombus11.EstadoDeCuentaCombustible11, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed edocuentacombus11.EstadoDeCuentaCombustible11
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	testing2.GenerateJSONFromStructure("estadodecuentacombustible11.json", parsed)
	return parsed, errUnmarshal
}

func TestEstadoDeCuentaCombustible11(t *testing.T) {
	estadoCuentaCombus11, _ := GetEstadoDeCuentaCombustible11ForTest("./estadodecuentacombustible11.xml", t)
	InternalTestFullAtributesEstadoDeCuentaCombus11(t, estadoCuentaCombus11)
	InternalTestFullAtributesConceptoEstadoCuentaCombus11(t, estadoCuentaCombus11.Conceptos.ConceptoEstadoDeCuentaCombustible[0])
	InternalTestFullAtributesTrasladoEstadoCuentaCombus11(t, estadoCuentaCombus11.Conceptos.ConceptoEstadoDeCuentaCombustible[0].Traslados.Traslado)
}

func InternalTestFullAtributesEstadoDeCuentaCombus11(t *testing.T, combustible11 edocuentacombus11.EstadoDeCuentaCombustible11) {
	assert.Equal(t, "1.1", combustible11.Version)
	assert.Equal(t, "Tarjeta", combustible11.TipoOperacion)
	assert.Equal(t, "11238", combustible11.NumeroDeCuenta)
	assert.Equal(t, 1000.0, combustible11.SubTotal)
	assert.Equal(t, 1160.0, combustible11.Total)
}

func InternalTestFullAtributesConceptoEstadoCuentaCombus11(t *testing.T, conceptocombustible11 edocuentacombus11.ConceptoEstadoDeCuentaCombustible11) {
	assert.Equal(t, "12234", conceptocombustible11.Identificador)
	assert.Equal(t, "2015-11-02T10:20:02", conceptocombustible11.FechaString)
	assert.Equal(t, "XAXX010101000", conceptocombustible11.Rfc)
	assert.Equal(t, "111", conceptocombustible11.ClaveEstacion)
	assert.Equal(t, "617", *conceptocombustible11.Tar)
	assert.Equal(t, "3.0", conceptocombustible11.Cantidad)
	assert.Equal(t, "32011", conceptocombustible11.NoIdentificacion)
	assert.Equal(t, "Litros", *conceptocombustible11.Unidad)
	assert.Equal(t, "DIESEL", conceptocombustible11.NombreCombustible)
	assert.Equal(t, "01", conceptocombustible11.FolioOperacion)
	assert.Equal(t, 14.50, conceptocombustible11.ValorUnitario)
	assert.Equal(t, 1000.00, conceptocombustible11.Importe)
}

func InternalTestFullAtributesTrasladoEstadoCuentaCombus11(t *testing.T, trasladocombustible11 []edocuentacombus11.TrasladoEstadoDeCuentaCombus11) {
	assert.Equal(t, "IVA", trasladocombustible11[0].Impuesto)
	assert.Equal(t, 0.16, trasladocombustible11[0].TasaoCuota)
	assert.Equal(t, 172.78, trasladocombustible11[0].Importe)

	assert.Equal(t, "IEPS", trasladocombustible11[1].Impuesto)
	assert.Equal(t, 0.08, trasladocombustible11[1].TasaoCuota)
	assert.Equal(t, 79.99, trasladocombustible11[1].Importe)

}
