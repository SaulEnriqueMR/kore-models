package intereseshipotecarios

import (
	"encoding/xml"

	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type InteresesHipotecarios10 struct {
	Version string `xml:"Version,attr" bson:"Version" json:"Version"`

	CreditoOtorgadoInstitucionFinanciera   string `xml:"CreditoDeInstFinanc,attr" bson:"CreditoOtorgadoInstitucionFinanciera" json:"CreditoOtorgadoInstitucionFinanciera"`
	EsCreditoOtorgadoInstitucionFinanciera bool   `bson:"EsCreditoOtorgadoInstitucionFinanciera" json:"EsCreditoOtorgadoInstitucionFinanciera"`

	SaldoInsoluto                                  float64  `xml:"SaldoInsoluto,attr" bson:"SaldoInsoluto" json:"SaldoInsoluto"`
	ProporcionDeducibleCredito                     *float64 `xml:"PropDeducDelCredit,attr" bson:"ProporcionDeducibleCredito,omitempty" json:"ProporcionDeducibleCredito,omitempty"`
	MontoTotalInteresesNominalesDevengados         *float64 `xml:"MontTotIntNominalesDev,attr" bson:"MontoTotalInteresesNominalesDevengados,omitempty" json:"MontoTotalInteresesNominalesDevengados,omitempty"`
	MontoTotalInteresesNominalesDevengadosYPagados *float64 `xml:"MontTotIntNominalesDevYPag,attr" bson:"MontoTotalInteresesNominalesDevengadosYPagados,omitempty" json:"MontoTotalInteresesNominalesDevengadosYPagados,omitempty"`
	MontoTotalInteresesRealesPagsDeducibles        *float64 `xml:"MontTotIntRealPagDeduc,attr" bson:"MontoTotalInteresesRealesPagsDeducibles,omitempty" json:"MontoTotalInteresesRealesPagsDeducibles,omitempty"`
	NoContrato                                     *string  `xml:"NumContrato,attr" bson:"NoContrato,omitempty" json:"NoContrato,omitempty"`
}

func (ih *InteresesHipotecarios10) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias InteresesHipotecarios10
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	*ih = InteresesHipotecarios10(aux)
	ih.EsCreditoOtorgadoInstitucionFinanciera = helpers.ResolveSatBoolean(ih.CreditoOtorgadoInstitucionFinanciera)

	return nil
}
