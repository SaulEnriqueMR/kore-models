package valesdespensa

import (
	"encoding/xml"
	"time"

	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type ValesDespensa10 struct {
	Version          string      `xml:"version,attr" bson:"Version"`
	TipoOperacion    string      `xml:"tipoOperacion,attr" bson:"TipoOperacion"`
	RegistroPatronal *string     `xml:"registroPatronal,attr" bson:"RegistroPatronal,omitempty"`
	NumeroCuenta     string      `xml:"numeroDeCuenta,attr" bson:"NumeroCuenta"`
	Total            float64     `xml:"total,attr" bson:"Total"`
	Conceptos        []Conceptos `xml:"Conceptos>Concepto" bson:"Conceptos"`
}

type Conceptos struct {
	Identificador      string    `xml:"identificador,attr" bson:"Identificador"`
	FechaString        string    `xml:"fecha,attr"`
	Fecha              time.Time `bson:"Fecha"`
	Rfc                string    `xml:"rfc,attr" bson:"Rfc"`
	Curp               string    `xml:"curp,attr" bson:"Curp"`
	Nombre             string    `xml:"nombre,attr" bson:"Nombre"`
	NumSeguridadSocial string    `xml:"numSeguridadSocial,attr" bson:"NumSeguridadSocial,omitempty"`
	Importe            float64   `xml:"importe,attr" bson:"Importe "`
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
