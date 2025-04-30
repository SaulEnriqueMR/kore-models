package valesdespensa

import (
	"encoding/xml"
	"time"

	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type ValesDespensa10 struct {
	Version          string      `xml:"version,attr" bson:"Version" json:"Version"`
	TipoOperacion    string      `xml:"tipoOperacion,attr" bson:"TipoOperacion" json:"TipoOperacion"`
	RegistroPatronal *string     `xml:"registroPatronal,attr" bson:"RegistroPatronal,omitempty" json:"RegistroPatronal,omitempty"`
	NoCuenta         string      `xml:"numeroDeCuenta,attr" bson:"NoCuenta" json:"NoCuenta"`
	Total            float64     `xml:"total,attr" bson:"Total" json:"Total"`
	Conceptos        []Conceptos `xml:"Conceptos>Concepto" bson:"Conceptos" json:"Conceptos"`
}

type Conceptos struct {
	Identificador     string    `xml:"identificador,attr" bson:"Identificador" json:"Identificador"`
	FechaString       string    `xml:"fecha,attr" bson:"FechaString" json:"FechaString"`
	Fecha             time.Time `bson:"Fecha" json:"Fecha"`
	Rfc               string    `xml:"rfc,attr" bson:"Rfc" json:"Rfc"`
	Curp              string    `xml:"curp,attr" bson:"Curp" json:"Curp"`
	Nombre            string    `xml:"nombre,attr" bson:"Nombre" json:"Nombre"`
	NoSeguridadSocial string    `xml:"numSeguridadSocial,attr" bson:"NoSeguridadSocial,omitempty" json:"NoSeguridadSocial,omitempty"`
	Importe           float64   `xml:"importe,attr" bson:"Importe" json:"Importe"`
}

func (c *Conceptos) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias Conceptos
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*c = Conceptos(aux)

	if aux.FechaString != "" {
		fecha, err := helpers.ParseDatetime(aux.FechaString)
		if err != nil {
			return err
		}
		c.Fecha = fecha
	}

	return nil
}
