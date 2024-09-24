package timbrefiscaldigital

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models"
	"time"
)

//TIP Namespace es <b>tfd</b>.

// TimbreFiscalDigital10 Versión 1.0 del complemento TimbreFiscalDigital.
// Solo aplica para versión Comprobante 3.2.
type TimbreFiscalDigital10 struct {
	XMLName xml.Name `xml:"TimbreFiscalDigital"`
	// Debería siempre estar fijo a 1.0
	Version string `xml:"version,attr" bson:"Version"`
	Uuid    string `xml:"UUID,attr" bson:"Uuid"`
	// Extracción a nivel string, debe convertise después para poder ser filtrable
	Fecha string `xml:"FechaTimbrado,attr"`
	// Conversión de campo XML FechaTimbrado
	FechaTimbrado    time.Time `bson:"FechaTimbrado"`
	SelloCfd         string    `xml:"selloCFD,attr" bson:"SelloCFD"`
	NoCertificadoSat string    `xml:"noCertificadoSAT,attr" bson:"NoCertificadoSAT"`
	SelloSat         string    `xml:"selloSAT,attr" bson:"SelloSAT"`
}

func (t *TimbreFiscalDigital10) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias TimbreFiscalDigital10
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
	*t = TimbreFiscalDigital10(aux)
	t.FechaTimbrado = fechaTimbrado

	return nil
}
