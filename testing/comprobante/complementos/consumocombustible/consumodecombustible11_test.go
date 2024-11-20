package complementos

import (
	"encoding/xml"
	"testing"

	consumoCombus11 "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/consumocombustible"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetConsumoDeCombustible11ForTest(filename string, t *testing.T) (consumoCombus11.ConsumoDeCombustible11, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed consumoCombus11.ConsumoDeCombustible11
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	testing2.GenerateJSONFromStructure("consumodecombustible11.json", parsed)
	return parsed, errUnmarshal
}

func TestConsumoDeCombustible11(t *testing.T) {
	consumoCombustible11, _ := GetConsumoDeCombustible11ForTest("./consumodecombustible11.xml", t)
	InternalTestFullAtributesConsumoCombus11(t, consumoCombustible11)
	InternalTestFullAtributesConceptoConsumoCombus11(t, consumoCombustible11.Conceptos.Concepto[0])
	InternalTestFullAtributesDeterminadoConsumoCombus11(t, consumoCombustible11.Conceptos.Concepto[0].Determinados.Determinado[0])
}

func InternalTestFullAtributesConsumoCombus11(t *testing.T, consumoCombustible11 consumoCombus11.ConsumoDeCombustible11) {
	assert.Equal(t, "1.1", consumoCombustible11.Version)
	assert.Equal(t, "monedero electrónico", consumoCombustible11.TipoOperacion)
	assert.Equal(t, "E1921212", consumoCombustible11.NoCuenta)
	assert.Equal(t, 10000.0, *consumoCombustible11.Subtotal)
	assert.Equal(t, 10000.0, consumoCombustible11.Total)
}

func InternalTestFullAtributesConceptoConsumoCombus11(t *testing.T, conceptoConsumoCombus11 consumoCombus11.ConceptoConsumoDeCombustibles11) {
	assert.Equal(t, "0001", conceptoConsumoCombus11.Identificador)
	assert.Equal(t, "2016-04-08T06:02:00", conceptoConsumoCombus11.FechaString)
	assert.Equal(t, "XAXX010101000", conceptoConsumoCombus11.Rfc)
	assert.Equal(t, "111", conceptoConsumoCombus11.ClaveEstacion)
	assert.Equal(t, float64(1), conceptoConsumoCombus11.Cantidad)
	assert.Equal(t, "Diesel", conceptoConsumoCombus11.NombreCombustible)
	assert.Equal(t, "1291398", conceptoConsumoCombus11.FolioOperacion)
	assert.Equal(t, 10000.0, conceptoConsumoCombus11.ValorUnitario)
	assert.Equal(t, 10000.0, conceptoConsumoCombus11.Importe)
	assert.Equal(t, "Gasolina", conceptoConsumoCombus11.TipoCombustible)
}

func InternalTestFullAtributesDeterminadoConsumoCombus11(t *testing.T, determinadoConsumoCombus11 consumoCombus11.DeterminadoConsumoDeCombus11) {
	assert.Equal(t, "IVA", determinadoConsumoCombus11.Impuesto)
	assert.Equal(t, float64(16), determinadoConsumoCombus11.TasaOCuota)
	assert.Equal(t, 160.0, determinadoConsumoCombus11.Importe)
}
