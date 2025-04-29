package leyendasfiscales

type LeyendasFiscales10 struct {
	Version string                 `xml:"version,attr" bson:"Version" json:"Version"`
	Leyenda []LeyendaLeyendasFis10 `xml:"Leyenda" bson:"Leyenda" json:"Leyenda"`
}

type LeyendaLeyendasFis10 struct {
	DisposicionFiscal *string `xml:"disposicionFiscal,attr" bson:"DisposicionFiscal,omitempty" json:"DisposicionFiscal,omitempty"`
	Norma             *string `xml:"norma,attr" bson:"Norma,omitempty" json:"Norma,omitempty"`
	TextoLeyenda      string  `xml:"textoLeyenda,attr" bson:"TextoLeyenda" json:"TextoLeyenda"`
}
