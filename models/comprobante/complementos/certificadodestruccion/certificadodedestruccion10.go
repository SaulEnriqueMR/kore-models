package certificadodestruccion

type CertificadoDeDestruccion10 struct {
	Version             string                           `xml:"Version,attr" bson:"Version"`
	Serie               string                           `xml:"Serie,attr" bson:"Serie"`
	NumFolDesVeh        string                           `xml:"NumFolDesVeh,attr" bson:"NumFolDesVeh"`
	VehiculoDestruido   *[]VehiculoDestruidoCertDest10   `xml:"VehiculoDestruido" bson:"VehiculoDestruido"`
	InformacinoAduanera *[]InformacionAduaneraCertDest10 `xml:"InformacionAduanera" bson:"InformacionAduanera"`
}

type VehiculoDestruidoCertDest10 struct {
	Marca         string  `xml:"Marca,attr" bson:"Marca"`
	TipoOClase    string  `xml:"TipooClase,attr" bson:"TipoOClase"`
	Anio          string  `xml:"Año,attr" bson:"Anio"`
	Modelo        *string `xml:"Modelo,attr" bson:"Modelo,omitempty"`
	Niv           *string `xml:"NIV,attr" bson:"NIV,omitempty"`
	NumSerie      *string `xml:"NumSerie,attr" bson:"NumSerie,omitempty"`
	NumPlacas     string  `xml:"NumPlacas,attr" bson:"NumPlacas"`
	NumMotor      *string `xml:"NumMotor,attr" bson:"NumMotor,omitempty"`
	NumFolTarjCir string  `xml:"NumFolTarjCir,attr" bson:"NumFolTarjCir"`
}

type InformacionAduaneraCertDest10 struct {
	NumPedImp string `xml:"NumPedImp,attr" bson:"NumPedImp"`
	Fecha     string `xml:"Fecha,attr" bson:"Fecha"`
	Aduana    string `xml:"Aduana,attr" bson:"Aduana"`
}
