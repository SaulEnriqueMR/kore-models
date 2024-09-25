package cfdiregistrofiscal

import (
	"encoding/xml"
	"strings"
)

type CfdiRegistroFiscal struct {
	CfdiRegistroFiscal10 *[]CfdiRegistroFiscal10 `bson:"CfdiRegistroFiscal10,omitempty"`
}

func (c *CfdiRegistroFiscal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var cfdi10 []CfdiRegistroFiscal10
		if err := d.DecodeElement(&cfdi10, &start); err != nil {
			return err
		}
		c.CfdiRegistroFiscal10 = &cfdi10
	}

	return nil
}
