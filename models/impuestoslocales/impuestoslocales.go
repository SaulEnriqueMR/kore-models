package impuestoslocales

import (
	"encoding/xml"
	"strings"
)

type ImpuestosLocales struct {
	ImpuestoLocales10 *[]ImpuestoLocales10 `bson:"ImpuestoLocales10,omitempty"`
}

func (c *ImpuestosLocales) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var implocales10 []ImpuestoLocales10
		if err := d.DecodeElement(&implocales10, &start); err != nil {
			return err
		}
		c.ImpuestoLocales10 = &implocales10
	}

	return nil
}