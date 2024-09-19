package complementos

type Divisas10 struct {
	Version       string `xml:"version,attr" bson:"Version"`
	TipoOperacion string `xml:"tipoOperacion,attr" bson:"TipoOperacion"`
}
