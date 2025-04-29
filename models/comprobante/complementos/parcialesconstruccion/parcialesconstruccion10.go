package parcialesconstruccion

type ParcialesConstruccion10 struct {
	Version                        string   `xml:"Version,attr" bson:"Version" json:"Version"`
	NoPermisoLicenciaOAutorizacion string   `xml:"NumPerLicoAut,attr" bson:"NoPermisoLicenciaOAutorizacion" json:"NoPermisoLicenciaOAutorizacion"`
	Inmueble                       Inmueble `xml:"Inmueble" bson:"Inmueble" json:"Inmueble"`
}

type Inmueble struct {
	Calle        string  `xml:"Calle,attr" bson:"Calle" json:"Calle"`
	NoExterior   *string `xml:"NoExterior,attr" bson:"NoExterior,omitempty" json:"NoExterior,omitempty"`
	NoInterior   *string `xml:"NoInterior,attr" bson:"NoInterior,omitempty" json:"NoInterior,omitempty"`
	Colonia      *string `xml:"Colonia,attr" bson:"Colonia,omitempty" json:"Colonia,omitempty"`
	Localidad    *string `xml:"Localidad,attr" bson:"Localidad,omitempty" json:"Localidad,omitempty"`
	Referencia   *string `xml:"Referencia,attr" bson:"Referencia,omitempty" json:"Referencia,omitempty"`
	Municipio    string  `xml:"Municipio,attr" bson:"Municipio" json:"Municipio"`
	Estado       string  `xml:"Estado,attr" bson:"Estado" json:"Estado"`
	CodigoPostal string  `xml:"CodigoPostal,attr" bson:"CodigoPostal" json:"CodigoPostal"`
}
