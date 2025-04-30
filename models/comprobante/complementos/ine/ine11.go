package ine

type INE11 struct {
	Version        string          `xml:"Version,attr" bson:"Version" json:"Version"`
	TipoProceso    string          `xml:"TipoProceso,attr" bson:"TipoProceso" json:"TipoProceso"`
	TipoComite     *string         `xml:"tipoComite,attr" bson:"TipoComite,omitempty" json:"TipoComite,omitempty"`
	IdContabilidad int32           `xml:"IdContabilidad,attr" bson:"IdContabilidad,omitempty" json:"IdContabilidad,omitempty"`
	Entidad        *[]EntidadINE11 `xml:"Entidad" bson:"Entidad,omitempty" json:"Entidad,omitempty"`
}

type EntidadINE11 struct {
	ClaveEntidad string               `xml:"ClaveEntidad,attr" bson:"ClaveEntidad" json:"ClaveEntidad"`
	Ambito       *string              `xml:"Ambito,attr" bson:"Ambito,omitempty" json:"Ambito,omitempty"`
	Contabilidad *[]ContabilidadINE11 `xml:"Contabilidad" bson:"Contabilidad,omitempty" json:"Contabilidad,omitempty"`
}

type ContabilidadINE11 struct {
	IdContabilidad int32 `xml:"IdContabilidad,attr" bson:"IdContabilidad" json:"IdContabilidad"`
}
