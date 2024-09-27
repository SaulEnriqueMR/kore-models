package impuestoslocales

type ImpuestoLocales10 struct {
	Version          string                        `xml:"version,attr" bson:"Version"`
	TotalRetenciones float64                       `xml:"TotaldeRetenciones,attr" bson:"TotalRetenciones"`
	TotalTraslados   float64                       `xml:"TotaldeTraslados,attr" bson:"TotalTraslados"`
	Retenciones      *[]RetencionesLocalesImpLoc10 `xml:"RetencionesLocales" bson:"Retenciones,omitempty"`
	Traslados        *[]TrasladosLocalesImpLoc10   `xml:"TrasladosLocales" bson:"Traslados,omitempty"`
}

type RetencionesLocalesImpLoc10 struct {
	Impuesto string  `xml:"ImpLocRetenido,attr" bson:"Impuesto"`
	Tasa     float64 `xml:"TasadeRetencion,attr" bson:"Tasa"`
	Importe  float64 `xml:"Importe,attr" bson:"Importe"`
}

type TrasladosLocalesImpLoc10 struct {
	Impuesto string  `xml:"ImpLocTrasladado,attr" bson:"Impuesto"`
	Tasa     float64 `xml:"TasadeTraslado,attr" bson:"Tasa"`
	Importe  float64 `xml:"Importe,attr" bson:"Importe"`
}
