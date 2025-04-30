package timbrefiscaldigital

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"time"
)

//TIP Namespace es <b>tfd</b>.

// TimbreFiscalDigital11 Versión 1.1 del complemento TimbreFiscalDigital.
// Aplica para versión Comprobante 3.3 y 4.0.
type TimbreFiscalDigital11 struct {
	// Debería siempre estar fijo a 1.0
	Version string `xml:"Version,attr" bson:"Version" json:"Version"`
	Uuid    string `xml:"UUID,attr" bson:"Uuid" json:"Uuid" xml:"Uuid"`
	// Extracción a nivel string, debe convertise después para poder ser filtrable
	Fecha string `xml:"FechaTimbrado,attr" json:"Fecha"`
	// Conversión de campo XML FechaTimbrado
	FechaTimbrado time.Time `bson:"FechaTimbrado" json:"FechaTimbrado"`
	// PAC que provee el certificado
	RfcProvCertif    string  `xml:"RfcProvCertif,attr" bson:"RfcProvCertif" json:"RfcProvCertif"`
	Leyenda          *string `xml:"Leyenda,attr" bson:"Leyenda,omitempty" json:"Leyenda,omitempty"`
	SelloCfd         string  `xml:"SelloCFD,attr" bson:"SelloCfd" json:"SelloCfd"`
	NoCertificadoSat string  `xml:"NoCertificadoSAT,attr" bson:"NoCertificadoSat" json:"NoCertificadoSat"`
	SelloSat         string  `xml:"SelloSAT,attr" bson:"SelloSat" json:"SelloSat"`
}

func (t *TimbreFiscalDigital11) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	// Create an alias to avoid recursion
	type Alias TimbreFiscalDigital11
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	// Parse the date from the 'Fecha' field
	fechaTimbrado, err := helpers.ParseDatetime(aux.Fecha)
	if err != nil {
		return err
	}

	// Assign the parsed date and other fields back to the original struct
	*t = TimbreFiscalDigital11(aux)
	t.FechaTimbrado = fechaTimbrado

	return nil
}
