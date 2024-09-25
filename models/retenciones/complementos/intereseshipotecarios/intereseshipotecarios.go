package intereseshipotecarios

import (
	"encoding/xml"
	"strings"
)

type InteresesHipotecarios struct {
	InteresesHipotecarios10 *[]InteresesHipotecarios10 `bson:"InteresesHipotecarios10,omitempty"`
}

func (c *InteresesHipotecarios) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var intereses10 []InteresesHipotecarios10
		if err := d.DecodeElement(&intereses10, &start); err != nil {
			return err
		}
		c.InteresesHipotecarios10 = &intereses10
	}

	return nil
}
