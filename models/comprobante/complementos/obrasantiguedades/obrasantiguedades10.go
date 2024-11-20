package obrasantiguedades

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"time"
)

type ObrasAntiguedades10 struct {
	Version                   string    `xml:"Version,attr" bson:"Version"`
	TipoBien                  string    `xml:"TipoBien,attr" bson:"TipoBien"`
	OtrosTipoBien             *string   `xml:"OtrosTipoBien,attr" bson:"OtrosTipoBien,omitempty"`
	TituloAdquirido           string    `xml:"TituloAdquirido,attr" bson:"TituloAdquirido"`
	OtrosTitulosAdquiridos    *string   `xml:"OtrosTitulosAdquiridos,attr" bson:"OtrosTitulosAdquiridos,omitempty"`
	Subtotal                  float64   `xml:"Subtotal,attr" bson:"Subtotal"`
	Iva                       float64   `xml:"IVA,attr" bson:"Iva"`
	FechaAdquisicionString    string    `xml:"FechaAdquisicion,attr" bson:"FechaAdquisicionString"`
	FechaAdquisicion          time.Time `bson:"FechaAdquisicion"`
	CaracteristicasObraOPieza string    `xml:"CaracterísticasDeObraoPieza,attr" bson:"CaracteristicasObraOPieza"`
}

func (t *ObrasAntiguedades10) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias ObrasAntiguedades10
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	// Parse the date from the 'Fecha' field
	fechaAdquisicion, err := helpers.ParseDatetime(aux.FechaAdquisicionString)
	if err != nil {
		return err
	}

	// Assign the parsed date and other fields back to the original struct
	*t = ObrasAntiguedades10(aux)
	t.FechaAdquisicion = fechaAdquisicion

	return nil
}
