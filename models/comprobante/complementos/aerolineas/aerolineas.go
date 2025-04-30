package aerolineas

import (
	"encoding/xml"
	"strings"
)

type Aerolineas struct {
	Aerolineas10 *[]Aerolineas10 `bson:"Aerolineas10,omitempty" json:"Aerolineas10,omitempty"`
}

func (c *Aerolineas) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var aerolineas10 []Aerolineas10
		if err := d.DecodeElement(&aerolineas10, &start); err != nil {
			return err
		}
		if c.Aerolineas10 == nil {
			c.Aerolineas10 = &[]Aerolineas10{}
		}
		*c.Aerolineas10 = append(*c.Aerolineas10, aerolineas10...)
	}

	return nil
}
