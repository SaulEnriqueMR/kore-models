package consumocombustible

import (
	"encoding/xml"
	"strings"
)

type ConsumoDeCombustible struct {
	ConsumoDeCombustible11 *[]ConsumoDeCombustible11 `bson:"ConsumoDeCombustible11,omitempty"`
}

func (c *ConsumoDeCombustible) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.1" {
		var consumo11 []ConsumoDeCombustible11
		if err := d.DecodeElement(&consumo11, &start); err != nil {
			return err
		}
		if c.ConsumoDeCombustible11 == nil {
			c.ConsumoDeCombustible11 = &[]ConsumoDeCombustible11{}
		}
		*c.ConsumoDeCombustible11 = append(*c.ConsumoDeCombustible11, consumo11...)
	}

	return nil
}
