package pagoespecie

import (
	"encoding/xml"
	"strings"
)

type PagoEnEspecie struct {
	PagoEspecie10 *[]PagoEspecie10 `bson:"PagoEspecie10,omitempty"`
}

func (v *PagoEnEspecie) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var pago10 []PagoEspecie10
		if err := d.DecodeElement(&pago10, &start); err != nil {
			return err
		}
		v.PagoEspecie10 = &pago10
	}

	return nil
}
