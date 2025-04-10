package estadodecuentacombustible

import (
	"encoding/xml"
	"strings"
)

type EstadoDeCuentaCombustible struct {
	EstadoDeCuentaCombustible10 *[]EstadoDeCuentaCombustible10 `bson:"EstadoDeCuentaCombustible10,omitempty"`
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
		if v.EstadoDeCuentaCombustible11 == nil {
			v.EstadoDeCuentaCombustible11 = &[]EstadoDeCuentaCombustible11{}
		}
		*v.EstadoDeCuentaCombustible11 = append(*v.EstadoDeCuentaCombustible11, edcc11...)

		return nil
	}

	if Version == "1.2" {
		var edcc12 []EstadoDeCuentaCombustible12
		if err := d.DecodeElement(&edcc12, &start); err != nil {
			return err
		}
		if v.EstadoDeCuentaCombustible12 == nil {
			v.EstadoDeCuentaCombustible12 = &[]EstadoDeCuentaCombustible12{}
		}
		*v.EstadoDeCuentaCombustible12 = append(*v.EstadoDeCuentaCombustible12, edcc12...)

		return nil
	}

	var edcc10 []EstadoDeCuentaCombustible10
	if err := d.DecodeElement(&edcc10, &start); err != nil {
		return err
	}
	if v.EstadoDeCuentaCombustible10 == nil {
		v.EstadoDeCuentaCombustible10 = &[]EstadoDeCuentaCombustible10{}
	}
	*v.EstadoDeCuentaCombustible10 = append(*v.EstadoDeCuentaCombustible10, edcc10...)

	return nil
}
