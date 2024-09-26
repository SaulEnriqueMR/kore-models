package planesretiro

import (
	"encoding/xml"
	"strings"
)

type PlanesDeRetiro struct {
	PlanesDeRetiro10 *[]PlanesDeRetiro10 `bson:"PlanesDeRetiro10,omitempty"`
	PlanesDeRetiro11 *[]PlanesDeRetiro11 `bson:"PlanesDeRetiro11,omitempty"`
}

func (p *PlanesDeRetiro) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var planes10 []PlanesDeRetiro10
		if err := d.DecodeElement(&planes10, &start); err != nil {
			return err
		}
		if p.PlanesDeRetiro10 == nil {
			p.PlanesDeRetiro10 = &[]PlanesDeRetiro10{}
		}
		*p.PlanesDeRetiro10 = append(*p.PlanesDeRetiro10, planes10...)
	}

	if Version == "1.1" {
		var planes11 []PlanesDeRetiro11
		if err := d.DecodeElement(&planes11, &start); err != nil {
			return err
		}
		if p.PlanesDeRetiro11 == nil {
			p.PlanesDeRetiro11 = &[]PlanesDeRetiro11{}
		}
		*p.PlanesDeRetiro11 = append(*p.PlanesDeRetiro11, planes11...)
	}

	return nil
}
