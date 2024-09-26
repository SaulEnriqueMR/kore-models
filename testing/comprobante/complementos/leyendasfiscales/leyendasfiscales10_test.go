package complementoconcepto

import (
	"encoding/xml"
	"testing"

	leyendasfiscales210 "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/leyendasfiscales"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetLeyendasFiscales10ForTest(filename string, t *testing.T) (leyendasfiscales210.LeyendasFiscales10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed leyendasfiscales210.LeyendasFiscales10
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	return parsed, errUnmarshal
}

func TestLeyendasFiscales10(t *testing.T) {
	leyendasFiscales10, _ := GetLeyendasFiscales10ForTest("./leyendasfiscales10.xml", t)
	InternalTestFullAtributesLeyendasFis10(t, leyendasFiscales10)
	InternalTestFullAtributesLeyendaLeyendasFis10(t, leyendasFiscales10.Leyenda)
}

func InternalTestFullAtributesLeyendasFis10(t *testing.T, fis10 leyendasfiscales210.LeyendasFiscales10) {
	assert.Equal(t, "1.0", fis10.Version)
}

func InternalTestFullAtributesLeyendaLeyendasFis10(t *testing.T, fis10 []leyendasfiscales210.LeyendaLeyendasFis10) {
	assert.Equal(t, 2, len(fis10))
	assert.Equal(t, "ISR", *fis10[0].DisposicionFiscal)
	assert.Equal(t, "133", *fis10[0].Norma)
	assert.Equal(t, "Efectos fiscales al pago", fis10[0].TextoLeyenda)

	assert.Equal(t, "IVA", *fis10[1].DisposicionFiscal)
	assert.Equal(t, "32", *fis10[1].Norma)
	assert.Equal(t, "Impuesto retenido de conformidad con la Ley del Impuesto al Valor Agregado", fis10[1].TextoLeyenda)
}
