package premios

type Premios10 struct {
	Version              string  `xml:"Version,attr" bson:"version" json:"version"`
	EntidadFederativa    string  `xml:"EntidadFederativa,attr" bson:"EntidadFederativa" json:"EntidadFederativa"`
	MontoTotalPago       float64 `xml:"MontTotPago,attr" bson:"MontoTotalPago" json:"MontoTotalPago"`
	MontTotalPagoGravado float64 `xml:"MontTotPagoGrav,attr" bson:"MontTotalPagoGravado" json:"MontTotalPagoGravado"`
	MontTotalPagoExento  float64 `xml:"MontTotPagoExent,attr" bson:"MontTotalPagoExento" json:"MontTotalPagoExento"`
}
