package pagoespecie

type PagoEspecie10 struct {
	Version                            string `xml:"Version,attr" bson:"Version"`
	ClavePadronInstitucionesCulturales string `xml:"CvePIC,attr" bson:"ClavePadronInstitucionesCulturales"`
	FolioSolicitudDonacion             string `xml:"FolioSolDon,attr" bson:"FolioSolicitudDonacion"`
	NombrePiezaArte                    string `xml:"PzaArtNombre,attr" bson:"NombrePiezaArte"`
	TecnicaProduccionPiezaArte         string `xml:"PzaArtTecn,attr" bson:"TecnicaProduccionPiezaArte"`
	AnioProduccionPiezaArte            string `xml:"PzaArtAProd,attr" bson:"AnioProduccionPiezaArte"`
	DimensionPiezaArte                 string `xml:"PzaArtDim,attr" bson:"DimensionPiezaArte"`
}
