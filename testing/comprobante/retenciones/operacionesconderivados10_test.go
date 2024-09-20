package retenciones

import (
	"encoding/xml"
	opderivados2 "github.com/SaulEnriqueMR/kore-models/models/comprobante/retenciones"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetOperacioneConDerivados10ForTest(filename string, t *testing.T) (opderivados2.OperacionesConDerivados10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed opderivados2.OperacionesConDerivados10
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	return parsed, errUnmarshal
}

func TestOperacionesConderivados10(t *testing.T) {
	operacionesConDerivados, _ := GetOperacioneConDerivados10ForTest("./operacionesconderivados10.xml", t)
	assert.Equal(t, "1.0", operacionesConDerivados.Version)
	assert.Equal(t, 1.00, operacionesConDerivados.MontGanAcum)
	assert.Equal(t, 1.00, operacionesConDerivados.MontPerdDed)
}