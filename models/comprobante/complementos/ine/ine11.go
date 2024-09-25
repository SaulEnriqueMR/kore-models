package ine

type INE11 struct {
	Version        string          `xml:"Version,attr" bson:"Version"`
	TipoProceso    string          `xml:"TipoProceso,attr" bson:"TipoProceso"`
	TipoComite     *string         `xml:"tipoComite,attr" bson:"TipoComite,omitempty"`
	IdContabilidad int32           `xml:"IdContabilidad,attr" bson:"IdContabilidad,omitempty"`
	Entidad        *[]EntidadINE11 `xml:"Entidad" bson:"Entidad,omitempty"`
}

type EntidadINE11 struct {
	ClaveEntidad string               `xml:"ClaveEntidad,attr" bson:"ClaveEntidad"`
	Ambito       *string              `xml:"Ambito,attr" bson:"Ambito,omitempty"`
	Contabilidad *[]ContabilidadINE11 `xml:"Contabilidad" bson:"Contabilidad,omitempty"`
}

type ContabilidadINE11 struct {
	IdContabilidad int32 `xml:"IdContabilidad,attr" bson:"IdContabilidad"`
}
