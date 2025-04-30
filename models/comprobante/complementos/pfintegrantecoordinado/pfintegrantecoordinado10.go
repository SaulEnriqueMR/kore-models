package pfintegrantecoordinado

type PFIntegranteCoordinado10 struct {
	Version        string `xml:"version,attr" bson:"Version" json:"Version"`
	ClaveVehicular string `xml:"ClaveVehicular,attr" bson:"ClaveVehicular" json:"ClaveVehicular"`
	Placa          string `xml:"Placa,attr" bson:"Placa" json:"Placa"`
	Rfc            string `xml:"RFCPF,attr" bson:"Rfc" json:"Rfc"`
}
