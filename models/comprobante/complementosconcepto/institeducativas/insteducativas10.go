package institeducativas

type InstitucioneEducativas10 struct {
	Version        string  `xml:"version,attr" bson:"Version" json:"Version"`
	NombreAlumno   string  `xml:"nombreAlumno,attr" bson:"NombreAlumno" json:"NombreAlumno"`
	Curp           string  `xml:"CURP,attr" bson:"Curp" json:"Curp"`
	NivelEducativo string  `xml:"nivelEducativo,attr" bson:"NivelEducativo" json:"NivelEducativo"`
	AutRvoe        string  `xml:"autRVOE,attr" bson:"AutRvoe" json:"AutRvoe"`
	RfcPago        *string `xml:"rfcPago,attr" bson:"RfcPago,omitempty" json:"RfcPago,omitempty"`
}
