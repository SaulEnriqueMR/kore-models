package donatorias

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"time"
)

type Donatarias11 struct {
	Version           string    `xml:"version,attr" bson:"Version"`
	NoAutorizacion    string    `xml:"noAutorizacion,attr" bson:"NoAutorizacion"`
	FechaAutString    string    `xml:"fechaAutorizacion,attr"`
	FechaAutorizacion time.Time `bson:"FechaAutorizacion"`
	Leyenda           string    `xml:"leyenda,attr" bson:"Leyenda"`
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
