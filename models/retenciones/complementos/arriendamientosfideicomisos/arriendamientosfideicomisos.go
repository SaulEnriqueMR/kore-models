package arriendamientosfideicomisos

import (
	"encoding/xml"
	"strings"
)

type ArriendamientosFideicomisos struct {
	ArriendamientosFideicomisos10 *[]ArriendamientosFideicomisos10 `bson:"ArriendamientosFideicomisos10,omitempty"`
}

func (a *ArriendamientosFideicomisos) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var arriendamientos10 []ArriendamientosFideicomisos10
		if err := d.DecodeElement(&arriendamientos10, &start); err != nil {
			return err
		}
		if a.ArriendamientosFideicomisos10 == nil {
			a.ArriendamientosFideicomisos10 = &[]ArriendamientosFideicomisos10{}
		}
		*a.ArriendamientosFideicomisos10 = append(*a.ArriendamientosFideicomisos10, arriendamientos10...)
	}

	return nil
}
