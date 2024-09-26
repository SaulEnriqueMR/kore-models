package divisas

import (
	"encoding/xml"
	"strings"
)

type Divisas struct {
	Divisas10 *[]Divisas10 `bson:"Divisas10,omitempty"`
}

func (v *Divisas) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var divisas10 []Divisas10
		if err := d.DecodeElement(&divisas10, &start); err != nil {
			return err
		}
		if v.Divisas10 == nil {
			v.Divisas10 = &[]Divisas10{}
		}
		*v.Divisas10 = append(*v.Divisas10, divisas10...)
	}

	return nil
}
