package parcialesconstruccion

type ParcialesConstruccion10 struct {
	Version			string `xml:"Version" bson:"Version"`
	NumPerLicoAut 	string `xml:"NumPerLicoAut" bson:"NumPerLicoAut"`
	Inmueble 		Inmueble `xml:"Inmueble" bson:"Inmueble"`
}

type Inmueble struct {
	Calle 			string  `xml:"Calle" bson:"Calle"`
	NoExterior 		*string `xml:"NoExterior" bson:"NoExterior,omitempty"`
	NoInterior 		*string `xml:"NoInterior" bson:"NoInterior,omitempty"`
	Colonia 		*string `xml:"Colonia" bson:"Colonia,omitempty"`
	Localidad 		*string `xml:"Localidad" bson:"Localidad,omitempty"`
	Referencia 		*string `xml:"Referencia" bson:"Referencia,omitempty"`
	Municipio 		string  `xml:"Municipio" bson:"Municipio"`
	Estado 			string  `xml:"Estado" bson:"Estado"`
	CodigoPostal 	string  `xml:"CodigoPostal" bson:"CodigoPostal"`
}
