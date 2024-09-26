package pagoaextranjeros

import (
	"encoding/xml"
	dividendos1 "github.com/SaulEnriqueMR/kore-models/models/retenciones/complementos/dividendos"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetTurisDividendosForTest(filename string, t *testing.T) (dividendos1.Dividendos10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed dividendos1.Dividendos10
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	return parsed, errUnmashal
}

func TestFullPagoATerceros10(t *testing.T) {

}
