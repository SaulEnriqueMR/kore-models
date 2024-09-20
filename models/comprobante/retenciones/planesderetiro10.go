package retenciones

type PlanesDeRetiro10 struct {
	Version                        string   `xml:"Version,attr" bson:"version"`
	SistemaFinanc                  string   `xml:"SistemaFinanc,attr" bson:"SistemaFinanc"`
	MontTotAportAnioInmAnterior    *float64 `xml:"MontTotAportAnioInmAnterior,attr" bson:"MontTotAportAnioInmAnterior,omitempty"`
	MontIntRealesDevengAniooInmAnt float64  `xml:"MontIntRealesDevengAniooInmAnt,attr" bson:"MontIntRealesDevengAniooInmAnt"`
	HuboRetirosAnioInmAntPer       string   `xml:"HuboRetirosAnioInmAntPer,attr" bson:"HuboRetirosAnioInmAntPer"`
	MontTotRetiradoAnioInmAntPer   *float64 `xml:"MontTotRetiradoAnioInmAntPer,attr" bson:"MontTotRetiradoAnioInmAntPer,omitempty"`
	MontTotExentRetiradoAnioInmAnt *float64 `xml:"MontTotExentRetiradoAnioInmAnt,attr" bson:"MontTotExentRetiradoAnioInmAnt,omitempty"`
	MontTotExedenteAnioInmAnt      *float64 `xml:"MontTotExedenteAnioInmAnt,attr" bson:"MontTotExedenteAnioInmAnt,omitempty"`
	HuboRetirosAnioInmAnt          string   `xml:"HuboRetirosAnioInmAnt,attr" bson:"HuboRetirosAnioInmAnt"`
	MontTotRetiradoAnioInmAnt      *float64 `xml:"MontTotRetiradoAnioInmAnt,attr" bson:"MontTotRetiradoAnioInmAnt,omitempty"`
}
