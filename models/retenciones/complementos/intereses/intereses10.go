package intereses

import (
	"encoding/xml"

	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type Intereses10 struct {
	Version string `xml:"Version,attr" bson:"Version" json:"Version"`

	SistemaFinanciero          string `xml:"SistFinanciero,attr" bson:"SistemaFinanciero" json:"SistemaFinanciero"`
	ProvienenSistemaFinanciero bool   `bson:"ProvienenSistemaFinanciero" json:"ProvienenSistemaFinanciero"`

	RetiroIntereses     string `xml:"RetiroAORESRetInt,attr" bson:"RetiroIntereses" json:"RetiroIntereses"`
	HuboRetiroIntereses bool   `bson:"HuboRetiroIntereses" json:"HuboRetiroIntereses"`

	OperacionFinancieraDerivada            string `xml:"OperFinancDerivad,attr" bson:"OperacionFinancieraDerivada" json:"OperacionFinancieraDerivada"`
	CorrespondeOperacionFinancieraDerivada bool   `bson:"CorrespondeOperacionFinancieraDerivada" json:"CorrespondeOperacionFinancieraDerivada"`

	MontoInteresNominal float64 `xml:"MontIntNominal,attr" bson:"MontoInteresNominal" json:"MontoInteresNominal"`
	MontoInteresReal    float64 `xml:"MontIntReal,attr" bson:"MontoInteresReal" json:"MontoInteresReal"`
	Perdida             float64 `xml:"Perdida,attr" bson:"Perdida" json:"Perdida"`
}

func (i *Intereses10) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias Intereses10
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	*i = Intereses10(aux)
	i.ProvienenSistemaFinanciero = helpers.ResolveSatBoolean(aux.SistemaFinanciero)
	i.HuboRetiroIntereses = helpers.ResolveSatBoolean(aux.RetiroIntereses)
	i.CorrespondeOperacionFinancieraDerivada = helpers.ResolveSatBoolean(aux.OperacionFinancieraDerivada)

	return nil
}
