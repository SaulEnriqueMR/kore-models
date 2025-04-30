package divisas

type Divisas10 struct {
	Version       string `xml:"version,attr" bson:"Version" json:"Version"`
	TipoOperacion string `xml:"tipoOperacion,attr" bson:"TipoOperacion" json:"TipoOperacion"`
}
