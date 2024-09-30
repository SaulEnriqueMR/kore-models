package planesretiro

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type PlanesDeRetiro10 struct {
	Version string `xml:"Version,attr" bson:"version"`

	SistemaFinanciero          string `xml:"SistemaFinanc,attr" bson:"SistemaFinanciero"`
	ProvienenSistemaFinanciero bool   `bson:"ProvienenSistemaFinanciero"`

	MontoTotalAportacionesAnioAnterior         *float64 `xml:"MontTotAportAnioInmAnterior,attr" bson:"MontoTotalAportacionesAnioAnterior,omitempty"`
	MontoInteresesRealesDevengadosAnioAnterior float64  `xml:"MontIntRealesDevengAniooInmAnt,attr" bson:"MontoInteresesRealesDevengadosAnioAnterior"`

	HuboRetirosAnioInmAntPer                string `xml:"HuboRetirosAnioInmAntPer,attr" bson:"HuboRetirosAnioInmAntPer"`
	HuboRetirosAnioAnteriorAntesPermanencia bool   `bson:"HuboRetirosAnioAnteriorAntesPermanencia"`

	MontoTotalRetiradoAnioAnteriorPermanencia *float64 `xml:"MontTotRetiradoAnioInmAntPer,attr" bson:"MontoTotalRetiradoAnioAnteriorPermanencia,omitempty"`
	MontoTotalExentoRetiradoAnioAnterior      *float64 `xml:"MontTotExentRetiradoAnioInmAnt,attr" bson:"MontoTotalExentoRetiradoAnioAnterior,omitempty"`
	MontoTotalExedenteAnioAnterior            *float64 `xml:"MontTotExedenteAnioInmAnt,attr" bson:"MontoTotalExedenteAnioAnterior,omitempty"`

	HuboRetirosAnioInmAnt   string `xml:"HuboRetirosAnioInmAnt,attr" bson:"HuboRetirosAnioInmAnt"`
	HuboRetirosAnioAnterior bool   `bson:"HuboRetirosAnioAnterior"`

	MontoTotalRetiradoAnioAnterior *float64 `xml:"MontTotRetiradoAnioInmAnt,attr" bson:"MontoTotalRetiradoAnioAnterior,omitempty"`
}

func (pr *PlanesDeRetiro10) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias PlanesDeRetiro10
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	*pr = PlanesDeRetiro10(aux)
	pr.ProvienenSistemaFinanciero = helpers.ResolveSatBoolean(aux.SistemaFinanciero)
	pr.HuboRetirosAnioAnteriorAntesPermanencia = helpers.ResolveSatBoolean(aux.HuboRetirosAnioInmAntPer)
	pr.HuboRetirosAnioAnterior = helpers.ResolveSatBoolean(aux.HuboRetirosAnioInmAnt)

	return nil
}
