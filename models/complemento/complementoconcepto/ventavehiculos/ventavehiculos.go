package ventavehiculos

import (
	"encoding/xml"
	"strings"
)

type VentaVehiculos struct {
	VentaVehiculos10 *[]VentaVehiculos10 `bson:"VentaVehiculos10"`
	VentaVehiculos11 *[]VentaVehiculos11 `bson:"VentaVehiculos11"`
}

func (v *VentaVehiculos) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var vv10 []VentaVehiculos10
		if err := d.DecodeElement(&vv10, &start); err != nil {
			return err
		}

	}

	if Version == "1.1" {
		var vv11 []VentaVehiculos11
		if err := d.DecodeElement(&vv11, &start); err != nil {
			return err
		}
	}

	return nil
}
