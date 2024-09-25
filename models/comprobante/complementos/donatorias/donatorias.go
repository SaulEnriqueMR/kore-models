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

	if Version == "1.0" {
		var donatorias10 []Donatarias11
		if err := d.DecodeElement(&donatorias10, &start); err != nil {
			return err
		}
		v.Donatarias11 = &donatorias10
	}

	return nil
}
