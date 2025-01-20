package addenda

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/addenda"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetBuzonFiscalForTest(filename string, t *testing.T) (addenda.BuzonFiscal, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed addenda.BuzonFiscal
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	testing2.GenerateJSONFromStructure("buzonfiscal.json", parsed)
	return parsed, errUnmashal
}

func TestFullComprobante40(t *testing.T) {
	_, _ = GetBuzonFiscalForTest("./buzonfiscal.xml", t)
}
