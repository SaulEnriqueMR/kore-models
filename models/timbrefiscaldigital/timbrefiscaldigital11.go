package timbrefiscaldigital

import (
	"encoding/xml"
	"github.com/saulenriquemr/koremodels/models"
	"time"
)

//TIP Namespace es <b>tfd</b>.

// TimbreFiscalDigital11 Versión 1.1 del complemento TimbreFiscalDigital.
// Aplica para versión Comprobante 3.3 y 4.0.
type TimbreFiscalDigital11 struct {
	XMLName xml.Name `xml:"TimbreFiscalDigital"`
	// Debería siempre estar fijo a 1.0
	Version string `xml:"Version,attr" bson:"Version"`
	Uuid    string `xml:"UUID,attr" bson:"Uuid"`
	// Extracción a nivel string, debe convertise después para poder ser filtrable
	Fecha string `xml:"FechaTimbrado,attr"`
	// Conversión de campo XML FechaTimbrado
	FechaTimbrado time.Time `bson:"FechaTimbrado"`
	// PAC que provee el certificado
	RfcProvCertif    string  `xml:"RfcProvCertif,attr" bson:"RfcProvCertif"`
	Leyenda          *string `xml:"Leyenda,attr" bson:"Leyenda,omitempty"`
	SelloCfd         string  `xml:"SelloCFD,attr" bson:"SelloCFD"`
	NoCertificadoSat string  `xml:"NoCertificadoSAT,attr" bson:"NoCertificadoSAT"`
	SelloSat         string  `xml:"SelloSAT,attr" bson:"SelloSAT"`
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
	fechaTimbrado, err := models.ParseDatetime(aux.Fecha)
	if err != nil {
		return err
	}

	// Assign the parsed date and other fields back to the original struct
	*t = TimbreFiscalDigital11(aux)
	t.FechaTimbrado = fechaTimbrado

	return nil
}
