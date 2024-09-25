package obrasantiguedades

import (
	"encoding/xml"
	"strings"
)

type ObrasAntiguedades struct {
	ObrasAntiguedades10 *[]ObrasAntiguedades10 `bson:"ObrasAntiguedades10,omitempty"`
}

func (v *ObrasAntiguedades) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var obras10 []ObrasAntiguedades10
		if err := d.DecodeElement(&obras10, &start); err != nil {
			return err
		}
		v.ObrasAntiguedades10 = &obras10
	}

	return nil
}
