package renovacionsustitucion

import (
	"encoding/xml"
	"time"

	"github.com/SaulEnriqueMR/kore-models/models"
)

type RenovacionSustitucion10 struct {
	Version            string              `xml:"Version,attr" bson:"Version"`
	TipoDecreto        string              `xml:"TipoDeDecreto,attr" bson:"TipoDecreto"`
	DecretoRenovacion  *DecretoRenovacion  `xml:"DecretoRenovVehicular" bson:"DecretoRenovacion,omitempty"`
	DecretoSustitucion *DecretoSustitucion `xml:"DecretoSustitVehicular" bson:"DecretoSustitucion,omitempty"`
}

type DecretoRenovacion struct {
	VehEnaj         string            `xml:"VehEnaj,attr" bson:"VehEnaj"`
	VehiculosUsados []VehiculosUsados `xml:"VehiculosUsadosEnajenadoPermAlFab" bson:"VehiculosUsados"`
	VehiculosNuevos []VehiculosNuevos `xml:"VehiculoNuvoSemEnajenadoFabAlPerm" bson:"VehiculosNuevos"`
}

type DecretoSustitucion struct {
	VehEnaj         string            `xml:"VehEnaj,attr" bson:"VehEnaj"`
	VehiculosUsados []VehiculosUsados `xml:"VehiculosUsadosEnajenadoPermAlFab" bson:"VehiculosUsados"`
	VehiculosNuevos []VehiculosNuevos `xml:"VehiculoNuvoSemEnajenadoFabAlPerm" bson:"VehiculosNuevos"`
}

type VehiculosUsados struct {
	PrecioVehUsado      float64    `xml:"PrecioVehUsado,attr" bson:"PrecioVehUsado"`
	TipoVeh             string     `xml:"TipoVeh,attr" bson:"TipoVeh"`
	Marca               string     `xml:"Marca,attr" bson:"Marca"`
	TipooClase          string     `xml:"TipooClase,attr" bson:""`
	Anio                int        `xml:"Año,attr" bson:"Anio"`
	Modelo              *string    `xml:"Modelo,attr" bson:"Modelo,omitempty"`
	NIV                 *string    `xml:"NIV,attr" bson:"NIV,omitempty"`
	NumSerie            *string    `xml:"NumSerie,attr" bson:"NumSerie,omitempty"`
	NumPlacas           string     `xml:"NumPlacas,attr" bson:"NumPlacas"`
	NumMotor            *string    `xml:"NumMotor,attr" bson:"NumMotor,omitempty"`
	NumFolTarjCir       string     `xml:"NumFolTarjCir,attr" bson:"NumFolTarjCir"`
	NumPedIm            *string    `xml:"NumPedIm,attr" bson:"NumPedIm,omitempty"`
	Aduana              *string    `xml:"Aduana,attr" bson:"Aduana,omitempty"`
	FechaRegulVehString *string    `xml:"FechaRegulVeh,attr"`
	FechaRegulVeh       *time.Time `bson:"FechaRegulVeh,omitempty"`
	Foliofiscal         string     `xml:"Foliofiscal,attr" bson:"Foliofiscal"`
}

type VehiculosNuevos struct {
	Anio      int     `xml:"Año,attr" bson:"Anio"`
	Modelo    *string `xml:"Modelo,attr" bson:"Modelo,omitempty"`
	NumPlacas string  `xml:"NumPlacas,attr" bson:"NumPlacas"`
	Rfc       *string `xml:"RFC,attr" bson:"Rfc,omitempty"`
}

func (c *RenovacionSustitucion10) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias RenovacionSustitucion10
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*c = RenovacionSustitucion10(aux)

	if aux.DecretoRenovacion != nil {
		for index, vehiculoUsado := range aux.DecretoRenovacion.VehiculosUsados {
			if vehiculoUsado.FechaRegulVehString != nil {
				fecha, err := models.ParseDatetime(*vehiculoUsado.FechaRegulVehString)
				if err != nil {
					return err
				}
				c.DecretoRenovacion.VehiculosUsados[index].FechaRegulVeh = &fecha
			}
		}
	}
	if aux.DecretoSustitucion != nil {
		for index, vehiculoUsado := range aux.DecretoSustitucion.VehiculosUsados {
			if vehiculoUsado.FechaRegulVehString != nil {
				fecha, err := models.ParseDatetime(*vehiculoUsado.FechaRegulVehString)
				if err != nil {
					return err
				}
				c.DecretoSustitucion.VehiculosUsados[index].FechaRegulVeh = &fecha
			}
		}
	}

	return nil
}
