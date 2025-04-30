package addenda

type BuzonFiscal struct {
	Version       string                   `xml:"version,attr" bson:"Version" json:"Version"`
	TipoDocumento TipoDocumentoBuzonFiscal `xml:"TipoDocumento" bson:"TipoDocumento" json:"TipoDocumento"`
	Receptor      ReceptorBuzonFiscal      `xml:"Receptor" bson:"Receptor" json:"Receptor"`
	Destino       DestinoBuzonFiscal       `xml:"Destino" bson:"Destino" json:"Destino"`
	Extra         []ExtraBuzonFiscal       `xml:"Extra" bson:"Extra" json:"Extra"`
}

type TipoDocumentoBuzonFiscal struct {
	Descripcion string `xml:"descripcion,attr" bson:"Descripcion" json:"Descripcion"`
	NombreCorto string `xml:"nombreCorto,attr" bson:"NombreCorto" json:"NombreCorto"`
}

type ReceptorBuzonFiscal struct {
	NoCliente string `xml:"noCliente,attr" bson:"NoCliente" json:"NoCliente"`
}

type DestinoBuzonFiscal struct {
	Nombre              string `xml:"nombre,attr" bson:"Nombre" json:"Nombre"`
	Estado              string `xml:"estado,attr" bson:"Estado" json:"Estado"`
	Municipio           string `xml:"municipio,attr" bson:"Municipio" json:"Municipio"`
	Calle               string `xml:"calle,attr" bson:"Calle" json:"Calle"`
	ClaveIdentificacion string `xml:"claveIdentificacion,attr" bson:"ClaveIdentificacion" json:"ClaveIdentificacion"`
}

type CfdBuzonFiscal struct {
	TotalConLetra string `xml:"totalConLetra,attr" bson:"TotalConLetra" json:"TotalConLetra"`
}

type ExtraBuzonFiscal struct {
	Valor    string `xml:"valor,attr" bson:"Valor" json:"Valor"`
	Atributo string `xml:"atributo,attr" bson:"Atributo" json:"Atributo"`
}
