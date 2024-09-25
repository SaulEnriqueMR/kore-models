package aerolineas

type Aerolineas10 struct {
	Version     string               `xml:"Version,attr" bson:"Version"`
	Tua         string               `xml:"TUA,attr" bson:"Tua"`
	Otroscargos *[]OtrosCargosAero10 `xml:"OtrosCargos" bson:"OtrosCargos,omitempty"`
}

type OtrosCargosAero10 struct {
	TotalCargos string        `xml:"TotalCargos,attr" bson:"TotalCargos"`
	Cargo       []CargoAero10 `xml:"Cargo" bson:"Cargo"`
}

type CargoAero10 struct {
	CodigoCargo string `xml:"CodigoCargo,attr" bson:"CodigoCargo"`
	Importe     string `xml:"Importe,attr" bson:"Importe"`
}
