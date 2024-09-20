package comercioexterior

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type ComercioExterior struct {
	ComercioExterior10 *[]ComercioExterior10 `bson:"ComercioExterior10,omitempty"`
	ComercioExterior11 *[]ComercioExterior11 `bson:"ComercioExterior11,omitempty"`
	ComercioExterior20 *[]ComercioExterior20 `bson:"ComercioExterior20,omitempty"`
}

func (c *ComercioExterior) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var comExt10 []ComercioExterior10
		if err := d.DecodeElement(&comExt10, &start); err != nil {
			return err
		}
		c.ComercioExterior10 = &comExt10
	}

	if Version == "1.1" {
		var comExt11 []ComercioExterior11
		if err := d.DecodeElement(&comExt11, &start); err != nil {
			return err
		}
		c.ComercioExterior11 = &comExt11
	}

	if Version == "2.0" {
		var comExt20 []ComercioExterior20
		if err := d.DecodeElement(&comExt20, &start); err != nil {
			fmt.Println(err)
			return err
		}
		c.ComercioExterior20 = &comExt20
	}

	return nil
}
