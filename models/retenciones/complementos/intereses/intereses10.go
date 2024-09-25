package intereses

type Intereses10 struct {
	Version           string  `xml:"Version,attr" bson:"Version"`
	SistFinanciero    string  `xml:"SistFinanciero,attr" bson:"SistFinanciero"`
	RetiroAORESRetInt string  `xml:"RetiroAORESRetInt,attr" bson:"RetiroAORESRetInt"`
	OperFinancDerivad string  `xml:"OperFinancDerivad,attr" bson:"OperFinancDerivad"`
	MontIntNominal    float64 `xml:"MontIntNominal,attr" bson:"MontIntNominal"`
	MontIntReal       float64 `xml:"MontIntReal,attr" bson:"MontIntReal"`
	Perdida           float64 `xml:"Perdida,attr" bson:"Perdida"`
}
