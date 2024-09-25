package ine

import (
	"encoding/xml"
	"strings"
)

type INE struct {
	INE11 *[]INE11 `bson:"INE11,omitempty"`
}

func (v *INE) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.1" {
		var ine11 []INE11
		if err := d.DecodeElement(&ine11, &start); err != nil {
			return err
		}
		v.INE11 = &ine11
	}

	return nil
}
