package addenda

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/addenda"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetRBoschSapForTest(filename string, t *testing.T) (addenda.RBoschSap, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed addenda.RBoschSap
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	testing2.GenerateJSONFromStructure("rboschsap.json", parsed)
	return parsed, errUnmashal
}

func TestRCoschSap(t *testing.T) {
	_, _ = GetRBoschSapForTest("./rboschsap.xml", t)
}
