package vehiculousado

import (
	"encoding/xml"
	"time"

	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type VehiculoUsado10 struct {
	Version                   string                 `xml:"Version,attr" bson:"Version" json:"Version"`
	MontoAdquisicion          float64                `xml:"montoAdquisicion,attr" bson:"MontoAdquisicion" json:"MontoAdquisicion"`
	MontoEnajenacion          float64                `xml:"montoEnajenacion,attr" bson:"MontoEnajenacion" json:"MontoEnajenacion"`
	ClaveVehicular            string                 `xml:"claveVehicular,attr" bson:"ClaveVehicular" json:"ClaveVehicular"`
	Marca                     string                 `xml:"marca,attr" bson:"Marca" json:"Marca"`
	Tipo                      string                 `xml:"tipo,attr" bson:"Tipo" json:"Tipo"`
	Modelo                    string                 `xml:"modelo,attr" bson:"Modelo" json:"Modelo"`
	NoMotor                   *string                `xml:"numeroMotor,attr" bson:"NoMotor,omitempty" json:"NoMotor,omitempty"`
	NoSerie                   *string                `xml:"numeroSerie,attr" bson:"NoSerie,omitempty" json:"NoSerie,omitempty"`
	NoIdentificacionVehicular *string                `xml:"NIV,attr" bson:"NoIdentificacionVehicular,omitempty" json:"NoIdentificacionVehicular,omitempty"`
	Valor                     float64                `xml:"valor,attr" bson:"Valor" json:"Valor"`
	InformacionAduanera       *[]InformacionAduanera `xml:"InformacionAduanera" bson:"InformacionAduanera,omitempty" json:"InformacionAduanera,omitempty"`
}

type InformacionAduanera struct {
	Numero      string    `xml:"numero,attr" bson:"Numero" json:"Numero"`
	FechaString string    `xml:"fecha,attr" bson:"FechaString" json:"FechaString"`
	Fecha       time.Time `bson:"Fecha" json:"Fecha"`
	Aduana      *string   `xml:"aduana,attr" bson:"Aduana,omitempty" json:"Aduana,omitempty"`
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
