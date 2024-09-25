package pagosaextranjeros

import (
	"encoding/xml"
	"strings"
)

type PagosAExtranjeros struct {
	PagosAExtranjeros10 *[]PagosAExtranjeros10 `bson:"PagosAExtranjeros10,omitempty"`
}

func (c *PagosAExtranjeros) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var pagoext10 []PagosAExtranjeros10
		if err := d.DecodeElement(&pagoext10, &start); err != nil {
			return err
		}
		c.PagosAExtranjeros10 = &pagoext10
	}

	return nil
}
