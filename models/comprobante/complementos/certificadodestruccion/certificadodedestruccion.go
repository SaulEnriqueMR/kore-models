package certificadodestruccion

import (
	"encoding/xml"
	"strings"
)

type CertificadoDeDestruccion struct {
	CertificadoDeDestruccion10 *[]CertificadoDeDestruccion10 `bson:"CertificadoDeDestruccion10,omitempty" json:"CertificadoDeDestruccion10,omitempty"`
}

func (c *CertificadoDeDestruccion) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var certi10 []CertificadoDeDestruccion10
		if err := d.DecodeElement(&certi10, &start); err != nil {
			return err
		}
		if c.CertificadoDeDestruccion10 == nil {
			c.CertificadoDeDestruccion10 = &[]CertificadoDeDestruccion10{}
		}
		*c.CertificadoDeDestruccion10 = append(*c.CertificadoDeDestruccion10, certi10...)
	}

	return nil
}
