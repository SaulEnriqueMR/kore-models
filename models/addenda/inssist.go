package addenda

type Inssist struct {
	Referencia                       *string `xml:"Referencia,attr" json:"Referencia,omitempty" bson:"Referencia,omitempty"`
	Reservacion                      *string `xml:"Reservacion,attr" json:"Reservacion,omitempty" bson:"Reservacion,omitempty"`
	OtrosCargos                      *string `xml:"OtrosCargos,attr" json:"OtrosCargos,omitempty" bson:"OtrosCargos,omitempty"`
	ImportePagar                     *string `xml:"importeaPagar,attr" json:"ImportePagar,omitempty" bson:"ImportePagar,omitempty"`
	ImporteRetenidoCuentaTerceros    *string `xml:"Terceros,attr" json:"Terceros,omitempty" bson:"Terceros,omitempty"`
	ImporteRetenidoaCuentadeTerceros *string `xml:"ImporteRetenidoaCuentadeTerceros,attr" json:"ImporteRetenidoaCuentadeTerceros" bson:"ImporteRetenidoaCuentadeTerceros"`
}
