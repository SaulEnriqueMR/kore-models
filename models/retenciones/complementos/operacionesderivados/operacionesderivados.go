package operacionesderivados

import (
	"encoding/xml"
	"strings"
)

type OperacionesDerivados struct {
	OperacionesDerivados10 *[]OperacionesDerivados10 `bson:"OperacionesDerivados10,omitempty"`
}

func (v *OperacionesDerivados) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var oper10 []OperacionesDerivados10
		if err := d.DecodeElement(&oper10, &start); err != nil {
			return err
		}
		v.OperacionesDerivados10 = &oper10
	}

	return nil
}
