package impuestoslocales

type ImpuestoLocales10 struct {
	Version          string                        `xml:"version,attr" bson:"Version" json:"Version"`
	TotalRetenciones float64                       `xml:"TotaldeRetenciones,attr" bson:"TotalRetenciones" json:"TotalRetenciones"`
	TotalTraslados   float64                       `xml:"TotaldeTraslados,attr" bson:"TotalTraslados" json:"TotalTraslados"`
	Retenciones      *[]RetencionesLocalesImpLoc10 `xml:"RetencionesLocales" bson:"Retenciones,omitempty" json:"Retenciones,omitempty"`
	Traslados        *[]TrasladosLocalesImpLoc10   `xml:"TrasladosLocales" bson:"Traslados,omitempty" json:"Traslados,omitempty"`
}

type RetencionesLocalesImpLoc10 struct {
	Impuesto string  `xml:"ImpLocRetenido,attr" bson:"Impuesto" json:"Impuesto"`
	Tasa     float64 `xml:"TasadeRetencion,attr" bson:"Tasa" json:"Tasa"`
	Importe  float64 `xml:"Importe,attr" bson:"Importe" json:"Importe"`
}

type TrasladosLocalesImpLoc10 struct {
	Impuesto string  `xml:"ImpLocTrasladado,attr" bson:"Impuesto" json:"Impuesto"`
	Tasa     float64 `xml:"TasadeTraslado,attr" bson:"Tasa" json:"Tasa"`
	Importe  float64 `xml:"Importe,attr" bson:"Importe" json:"Importe"`
}
