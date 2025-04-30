package certificadodestruccion

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"time"
)

type CertificadoDeDestruccion10 struct {
	Version                  string                           `xml:"Version,attr" bson:"Version" json:"Version"`
	Serie                    string                           `xml:"Serie,attr" bson:"Serie" json:"Serie"`
	FolioDestruccionVehiculo string                           `xml:"NumFolDesVeh,attr" bson:"FolioDestruccionVehiculo" json:"FolioDestruccionVehiculo"`
	VehiculoDestruido        *[]VehiculoDestruidoCertDest10   `xml:"VehiculoDestruido" bson:"VehiculoDestruido" json:"VehiculoDestruido"`
	InformacinoAduanera      *[]InformacionAduaneraCertDest10 `xml:"InformacionAduanera" bson:"InformacionAduanera" json:"InformacionAduanera"`
}

type VehiculoDestruidoCertDest10 struct {
	Marca                     string  `xml:"Marca,attr" bson:"Marca" json:"Marca"`
	TipoOClase                string  `xml:"TipooClase,attr" bson:"TipoOClase" json:"TipoOClase"`
	Anio                      string  `xml:"Año,attr" bson:"Anio" json:"Anio"`
	Modelo                    *string `xml:"Modelo,attr" bson:"Modelo,omitempty" json:"Modelo,omitempty"`
	NoIdentificacionVehicular *string `xml:"NIV,attr" bson:"NoIdentificacionVehicular,omitempty" json:"NoIdentificacionVehicular,omitempty"`
	NoSerie                   *string `xml:"NumSerie,attr" bson:"NoSerie,omitempty" json:"NoSerie,omitempty"`
	NoPlacas                  string  `xml:"NumPlacas,attr" bson:"NoPlacas" json:"NoPlacas"`
	NoMotor                   *string `xml:"NumMotor,attr" bson:"NoMotor,omitempty" json:"NoMotor,omitempty"`
	FolioTarjetaCirculacion   string  `xml:"NumFolTarjCir,attr" bson:"FolioTarjetaCirculacion" json:"FolioTarjetaCirculacion"`
}

type InformacionAduaneraCertDest10 struct {
	NumeroPedimentoImportacion string    `xml:"NumPedImp,attr" bson:"NumeroPedimentoImportacion" json:"NumeroPedimentoImportacion"`
	FechaString                string    `xml:"Fecha,attr" bson:"FechaString" json:"FechaString"`
	Fecha                      time.Time `bson:"Fecha" json:"Fecha"`
	Aduana                     string    `xml:"Aduana,attr" bson:"Aduana" json:"Aduana"`
}

func (iacd10 *InformacionAduaneraCertDest10) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias InformacionAduaneraCertDest10
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*iacd10 = InformacionAduaneraCertDest10(aux)

	if iacd10 != nil {
		fecha, err := helpers.ParseDatetime(iacd10.FechaString)
		if err != nil {
			return err
		}
		iacd10.Fecha = fecha
	}

	return nil
}
