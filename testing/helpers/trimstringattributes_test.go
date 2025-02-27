package helpers

import (
	comprobante2 "github.com/SaulEnriqueMR/kore-models/models/comprobante"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
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
		<cfdi:Comprobante 
			xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
			xsi:schemaLocation="http://www.sat.gob.mx/cfd/4 
			http://www.sat.gob.mx/sitio_internet/cfd/4/cfdv40.xsd" 
			Total="100" 
			Subtotal=" 50       ">
		</cfdi:Comprobante>`)
	// Trim the XML attributes and text
	outputXML, err := helpers.TrimXml(inputXML)
	assert.NoError(t, err)
	expectedXML := `<cfdi:Comprobante xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.sat.gob.mx/cfd/4 http://www.sat.gob.mx/sitio_internet/cfd/4/cfdv40.xsd" Total="100" Subtotal="50"></cfdi:Comprobante>`
	assert.Equal(t, expectedXML, string(outputXML))
}
