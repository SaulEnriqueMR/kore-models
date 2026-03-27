package hidrocarburospetroliferos

import (
	"encoding/xml"
	"strings"
)

type HidrocarburosPetroliferos struct {
	HidroCarburosPetroliferos10 *[]HidrocarburosPetroliferos10 `bson:"HidroCarburosPetroliferos10" json:"HidroCarburosPetroliferos10"`
}

func (a *HidrocarburosPetroliferos) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var hp10 []HidrocarburosPetroliferos10
		if err := d.DecodeElement(&hp10, &start); err != nil {
			return err
		}
		if a.HidroCarburosPetroliferos10 == nil {
			a.HidroCarburosPetroliferos10 = &[]HidrocarburosPetroliferos10{}
		}
		*a.HidroCarburosPetroliferos10 = append(*a.HidroCarburosPetroliferos10, hp10...)
	}

	return nil
}
