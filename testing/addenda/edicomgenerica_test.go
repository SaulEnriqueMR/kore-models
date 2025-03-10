package addenda

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/addenda"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetEdicomGenericaForTest(filename string, t *testing.T) (addenda.EdicomGenerica, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed addenda.EdicomGenerica
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	testing2.GenerateJSONFromStructure("edicomgenerica.json", parsed)
	return parsed, errUnmashal
}

func TestEdicomGenerica(t *testing.T) {
	_, _ = GetEdicomGenericaForTest("./edicomgenerica.xml", t)
}
