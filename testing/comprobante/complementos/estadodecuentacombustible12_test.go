package complementos

import (
	"encoding/xml"
	edocuentacombus12 "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetEstadoDeCuentaCombustible12ForTest(filename string, t *testing.T) (edocuentacombus12.EstadoDeCuentaCombustible12, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed edocuentacombus12.EstadoDeCuentaCombustible12
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	return parsed, errUnmarshal
}

func TestEstadoDeCuentaCombustible12(t *testing.T) {
	estadoCuentaCombus12, _ := GetEstadoDeCuentaCombustible12ForTest("./estadodecuentacombustible12.xml", t)
	InternalTestFullAtributesEstadoDeCuentaCombus12(t, estadoCuentaCombus12)
	InternalTestFullAtributesConceptoEstadoCuentaCombus12(t, estadoCuentaCombus12.Conceptos.ConceptoEstadoDeCuentaCombustible[0])
	InternalTestFullAtributesTrasladoEstadoCuentaCombus12(t, estadoCuentaCombus12.Conceptos.ConceptoEstadoDeCuentaCombustible[0].Traslados.Traslado)
}

func InternalTestFullAtributesEstadoDeCuentaCombus12(t *testing.T, combustible12 edocuentacombus12.EstadoDeCuentaCombustible12) {
	assert.Equal(t, "1.2", combustible12.Version)
	assert.Equal(t, "Tarjeta", combustible12.TipoOperacion)
	assert.Equal(t, "11238", combustible12.NumeroDeCuenta)
	assert.Equal(t, 1000.0, combustible12.SubTotal)
	assert.Equal(t, 1160.0, combustible12.Total)
}

func InternalTestFullAtributesConceptoEstadoCuentaCombus12(t *testing.T, conceptocombustible12 edocuentacombus12.ConceptoEstadoDeCuentaCombustible12) {
	assert.Equal(t, "12234", conceptocombustible12.Identificador)
	assert.Equal(t, "2015-11-02T10:20:02", conceptocombustible12.FechaString)
	assert.Equal(t, "XAXX010101000", conceptocombustible12.Rfc)
	assert.Equal(t, "111", conceptocombustible12.ClaveEstacion)
	assert.Equal(t, "3.0", conceptocombustible12.Cantidad)
	assert.Equal(t, "32011", conceptocombustible12.NoIdentificacion)
	assert.Equal(t, "Litros", *conceptocombustible12.Unidad)
	assert.Equal(t, "DIESEL", conceptocombustible12.NombreCombustible)
	assert.Equal(t, "01", conceptocombustible12.FolioOperacion)
	assert.Equal(t, 14.50, conceptocombustible12.ValorUnitario)
	assert.Equal(t, 1000.00, conceptocombustible12.Importe)
}

func InternalTestFullAtributesTrasladoEstadoCuentaCombus12(t *testing.T, trasladocombustible12 []edocuentacombus12.TrasladoEstadoDeCuentaCombus12) {
	assert.Equal(t, "IVA", trasladocombustible12[0].Impuesto)
	assert.Equal(t, 0.16, trasladocombustible12[0].TasaOCuota)
	assert.Equal(t, 172.78, trasladocombustible12[0].Importe)

	assert.Equal(t, "IEPS", trasladocombustible12[1].Impuesto)
	assert.Equal(t, 0.08, trasladocombustible12[1].TasaOCuota)
	assert.Equal(t, 79.99, trasladocombustible12[1].Importe)

}
