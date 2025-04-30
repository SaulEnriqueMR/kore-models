package pagoespecie

type PagoEspecie10 struct {
	Version                            string `xml:"Version,attr" bson:"Version" json:"Version"`
	ClavePadronInstitucionesCulturales string `xml:"CvePIC,attr" bson:"ClavePadronInstitucionesCulturales" json:"ClavePadronInstitucionesCulturales"`
	FolioSolicitudDonacion             string `xml:"FolioSolDon,attr" bson:"FolioSolicitudDonacion" json:"FolioSolicitudDonacion"`
	NombrePiezaArte                    string `xml:"PzaArtNombre,attr" bson:"NombrePiezaArte" json:"NombrePiezaArte"`
	TecnicaProduccionPiezaArte         string `xml:"PzaArtTecn,attr" bson:"TecnicaProduccionPiezaArte" json:"TecnicaProduccionPiezaArte"`
	AnioProduccionPiezaArte            string `xml:"PzaArtAProd,attr" bson:"AnioProduccionPiezaArte" json:"AnioProduccionPiezaArte"`
	DimensionPiezaArte                 string `xml:"PzaArtDim,attr" bson:"DimensionPiezaArte" json:"DimensionPiezaArte"`
}
