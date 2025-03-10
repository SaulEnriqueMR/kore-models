package addenda

type EdicomGenerica struct {
	Elementos []EdicomGenericaElemento `xml:"elementos>elemento" bson:"Elementos"`
}

type EdicomGenericaElemento struct {
	Etiqueta    string `xml:"etiqueta"`
	Descripcion string `xml:"descripcion"`
	Valor       string `xml:"valor"`
}
