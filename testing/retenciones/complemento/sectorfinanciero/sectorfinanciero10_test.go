package retenciones

import (
	"encoding/xml"
	"testing"

	sectfinan2 "github.com/SaulEnriqueMR/kore-models/models/retenciones/complementos/sectorfinanciero"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetSectorFinanciero10ForTest(filename string, t *testing.T) (sectfinan2.SectorFinanciero10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed sectfinan2.SectorFinanciero10
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	testing2.GenerateJSONFromStructure("sectorfinanciero.json", parsed)
	return parsed, errUnmarshal
}

func TestSectorFinanciero10(t *testing.T) {
	sectorFinanciero10, _ := GetSectorFinanciero10ForTest("./sectorfinanciero10.xml", t)
	assert.Equal(t, "1.0", sectorFinanciero10.Version)
	assert.Equal(t, "111111", sectorFinanciero10.IdFideicomiso)
	assert.Equal(t, "NOMBRE FIDEICOMISO", *sectorFinanciero10.NombreFideicomiso)
	assert.Equal(t, "OBJETO/FIN FIDEICOMISO", sectorFinanciero10.DescripcionFideicomiso)
}
