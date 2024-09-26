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
		if v.INE11 == nil {
			v.INE11 = &[]INE11{}
		}
		*v.INE11 = append(*v.INE11, ine11...)
	}

	return nil
}
