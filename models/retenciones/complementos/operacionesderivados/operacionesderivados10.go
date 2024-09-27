package operacionesderivados

type OperacionesDerivados10 struct {
	Version                 string  `xml:"Version,attr" bson:"Version"`
	MontoGananciaAcumulable float64 `xml:"MontGanAcum,attr" bson:"MontoGananciaAcumulable"`
	MontoPerdidaDeducible   float64 `xml:"MontPerdDed,attr" bson:"MontoPerdidaDeducible"`
}
