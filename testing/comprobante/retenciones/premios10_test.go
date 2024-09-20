package retenciones

import (
	"encoding/xml"
	premios2 "github.com/SaulEnriqueMR/kore-models/models/comprobante/retenciones"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetPremios10ForTest(filename string, t *testing.T) (premios2.Premios10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed premios2.Premios10
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	return parsed, errUnmarshal
}

func TestPremios10(t *testing.T) {
	premios10, _ := GetPremios10ForTest("./premios10.xml", t)
	assert.Equal(t, "1.0", premios10.Version)
	assert.Equal(t, "19", premios10.EntidadFederativa)
	assert.Equal(t, 28263.91, premios10.MontTotPago)
	assert.Equal(t, 28263.91, premios10.MontTotPagoGrav)
	assert.Equal(t, 0.0, premios10.MontTotPagoExent)
}
