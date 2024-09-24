package dividendos

import (
	"encoding/xml"
	"strings"
)

type Dividendos struct {
	Dividendos10 *[]Dividendos10 `bson:"Dividendos10,omitempty"`
}

func (c *Dividendos) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

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
		if err := d.DecodeElement(&dividendos10, &start); err != nil {
			return err
		}
		c.Dividendos10 = &dividendos10
	}

	return nil
}
