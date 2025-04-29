package renovacionsustitucion

import (
	"encoding/xml"
	"time"

	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type RenovacionSustitucion10 struct {
	Version            string              `xml:"Version,attr" bson:"Version" json:"Version"`
	TipoDecreto        string              `xml:"TipoDeDecreto,attr" bson:"TipoDecreto" json:"TipoDecreto"`
	DecretoRenovacion  *DecretoRenovacion  `xml:"DecretoRenovVehicular" bson:"DecretoRenovacion,omitempty" json:"DecretoRenovacion,omitempty"`
	DecretoSustitucion *DecretoSustitucion `xml:"DecretoSustitVehicular" bson:"DecretoSustitucion,omitempty" json:"DecretoSustitucion,omitempty"`
}

type DecretoRenovacion struct {
	VehiculoEnajenado string            `xml:"VehEnaj,attr" bson:"VehiculoEnajenado" json:"VehiculoEnajenado"`
	VehiculosUsados   []VehiculosUsados `xml:"VehiculosUsadosEnajenadoPermAlFab" bson:"VehiculosUsados" json:"VehiculosUsados"`
	VehiculosNuevos   []VehiculosNuevos `xml:"VehiculoNuvoSemEnajenadoFabAlPerm" bson:"VehiculosNuevos" json:"VehiculosNuevos"`
}

type DecretoSustitucion struct {
	VehiculoEnajenado string            `xml:"VehEnaj,attr" bson:"VehiculoEnajenado" json:"VehiculoEnajenado"`
	VehiculosUsados   []VehiculosUsados `xml:"VehiculosUsadosEnajenadoPermAlFab" bson:"VehiculosUsados" json:"VehiculosUsados"`
	VehiculosNuevos   []VehiculosNuevos `xml:"VehiculoNuvoSemEnajenadoFabAlPerm" bson:"VehiculosNuevos" json:"VehiculosNuevos"`
}

type VehiculosUsados struct {
	PrecioVehUsado             float64    `xml:"PrecioVehUsado,attr" bson:"Precio" json:"Precio"`
	Tipo                       string     `xml:"TipoVeh,attr" bson:"Tipo" json:"Tipo"`
	Marca                      string     `xml:"Marca,attr" bson:"Marca" json:"Marca"`
	TipoOClase                 string     `xml:"TipooClase,attr" bson:"TipoOClase" json:"TipoOClase"`
	Anio                       string     `xml:"Año,attr" bson:"Anio" json:"Anio"`
	Modelo                     *string    `xml:"Modelo,attr" bson:"Modelo,omitempty" json:"Modelo,omitempty"`
	NoIdentificacionVehicular  *string    `xml:"NIV,attr" bson:"NoIdentificacionVehicular,omitempty" json:"NoIdentificacionVehicular,omitempty"`
	NoSerie                    *string    `xml:"NumSerie,attr" bson:"NoSerie,omitempty" json:"NoSerie,omitempty"`
	NoPlacas                   string     `xml:"NumPlacas,attr" bson:"NoPlacas" json:"NoPlacas"`
	NoMotor                    *string    `xml:"NumMotor,attr" bson:"NoMotor,omitempty" json:"NoMotor,omitempty"`
	NoFolioTarjetaCirculacion  string     `xml:"NumFolTarjCir,attr" bson:"NoFolioTarjetaCirculacion" json:"NoFolioTarjetaCirculacion"`
	NumeroPedimentoImportacion *string    `xml:"NumPedIm,attr" bson:"NumeroPedimentoImportacion,omitempty" json:"NumeroPedimentoImportacion,omitempty"`
	Aduana                     *string    `xml:"Aduana,attr" bson:"Aduana,omitempty" json:"Aduana,omitempty"`
	FechaRegulVeh              *string    `xml:"FechaRegulVeh,attr" bson:"FechaRegulVeh" json:"FechaRegulVeh"`
	FechaRegulacionImportacion *time.Time `bson:"FechaRegulacionImportacion,omitempty" json:"FechaRegulacionImportacion,omitempty"`
	FolioFiscal                string     `xml:"Foliofiscal,attr" bson:"FolioFiscal" json:"FolioFiscal"`
}

type VehiculosNuevos struct {
	Anio     string  `xml:"Año,attr" bson:"Anio" json:"Anio"`
	Modelo   *string `xml:"Modelo,attr" bson:"Modelo,omitempty" json:"Modelo,omitempty"`
	NoPlacas string  `xml:"NumPlacas,attr" bson:"NoPlacas" json:"NoPlacas"`
	Rfc      *string `xml:"RFC,attr" bson:"Rfc,omitempty" json:"Rfc,omitempty"`
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
