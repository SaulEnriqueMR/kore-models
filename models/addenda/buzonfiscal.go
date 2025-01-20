package addenda

type BuzonFiscal struct {
	Version       string                   `xml:"version,attr" bson:"Version"`
	TipoDocumento TipoDocumentoBuzonFiscal `xml:"TipoDocumento" bson:"TipoDocumento"`
	Receptor      ReceptorBuzonFiscal      `xml:"Receptor" bson:"Receptor"`
	Destino       DestinoBuzonFiscal       `xml:"Destino" bson:"Destino"`
	Extra         []ExtraBuzonFiscal       `xml:"Extra" bson:"Extra"`
}

type TipoDocumentoBuzonFiscal struct {
	Descripcion string `xml:"descripcion,attr" bson:"Descripcion"`
	NombreCorto string `xml:"nombreCorto,attr" bson:"NombreCorto"`
}

type ReceptorBuzonFiscal struct {
	NoCliente string `xml:"noCliente,attr" bson:"NoCliente"`
}

type DestinoBuzonFiscal struct {
	Nombre              string `xml:"nombre,attr" bson:"Nombre"`
	Estado              string `xml:"estado,attr" bson:"Estado"`
	Municipio           string `xml:"municipio,attr" bson:"Municipio"`
	Calle               string `xml:"calle,attr" bson:"Calle"`
	ClaveIdentificacion string `xml:"claveIdentificacion,attr" bson:"ClaveIdentificacion"`
}

type CfdBuzonFiscal struct {
	TotalConLetra string `xml:"totalConLetra,attr" bson:"TotalConLetra"`
}

type ExtraBuzonFiscal struct {
	Valor    string `xml:"valor,attr" bson:"Valor"`
	Atributo string `xml:"atributo,attr" bson:"Atributo"`
}
