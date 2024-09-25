package premios

type Premios10 struct {
	Version           string  `xml:"Version,attr" bson:"version"`
	EntidadFederativa string  `xml:"EntidadFederativa,attr" bson:"EntidadFederativa"`
	MontTotPago       float64 `xml:"MontTotPago,attr" bson:"MontTotPago"`
	MontTotPagoGrav   float64 `xml:"MontTotPagoGrav,attr" bson:"MontTotPagoGrav"`
	MontTotPagoExent  float64 `xml:"MontTotPagoExent,attr" bson:"MontTotPagoExent"`
}
