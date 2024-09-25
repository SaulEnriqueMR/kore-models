package enajenaciondeacciones

import (
	"encoding/xml"
	"strings"
)

type EnajenacionDeAcciones struct {
	EnajenacionDeAcciones10 *[]EnajenacionDeAcciones10 `bson:"EnajenacionDeAcciones10,omitempty"`
}

func (c *EnajenacionDeAcciones) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var enajenacion10 []EnajenacionDeAcciones10
		if err := d.DecodeElement(&enajenacion10, &start); err != nil {
			return err
		}
		c.EnajenacionDeAcciones10 = &enajenacion10
	}

	return nil
}