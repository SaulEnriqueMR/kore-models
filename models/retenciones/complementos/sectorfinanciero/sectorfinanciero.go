package sectorfinanciero

import (
	"encoding/xml"
	"strings"
)

type SectorFinanciero struct {
	SectorFinanciero10 *[]SectorFinanciero10 `bson:"SectorFinanciero,omitempty" json:"SectorFinanciero,omitempty"`
}

func (s *SectorFinanciero) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var sector10 []SectorFinanciero10
		if err := d.DecodeElement(&sector10, &start); err != nil {
			return err
		}
		if s.SectorFinanciero10 == nil {
			s.SectorFinanciero10 = &[]SectorFinanciero10{}
		}
		*s.SectorFinanciero10 = append(*s.SectorFinanciero10, sector10...)
	}

	return nil
}
