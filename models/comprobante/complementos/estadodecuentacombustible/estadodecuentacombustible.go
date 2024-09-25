package estadodecuentacombustible

import (
	"encoding/xml"
	"strings"
)

type EstadoDeCuentaCombustible struct {
	EstadoDeCuentaCombustible11 *[]EstadoDeCuentaCombustible11 `bson:"EstadoDeCuentaCombustible11,omitempty"`
	EstadoDeCuentaCombustible12 *[]EstadoDeCuentaCombustible12 `bson:"EstadoDeCuentaCombustible12,omitempty"`
}

func (v *EstadoDeCuentaCombustible) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.1" {
		var edcc11 []EstadoDeCuentaCombustible11
		if err := d.DecodeElement(&edcc11, &start); err != nil {
			return err
		}
		v.EstadoDeCuentaCombustible11 = &edcc11
	}

	if Version == "1.2" {
		var edcc12 []EstadoDeCuentaCombustible12
		if err := d.DecodeElement(&edcc12, &start); err != nil {
			return err
		}
		v.EstadoDeCuentaCombustible12 = &edcc12
	}

	return nil
}
