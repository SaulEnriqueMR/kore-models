package turistapasajeroextranjero

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"time"
)

type TuristaPasajeroExtranjero10 struct {
	Version             string        `xml:"version,attr" bson:"Version" json:"Version"`
	FechaTransitoString string        `xml:"fechadeTransito,attr" bson:"FechaTransitoString" json:"FechaTransitoString"`
	FechaTransito       time.Time     `bson:"FechaTransito" json:"FechaTransito"`
	TipoTransito        string        `xml:"tipoTransito,attr" bson:"TipoTransito" json:"TipoTransito"`
	DatosTransito       DatosTransito `xml:"datosTransito" bson:"DatosTransito" json:"DatosTransito"`
}

type DatosTransito struct {
	Via               string  `xml:"Via,attr" bson:"Via" json:"Via"`
	TipoId            string  `xml:"TipoId,attr" bson:"TipoId" json:"TipoId"`
	NumeroId          string  `xml:"NumeroId,attr" bson:"NumeroId" json:"NumeroId"`
	Nacionalidad      string  `xml:"Nacionalidad,attr" bson:"Nacionalidad" json:"Nacionalidad"`
	EmpresaTransporte string  `xml:"EmpresaTransporte,attr" bson:"EmpresaTransporte" json:"EmpresaTransporte"`
	IdTransporte      *string `xml:"IdTransporte,attr" bson:"IdTransporte,omitempty" json:"IdTransporte,omitempty"`
}

func (c *TuristaPasajeroExtranjero10) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias TuristaPasajeroExtranjero10
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	if aux.FechaTransitoString != "" {
		fecha, err := helpers.ParseDatetime(aux.FechaTransitoString)
		if err != nil {
			return err
		}
		*c = TuristaPasajeroExtranjero10(aux)
		c.FechaTransito = fecha
	}

	return nil
}
