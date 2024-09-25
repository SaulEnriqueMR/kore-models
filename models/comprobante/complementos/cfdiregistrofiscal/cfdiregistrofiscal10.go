package cfdiregistrofiscal

type CfdiRegistroFiscal10 struct {
	Version string `xml:"Version,attr" bson:"Version"`
	Folio   string `xml:"Folio,attr" bson:"Folio"`
}
