package premios

import (
	"encoding/xml"
	"strings"
)

type Premios struct {
	Premios10 *[]Premios10 `bson:"Premios10,omitempty"`
}

func (v *Premios) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var premios10 []Premios10
		if err := d.DecodeElement(&premios10, &start); err != nil {
			return err
		}
		v.Premios10 = &premios10
	}

	return nil
}
