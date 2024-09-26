package complementoconcepto

import (
	"encoding/xml"
	"testing"

	pagoespecie2 "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/pagoespecie"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetPagoEnEspecie10ForTest(filename string, t *testing.T) (pagoespecie2.PagoEspecie10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed pagoespecie2.PagoEspecie10
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	testing2.GenerateJSONFromStructure("pagoenespecie10.json", parsed)
	return parsed, errUnmarshal
}

func TestPagoEnEspecie10(t *testing.T) {
	pagoEnEspecie, _ := GetPagoEnEspecie10ForTest("./pagoenespecie10.xml", t)
	assert.Equal(t, "1.0", pagoEnEspecie.Version)
	assert.Equal(t, "Aamp;C8317286A1-18000101-020", pagoEnEspecie.CvePIC)
	assert.Equal(t, "PE-12-45554", pagoEnEspecie.FolioSolDon)
	assert.Equal(t, "PINTURA ARBOL ROJO", pagoEnEspecie.PzaArtNombre)
	assert.Equal(t, "PINTURA AL OLEO", pagoEnEspecie.PzaArtTecn)
	assert.Equal(t, "2000", pagoEnEspecie.PzaArtAProd)
	assert.Equal(t, "2 x 4 mts", pagoEnEspecie.PzaArtDim)
}
