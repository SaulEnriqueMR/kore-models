package hidrocarburospetroliferos

import (
	"encoding/xml"
	"testing"

	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementosconcepto/hidrocarburospetroliferos"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetHidrocarburosPetroliferosForTest(filename string, t *testing.T) (hidrocarburospetroliferos.HidrocarburosPetroliferos10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed hidrocarburospetroliferos.HidrocarburosPetroliferos10
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	testing2.GenerateJSONFromStructure("hidrocarburospetroliferos10.json", parsed)
	return parsed, errUnmashal
}

func TestFullHidrocarburosPetroliferos(t *testing.T) {
	hyp, _ := GetHidrocarburosPetroliferosForTest("./hidrocarburospetroliferos10.xml", t)
	assert.Equal(t, "1.0", hyp.Version)
	assert.Equal(t, "PER01", hyp.TipoPermiso)
	assert.Equal(t, "NO PERMISO", hyp.NoPermiso)
	assert.Equal(t, "15101505", hyp.Clave)
	assert.Equal(t, "SP16", hyp.Subproducto)
}
