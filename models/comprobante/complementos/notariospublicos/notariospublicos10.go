package notariospublicos

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"time"
)

type NotariosPublicos10 struct {
	Version          string                        `xml:"Version,attr" bson:"Version"`
	DescInmueble     *[]DescInmuebleNotariosPub10  `xml:"DescInmuebles>DescInmueble" bson:"DescInmuebles"`
	DatosOperacion   DatosOperacionNotariosPub10   `xml:"DatosOperacion" bson:"DatosOperacion"`
	DatosNotario     DatosNotarioNotariosPub10     `xml:"DatosNotario" bson:"DatosNotario"`
	DatosEnajenante  DatosEnajenanteNotariosPub10  `xml:"DatosEnajenante" bson:"DatosEnajenante"`
	DatosAdquiriente DatosAdquirienteNotariosPub10 `xml:"DatosAdquiriente" bson:"DatosAdquiriente"`
}

type DescInmuebleNotariosPub10 struct {
	TipoInmueble string  `xml:"TipoInmueble,attr" bson:"TipoInmueble"`
	Calle        string  `xml:"Calle,attr" bson:"Calle"`
	NoExterior   *string `xml:"NoExterior,attr" bson:"NoExterior,omitempty"`
	NoInterior   *string `xml:"NoInterior,attr" bson:"NoInterior,omitempty"`
	Colonia      *string `xml:"Colonia,attr" bson:"Colonia,omitempty"`
	Localidad    *string `xml:"Localidad,attr" bson:"Localidad,omitempty"`
	Referencia   *string `xml:"Referencia,attr" bson:"Referencia,omitempty"`
	Municipio    string  `xml:"Municipio,attr" bson:"Municipio"`
	Estado       string  `xml:"Estado,attr" bson:"Estado,omitempty"`
	Pais         string  `xml:"Pais,attr" bson:"Pais,omitempty"`
	CodigoPostal string  `xml:"CodigoPostal,attr" bson:"CodigoPostal,omitempty"`
}

type DatosOperacionNotariosPub10 struct {
	NumInstrumentoNotarial  int64     `xml:"NumInstrumentoNotarial,attr" bson:"NumInstrumentoNotarial"`
	FechaInstNotarialString string    `xml:"FechaInstNotarial,attr"`
	FechaInstNotarial       time.Time `bson:"FechaInstNotarial"`
	MontoOperacion          float64   `xml:"MontoOperacion,attr" bson:"MontoOperacion"`
	Subtotal                float64   `xml:"Subtotal,attr" bson:"Subtotal"`
	IVA                     float64   `xml:"IVA,attr" bson:"IVA"`
}

type DatosNotarioNotariosPub10 struct {
	Curp              string  `xml:"CURP,attr" bson:"Curp"`
	NumNotaria        int16   `xml:"NumNotaria,attr" bson:"NumNotaria"`
	EntidadFederativa string  `xml:"EntidadFederativa,attr" bson:"EntidadFederativa"`
	Adscripcion       *string `xml:"Adscripcion,attr" bson:"Adscripcion,omitempty"`
}

type DatosEnajenanteNotariosPub10 struct {
	CoproSocConyugalE    string                               `xml:"CoproSocConyugalE,attr" bson:"CoproSocConyugalE"`
	DatosUnEnajenante    *DatosUnEnajenanteNotariosPub10      `xml:"DatosUnEnajenante" bson:"DatosUnEnajenante"`
	DatosEnajenanteCopSC *[]DatosEnajenanteCopSCNotariosPub10 `xml:"DatosEnajenantesCopSC>DatosEnajenanteCopSC" bson:"DatosEnajenantesCopSC"`
}

type DatosUnEnajenanteNotariosPub10 struct {
	Nombre          string `xml:"Nombre,attr" bson:"Nombre"`
	ApellidoPaterno string `xml:"ApellidoPaterno,attr" bson:"ApellidoPaterno"`
	ApellidoMaterno string `xml:"ApellidoMaterno,attr" bson:"ApellidoMaterno"`
	Rfc             string `xml:"RFC,attr" bson:"RFC"`
	Curp            string `xml:"CURP,attr" bson:"Curp"`
}

type DatosEnajenanteCopSCNotariosPub10 struct {
	Nombre          string  `xml:"Nombre,attr" bson:"Nombre"`
	ApellidoPaterno *string `xml:"ApellidoPaterno,attr" bson:"ApellidoPaterno"`
	ApellidoMaterno *string `xml:"ApellidoMaterno,attr" bson:"ApellidoMaterno"`
	Rfc             string  `xml:"RFC,attr" bson:"RFC"`
	Curp            *string `xml:"CURP,attr" bson:"Curp"`
	Porcentaje      float64 `xml:"Porcentaje,attr" bson:"Porcentaje"`
}

type DatosAdquirienteNotariosPub10 struct {
	CoproSocConyugalE      string                                 `xml:"CoproSocConyugalE,attr" bson:"CoproSocConyugalE"`
	DatosUnAdquiriente     *DatosUnAdquirienteNotariosPub10       `xml:"DatosUnAdquiriente" bson:"DatosUnAdquiriente"`
	DatosAdquirientesCopSC *[]DatosAdquirientesCopSCNotariosPub10 `xml:"DatosAdquirientesCopSC>DatosAdquirienteCopSC" bson:"DatosAdquirientesCopSC"`
}

type DatosUnAdquirienteNotariosPub10 struct {
	Nombre          string  `xml:"Nombre,attr" bson:"Nombre"`
	ApellidoPaterno string  `xml:"ApellidoPaterno,attr" bson:"ApellidoPaterno"`
	ApellidoMaterno string  `xml:"ApellidoMaterno,attr" bson:"ApellidoMaterno"`
	Rfc             string  `xml:"RFC,attr" bson:"RFC"`
	Curp            *string `xml:"CURP,attr" bson:"Curp"`
}

type DatosAdquirientesCopSCNotariosPub10 struct {
	Nombre          string  `xml:"Nombre,attr" bson:"Nombre"`
	ApellidoPaterno *string `xml:"ApellidoPaterno,attr" bson:"ApellidoPaterno"`
	ApellidoMaterno *string `xml:"ApellidoMaterno,attr" bson:"ApellidoMaterno"`
	Rfc             string  `xml:"RFC,attr" bson:"RFC"`
	Curp            *string `xml:"CURP,attr" bson:"Curp"`
	Porcentaje      float64 `xml:"Porcentaje,attr" bson:"Porcentaje"`
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
	fechaInstNotarial, err := helpers.ParseDatetime(aux.FechaInstNotarialString)
	if err != nil {
		return err
	}

	// Assign the parsed date and other fields back to the original struct
	*t = DatosOperacionNotariosPub10(aux)
	t.FechaInstNotarial = fechaInstNotarial

	return nil
}
