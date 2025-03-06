package comprobante

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"strings"
)

type Comprobante struct {
	Comprobante32 *Comprobante32 `bson:"Comprobante32"`
	Comprobante33 *Comprobante33 `bson:"Comprobante33"`
	Comprobante40 *Comprobante40 `bson:"Comprobante40"`
}

func SerializeComprobanteFromXml(inputXml []byte) (Comprobante, error) {
	var parsed Comprobante
	
	trimmedXml, errOnTrim := helpers.TrimXml(inputXml)
	if errOnTrim != nil {
		return parsed, errOnTrim
	}

	errUnmarshal := helpers.UnmarshalXMLWithEncoding(trimmedXml, &parsed)
	if errUnmarshal != nil {
		return parsed, errUnmarshal
	}
	return parsed, nil
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

func (c Comprobante32) GetUuid() string {
	return c.Uuid
}
func (c Comprobante33) GetUuid() string {
	return c.Uuid
}
func (c Comprobante40) GetUuid() string {
	return c.Uuid
}
