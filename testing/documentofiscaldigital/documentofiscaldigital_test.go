package documentofiscaldigital

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/SaulEnriqueMR/kore-models/models/comprobante"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func TestDocumentoFiscalDigitalComprobante40(t *testing.T) {
	data := testing2.GetFileContentForTest("./df.xml", t)
	var parsed comprobante.Comprobante
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	assert.NotNil(t, parsed.Comprobante40)

	comprobante40 := parsed.Comprobante40
	assert.NotNil(t, comprobante40.Uuid)

	assert.Equal(t, strings.ToUpper("3ea43e97-71bf-4ac5-a28e-9374ea9f8b45"), comprobante40.Uuid)
}
