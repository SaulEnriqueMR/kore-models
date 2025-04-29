package notariospublicos

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"time"
)

type NotariosPublicos10 struct {
	Version                string                        `xml:"Version,attr" bson:"Version" json:"Version"`
	DescripcionesInmuebles *[]DescInmuebleNotariosPub10  `xml:"DescInmuebles>DescInmueble" bson:"DescripcionesInmuebles" json:"DescripcionesInmuebles"`
	DatosOperacion         DatosOperacionNotariosPub10   `xml:"DatosOperacion" bson:"DatosOperacion" json:"DatosOperacion"`
	DatosNotario           DatosNotarioNotariosPub10     `xml:"DatosNotario" bson:"DatosNotario" json:"DatosNotario"`
	DatosEnajenante        DatosEnajenanteNotariosPub10  `xml:"DatosEnajenante" bson:"DatosEnajenante" json:"DatosEnajenante"`
	DatosAdquiriente       DatosAdquirienteNotariosPub10 `xml:"DatosAdquiriente" bson:"DatosAdquiriente" json:"DatosAdquiriente"`
}

type DescInmuebleNotariosPub10 struct {
	Tipo         string  `xml:"TipoInmueble,attr" bson:"Tipo" json:"Tipo"`
	Calle        string  `xml:"Calle,attr" bson:"Calle" json:"Calle"`
	NoExterior   *string `xml:"NoExterior,attr" bson:"NoExterior,omitempty" json:"NoExterior,omitempty"`
	NoInterior   *string `xml:"NoInterior,attr" bson:"NoInterior,omitempty" json:"NoInterior,omitempty"`
	Colonia      *string `xml:"Colonia,attr" bson:"Colonia,omitempty" json:"Colonia,omitempty"`
	Localidad    *string `xml:"Localidad,attr" bson:"Localidad,omitempty" json:"Localidad,omitempty"`
	Referencia   *string `xml:"Referencia,attr" bson:"Referencia,omitempty" json:"Referencia,omitempty"`
	Municipio    string  `xml:"Municipio,attr" bson:"Municipio" json:"Municipio"`
	Estado       string  `xml:"Estado,attr" bson:"Estado,omitempty" json:"Estado,omitempty"`
	Pais         string  `xml:"Pais,attr" bson:"Pais,omitempty" json:"Pais,omitempty"`
	CodigoPostal string  `xml:"CodigoPostal,attr" bson:"CodigoPostal,omitempty" json:"CodigoPostal,omitempty"`
}

type DatosOperacionNotariosPub10 struct {
	NoInstrumentoNotarial         string    `xml:"NumInstrumentoNotarial,attr" bson:"NoInstrumentoNotarial" json:"NoInstrumentoNotarial"`
	FechaInstNotarial             string    `xml:"FechaInstNotarial,attr" bson:"FechaInstNotarial" json:"FechaInstNotarial"`
	FechaFirmaInstrumentoNotarial time.Time `bson:"FechaFirmaInstrumentoNotarial" json:"FechaFirmaInstrumentoNotarial"`
	MontoOperacion                float64   `xml:"MontoOperacion,attr" bson:"MontoOperacion" json:"MontoOperacion"`
	Subtotal                      float64   `xml:"Subtotal,attr" bson:"Subtotal" json:"Subtotal"`
	Iva                           float64   `xml:"IVA,attr" bson:"Iva" json:"Iva"`
}

type DatosNotarioNotariosPub10 struct {
	Curp              string  `xml:"CURP,attr" bson:"Curp" json:"Curp"`
	NoNotaria         int16   `xml:"NumNotaria,attr" bson:"NoNotaria" json:"NoNotaria"`
	EntidadFederativa string  `xml:"EntidadFederativa,attr" bson:"EntidadFederativa" json:"EntidadFederativa"`
	Adscripcion       *string `xml:"Adscripcion,attr" bson:"Adscripcion,omitempty" json:"Adscripcion,omitempty"`
}

type DatosEnajenanteNotariosPub10 struct {
	CoproSocConyugalE                           string                               `xml:"CoproSocConyugalE,attr" bson:"CoproSocConyugalE" json:"CoproSocConyugalE"`
	EsCopropiedadOSociedadConyugal              bool                                 `bson:"EsCopropiedadOSociedadConyugal" json:"EsCopropiedadOSociedadConyugal"`
	DatosUnEnajenante                           *DatosUnEnajenanteNotariosPub10      `xml:"DatosUnEnajenante" bson:"DatosUnEnajenante" json:"DatosUnEnajenante"`
	DatosEnajenanteCopropiedadOSociedadConyugal *[]DatosEnajenanteCopSCNotariosPub10 `xml:"DatosEnajenantesCopSC>DatosEnajenanteCopSC" bson:"DatosEnajenantesCopropiedadOSociedadConyugal" json:"DatosEnajenantesCopropiedadOSociedadConyugal"`
}

