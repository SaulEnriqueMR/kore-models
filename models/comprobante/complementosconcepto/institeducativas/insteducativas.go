package institeducativas

import (
	"encoding/xml"
	"strings"
)

type InstitucioneEducativas struct {
	InstitucioneEducativas10 *[]InstitucioneEducativas10 `bson:"InstitucioneEducativas10,omitempty" json:"InstitucioneEducativas10,omitempty"`
}

func (i *InstitucioneEducativas) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var instadu10 []InstitucioneEducativas10
		if err := d.DecodeElement(&instadu10, &start); err != nil {
			return err
		}
		if i.InstitucioneEducativas10 == nil {
			i.InstitucioneEducativas10 = &[]InstitucioneEducativas10{}
		}
		*i.InstitucioneEducativas10 = append(*i.InstitucioneEducativas10, instadu10...)
	}

	return nil
}
