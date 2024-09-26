package fideicomisonoempresarial

import (
	"encoding/xml"
	"strings"
)

type FideicomisoEmpresarial struct {
	FideicomisoEmpresarial10 *[]FideicomisoEmpresarial10 `bson:"FideicomisoEmpresarial10,omitempty"`
}

func (f *FideicomisoEmpresarial) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var fideicomiso10 []FideicomisoEmpresarial10
		if err := d.DecodeElement(&fideicomiso10, &start); err != nil {
			return err
		}
		if f.FideicomisoEmpresarial10 == nil {
			f.FideicomisoEmpresarial10 = &[]FideicomisoEmpresarial10{}
		}
		*f.FideicomisoEmpresarial10 = append(*f.FideicomisoEmpresarial10, fideicomiso10...)
	}

	return nil
}
