package pagosaextranjeros

import (
	"encoding/xml"

	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type PagosAExtranjeros10 struct {
	Version                     string                       `xml:"Version,attr" bson:"Version" json:"Version"`
	EsBenefEfectDelCobro        string                       `xml:"EsBenefEfectDelCobro,attr" bson:"EsBenefEfectDelCobro" json:"EsBenefEfectDelCobro"`
	EsBeneficiarioEfectivoCobro bool                         `bson:"EsBeneficiarioEfectivoCobro" json:"EsBeneficiarioEfectivoCobro"`
	NoBeneficiario              *NoBeneficiarioPagosExtran10 `xml:"NoBeneficiario" bson:"NoBeneficiario,omitempty" json:"NoBeneficiario,omitempty"`
	Beneficiario                *BeneficiarioPagosExtran10   `xml:"Beneficiario" bson:"Beneficiario,omitempty" json:"Beneficiario,omitempty"`
}

func (pae *PagosAExtranjeros10) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias PagosAExtranjeros10
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	*pae = PagosAExtranjeros10(aux)
	pae.EsBeneficiarioEfectivoCobro = helpers.ResolveSatBoolean(pae.EsBenefEfectDelCobro)

	return nil
}

type NoBeneficiarioPagosExtran10 struct {
	ResidenciaFiscal    string `xml:"PaisDeResidParaEfecFisc,attr" bson:"ResidenciaFiscal" json:"ResidenciaFiscal"`
	ConceptoPago        string `xml:"ConceptoPago,attr" bson:"ConceptoPago" json:"ConceptoPago"`
	DescripcionConcepto string `xml:"DescripcionConcepto,attr" bson:"DescripcionConcepto" json:"DescripcionConcepto"`
}

type BeneficiarioPagosExtran10 struct {
	Rfc                 string `xml:"RFC,attr" bson:"Rfc" json:"Rfc"`
	Curp                string `xml:"CURP,attr" bson:"Curp" json:"Curp"`
	Nombre              string `xml:"NomDenRazSocB,attr" bson:"Nombre" json:"Nombre"`
	ConceptoPago        string `xml:"ConceptoPago,attr" bson:"ConceptoPago" json:"ConceptoPago"`
	DescripcionConcepto string `xml:"DescripcionConcepto,attr" bson:"DescripcionConcepto" json:"DescripcionConcepto"`
}
