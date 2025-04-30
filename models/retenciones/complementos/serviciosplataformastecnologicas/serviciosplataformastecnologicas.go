package serviciosplataformastecnologicas

import (
	"encoding/xml"
	"strings"
)

type ServiciosPlataformasTecnologicas struct {
	ServiciosPlataformasTecnologicas10 *[]ServiciosPlataformasTecnologicas10 `bson:"ServiciosPlataformasTecnologicas10,omitempty" json:"ServiciosPlataformasTecnologicas10,omitempty"`
}

func (s *ServiciosPlataformasTecnologicas) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Version string

	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var servicios10 []ServiciosPlataformasTecnologicas10
		if err := d.DecodeElement(&servicios10, &start); err != nil {
			return err
		}
		if s.ServiciosPlataformasTecnologicas10 == nil {
			s.ServiciosPlataformasTecnologicas10 = &[]ServiciosPlataformasTecnologicas10{}
		}
		*s.ServiciosPlataformasTecnologicas10 = append(*s.ServiciosPlataformasTecnologicas10, servicios10...)
	}

	return nil
}
