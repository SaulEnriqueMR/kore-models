package acuentaterceros

import (
	"encoding/xml"
	"strings"
)

type ACuentaTerceros struct {
	ACuentaTerceros11 *[]ACuentaTerceros11 `bson:"ACuentaTerceros11,omitempty" json:"ACuentaTerceros11,omitempty"`
}

func (a *ACuentaTerceros) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.1" {
		var acuenta11 []ACuentaTerceros11
		if err := d.DecodeElement(&acuenta11, &start); err != nil {
			return err
		}
		if a.ACuentaTerceros11 == nil {
			a.ACuentaTerceros11 = &[]ACuentaTerceros11{}
		}
		*a.ACuentaTerceros11 = append(*a.ACuentaTerceros11, acuenta11...)
	}

	return nil
}
