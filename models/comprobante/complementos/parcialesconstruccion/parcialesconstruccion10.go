package parcialesconstruccion

type ParcialesConstruccion10 struct {
	Version       string   `xml:"Version,attr" bson:"Version"`
	NumPerLicoAut string   `xml:"NumPerLicoAut,attr" bson:"NumPerLicoAut"`
	Inmueble      Inmueble `xml:"Inmueble" bson:"Inmueble"`
}

type Inmueble struct {
	Calle        string  `xml:"Calle,attr" bson:"Calle"`
	NoExterior   *string `xml:"NoExterior,attr" bson:"NoExterior,omitempty"`
	NoInterior   *string `xml:"NoInterior,attr" bson:"NoInterior,omitempty"`
	Colonia      *string `xml:"Colonia,attr" bson:"Colonia,omitempty"`
	Localidad    *string `xml:"Localidad,attr" bson:"Localidad,omitempty"`
	Referencia   *string `xml:"Referencia,attr" bson:"Referencia,omitempty"`
	Municipio    string  `xml:"Municipio,attr" bson:"Municipio"`
	Estado       string  `xml:"Estado,attr" bson:"Estado"`
	CodigoPostal string  `xml:"CodigoPostal,attr" bson:"CodigoPostal"`
}
