package retenciones

import (
	"encoding/xml"
	sectfinan2 "github.com/SaulEnriqueMR/kore-models/models/comprobante/retenciones"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetSectorFinanciero10ForTest(filename string, t *testing.T) (sectfinan2.SectorFinanciero10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed sectfinan2.SectorFinanciero10
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	return parsed, errUnmarshal
}

func TestSectorFinanciero10(t *testing.T) {
	sectorFinanciero10, _ := GetSectorFinanciero10ForTest("./sectorfinanciero10.xml", t)
	assert.Equal(t, "1.0", sectorFinanciero10.Version)
	assert.Equal(t, "111111", sectorFinanciero10.IdFideicom)
	assert.Equal(t, "NOMBRE FIDEICOMISO", *sectorFinanciero10.NomFideicom)
	assert.Equal(t, "OBJETO/FIN FIDEICOMISO", sectorFinanciero10.DescripFideicom)
}
