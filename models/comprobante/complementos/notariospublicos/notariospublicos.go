package notariospublicos

import (
	"encoding/xml"
	"strings"
)

type NotariosPublicos struct {
	NotariosPublicos10 *[]NotariosPublicos10 `bson:"NotariosPublicos10,omitempty"`
}

func (v *NotariosPublicos) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var nota10 []NotariosPublicos10
		if err := d.DecodeElement(&nota10, &start); err != nil {
			return err
		}
		if v.NotariosPublicos10 == nil {
			v.NotariosPublicos10 = &[]NotariosPublicos10{}
		}
		*v.NotariosPublicos10 = append(*v.NotariosPublicos10, nota10...)
	}

	return nil
}
