package addenda

import (
	"encoding/xml"
	"testing"

	"github.com/SaulEnriqueMR/kore-models/models/addenda"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetInssistForTest(filename string, t *testing.T) (addenda.Inssist, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed addenda.Inssist
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	testing2.GenerateJSONFromStructure("innsist.json", parsed)
	return parsed, errUnmashal
}

func TestInnsist(t *testing.T) {
	_, _ = GetInssistForTest("./innsist.xml", t)
}
