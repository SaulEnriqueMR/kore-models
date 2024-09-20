package intereses

import (
	"encoding/xml"
	"testing"

	intereses1 "github.com/SaulEnriqueMR/kore-models/models/intereses"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetTurisInteresesForTest(filename string, t *testing.T) (intereses1.Intereses10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed intereses1.Intereses10
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	return parsed, errUnmashal
}

func TestFullIntereses(t *testing.T) {
	fideicomiso, _ := GetTurisInteresesForTest("./intereses10.xml", t)
	assert.NotNil(t, fideicomiso)
	InternalTestIntereses(t, fideicomiso)
}

func InternalTestIntereses(t *testing.T, intereses intereses1.Intereses10) {
	assert.Equal(t, "1.0", intereses.Version)
	assert.Equal(t, "SI", intereses.SistFinanciero)
	assert.Equal(t, "NO", intereses.RetiroAORESRetInt)
	assert.Equal(t, "NO", intereses.OperFinancDerivad)
	assert.Equal(t, 1500.00, intereses.MontIntNominal)
	assert.Equal(t, 1200.00, intereses.MontIntReal)
	assert.Equal(t, 100.00, intereses.Perdida)
}
