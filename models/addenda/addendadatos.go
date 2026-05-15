package addenda

type AddendaDatos struct {
	Detalles []AddendaDatosDetalles `xml:"Detalles" bson:"Detalles" json:"Detalles"`
}

type AddendaDatosDetalles struct {
	IdRastreo string  `xml:"IDRastreo" bson:"IdRastreo" json:"IdRastreo"`
	Moneda    string  `xml:"Moneda" bson:"Moneda" json:"Moneda"`
	Monto     float64 `xml:"Monto" bson:"Monto" json:"Monto"`
	Tipo      string  `xml:"Tipo" bson:"Tipo" json:"Tipo"`
}
