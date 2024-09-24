package pagos

import (
	"encoding/xml"
	"strings"
)

type Pagos struct {
	Pagos10 *[]Pagos10 `bson:"Pagos10,omitempty"`
	Pagos20 *[]Pagos20 `bson:"Pagos20,omitempty"`
}

func (p *Pagos) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var pagos10 []Pagos10
		if err := d.DecodeElement(&pagos10, &start); err != nil {
			return err
		}
		p.Pagos10 = &pagos10
	}

	if Version == "2.0" {
		var pagos20 []Pagos20
		if err := d.DecodeElement(&pagos20, &start); err != nil {
			return err
		}
		p.Pagos20 = &pagos20
	}

	return nil

}
