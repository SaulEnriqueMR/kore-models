package detallista

import (
	"encoding/xml"
	"strings"
)

type Detallista struct {
	Detallista10 *[]Detallista10 `bson:"Detallista10,omitempty"`
}

func (c *Detallista) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "contentVersion" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.3.1" {
		var detallista10 []Detallista10
		if err := d.DecodeElement(&detallista10, &start); err != nil {
			return err
		}
		c.Detallista10 = &detallista10
	}

	return nil
}
