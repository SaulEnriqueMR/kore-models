package pagoaextranjeros

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/retenciones/complementos/pagosaextranjeros"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetPagoAExtranjeros10ForTest(filename string, t *testing.T) (pagosaextranjeros.PagosAExtranjeros10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed pagosaextranjeros.PagosAExtranjeros10
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	return parsed, errUnmashal
}

func TestFullPagoATerceros10(t *testing.T) {
	pagoExtranjeros, _ := GetPagoAExtranjeros10ForTest("./pagoaextranjeros10.xml", t)

	assert.Equal(t, "1.0", pagoExtranjeros.Version)
	assert.Equal(t, "NO", pagoExtranjeros.EsBenefEfectDelCobro)

	InternalTestNoBeneficiario10(t, pagoExtranjeros.NoBeneficiario)
	InternalTestBeneficiario10(t, pagoExtranjeros.Beneficiario)
}

func InternalTestNoBeneficiario10(t *testing.T, noBeneficiario10 *pagosaextranjeros.NoBeneficiarioPagosExtran10) {
	assert.NotNil(t, noBeneficiario10)

	assert.Equal(t, "US", noBeneficiario10.PaisDeResidParaEfecFisc)
	assert.Equal(t, "Artistas, deportistas y espectáculos públicos", noBeneficiario10.DescripcionConcepto)
	assert.Equal(t, "1", noBeneficiario10.ConceptoPago)
}

func InternalTestBeneficiario10(t *testing.T, beneficiario10 *pagosaextranjeros.BeneficiarioPagosExtran10) {
	assert.NotNil(t, beneficiario10)

	assert.Equal(t, "EKU9003173C9", beneficiario10.Rfc)
	assert.Equal(t, "EKUW9003173C9RPB01", beneficiario10.Curp)
	assert.Equal(t, "ESCUELA KEMPER URGATE", beneficiario10.Nombre)
	assert.Equal(t, "1", beneficiario10.ConceptoPago)
	assert.Equal(t, "Artistas, deportistas y espectáculos públicos", beneficiario10.DescripcionConcepto)
}
