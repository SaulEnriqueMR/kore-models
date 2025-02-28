package addenda

type AddendaFacto struct {
	GranTotal           float64             `xml:"granTotal"`
	CargosNoFacturables []CargoNoFacturable `xml:"cargoNoFacturable>cargoNoFacturable" bson:"CargosNoFacturables"`
	Hospedaje           Hospedaje           `xml:"hospedaje" bson:"Hospedaje"`
}

type CargoNoFacturable struct {
	Cargo   string  `xml:"cargo,attr" bson:"Cargo"`
	Importe float64 `xml:"importe,attr" bson:"Importe"`
}

type Hospedaje struct {
	Folio      string `xml:"folio" bson:"Folio"`
	Huesped    string `xml:"huesped" bson:"Huesped"`
	Habitacion string `xml:"habitacion" bson:"Habitacion"`
	CheckIn    string `xml:"chekIn" bson:"CheckIn"`
	CheckOut   string `xml:"chekOut" bson:"CheckOut"`
	Reserva    string `xml:"reserva" bson:"Reserva"`
	NoCoches   string `xml:"numeroNoches" bson:"NoCoches"`
}
