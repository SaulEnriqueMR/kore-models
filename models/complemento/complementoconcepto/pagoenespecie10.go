package complementoconcepto

type PagoEnEspecie10 struct {
	Version      string `xml:"Version,attr" bson:"Version"`
	CvePIC       string `xml:"CvePIC,attr" bson:"CvePIC"`
	FolioSolDon  string `xml:"FolioSolDon,attr" bson:"FolioSolDon"`
	PzaArtNombre string `xml:"PzaArtNombre,attr" bson:"PzaArtNombre"`
	PzaArtTecn   string `xml:"PzaArtTecn,attr" bson:"PzaArtTecn"`
	PzaArtAProd  string `xml:"PzaArtAProd,attr" bson:"PzaArtAProd"`
	PzaArtDim    string `xml:"PzaArtDim,attr" bson:"PzaArtDim"`
}
