package donatarias

import (
	"encoding/xml"
	"strings"
)

type Donatarias struct {
	Donatarias11 *[]Donatarias11 `bson:"Donatarias11,omitempty"`
}

func (d *Donatarias) UnmarshalXML(x *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.1" {
		var donatarias11 []Donatarias11
		if err := x.DecodeElement(&donatarias11, &start); err != nil {
			return err
		}
		if d.Donatarias11 == nil {
			d.Donatarias11 = &[]Donatarias11{}
		}
		*d.Donatarias11 = append(*d.Donatarias11, donatarias11...)
	}

	return nil
}
