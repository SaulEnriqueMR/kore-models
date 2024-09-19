package complementos

//xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xs="http://www.w3.org/2001/XMLSchema" xsi:schemaLocation="http://www.sat.gob.mx/cfd/3 http://www.sat.gob.mx/sitio_internet/cfd/3/cfdv32.xsd http://www.sat.gob.mx/divisas http://www.sat.gob.mx/sitio_internet/cfd/divisas/divisas.xsd"

import (
	"encoding/xml"
	divisas2 "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetDivisas10ForTesting(filename string, t *testing.T) (divisas2.Divisas10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed divisas2.Divisas10
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	return parsed, errUnmarshal
}

func TestDivisas10(t *testing.T) {
	divisas10, _ := GetDivisas10ForTesting("./divisas10.xml", t)
	InternalTestFullAtributesDivisas10(t, divisas10)
}

func InternalTestFullAtributesDivisas10(t *testing.T, divisas divisas2.Divisas10) {
	assert.Equal(t, "1.0", divisas.Version)
	assert.Equal(t, "compra", divisas.TipoOperacion)
}
