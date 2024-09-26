package vehiculousado

import (
	"encoding/xml"
	"strings"
)

type VehiculoUsado struct {
	VehiculoUsado10 *[]VehiculoUsado10 `bson:"VehiculoUsado10"`
}

func (v *VehiculoUsado) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var vu10 []VehiculoUsado10
		if err := d.DecodeElement(&vu10, &start); err != nil {
			return err
		}
		if v.VehiculoUsado10 == nil {
			v.VehiculoUsado10 = &[]VehiculoUsado10{}
		}
		*v.VehiculoUsado10 = append(*v.VehiculoUsado10, vu10...)
	}

	return nil
}
