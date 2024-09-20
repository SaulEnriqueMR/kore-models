package retenciones

type OperacionesConDerivados10 struct {
	Version     string  `xml:"Version,attr" bson:"version"`
	MontGanAcum float64 `xml:"MontGanAcum,attr" bson:"MontGanAcum"`
	MontPerdDed float64 `xml:"MontPerdDed,attr" bson:"MontPerdDed"`
}
