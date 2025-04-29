package pfintegrantecoordinado

import (
	"encoding/xml"
	"strings"
)

type PFIntegranteCoordinado struct {
	PFIntegranteCoordinado10 *[]PFIntegranteCoordinado10 `bson:"PFIntegranteCoordinado10,omitempty" json:"PFIntegranteCoordinado10,omitempty"`
}

func (p *PFIntegranteCoordinado) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var pfintegrante10 []PFIntegranteCoordinado10
		if err := d.DecodeElement(&pfintegrante10, &start); err != nil {
			return err
		}
		if p.PFIntegranteCoordinado10 == nil {
			p.PFIntegranteCoordinado10 = &[]PFIntegranteCoordinado10{}
		}
		*p.PFIntegranteCoordinado10 = append(*p.PFIntegranteCoordinado10, pfintegrante10...)
	}

	return nil
}