func (de *DatosEnajenanteNotariosPub10) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias DatosEnajenanteNotariosPub10
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	*de = DatosEnajenanteNotariosPub10(aux)
	de.EsCopropiedadOSociedadConyugal = helpers.ResolveSatBoolean(de.CoproSocConyugalE)
	return nil
}

type DatosUnEnajenanteNotariosPub10 struct {
	Nombre          string `xml:"Nombre,attr" bson:"Nombre" json:"Nombre"`
	ApellidoPaterno string `xml:"ApellidoPaterno,attr" bson:"ApellidoPaterno" json:"ApellidoPaterno"`
	ApellidoMaterno string `xml:"ApellidoMaterno,attr" bson:"ApellidoMaterno" json:"ApellidoMaterno"`
	Rfc             string `xml:"RFC,attr" bson:"Rfc" json:"Rfc"`
	Curp            string `xml:"CURP,attr" bson:"Curp" json:"Curp"`
}

type DatosEnajenanteCopSCNotariosPub10 struct {
	Nombre          string  `xml:"Nombre,attr" bson:"Nombre" json:"Nombre"`
	ApellidoPaterno *string `xml:"ApellidoPaterno,attr" bson:"ApellidoPaterno" json:"ApellidoPaterno"`
	ApellidoMaterno *string `xml:"ApellidoMaterno,attr" bson:"ApellidoMaterno" json:"ApellidoMaterno"`
	Rfc             string  `xml:"RFC,attr" bson:"Rfc" json:"Rfc"`
	Curp            *string `xml:"CURP,attr" bson:"Curp" json:"Curp"`
	Porcentaje      float64 `xml:"Porcentaje,attr" bson:"Porcentaje" json:"Porcentaje"`
}

type DatosAdquirienteNotariosPub10 struct {
	CoproSocConyugalE                             string                                 `xml:"CoproSocConyugalE,attr" bson:"CoproSocConyugalE" json:"CoproSocConyugalE"`
	EsCopropiedadOSociedadConyugal                bool                                   `bson:"EsCopropiedadOSociedadConyugal" json:"EsCopropiedadOSociedadConyugal"`
	DatosUnAdquiriente                            *DatosUnAdquirienteNotariosPub10       `xml:"DatosUnAdquiriente" bson:"DatosUnAdquiriente" json:"DatosUnAdquiriente"`
	DatosAdquirientesCopropiedadOSociedadConyugal *[]DatosAdquirientesCopSCNotariosPub10 `xml:"DatosAdquirientesCopSC>DatosAdquirienteCopSC" bson:"DatosAdquirientesCopropiedadOSociedadConyugal" json:"DatosAdquirientesCopropiedadOSociedadConyugal"`
}

func (de *DatosAdquirienteNotariosPub10) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias DatosAdquirienteNotariosPub10
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	*de = DatosAdquirienteNotariosPub10(aux)
	de.EsCopropiedadOSociedadConyugal = helpers.ResolveSatBoolean(de.CoproSocConyugalE)
	return nil
}

type DatosUnAdquirienteNotariosPub10 struct {
	Nombre          string  `xml:"Nombre,attr" bson:"Nombre" json:"Nombre"`
	ApellidoPaterno string  `xml:"ApellidoPaterno,attr" bson:"ApellidoPaterno" json:"ApellidoPaterno"`
	ApellidoMaterno string  `xml:"ApellidoMaterno,attr" bson:"ApellidoMaterno" json:"ApellidoMaterno"`
	Rfc             string  `xml:"RFC,attr" bson:"Rfc" json:"Rfc"`
	Curp            *string `xml:"CURP,attr" bson:"Curp" json:"Curp"`
}

type DatosAdquirientesCopSCNotariosPub10 struct {
	Nombre          string  `xml:"Nombre,attr" bson:"Nombre" json:"Nombre"`
	ApellidoPaterno *string `xml:"ApellidoPaterno,attr" bson:"ApellidoPaterno" json:"ApellidoPaterno"`
	ApellidoMaterno *string `xml:"ApellidoMaterno,attr" bson:"ApellidoMaterno" json:"ApellidoMaterno"`
	Rfc             string  `xml:"RFC,attr" bson:"Rfc" json:"Rfc"`
	Curp            *string `xml:"CURP,attr" bson:"Curp" json:"Curp"`
	Porcentaje      float64 `xml:"Porcentaje,attr" bson:"Porcentaje" json:"Porcentaje"`
}

func (t *DatosOperacionNotariosPub10) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias DatosOperacionNotariosPub10
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	// Parse the date from the 'Fecha' field
	fechaInstNotarial, err := helpers.ParseDatetime(aux.FechaInstNotarial)
	if err != nil {
		return err
	}

	// Assign the parsed date and other fields back to the original struct
	*t = DatosOperacionNotariosPub10(aux)
	t.FechaFirmaInstrumentoNotarial = fechaInstNotarial

	return nil
}
