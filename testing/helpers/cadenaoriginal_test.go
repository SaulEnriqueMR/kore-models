package helpers

import (
	"encoding/xml"
	comprobante2 "github.com/SaulEnriqueMR/kore-models/models/comprobante"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetComprobanteForTest(filename string, t *testing.T) (comprobante2.Comprobante40, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed comprobante2.Comprobante40
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	return parsed, errUnmashal
}

func TestCreateCadenaOriginal(t *testing.T) {
	comprobante, _ := GetComprobanteForTest("./comprobantecadenaoriginal.xml", t)
	result := helpers.CreateCadenaOriginal(comprobante)
	assert.Equal(t, "||4.0|||2024-04-29T00:00:55|||200||MXN||199.96|I|01||20000||||||EKU9003173C9|ESCUELA KEMPER URGATE|601||||||URE180429TM6|UNIVERSIDAD ROBOTICA ESPAÑOLA|86991|||601|G01|||SPR190613I52|||50211503||1|H87||Cigarros|200|200||02||||||||||||||||||||||||||||||||||||||||||||||1.1|3ea43e97-71bf-4ac5-a28e-9374ea9f8b45|2024-04-29T10:46:30||||||||SPR190613I52||||||||||||||||||", result)
}
