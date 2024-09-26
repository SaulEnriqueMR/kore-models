package pfintegrantecoordinado

import (
	"encoding/xml"
	"testing"

	pfintegrante1 "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/pfintegrantecoordinado"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetPFIntegranteForTest(filename string, t *testing.T) (pfintegrante1.PFIntegranteCoordinado10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed pfintegrante1.PFIntegranteCoordinado10
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	testing2.GenerateJSONFromStructure("pfintegrante10.json", parsed)
	return parsed, errUnmashal
}

func TestFullPFIntegrante(t *testing.T) {
	pfintegrante, _ := GetPFIntegranteForTest("./pfintegrante.xml", t)
	assert.NotNil(t, pfintegrante)
	InternalTestBasePFIntegrante(t, pfintegrante)
}

func InternalTestBasePFIntegrante(t *testing.T, pfIntegrante pfintegrante1.PFIntegranteCoordinado10) {
	assert.Equal(t, "1.0", pfIntegrante.Version)
	assert.Equal(t, "SEDAN2021", pfIntegrante.ClaveVehicular)
	assert.Equal(t, "ABC1234", pfIntegrante.Placa)
	assert.Equal(t, "XAXX010101000", pfIntegrante.RfcPF)
}
