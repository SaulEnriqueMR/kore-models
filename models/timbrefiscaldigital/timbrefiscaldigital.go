package timbrefiscaldigital

import (
	"encoding/xml"
	"strings"
)

type TimbreFiscalDigital struct {
	TimbreFiscalDigital11 *TimbreFiscalDigital11 `bson:"TimbreFiscalDigital11,omitempty" json:"TimbreFiscalDigital11"`
	TimbreFiscalDigital10 *TimbreFiscalDigital10 `bson:"TimbreFiscalDigital10,omitempty" json:"TimbreFiscalDigital10"`
}

func (t *TimbreFiscalDigital) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" || attributes.Name.Local == "version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var timb10 TimbreFiscalDigital10
		if err := d.DecodeElement(&timb10, &start); err != nil {
			return err
		}
		t.TimbreFiscalDigital10 = &timb10
	}

	if Version == "1.1" {
		var timb11 TimbreFiscalDigital11
		if err := d.DecodeElement(&timb11, &start); err != nil {
			return err
		}
		t.TimbreFiscalDigital11 = &timb11
	}

	return nil
}
