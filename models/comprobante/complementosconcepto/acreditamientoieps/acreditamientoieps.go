package acreditamientoieps

import (
	"encoding/xml"
	"strings"
)

type AcreditamientoIeps struct {
	AcreditamientoIeps10 *[]AcreditamientoIeps10 `xml:"acreditamientoIEPS" bson:"AcreditamientoIeps10,omitempty" json:"AcreditamientoIeps10,omitempty"`
}

func (ai *AcreditamientoIeps) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var ai10 []AcreditamientoIeps10
		if err := d.DecodeElement(&ai10, &start); err != nil {
			return err
		}
		if ai.AcreditamientoIeps10 == nil {
			ai.AcreditamientoIeps10 = &[]AcreditamientoIeps10{}
		}
		*ai.AcreditamientoIeps10 = append(*ai.AcreditamientoIeps10, ai10...)
	}

	return nil
}
