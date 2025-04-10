package complementos

import (
	"encoding/xml"
	edocuentacombus11 "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/estadodecuentacombustible"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetEstadoDeCuentaCombustible10ForTest(filename string, t *testing.T) (edocuentacombus11.EstadoDeCuentaCombustible10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed edocuentacombus11.EstadoDeCuentaCombustible10
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	testing2.GenerateJSONFromStructure("estadodecuentacombustible10.json", parsed)
	return parsed, errUnmarshal
}

func TestEstadoDeCuentaCombustible10(t *testing.T) {
	estadoCuentaCombus10, _ := GetEstadoDeCuentaCombustible10ForTest("./estadodecuentacombustible10.xml", t)
	InternalTestFullAtributesEstadoDeCuentaCombus10(t, estadoCuentaCombus10)
	InternalTestFullAtributesConceptoEstadoCuentaCombus10(t, estadoCuentaCombus10.Conceptos.Concepto[0])
	InternalTestFullAtributesTrasladoEstadoCuentaCombus10(t, estadoCuentaCombus10.Conceptos.Concepto[0].Traslados.Traslado)
}

func InternalTestFullAtributesEstadoDeCuentaCombus10(t *testing.T, combustible10 edocuentacombus11.EstadoDeCuentaCombustible10) {
	assert.Equal(t, "Tarjeta", combustible10.TipoOperacion)
	assert.Equal(t, "38572-3", combustible10.NoCuenta)
	assert.Equal(t, 10.0, *combustible10.Subtotal)
	assert.Equal(t, 77061.94, combustible10.Total)
}

func InternalTestFullAtributesConceptoEstadoCuentaCombus10(t *testing.T, conceptocombustible10 edocuentacombus11.ConceptoEstadoDeCuentaCombustible10) {
	assert.Equal(t, "244086", conceptocombustible10.Identificador)
	assert.Equal(t, "2014-01-02T08:50:25", conceptocombustible10.FechaString)
	assert.Equal(t, "CGP970522EE4", conceptocombustible10.Rfc)
	assert.Equal(t, "76", conceptocombustible10.ClaveEstacion)
	assert.Equal(t, "141.40", conceptocombustible10.Cantidad)
	assert.Equal(t, "Diesel", conceptocombustible10.NombreCombustible)
	assert.Equal(t, "1042", conceptocombustible10.FolioOperacion)
	assert.Equal(t, 12.73, conceptocombustible10.ValorUnitario)
	assert.Equal(t, 1800.00, conceptocombustible10.Importe)
}

func InternalTestFullAtributesTrasladoEstadoCuentaCombus10(t *testing.T, trasladocombustible10 []edocuentacombus11.TrasladoEstadoDeCuentaCombustible10) {
	assert.Equal(t, "IVA", trasladocombustible10[0].Impuesto)
	assert.Equal(t, 16.0, trasladocombustible10[0].TasaOCuota)
	assert.Equal(t, 248.28, trasladocombustible10[0].Importe)
}
