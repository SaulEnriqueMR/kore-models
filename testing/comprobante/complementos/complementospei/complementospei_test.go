package complementos

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"os"
	"testing"

	compspei2 "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/complementospei"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetComplementoSpeiForTest(filename string, t *testing.T) (compspei2.ComplementoSpei, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed compspei2.ComplementoSpei
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	GenerateJSONFromXMLComplementoSpei("complementospei.json", parsed)
	return parsed, errUnmarshal
}

func TestComplementoSpei(t *testing.T) {
	complementoSpei, _ := GetComplementoSpeiForTest("./complementospei.xml", t)
	speiTercero := complementoSpei.SpeiTercero
	InternalTestFullAtributesCompSpei(t, speiTercero[0])
	InternalTestFullAtributesOrdenanteSPEI(t, speiTercero[0].Ordenante[0])
	InternalTestFullAtributesBeneficiario(t, speiTercero[0].Beneficiario[0])
}

func InternalTestFullAtributesCompSpei(t *testing.T, speiTercero compspei2.SpeiTerceroCompSpei) {
	assert.Equal(t, "2016-04-07", speiTercero.FechaOperacion)
	assert.Equal(t, "07:20:17.635", speiTercero.Hora)
	assert.Equal(t, "23837", speiTercero.ClaveSPEI)
	sello := "v7Uq8JveMHDigYBN1mgeUqtAeuCRnTIDE1IpvEQWCCwttZkEOA7OZe+JzdgkLGMlhCzlFMTp3aOZkhhK7WPZXHeov1XlxevfctNqGl1sQWOyZjqqEmngliZppl4dr"
	assert.Equal(t, sello, speiTercero.Sello)
	assert.Equal(t, "20001000000200000876", speiTercero.NumeroCertificado)
	assert.Equal(t, "cadena", speiTercero.CadenaCDA)
}

func InternalTestFullAtributesOrdenanteSPEI(t *testing.T, ordenante compspei2.OrdenanteCompSpei) {
	assert.Equal(t, "Banco Emisor S.A.", ordenante.BancoEmisor)
	assert.Equal(t, "Francisco Bautista", ordenante.Nombre)
	assert.Equal(t, "43328545620125000237", ordenante.Cuenta)
	assert.Equal(t, "AAA010101AAA", ordenante.RFC)
	assert.Equal(t, "01", ordenante.TipoCuenta)
}

func InternalTestFullAtributesBeneficiario(t *testing.T, beneficiario compspei2.BeneficiarioCompSpei) {
	assert.Equal(t, "Banco Receptor S.A.", beneficiario.BancoReceptor)
	assert.Equal(t, "Juan Perez Santiago", beneficiario.Nombre)
	assert.Equal(t, "98363842374909422123", beneficiario.Cuenta)
	assert.Equal(t, "03", beneficiario.TipoCuenta)
	assert.Equal(t, "AAA010101AAA", beneficiario.RFC)
	assert.Equal(t, "Honorarios", beneficiario.Concepto)
	assert.Equal(t, "0.00", *beneficiario.IVA)
	assert.Equal(t, "10000.00", beneficiario.MontoPago)
}

func GenerateJSONFromXMLComplementoSpei(namefile string, data compspei2.ComplementoSpei) {
	jsonData, err := json.MarshalIndent(data, "", "	")
	err = os.WriteFile(namefile, jsonData, 0644)
	if err != nil {
		log.Println(err)
	}
}
