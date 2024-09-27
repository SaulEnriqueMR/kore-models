package dividendos

import (
	"encoding/xml"
	"testing"

	dividendos1 "github.com/SaulEnriqueMR/kore-models/models/retenciones/complementos/dividendos"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetTurisDividendosForTest(filename string, t *testing.T) (dividendos1.Dividendos10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed dividendos1.Dividendos10
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	return parsed, errUnmashal
}

func TestFullDividendos(t *testing.T) {
	dividendos, _ := GetTurisDividendosForTest("./dividendos10.xml", t)
	testing2.GenerateJSONFromStructure("dividendos10.json", dividendos)
	assert.NotNil(t, dividendos)
	InternalTestDividendos(t, dividendos)
	InternalTestDividOUtil(t, dividendos.DividendoOUtilidad)
	InternalTestRemanente(t, dividendos.Remanente)
}

func InternalTestDividendos(t *testing.T, dividendos dividendos1.Dividendos10) {
	assert.Equal(t, "1.0", dividendos.Version)
}

func InternalTestDividOUtil(t *testing.T, dividOUtil *dividendos1.DividOUtilDividendos10) {
	assert.Equal(t, "DIV123", dividOUtil.ClaveTipoDivendoOUtilidad)
	assert.Equal(t, 1500.50, dividOUtil.MontoIsrRetenidoMexico)
	assert.Equal(t, 2000.75, dividOUtil.MontoIsrRetenidoExtranjero)
	assert.Equal(t, 500.25, *dividOUtil.MontoRetencionDividendoExtranjero)
	assert.Equal(t, "Sociedad Nacional", dividOUtil.TipoSociedadDistribucionDividendo)
	assert.Equal(t, 300.00, *dividOUtil.MontoIsrAcreditableNacional)
	assert.Equal(t, 1200.00, *dividOUtil.MontoDividendoAcumulableNacional)
	assert.Equal(t, 800.00, *dividOUtil.MontoDividendoAcumulableExtranjero)
}

func InternalTestRemanente(t *testing.T, remanente *dividendos1.RemanenteDividendos10) {
	assert.Equal(t, 10.503456, *remanente.ProporcionRemanente)
}
