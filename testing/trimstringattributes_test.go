package testing

import (
	comprobante2 "github.com/SaulEnriqueMR/kore-models/app/comprobante"
	"github.com/SaulEnriqueMR/kore-models/app/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrimStringAttributes(t *testing.T) {
	comprobante := comprobante2.Comprobante40{
		Version:         " 1.0 ",
		Sello:           " SELLO",
		NoCertificado:   "NO CERTIFICADO ",
		Certificado:     "   CERTIFICADO",
		Moneda:          "MONEDA   ",
		TipoComprobante: "   TIPO COMPROBANTE   ",
		Exportacion:     "EXPORTACION",
		Emisor: comprobante2.Emisor40{
			Rfc:    " RFC ",
			Nombre: "NOMBRE",
		},
		Conceptos: []comprobante2.Concepto40{
			{
				ClaveProdServ: "0000000   ",
			},
		},
	}
	helpers.TrimStringAttributes(&comprobante)
	assert.Equal(t, "1.0", comprobante.Version)
	assert.Equal(t, "SELLO", comprobante.Sello)
	assert.Equal(t, "NO CERTIFICADO", comprobante.NoCertificado)
	assert.Equal(t, "CERTIFICADO", comprobante.Certificado)
	assert.Equal(t, "MONEDA", comprobante.Moneda)
	assert.Equal(t, "TIPO COMPROBANTE", comprobante.TipoComprobante)
	assert.Equal(t, "EXPORTACION", comprobante.Exportacion)
	expectedEmisor := comprobante2.Emisor40{
		Rfc:    "RFC",
		Nombre: "NOMBRE",
	}
	assert.Equal(t, expectedEmisor, comprobante.Emisor)
	expectedConceptos := []comprobante2.Concepto40{
		{
			ClaveProdServ: "0000000",
		},
	}
	assert.Equal(t, expectedConceptos, comprobante.Conceptos)
}

func TestTrimAttributesAndText(t *testing.T) {
	// Sample XML input
	inputXML := []byte(`
		<root>
			<element attribute="   some value   ">   some text   </element>
		</root>`)
	// Trim the XML attributes and text
	outputXML, err := helpers.TrimAttributesAndText(inputXML)
	assert.NoError(t, err)
	expectedXML := []byte(`<root><element attribute="some value">some text</element></root>`)
	assert.Equal(t, string(outputXML), string(expectedXML))
}
