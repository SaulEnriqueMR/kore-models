package addenda

import (
	"encoding/xml"
	"testing"

	"github.com/SaulEnriqueMR/kore-models/models/addenda"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetInformacionAdicionalForTest(filename string, t *testing.T) (addenda.AdditionalInfo, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed addenda.AdditionalInfo
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	testing2.GenerateJSONFromStructure("informacionAdicional.json", parsed)
	return parsed, errUnmarshal
}

func TestInformacionAdicional(t *testing.T) {
	_, _ = GetInformacionAdicionalForTest("./InformacionAdicional.xml", t)
}
