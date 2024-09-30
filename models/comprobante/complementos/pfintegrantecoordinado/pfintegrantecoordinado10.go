package pfintegrantecoordinado

type PFIntegranteCoordinado10 struct {
	Version        string `xml:"version,attr" bson:"Version"`
	ClaveVehicular string `xml:"ClaveVehicular,attr" bson:"ClaveVehicular"`
	Placa          string `xml:"Placa,attr" bson:"Placa"`
	Rfc            string `xml:"RFCPF,attr" bson:"Rfc"`
}
