package gastohidrocarburos

import (
	"encoding/xml"
	"strings"
)

type GastoHidrocarburos struct {
	GastoHidrocarburos10 *[]GastoHidrocarburos10 `bson:"GastoHidrocarburos10,omitempty"`
}

func (v *GastoHidrocarburos) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var gastos10 []GastoHidrocarburos10
		if err := d.DecodeElement(&gastos10, &start); err != nil {
			return err
		}
		v.GastoHidrocarburos10 = &gastos10
	}

	return nil
}
