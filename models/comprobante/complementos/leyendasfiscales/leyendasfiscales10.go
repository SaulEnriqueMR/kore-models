package leyendasfiscales

type LeyendasFiscales10 struct {
	Version string                 `xml:"version,attr" bson:"Version"`
	Leyenda []LeyendaLeyendasFis10 `xml:"Leyenda" bson:"Leyenda"`
}

type LeyendaLeyendasFis10 struct {
	DisposicionFiscal *string `xml:"disposicionFiscal,attr" bson:"DisposicionFiscal,omitempty"`
	Norma             *string `xml:"norma,attr" bson:"Norma,omitempty"`
	TextoLeyenda      string  `xml:"textoLeyenda,attr" bson:"TextoLeyenda"`
}
