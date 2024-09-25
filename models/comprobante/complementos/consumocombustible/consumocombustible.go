package consumocombustible

import (
	"encoding/xml"
	"strings"
)

type ConsumoDeCombustible struct {
	ConsumoDeCombustible11 *[]ConsumoDeCombustible11 `bson:"ConsumoDeCombustible11,omitempty"`
}

func (v *ConsumoDeCombustible) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.1" {
		var nota10 []ConsumoDeCombustible11
		if err := d.DecodeElement(&nota10, &start); err != nil {
			return err
		}
		v.ConsumoDeCombustible11 = &nota10
	}

	return nil
}
