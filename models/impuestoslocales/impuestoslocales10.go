package impuestoslocales

type ImpuestoLocales10 struct {
	Version            string                        `xml:"version,attr" bson:"Version"`
	TotalRetenciones   float64                       `xml:"TotaldeRetenciones,attr" bson:"TotalRetenciones"`
	TotalTraslados     float64                       `xml:"TotaldeTraslados,attr" bson:"TotalTraslados"`
	RetencionesLocales *[]RetencionesLocalesImpLoc10 `xml:"RetencionesLocales" bson:"RetencionesLocales,omitempty"`
	TrasladosLocales   *[]TrasladosLocalesImpLoc10   `xml:"TrasladosLocales" bson:"TrasladosLocales,omitempty"`
}

type RetencionesLocalesImpLoc10 struct {
	ImpLocRetenido string  `xml:"ImpLocRetenido,attr" bson:"ImpLocRetenido"`
	TasaRetencion  float64 `xml:"TasadeRetencion,attr" bson:"TasaRetencion"`
	Importe        float64 `xml:"Importe,attr" bson:"Importe"`
}

type TrasladosLocalesImpLoc10 struct {
	ImpLocTrasladado string  `xml:"ImpLocTrasladado,attr" bson:"ImpLocTrasladado"`
	TasaTraslado     float64 `xml:"TasadeTraslado,attr" bson:"TasaTraslado"`
	Importe          float64 `xml:"Importe,attr" bson:"Importe"`
}
