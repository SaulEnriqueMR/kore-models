package institeducativas

type InstitucioneEducativas10 struct {
	Version        string  `xml:"version,attr" bson:"Version"`
	NombreAlumno   string  `xml:"nombreAlumno,attr" bson:"NombreAlumno"`
	Curp           string  `xml:"CURP,attr" bson:"Curp"`
	NivelEducativo string  `xml:"nivelEducativo,attr" bson:"NivelEducativo"`
	AutRvoe        string  `xml:"autRVOE,attr" bson:"AutRvoe"`
	RfcPago        *string `xml:"rfcPago,attr" bson:"RfcPago,omitempty"`
}
