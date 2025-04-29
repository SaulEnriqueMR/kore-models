package leyendasfiscales

import (
	"encoding/xml"
	"strings"
)

type LeyendasFiscales struct {
	LeyendasFiscales10 *[]LeyendasFiscales10 `bson:"LeyendasFiscales10,omitempty" json:"LeyendasFiscales10,omitempty"`
}

func (v *LeyendasFiscales) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var leyendas10 []LeyendasFiscales10
		if err := d.DecodeElement(&leyendas10, &start); err != nil {
			return err
		}
		if v.LeyendasFiscales10 == nil {
			v.LeyendasFiscales10 = &[]LeyendasFiscales10{}
		}
		*v.LeyendasFiscales10 = append(*v.LeyendasFiscales10, leyendas10...)
	}

	return nil
}
