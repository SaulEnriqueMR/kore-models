package renovacionsustitucion

import (
	"encoding/xml"
	"time"

	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type RenovacionSustitucion10 struct {
	Version            string              `xml:"Version,attr" bson:"Version"`
	TipoDecreto        string              `xml:"TipoDeDecreto,attr" bson:"TipoDecreto"`
	DecretoRenovacion  *DecretoRenovacion  `xml:"DecretoRenovVehicular" bson:"DecretoRenovacion,omitempty"`
	DecretoSustitucion *DecretoSustitucion `xml:"DecretoSustitVehicular" bson:"DecretoSustitucion,omitempty"`
}

type DecretoRenovacion struct {
	VehiculoEnajenado string            `xml:"VehEnaj,attr" bson:"VehiculoEnajenado"`
	VehiculosUsados   []VehiculosUsados `xml:"VehiculosUsadosEnajenadoPermAlFab" bson:"VehiculosUsados"`
	VehiculosNuevos   []VehiculosNuevos `xml:"VehiculoNuvoSemEnajenadoFabAlPerm" bson:"VehiculosNuevos"`
}

type DecretoSustitucion struct {
	VehiculoEnajenado string            `xml:"VehEnaj,attr" bson:"VehiculoEnajenado"`
	VehiculosUsados   []VehiculosUsados `xml:"VehiculosUsadosEnajenadoPermAlFab" bson:"VehiculosUsados"`
	VehiculosNuevos   []VehiculosNuevos `xml:"VehiculoNuvoSemEnajenadoFabAlPerm" bson:"VehiculosNuevos"`
}

type VehiculosUsados struct {
	PrecioVehUsado             float64    `xml:"PrecioVehUsado,attr" bson:"Precio"`
	Tipo                       string     `xml:"TipoVeh,attr" bson:"Tipo"`
	Marca                      string     `xml:"Marca,attr" bson:"Marca"`
	TipoOClase                 string     `xml:"TipooClase,attr" bson:"TipoOClase"`
	Anio                       string     `xml:"Año,attr" bson:"Anio"`
	Modelo                     *string    `xml:"Modelo,attr" bson:"Modelo,omitempty"`
	NoIdentificacionVehicular  *string    `xml:"NIV,attr" bson:"NoIdentificacionVehicular,omitempty"`
	NoSerie                    *string    `xml:"NumSerie,attr" bson:"NoSerie,omitempty"`
	NoPlacas                   string     `xml:"NumPlacas,attr" bson:"NoPlacas"`
	NoMotor                    *string    `xml:"NumMotor,attr" bson:"NoMotor,omitempty"`
	NoFolioTarjetaCirculacion  string     `xml:"NumFolTarjCir,attr" bson:"NoFolioTarjetaCirculacion"`
	NumeroPedimentoImportacion *string    `xml:"NumPedIm,attr" bson:"NumeroPedimentoImportacion,omitempty"`
	Aduana                     *string    `xml:"Aduana,attr" bson:"Aduana,omitempty"`
	FechaRegulVeh              *string    `xml:"FechaRegulVeh,attr" bson:"FechaRegulVeh"`
	FechaRegulacionImportacion *time.Time `bson:"FechaRegulacionImportacion,omitempty"`
	FolioFiscal                string     `xml:"Foliofiscal,attr" bson:"FolioFiscal"`
}

type VehiculosNuevos struct {
	Anio     string  `xml:"Año,attr" bson:"Anio"`
	Modelo   *string `xml:"Modelo,attr" bson:"Modelo,omitempty"`
	NoPlacas string  `xml:"NumPlacas,attr" bson:"NoPlacas"`
	Rfc      *string `xml:"RFC,attr" bson:"Rfc,omitempty"`
}

func (c *VehiculosUsados) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias VehiculosUsados
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*c = VehiculosUsados(aux)

	if aux.FechaRegulVeh != nil {
		fecha, err := helpers.ParseDatetime(*aux.FechaRegulVeh)
		if err != nil {
			return err
		}
		c.FechaRegulacionImportacion = &fecha
	}

	return nil
}
