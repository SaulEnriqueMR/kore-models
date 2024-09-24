package pagosaextranjeros

import (
	"encoding/xml"
	"strings"
)

type Pagosaextranjeros struct {
	Pagosaextranjeros10 *[]Pagosaextranjeros10 `bson:"Pagosaextranjeros10,omitempty"`
}

func (c *Pagosaextranjeros) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var pagoext10 []Pagosaextranjeros10
		if err := d.DecodeElement(&pagoext10, &start); err != nil {
			return err
		}
		c.Pagosaextranjeros10 = &pagoext10
	}

	return nil
}
