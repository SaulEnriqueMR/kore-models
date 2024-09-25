package obrasantiguedades

import (
	"encoding/xml"
	"time"

	"github.com/SaulEnriqueMR/kore-models/models"
)

type ObrasAntiguedades10 struct {
	Version                     string    `xml:"Version,attr" bson:"version"`
	TipoBien                    string    `xml:"TipoBien,attr" bson:"TipoBien"`
	OtrosTipoBien               *string   `xml:"OtrosTipoBien,attr" bson:"OtrosTipoBien,omitempty"`
	TituloAdquirido             string    `xml:"TituloAdquirido,attr" bson:"TituloAdquirido"`
	OtrosTitulosAdquiridos      *string   `xml:"OtrosTitulosAdquiridos,attr" bson:"OtrosTitulosAdquiridos,omitempty"`
	Subtotal                    float64   `xml:"Subtotal,attr" bson:"Subtotal"`
	IVA                         float64   `xml:"IVA,attr" bson:"IVA"`
	FechaAdquisicionString      string    `xml:"FechaAdquisicion,attr"`
	FechaAdquisicion            time.Time `bson:"FechaAdquisicion"`
	CaracterísticasDeObraoPieza string    `xml:"CaracterísticasDeObraoPieza,attr" bson:"CaracteristicasDeObraoPieza"`
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
	fechaAdquisicion, err := models.ParseDatetime(aux.FechaAdquisicionString)
	if err != nil {
		return err
	}

	// Assign the parsed date and other fields back to the original struct
	*t = ObrasAntiguedades10(aux)
	t.FechaAdquisicion = fechaAdquisicion

	return nil
}