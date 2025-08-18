package addenda

import (
	"encoding/xml"
	"testing"

	"github.com/SaulEnriqueMR/kore-models/models/addenda"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetAddendaFactoForTest(filename string, t *testing.T) (addenda.AddendaFacto, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed addenda.AddendaFacto
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	testing2.GenerateJSONFromStructure("addendafacto.json", parsed)
	return parsed, errUnmashal
}

func TestAddendaFacto(t *testing.T) {
	_, _ = GetAddendaFactoForTest("./addendafacto.xml", t)
}
