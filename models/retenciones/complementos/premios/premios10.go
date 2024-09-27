package premios

type Premios10 struct {
	Version              string  `xml:"Version,attr" bson:"version"`
	EntidadFederativa    string  `xml:"EntidadFederativa,attr" bson:"EntidadFederativa"`
	MontoTotalPago       float64 `xml:"MontTotPago,attr" bson:"MontoTotalPago"`
	MontTotalPagoGravado float64 `xml:"MontTotPagoGrav,attr" bson:"MontTotalPagoGravado"`
	MontTotalPagoExento  float64 `xml:"MontTotPagoExent,attr" bson:"MontTotalPagoExento"`
}
