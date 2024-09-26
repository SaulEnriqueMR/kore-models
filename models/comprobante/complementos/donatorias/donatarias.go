package donatorias

import (
	"encoding/xml"
	"strings"
)

type Donatarias struct {
	Donatarias11 *[]Donatarias11 `bson:"Donatarias11,omitempty"`
}

func (v *Donatarias) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.1" {
		var donatorias11 []Donatarias11
		if err := d.DecodeElement(&donatorias11, &start); err != nil {
			return err
		}
		v.Donatarias11 = &donatorias11
	}

	return nil
}
