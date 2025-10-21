package addenda

type AdditionalInfo struct {
	Conector *AddInfoConector `xml:"Conector,omitempty" bson:"Conector" json:"Conector,omitempty"`
	Cfdi     *AddInfoCfdi     `xml:"CFDI,omitempty" bson:"Cfdi,omitempty" json:"Cfdi,omitempty"`
	Emsr     *AddInfoEmisor   `xml:"EMSR,omitempty" bson:"Emsr,omitempty" json:"Emsr,omitempty"`
	R        *AddInfoReceptor `xml:"R,omitempty" bson:"R,omitempty" json:"R,omitempty"`
	EXP      *AddInfoExp      `xml:"EXP,omitempty" bson:"exp,omitempty" json:"Exp,omitempty"`
	PTDAo    *[]AddInfoPtdao  `xml:"PTDAo,omitempty" bson:"PTDAo,omitempty" json:"PTDAo,omitempty"`
}

type AddInfoConector struct {
	Value string `xml:",innerxml" bson:"Value" json:"Value"`
}

type AddInfoCfdi struct {
	TipoDocumento    *string `xml:"tipoDocumento,attr,omitempty" bson:"tipoDocumento,omitempty" json:"TipoDocumento,omitempty"`
	NumeroDocumentos *string `xml:"numeroDocumentos,attr,omitempty" bson:"numeroDocumentos,omitempty" json:"NumeroDocumentos,omitempty"`
	Etiqueta1C       *string `xml:"etiqueta1C,attr,omitempty" bson:"etiqueta1C,omitempty,omitempty" json:"Etiqueta1C,omitempty,omitempty"`
	Valor1C          *string `xml:"valor1C,attr,omitempty" bson:"valor1C,omitempty" json:"Valor1C,omitempty"`
	Valor10C         *string `xml:"valor10C,attr,omitempty" bson:"valor10C,omitempty" json:"Valor10C,omitempty"`
}

type AddInfoEmisor struct {
	Calle        *string `xml:"calle,attr,omitempty" bson:"calle,omitempty" json:"Calle,omitempty"`
	NoExterior   *string `xml:"noExterior,attr,omitempty" bson:"noExterior,omitempty" json:"NoExterior,omitempty,omitempty"`
	NoInterior   *string `xml:"noInterior,attr,omitempty" bson:"noInterior,omitempty" json:"NoInterior,omitempty,omitempty"`
	Colonia      *string `xml:"colonia,attr,omitempty" bson:"colonia,omitempty" json:"Colonia,omitempty"`
	Municipio    *string `xml:"municipio,attr,omitempty" bson:"municipio,omitempty" json:"Municipio,omitempty"`
	Estado       *string `xml:"estado,attr,omitempty" bson:"estado,omitempty" json:"Estado,omitempty"`
	Pais         *string `xml:"pais,attr,omitempty" bson:"pais,omitempty" json:"Pais,omitempty"`
	CodigoPostal *string `xml:"codigoPostal,attr,omitempty" bson:"codigoPostal,omitempty" json:"CodigoPostal,omitempty"`
}

type AddInfoReceptor struct {
	Calle        *string `xml:"calle,attr,omitempty" bson:"calle,omitempty" json:"Calle,omitempty"`
	NoInterior   *string `xml:"noInterior,attr,omitempty" bson:"noInterior,omitempty" json:"NoInterior,omitempty"`
	Colonia      *string `xml:"colonia,attr,omitempty" bson:"colonia,omitempty" json:"Colonia,omitempty"`
	Referencia   *string `xml:"referencia,attr,omitempty" bson:"referencia,omitempty" json:"Referencia,omitempty"`
	Municipio    *string `xml:"municipio,attr,omitempty" bson:"municipio,omitempty" json:"Municipio,omitempty"`
	Estado       *string `xml:"estado,attr,omitempty" bson:"estado,omitempty" json:"Estado,omitempty"`
	Pais         *string `xml:"pais,attr,omitempty" bson:"pais,omitempty" json:"Pais,omitempty"`
	CodigoPostal *string `xml:"codigoPostal,attr,omitempty" bson:"codigoPostal,omitempty" json:"CodigoPostal,omitempty"`
	Conector     *string `xml:"conector,attr,omitempty" bson:"conector,omitempty" json:"Conector,omitempty"`
	Valor13R     *string `xml:"valor13R,attr,omitempty" bson:"valor13R,omitempty" json:"Valor13R,omitempty"`
}

type AddInfoPtdao struct {
	Etiqueta1Po *string `xml:"etiqueta1PO,attr,omitempty" bson:"Etiqueta1Po,omitempty" json:"Etiqueta1Po,omitempty"`
	Valor1Po    *string `xml:"valor1PO,attr,omitempty" bson:"Valor1Po,omitempty" json:"Valor1Po,omitempty"`
	Etiqueta2Po *string `xml:"etiqueta2PO,attr,omitempty" bson:"Etiqueta2Po,omitempty" json:"Etiqueta2Po,omitempty"`
	Valor2Po    *string `xml:"valor2PO,attr,omitempty" bson:"Valor2Po,omitempty" json:"Valor2Po,omitempty"`
	Etiqueta3Po *string `xml:"etiqueta3PO,attr,omitempty" bson:"Etiqueta3Po,omitempty" json:"Etiqueta3Po,omitempty"`
	Valor3Po    *string `xml:"valor3PO,attr,omitempty" bson:"Valor3Po,omitempty" json:"Valor3Po,omitempty"`
	Etiqueta4Po *string `xml:"etiqueta4PO,attr,omitempty" bson:"Etiqueta4Po,omitempty" json:"Etiqueta4Po,omitempty"`
	Valor4Po    *string `xml:"valor4PO,attr,omitempty" bson:"Valor4Po,omitempty" json:"Valor4Po,omitempty"`
	Etiqueta5Po *string `xml:"etiqueta5PO,attr,omitempty" bson:"Etiqueta5Po,omitempty" json:"Etiqueta5Po,omitempty"`
}

type AddInfoExp struct {
	Calle        *string `xml:"calle,attr,omitempty" bson:"calle,omitempty" json:"Calle,omitempty"`
	Colonia      *string `xml:"colonia,attr,omitempty" bson:"colonia,omitempty" json:"Colonia,omitempty"`
	Municipio    *string `xml:"municipio,attr,omitempty" bson:"municipio,omitempty" json:"Municipio,omitempty"`
	Estado       *string `xml:"estado,attr,omitempty" bson:"estado,omitempty" json:"Estado,omitempty"`
	Pais         *string `xml:"pais,attr,omitempty" bson:"pais,omitempty" json:"Pais,omitempty"`
	CodigoPostal *string `xml:"codigoPostal,attr,omitempty" bson:"codigoPostal,omitempty" json:"CodigoPostal,omitempty"`
}
