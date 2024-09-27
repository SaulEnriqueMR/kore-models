package vehiculousado

import (
	"encoding/xml"
	"time"

	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type VehiculoUsado10 struct {
	Version             string                 `xml:"Version,attr" bson:"Version"`
	MontoAdquisicion    float64                `xml:"montoAdquisicion,attr" bson:"MontoAdquisicion"`
	MontoEnajenacion    float64                `xml:"montoEnajenacion,attr" bson:"MontoEnajenacion"`
	ClaveVehicular      string                 `xml:"claveVehicular,attr" bson:"ClaveVehicular"`
	Marca               string                 `xml:"marca,attr" bson:"Marca"`
	Tipo                string                 `xml:"tipo,attr" bson:"Tipo"`
	Modelo              string                 `xml:"modelo,attr" bson:"Modelo"`
	NumeroMotor         *string                `xml:"numeroMotor,attr" bson:"NumeroMotor,omitempty"`
	NumeroSerie         *string                `xml:"numeroSerie,attr" bson:"NumeroSerie,omitempty"`
	NIV                 *string                `xml:"NIV,attr" bson:"NIV,omitempty"`
	Valor               float64                `xml:"valor,attr" bson:"Valor"`
	InformacionAduanera *[]InformacionAduanera `xml:"InformacionAduanera" bson:"InformacionAduanera,omitempty"`
}

type InformacionAduanera struct {
	Numero      string    `xml:"numero,attr" bson:"Numero"`
	FechaString string    `xml:"fecha,attr"`
	Fecha       time.Time `bson:"Fecha"`
	Aduana      *string   `xml:"aduana,attr" bson:"Aduana,omitempty"`
}

func (c *InformacionAduanera) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias InformacionAduanera
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*c = InformacionAduanera(aux)

	if aux.FechaString != "" {
		fecha, err := helpers.ParseDatetime(aux.FechaString)
		if err != nil {
			return err
		}
		c.Fecha = fecha
	}

	return nil
}
