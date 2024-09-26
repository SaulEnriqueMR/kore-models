package cartaporte

import (
	"encoding/xml"
	"log"
	"strings"
)

type CartaPorte struct {
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

	if Version == "2.0" {
		var cp20 []CartaPorte20
		log.Println(len(cp20))
		if err := d.DecodeElement(&cp20, &start); err != nil {
			return err
		}
		if c.CartaPorte20 == nil {
			c.CartaPorte20 = &[]CartaPorte20{}
		}
		*c.CartaPorte20 = append(*c.CartaPorte20, cp20...)
	}

	if Version == "3.0" {
		var cp30 []CartaPorte30
		for i := 0; i < len(cp30); i++ {
			log.Println(cp30[i])
		}
		if err := d.DecodeElement(&cp30, &start); err != nil {
			return err
		}
		if c.CartaPorte30 == nil {
			c.CartaPorte30 = &[]CartaPorte30{}
		}
		*c.CartaPorte30 = append(*c.CartaPorte30, cp30...)
	}

	if Version == "3.1" {
		var cp31 []CartaPorte31
		for i := 0; i < len(cp31); i++ {
			log.Println(cp31[i])
		}
		if err := d.DecodeElement(&cp31, &start); err != nil {
			return err
		}
		if c.CartaPorte31 == nil {
			c.CartaPorte31 = &[]CartaPorte31{}
		}
		*c.CartaPorte31 = append(*c.CartaPorte31, cp31...)
	}

	return nil
}
