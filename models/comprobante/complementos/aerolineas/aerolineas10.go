package aerolineas

type Aerolineas10 struct {
	Version     string               `xml:"Version,attr" bson:"Version" json:"Version"`
	Tua         string               `xml:"TUA,attr" bson:"Tua" json:"Tua"`
	Otroscargos *[]OtrosCargosAero10 `xml:"OtrosCargos" bson:"OtrosCargos,omitempty" json:"OtrosCargos,omitempty"`
}

type OtrosCargosAero10 struct {
	TotalCargos string        `xml:"TotalCargos,attr" bson:"TotalCargos" json:"TotalCargos"`
	Cargo       []CargoAero10 `xml:"Cargo" bson:"Cargo" json:"Cargo"`
}

type CargoAero10 struct {
	CodigoCargo string `xml:"CodigoCargo,attr" bson:"CodigoCargo" json:"CodigoCargo"`
	Importe     string `xml:"Importe,attr" bson:"Importe" json:"Importe"`
}
