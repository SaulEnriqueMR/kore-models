package addenda

type AddendaFacto struct {
	GranTotal           float64             `xml:"granTotal"`
	CargosNoFacturables []CargoNoFacturable `xml:"cargoNoFacturable>cargoNoFacturable" bson:"CargosNoFacturables" json:"CargosNoFacturables"`
	Hospedaje           Hospedaje           `xml:"hospedaje" bson:"Hospedaje" json:"Hospedaje"`
}

type CargoNoFacturable struct {
	Cargo   string  `xml:"cargo,attr" bson:"Cargo" json:"Cargo"`
	Importe float64 `xml:"importe,attr" bson:"Importe" json:"Importe"`
}

type Hospedaje struct {
	Folio      string `xml:"folio" bson:"Folio" json:"Folio"`
	Huesped    string `xml:"huesped" bson:"Huesped" json:"Huesped"`
	Habitacion string `xml:"habitacion" bson:"Habitacion" json:"Habitacion"`
	CheckIn    string `xml:"chekIn" bson:"CheckIn" json:"CheckIn"`
	CheckOut   string `xml:"chekOut" bson:"CheckOut" json:"CheckOut"`
	Reserva    string `xml:"reserva" bson:"Reserva" json:"Reserva"`
	NoCoches   string `xml:"numeroNoches" bson:"NoCoches" json:"NoCoches"`
}
