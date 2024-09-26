package dividendos

import (
	"encoding/xml"
	"strings"
)

type Dividendos struct {
	Dividendos10 *[]Dividendos10 `bson:"Dividendos10,omitempty"`
}

func (d *Dividendos) UnmarshalXML(x *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var dividendos10 []Dividendos10
		if err := x.DecodeElement(&dividendos10, &start); err != nil {
			return err
		}
		if d.Dividendos10 == nil {
			d.Dividendos10 = &[]Dividendos10{}
		}
		*d.Dividendos10 = append(*d.Dividendos10, dividendos10...)
	}

	return nil
}
