package operacionesderivados

type OperacionesDerivados10 struct {
	Version     string  `xml:"Version,attr" bson:"Version"`
	MontGanAcum float64 `xml:"MontGanAcum,attr" bson:"MontGanAcum"`
	MontPerdDed float64 `xml:"MontPerdDed,attr" bson:"MontPerdDed"`
}
