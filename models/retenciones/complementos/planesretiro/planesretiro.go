package planesretiro

import (
	"encoding/xml"
	"strings"
)

type PlanesDeRetiro struct {
	PlanesDeRetiro10 *[]PlanesDeRetiro10 `bson:"PlanesDeRetiro10,omitempty"`
	PlanesDeRetiro11 *[]PlanesDeRetiro11 `bson:"PlanesDeRetiro11,omitempty"`
}

func (c *PlanesDeRetiro) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var planes10 []PlanesDeRetiro10
		if err := d.DecodeElement(&planes10, &start); err != nil {
			return err
		}
		c.PlanesDeRetiro10 = &planes10
	}

	if Version == "1.1" {
		var planes11 []PlanesDeRetiro11
		if err := d.DecodeElement(&planes11, &start); err != nil {
			return err
		}
		c.PlanesDeRetiro11 = &planes11
	}

	return nil
}
