package ventavehiculos

type VentaVehiculos10 struct {
	Version        string             `xml:"version,attr" bson:"Version"`
	ClaveVehicular string             `xml:"ClaveVehicular,attr" bson:"ClaveVehicular"`
	Parte          *[]ParteVentVehi10 `xml:"Parte" bson:"Parte,omitempty"`
}

type ParteVentVehi10 struct {
	InformacionAduanera *[]InformacionAduaneraVentVehi10 `xml:"InformacionAduanera" bson:"InformacionAduanera,omitempty"`
	Cantidad            string                           `xml:"cantidad,attr" bson:"Cantidad"`
	Unidad              *string                          `xml:"unidad,attr" bson:"Unidad,omitempty"`
	NoIdentificacion    *string                          `xml:"noIdentificacion,attr" bson:"NoIdentificacion,omitempty"`
	Descripcion         string                           `xml:"descripcion,attr" bson:"Descripcion"`
	ValorUnitario       *string                          `xml:"valorUnitario,attr" bson:"ValorUnitario,omitempty"`
	Importe             *string                          `xml:"importe,attr" bson:"Importe,omitempty"`
}

type InformacionAduaneraVentVehi10 struct {
	Numero string  `xml:"numero,attr" bson:"Numero"`
	Fecha  string  `xml:"fecha,attr" bson:"Fecha"`
	Aduana *string `xml:"aduana,attr" bson:"Aduana,omitempty"`
}
