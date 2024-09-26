package parcialesconstruccion

import (
	"encoding/xml"
	"strings"
)

type ParcialesConstruccion struct {
	ParcialesConstruccion10 *[]ParcialesConstruccion10 `bson:"ParcialesConstruccion10,omitempty"`
}

func (p *ParcialesConstruccion) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var parciales10 []ParcialesConstruccion10
		if err := d.DecodeElement(&parciales10, &start); err != nil {
			return err
		}
		if p.ParcialesConstruccion10 == nil {
			p.ParcialesConstruccion10 = &[]ParcialesConstruccion10{}
		}
		*p.ParcialesConstruccion10 = append(*p.ParcialesConstruccion10, parciales10...)
	}

	return nil
}
