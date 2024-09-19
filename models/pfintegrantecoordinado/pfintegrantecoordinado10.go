package pfintegrantecoordinado

type PFIntegranteCoordinado10 struct {
	Version        string `xml:"version,attr" bson:""`
	ClaveVehicular string `xml:"ClaveVehicular,attr" bson:""`
	Placa          string `xml:"Placa,attr" bson:""`
	RfcPF          string `xml:"RFCPF,attr" bson:""`
}
