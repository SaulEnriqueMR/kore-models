package nomina

import (
	"encoding/xml"
	"strings"
)

type Nomina struct {
	Nomina11 *[]Nomina11 `bson:"Nomina11,omitempty"`
	Nomina12 *[]Nomina12 `bson:"Nomina12,omitempty"`
}

func (n *Nomina) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.1" {
		var nom11 []Nomina11
		if err := d.DecodeElement(&nom11, &start); err != nil {
			return err
		}
		if n.Nomina11 == nil {
			n.Nomina11 = &[]Nomina11{}
		}
		*n.Nomina11 = append(*n.Nomina11, nom11...)
	}

	if Version == "1.2" {
		var nom12 []Nomina12
		if err := d.DecodeElement(&nom12, &start); err != nil {
			return err
		}
		if n.Nomina12 == nil {
			n.Nomina12 = &[]Nomina12{}
		}
		*n.Nomina12 = append(*n.Nomina12, nom12...)
	}

	return nil
}
