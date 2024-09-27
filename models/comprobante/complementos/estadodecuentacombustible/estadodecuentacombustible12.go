package estadodecuentacombustible

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"time"
)

type EstadoDeCuentaCombustible12 struct {
	Version       string                          `xml:"Version,attr" bson:"Version"`
	TipoOperacion string                          `xml:"TipoOperacion,attr" bson:"TipoOperacion"`
	NoCuenta      string                          `xml:"NumeroDeCuenta,attr" bson:"NoCuenta"`
	Subtotal      float64                         `xml:"SubTotal,attr" bson:"Subtotal"`
	Total         float64                         `xml:"Total,attr" bson:"Total"`
	Conceptos     ConceptosEstadoDeCuentaCombus12 `xml:"Conceptos" bson:"Conceptos"`
}

type ConceptosEstadoDeCuentaCombus12 struct {
	Concepto []ConceptoEstadoDeCuentaCombustible12 `xml:"ConceptoEstadoDeCuentaCombustible" bson:"Concepto"`
}

type ConceptoEstadoDeCuentaCombustible12 struct {
	Identificador     string                          `xml:"Identificador,attr" bson:"Identificador"`
	FechaString       string                          `xml:"Fecha,attr"`
	Fecha             time.Time                       `bson:"Fecha"`
	Rfc               string                          `xml:"Rfc,attr" bson:"Rfc"`
	ClaveEstacion     string                          `xml:"ClaveEstacion,attr" bson:"ClaveEstacion"`
	Cantidad          string                          `xml:"Cantidad,attr" bson:"Cantidad"`
	TipoCombustible   string                          `xml:"TipoCombustible,attr" bson:"TipoCombustible"`
	NoIdentificacion  string                          `xml:"NoIdentificacion,attr" bson:"NoIdentificacion"`
	Unidad            *string                         `xml:"Unidad,attr" bson:"Unidad,omitempty"`
	NombreCombustible string                          `xml:"NombreCombustible,attr" bson:"NombreCombustible"`
	FolioOperacion    string                          `xml:"FolioOperacion,attr" bson:"FolioOperacion"`
	ValorUnitario     float64                         `xml:"ValorUnitario,attr" bson:"ValorUnitario"`
	Importe           float64                         `xml:"Importe,attr" bson:"Importe"`
	Traslados         TrasladosEstadoDeCuentaCombus12 `xml:"Traslados" bson:"Traslados"`
}

type TrasladosEstadoDeCuentaCombus12 struct {
	Traslado []TrasladoEstadoDeCuentaCombus12 `xml:"Traslado" bson:"Traslado"`
}

type TrasladoEstadoDeCuentaCombus12 struct {
	Impuesto   string  `xml:"Impuesto,attr" bson:"Impuesto"`
	TasaOCuota float64 `xml:"TasaOCuota,attr" bson:"TasaOCuota"`
	Importe    float64 `xml:"Importe,attr" bson:"Importe"`
}

func (c *ConceptoEstadoDeCuentaCombustible12) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias ConceptoEstadoDeCuentaCombustible12
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	fechaAutorizacion, err := helpers.ParseDatetime(aux.FechaString)
	if err != nil {
		return err
	}

	*c = ConceptoEstadoDeCuentaCombustible12(aux)
	c.Fecha = fechaAutorizacion

	return nil
}
