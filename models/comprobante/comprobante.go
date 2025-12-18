package comprobante

import (
	"encoding/xml"
	"strings"
)

type Comprobante struct {
	Comprobante30 *Comprobante30 `bson:"Comprobante30" json:"Comprobante30"`
	Comprobante32 *Comprobante32 `bson:"Comprobante32" json:"Comprobante32"`
	Comprobante33 *Comprobante33 `bson:"Comprobante33" json:"Comprobante33"`
	Comprobante40 *Comprobante40 `bson:"Comprobante40" json:"Comprobante40"`
}

func (c *Comprobante) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if strings.ToLower(attributes.Name.Local) == "version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "3.0" {
		var c30 Comprobante30
		if err := d.DecodeElement(&c30, &start); err != nil {
			return err
		}
		c.Comprobante30 = &c30
	}

	if Version == "3.2" {
		var c32 Comprobante32
		if err := d.DecodeElement(&c32, &start); err != nil {
			return err
		}
		c.Comprobante32 = &c32
	}

	if Version == "3.3" {
		var c33 Comprobante33
		if err := d.DecodeElement(&c33, &start); err != nil {
			return err
		}
		c.Comprobante33 = &c33
	}

	if Version == "4.0" {
		var c40 Comprobante40
		if err := d.DecodeElement(&c40, &start); err != nil {
			return err
		}
		c.Comprobante40 = &c40
	}

	return nil
}

type HasUuid interface {
	GetUuid() string
}

func (c Comprobante30) GetUuid() string {
	return c.Uuid
}
func (c Comprobante32) GetUuid() string {
	return c.Uuid
}
func (c Comprobante33) GetUuid() string {
	return c.Uuid
}
func (c Comprobante40) GetUuid() string {
	return c.Uuid
}
