package addenda

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/addenda"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetRequestForPaymentForTest(filename string, t *testing.T) (addenda.RequestForPayment, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed addenda.RequestForPayment
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	testing2.GenerateJSONFromStructure("requestforpayment.json", parsed)
	return parsed, errUnmashal
}

func TestRequestForPayment(t *testing.T) {
	_, _ = GetRequestForPaymentForTest("./requestforpayment.xml", t)
}
