package enajenaciondeacciones

type EnajenacionDeAcciones10 struct {
	Version                string  `xml:"Version,attr" bson:"Version" json:"Version"`
	ContratoIntermediacion string  `xml:"ContratoIntermediacion,attr" bson:"ContratoIntermediacion" json:"ContratoIntermediacion"`
	Ganancia               float64 `xml:"Ganancia,attr" bson:"Ganancia" json:"Ganancia"`
	Perdida                float64 `xml:"Perdida,attr" bson:"Perdida" json:"Perdida"`
}
