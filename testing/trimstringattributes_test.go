package testing

import (
	comprobante2 "github.com/SaulEnriqueMR/kore-models/models/comprobante"
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
			comprobante2.Concepto40{
				ClaveProdServ: "0000000   ",
			},
		},
	}
	comprobante2.TrimStringAttributes(&comprobante)

	expectedVersion := "1.0"
	if comprobante.Version != expectedVersion {
		t.Errorf("Wrong version. Expected: %s, got: %s", expectedVersion, comprobante.Version)
	}
	expectedSello := "SELLO"
	if comprobante.Sello != expectedSello {
		t.Errorf("Wrong sello. Expected: %s, got: %s", expectedSello, comprobante.Sello)
	}
	expectedNoCertificado := "NO CERTIFICADO"
	if comprobante.NoCertificado != expectedNoCertificado {
		t.Errorf("Wrong NoCertificado. Expected: %s, got: %s", expectedNoCertificado, comprobante.NoCertificado)
	}
	expectedCertificado := "CERTIFICADO"
	if comprobante.Certificado != expectedCertificado {
		t.Errorf("Wrong Certificado. Expected: %s, got: %s", expectedCertificado, comprobante.Certificado)
	}
	expectedMoneda := "MONEDA"
	if comprobante.Moneda != expectedMoneda {
		t.Errorf("Wrong Moneda. Expected: %s, got: %s", expectedMoneda, comprobante.Moneda)
	}
	expectedTipoComprobante := "TIPO COMPROBANTE"
	if comprobante.TipoComprobante != expectedTipoComprobante {
		t.Errorf("Wrong TipoComprobante. Expected: %s, got: %s", expectedTipoComprobante, comprobante.TipoComprobante)
	}
	expectedExportacion := "EXPORTACION"
	if comprobante.Exportacion != expectedExportacion {
		t.Errorf("Wrong Exportacion. Expected: %s, got: %s", expectedExportacion, comprobante.Exportacion)
	}

	expectedEmisor := comprobante2.Emisor40{
		Rfc:    "RFC",
		Nombre: "NOMBRE",
	}
	if comprobante.Emisor != expectedEmisor {
		t.Errorf("Wrong Emisor. Expected: %v, got: %v", expectedEmisor, comprobante.Emisor)
	}

	expectedConcepto := comprobante2.Concepto40{
		ClaveProdServ: "0000000",
	}
	concepto1 := comprobante.Conceptos[0]
	if expectedConcepto != concepto1 {
		t.Errorf("Wrong Concepto. Expected: %v, got: %v", expectedConcepto, concepto1)
	}
}
