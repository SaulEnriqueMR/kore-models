package renovacionsustitucion

import (
	"encoding/xml"
	"strings"
)

type RenovacionSustitucion struct {
	RenovacionSustitucion10 *[]RenovacionSustitucion10 `bson:"RenovacionSustitucion10,omitempty" json:"RenovacionSustitucion10,omitempty"`
}

func (r *RenovacionSustitucion) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var renova10 []RenovacionSustitucion10
		if err := d.DecodeElement(&renova10, &start); err != nil {
			return err
		}
		if r.RenovacionSustitucion10 == nil {
			r.RenovacionSustitucion10 = &[]RenovacionSustitucion10{}
		}
		*r.RenovacionSustitucion10 = append(*r.RenovacionSustitucion10, renova10...)
	}

	return nil
}
