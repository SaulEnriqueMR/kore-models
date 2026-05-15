package addenda

import (
	"encoding/xml"
	"testing"

	"github.com/SaulEnriqueMR/kore-models/models/addenda"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetAddendaDatosForTest(filename string, t *testing.T) (addenda.AddendaDatos, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed addenda.AddendaDatos
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	testing2.GenerateJSONFromStructure("addendadatos.json", parsed)
	return parsed, errUnmashal
}

func TestFullAddendaDatos(t *testing.T) {
	_, _ = GetAddendaDatosForTest("./addendadatos.xml", t)
}
