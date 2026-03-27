package hidrocarburospetroliferos

type HidrocarburosPetroliferos10 struct {
	Version     string `xml:"Version,attr" bson:"Version" json:"Version"`
	TipoPermiso string `xml:"TipoPermiso,attr" bson:"TipoPermiso" json:"TipoPermiso"`
	NoPermiso   string `xml:"NumeroPermiso,attr" bson:"NoPermiso" json:"NoPermiso"`
	Clave       string `xml:"ClaveHYP,attr" bson:"Clave" json:"Clave"`
	Subproducto string `xml:"SubProductoHYP,attr" bson:"Subproducto" json:"Subproducto"`
}
