package turistapasajeroextranjero

import (
	"encoding/xml"
	"strings"
)

type TuristaPasajeroExtranjero struct {
	TuristaPasajeroExtranjero10 *[]TuristaPasajeroExtranjero10 `bson:"TuristaPasajeroExtranjero10,omitempty"`
}

func (v *TuristaPasajeroExtranjero) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var tpe10 []TuristaPasajeroExtranjero10
		if err := d.DecodeElement(&tpe10, &start); err != nil {
			return err
		}
		v.TuristaPasajeroExtranjero10 = &tpe10
	}

	return nil
}
