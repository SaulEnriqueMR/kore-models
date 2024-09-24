package comprobante

import (
	"encoding/xml"
	"strings"
)

type Comprobante struct {
	Comprobante32 *Comprobante32 `bson:"Comprobante32,omitempty"`
	Comprobante33 *Comprobante33 `bson:"Comprobante33,omitempty"`
	Comprobante40 *Comprobante40 `bson:"Comprobante40,omitempty"`
}

func (c *Comprobante) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var Version string
	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" || attributes.Name.Local == "version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}
	if Version == "3.2" {
		var comp32 Comprobante32
		if err := d.DecodeElement(&comp32, &start); err != nil {
			return err
		}
		c.Comprobante32 = &comp32
	}
	if Version == "3.3" {
		var comp33 Comprobante33
		if err := d.DecodeElement(&comp33, &start); err != nil {
			return err
		}
		c.Comprobante33 = &comp33
	}
	if Version == "4.0" {
		var comp40 Comprobante40
		if err := d.DecodeElement(&comp40, &start); err != nil {
			return err
		}
		c.Comprobante40 = &comp40
	}

	return nil
}
