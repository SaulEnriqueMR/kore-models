package cartaporte

import (
	"encoding/xml"
	"strings"
)

type CartaPorte struct {
	CartaPorte10 *[]CartaPorte10 `bson:"CartaPorte10,omitempty"`
	CartaPorte20 *[]CartaPorte20 `bson:"CartaPorte20,omitempty"`
	CartaPorte30 *[]CartaPorte30 `bson:"CartaPorte30,omitempty"`
	CartaPorte31 *[]CartaPorte31 `bson:"CartaPorte31,omitempty"`
}

func (c *CartaPorte) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var cp10 []CartaPorte10
		if err := d.DecodeElement(&cp10, &start); err != nil {
			return err
		}
		c.CartaPorte10 = &cp10
	}

	if Version == "2.0" {
		var cp20 []CartaPorte20
		if err := d.DecodeElement(&cp20, &start); err != nil {
			return err
		}
		c.CartaPorte20 = &cp20
	}

	if Version == "3.0" {
		var cp30 []CartaPorte30
		if err := d.DecodeElement(&cp30, &start); err != nil {
			return err
		}
		c.CartaPorte30 = &cp30
	}

	if Version == "3.1" {
		var cp31 []CartaPorte31
		if err := d.DecodeElement(&cp31, &start); err != nil {
			return err
		}
		c.CartaPorte31 = &cp31
	}

	return nil
}
