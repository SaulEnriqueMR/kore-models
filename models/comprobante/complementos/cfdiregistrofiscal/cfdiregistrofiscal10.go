package cfdiregistrofiscal

type CfdiRegistroFiscal10 struct {
	Version string `xml:"Version,attr" bson:"Version" json:"Version"`
	Folio   string `xml:"Folio,attr" bson:"Folio" json:"Folio"`
}
