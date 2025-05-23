package donatarias

import (
	"encoding/xml"
	"time"

	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type Donatarias11 struct {
	Version           string    `xml:"version,attr" bson:"Version" json:"Version"`
	NoAutorizacion    string    `xml:"noAutorizacion,attr" bson:"NoAutorizacion" json:"NoAutorizacion"`
	FechaAutString    string    `xml:"fechaAutorizacion,attr" bson:"FechaAutString" json:"FechaAutString"`
	FechaAutorizacion time.Time `bson:"FechaAutorizacion" json:"FechaAutorizacion"`
	Leyenda           string    `xml:"leyenda,attr" bson:"Leyenda" json:"Leyenda"`
}

func (do *Donatarias11) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias Donatarias11
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	fechaAutorizacion, err := helpers.ParseDatetime(aux.FechaAutString)
	if err != nil {
		return err
	}

	*do = Donatarias11(aux)
	do.FechaAutorizacion = fechaAutorizacion

	return nil
}
