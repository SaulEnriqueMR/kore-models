package complementoconcepto

import "encoding/xml"

type InstitucioneEducativas10 struct {
	XMLName        xml.Name `xml:"instEducativas"`
	Version        string   `xml:"version,attr" bson:"Version"`
	NombreAlumno   string   `xml:"nombreAlumno,attr" bson:"NombreAlumno"`
	Curp           string   `xml:"CURP,attr" bson:"CURP"`
	NivelEducativo string   `xml:"nivelEducativo,attr" bson:"NivelEducativo"`
	AutRVOE        string   `xml:"autRVOE,attr" bson:"autRVOE"`
	RfcPago        *string  `xml:"rfcPago,attr" bson:"rfcPago,omitempty"`
}
