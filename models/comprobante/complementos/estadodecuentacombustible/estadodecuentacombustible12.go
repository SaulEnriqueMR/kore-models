package estadodecuentacombustible

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"time"
)

type EstadoDeCuentaCombustible12 struct {
	Version       string                          `xml:"Version,attr" bson:"Version" json:"Version"`
	TipoOperacion string                          `xml:"TipoOperacion,attr" bson:"TipoOperacion" json:"TipoOperacion"`
	NoCuenta      string                          `xml:"NumeroDeCuenta,attr" bson:"NoCuenta" json:"NoCuenta"`
	Subtotal      float64                         `xml:"SubTotal,attr" bson:"Subtotal" json:"Subtotal"`
	Total         float64                         `xml:"Total,attr" bson:"Total" json:"Total"`
	Conceptos     ConceptosEstadoDeCuentaCombus12 `xml:"Conceptos" bson:"Conceptos" json:"Conceptos"`
}

type ConceptosEstadoDeCuentaCombus12 struct {
	Concepto []ConceptoEstadoDeCuentaCombustible12 `xml:"ConceptoEstadoDeCuentaCombustible" bson:"Concepto" json:"Concepto"`
}

type ConceptoEstadoDeCuentaCombustible12 struct {
	Identificador     string                          `xml:"Identificador,attr" bson:"Identificador" json:"Identificador"`
	FechaString       string                          `xml:"Fecha,attr" bson:"FechaString" json:"FechaString"`
	Fecha             time.Time                       `bson:"Fecha" json:"Fecha"`
	Rfc               string                          `xml:"Rfc,attr" bson:"Rfc" json:"Rfc"`
	ClaveEstacion     string                          `xml:"ClaveEstacion,attr" bson:"ClaveEstacion" json:"ClaveEstacion"`
	Cantidad          string                          `xml:"Cantidad,attr" bson:"Cantidad" json:"Cantidad"`
	TipoCombustible   string                          `xml:"TipoCombustible,attr" bson:"TipoCombustible" json:"TipoCombustible"`
	NoIdentificacion  string                          `xml:"NoIdentificacion,attr" bson:"NoIdentificacion" json:"NoIdentificacion"`
	Unidad            *string                         `xml:"Unidad,attr" bson:"Unidad,omitempty" json:"Unidad,omitempty"`
	NombreCombustible string                          `xml:"NombreCombustible,attr" bson:"NombreCombustible" json:"NombreCombustible"`
	FolioOperacion    string                          `xml:"FolioOperacion,attr" bson:"FolioOperacion" json:"FolioOperacion"`
	ValorUnitario     float64                         `xml:"ValorUnitario,attr" bson:"ValorUnitario" json:"ValorUnitario"`
	Importe           float64                         `xml:"Importe,attr" bson:"Importe" json:"Importe"`
	Traslados         TrasladosEstadoDeCuentaCombus12 `xml:"Traslados" bson:"Traslados" json:"Traslados"`
}

type TrasladosEstadoDeCuentaCombus12 struct {
	Traslado []TrasladoEstadoDeCuentaCombus12 `xml:"Traslado" bson:"Traslado" json:"Traslado"`
}

type TrasladoEstadoDeCuentaCombus12 struct {
	Impuesto   string  `xml:"Impuesto,attr" bson:"Impuesto" json:"Impuesto"`
	TasaOCuota float64 `xml:"TasaOCuota,attr" bson:"TasaOCuota" json:"TasaOCuota"`
	Importe    float64 `xml:"Importe,attr" bson:"Importe" json:"Importe"`
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
