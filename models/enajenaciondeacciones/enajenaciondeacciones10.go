package enajenaciondeacciones

type EnajenacionDeAcciones10 struct {
	Version                string  `xml:"Version,attr" bson:"Version"`
	ContratoIntermediacion string  `xml:"ContratoIntermediacion,attr" bson:"ContratoIntermediacion"`
	Ganancia               float64 `xml:"Ganancia,attr" bson:"Ganancia"`
	Perdida                float64 `xml:"Perdida,attr" bson:"Perdida"`
}
