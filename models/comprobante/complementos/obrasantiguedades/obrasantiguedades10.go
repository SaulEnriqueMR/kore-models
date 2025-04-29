package obrasantiguedades

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"time"
)

type ObrasAntiguedades10 struct {
	Version                   string    `xml:"Version,attr" bson:"Version" json:"Version"`
	TipoBien                  string    `xml:"TipoBien,attr" bson:"TipoBien" json:"TipoBien"`
	OtrosTipoBien             *string   `xml:"OtrosTipoBien,attr" bson:"OtrosTipoBien,omitempty" json:"OtrosTipoBien,omitempty"`
	TituloAdquirido           string    `xml:"TituloAdquirido,attr" bson:"TituloAdquirido" json:"TituloAdquirido"`
	OtrosTitulosAdquiridos    *string   `xml:"OtrosTitulosAdquiridos,attr" bson:"OtrosTitulosAdquiridos,omitempty" json:"OtrosTitulosAdquiridos,omitempty"`
	Subtotal                  float64   `xml:"Subtotal,attr" bson:"Subtotal" json:"Subtotal"`
	Iva                       float64   `xml:"IVA,attr" bson:"Iva" json:"Iva"`
	FechaAdquisicionString    string    `xml:"FechaAdquisicion,attr" bson:"FechaAdquisicionString" json:"FechaAdquisicionString"`
	FechaAdquisicion          time.Time `bson:"FechaAdquisicion" json:"FechaAdquisicion"`
	CaracteristicasObraOPieza string    `xml:"CaracterísticasDeObraoPieza,attr" bson:"CaracteristicasObraOPieza" json:"CaracteristicasObraOPieza"`
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
