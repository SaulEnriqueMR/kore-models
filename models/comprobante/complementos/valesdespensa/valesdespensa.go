package valesdespensa

import (
	"encoding/xml"
	"strings"
)

type ValesDespensa struct {
	ValesDespensa10 *[]ValesDespensa10 `bson:"ValesDespensa10,omitempty"`
}

func (v *ValesDespensa) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var vd10 []ValesDespensa10
		if err := d.DecodeElement(&vd10, &start); err != nil {
			return err
		}
		if v.ValesDespensa10 == nil {
			v.ValesDespensa10 = &[]ValesDespensa10{}
		}
		*v.ValesDespensa10 = append(*v.ValesDespensa10, vd10...)
	}

	return nil
}
