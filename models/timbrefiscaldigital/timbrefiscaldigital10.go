package timbrefiscaldigital

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"time"
)

//TIP Namespace es <b>tfd</b>.

// TimbreFiscalDigital10 Versión 1.0 del complemento TimbreFiscalDigital.
// Solo aplica para versión Comprobante 3.2.
type TimbreFiscalDigital10 struct {
	// Debería siempre estar fijo a 1.0
	Version string `xml:"version,attr" bson:"Version" json:"Version"`
	Uuid    string `xml:"UUID,attr" bson:"Uuid" json:"Uuid"`
	// Extracción a nivel string, debe convertise después para poder ser filtrable
	Fecha string `xml:"FechaTimbrado,attr" json:"Fecha"`
	// Conversión de campo XML FechaTimbrado
	FechaTimbrado    time.Time `bson:"FechaTimbrado" json:"FechaTimbrado"`
	SelloCfd         string    `xml:"selloCFD,attr" bson:"SelloCfd" json:"SelloCfd"`
	NoCertificadoSat string    `xml:"noCertificadoSAT,attr" bson:"NoCertificadoSat" json:"NoCertificadoSat"`
	SelloSat         string    `xml:"selloSAT,attr" bson:"SelloSat" json:"SelloSat"`
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
	fechaTimbrado, err := helpers.ParseDatetime(aux.Fecha)
	if err != nil {
		return err
	}

	// Assign the parsed date and other fields back to the original struct
	*t = TimbreFiscalDigital10(aux)
	t.FechaTimbrado = fechaTimbrado

	return nil
}
