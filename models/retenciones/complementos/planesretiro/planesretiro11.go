package planesretiro

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type PlanesDeRetiro11 struct {
	Version string `xml:"Version,attr" bson:"version" json:"version"`

	SistemaFinanciero          string `xml:"SistemaFinanc,attr" bson:"SistemaFinanciero" json:"SistemaFinanciero"`
	ProvienenSistemaFinanciero bool   `bson:"ProvienenSistemaFinanciero" json:"ProvienenSistemaFinanciero"`

	MontoTotalAportacionesAnioAnterior         *float64 `xml:"MontTotAportAnioInmAnterior,attr" bson:"MontoTotalAportacionesAnioAnterior,omitempty" json:"MontoTotalAportacionesAnioAnterior,omitempty"`
	MontoInteresesRealesDevengadosAnioAnterior float64  `xml:"MontIntRealesDevengAniooInmAnt,attr" bson:"MontoInteresesRealesDevengadosAnioAnterior" json:"MontoInteresesRealesDevengadosAnioAnterior"`

	HuboRetirosAnioInmAntPer                string `xml:"HuboRetirosAnioInmAntPer,attr" bson:"HuboRetirosAnioInmAntPer" json:"HuboRetirosAnioInmAntPer"`
	HuboRetirosAnioAnteriorAntesPermanencia bool   `bson:"HuboRetirosAnioAnteriorAntesPermanencia" json:"HuboRetirosAnioAnteriorAntesPermanencia"`

	MontoTotalRetiradoAnioAnteriorPermanencia *float64 `xml:"MontTotRetiradoAnioInmAntPer,attr" bson:"MontTotRetiradoAnioInmAntPer,omitempty" json:"MontTotRetiradoAnioInmAntPer,omitempty"`
	MontoTotalExentoRetiradoAnioAnterior      *float64 `xml:"MontTotExentRetiradoAnioInmAnt,attr" bson:"MontoTotalExentoRetiradoAnioAnterior,omitempty" json:"MontoTotalExentoRetiradoAnioAnterior,omitempty"`
	MontoTotalExedenteAnioAnterior            *float64 `xml:"MontTotExedenteAnioInmAnt,attr" bson:"MontoTotalExedenteAnioAnterior,omitempty" json:"MontoTotalExedenteAnioAnterior,omitempty"`

	HuboRetirosAnioInmAnt   string `xml:"HuboRetirosAnioInmAnt,attr" bson:"HuboRetirosAnioInmAnt" json:"HuboRetirosAnioInmAnt"`
	HuboRetirosAnioAnterior bool   `bson:"HuboRetirosAnioAnterior" json:"HuboRetirosAnioAnterior"`

	MontoTotalRetiradoAnioAnterior *float64                          `xml:"MontTotRetiradoAnioInmAnt,attr" bson:"MontoTotalRetiradoAnioAnterior,omitempty" json:"MontoTotalRetiradoAnioAnterior,omitempty"`
	NoReferencia                   *string                           `xml:"NumReferencia,attr" bson:"NoReferencia,omitempty" json:"NoReferencia,omitempty"`
	AportacionesODepositos         *[]AportacionesODepositosPlanes11 `xml:"AportacionesODepositos" bson:"AportacionesODepositos" json:"AportacionesODepositos"`
}

func (pr *PlanesDeRetiro11) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias PlanesDeRetiro11
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	*pr = PlanesDeRetiro11(aux)
	pr.ProvienenSistemaFinanciero = helpers.ResolveSatBoolean(aux.SistemaFinanciero)
	pr.HuboRetirosAnioAnteriorAntesPermanencia = helpers.ResolveSatBoolean(aux.HuboRetirosAnioInmAntPer)
	pr.HuboRetirosAnioAnterior = helpers.ResolveSatBoolean(aux.HuboRetirosAnioInmAnt)

	return nil
}

type AportacionesODepositosPlanes11 struct {
	Tipo          string  `xml:"TipoAportacionODeposito,attr" bson:"Tipo" json:"Tipo"`
	Monto         float64 `xml:"MontAportODep,attr" bson:"Monto" json:"Monto"`
	RfcFiduciaria *string `xml:"RFCFiduciaria,attr" bson:"RfcFiduciaria,omitempty" json:"RfcFiduciaria,omitempty"`
}
