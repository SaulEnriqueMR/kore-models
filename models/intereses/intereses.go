package intereses

import (
	"encoding/xml"
	"strings"
)

type Intereses struct {
	Intereses10 *[]Intereses10 `bson:"Intereses10,omitempty"`
}

func (c *Intereses) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var intereses10 []Intereses10
		if err := d.DecodeElement(&intereses10, &start); err != nil {
			return err
		}
		c.Intereses10 = &intereses10
	}

	return nil
}
