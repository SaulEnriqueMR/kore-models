package retenciones

import (
	"encoding/xml"
	"strings"
)

type Retenciones struct {
	Retenciones10 *Retenciones10 `bson:"Retenciones10,omitempty"`
	Retenciones20 *Retenciones20 `bson:"Retenciones20,omitempty"`
}

func (r *Retenciones) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var Version string
	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}
	if Version == "1.0" {
		var ret10 Retenciones10
		if err := d.DecodeElement(&ret10, &start); err != nil {
			return err
		}
		r.Retenciones10 = &ret10
	}
	if Version == "2.0" {
		var ret20 Retenciones20
		if err := d.DecodeElement(&ret20, &start); err != nil {
			return err
		}
		r.Retenciones20 = &ret20
	}
	return nil
}
