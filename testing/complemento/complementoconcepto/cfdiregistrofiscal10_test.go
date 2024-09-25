package complementoconcepto

import (
	"encoding/xml"
	"testing"

	registrofiscal "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/cfdiregistrofiscal"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetCfdiRegistroFiscalForTest(filename string, t *testing.T) (registrofiscal.CfdiRegistroFiscal10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed registrofiscal.CfdiRegistroFiscal10
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	return parsed, errUnmarshal
}

func TestCfdiRegistroFiscal10(t *testing.T) {
	cfdiRegistroFiscal, _ := GetCfdiRegistroFiscalForTest("./cfdiregistrofiscal10.xml", t)
	InternalTestFullAtributesCfdiRegistroFiscal10(t, cfdiRegistroFiscal)
}

func InternalTestFullAtributesCfdiRegistroFiscal10(t *testing.T, registroFiscal registrofiscal.CfdiRegistroFiscal10) {
	assert.Equal(t, "1.0", registroFiscal.Version)
	assert.Equal(t, "8123831283128383", registroFiscal.Folio)
}
