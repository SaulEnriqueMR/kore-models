package vehiculousado

import (
	"encoding/xml"
	"time"

	"github.com/SaulEnriqueMR/kore-models/models"
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

func (c *VehiculoUsado10) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias VehiculoUsado10
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*c = VehiculoUsado10(aux)

	if aux.InformacionAduanera != nil {
		for index, infoaduanera := range *aux.InformacionAduanera {
			if infoaduanera.FechaString != "" {
				fecha, err := models.ParseDatetime(infoaduanera.FechaString)
				if err != nil {
					return err
				}
				(*c.InformacionAduanera)[index].Fecha = fecha
			}
		}
	}
	return nil
}
