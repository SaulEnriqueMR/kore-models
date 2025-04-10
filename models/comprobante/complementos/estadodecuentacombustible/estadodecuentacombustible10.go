package estadodecuentacombustible

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"time"
)

type EstadoDeCuentaCombustible10 struct {
	TipoOperacion string                          `xml:"tipoOperacion,attr" bson:"TipoOperacion"`
	NoCuenta      string                          `xml:"numeroDeCuenta,attr" bson:"NoCuenta"`
	Subtotal      *float64                        `xml:"subTotal,attr" bson:"Subtotal,omitempty"`
	Total         float64                         `xml:"total,attr" bson:"Total"`
	Conceptos     ConceptosEstadoDeCuentaCombus10 `xml:"Conceptos" bson:"Conceptos"`
}

type ConceptosEstadoDeCuentaCombus10 struct {
	Concepto []ConceptoEstadoDeCuentaCombustible10 `xml:"ConceptoEstadoDeCuentaCombustible" bson:"Concepto"`
}

type ConceptoEstadoDeCuentaCombustible10 struct {
	Identificador     string                               `xml:"identificador,attr" bson:"Identificador"`
	FechaString       string                               `xml:"fecha,attr" bson:"FechaString"`
	Fecha             time.Time                            `bson:"Fecha"`
	Rfc               string                               `xml:"rfc,attr" bson:"Rfc"`
	ClaveEstacion     string                               `xml:"claveEstacion,attr" bson:"ClaveEstacion"`
	Cantidad          string                               `xml:"cantidad,attr" bson:"Cantidad"`
	NombreCombustible string                               `xml:"nombreCombustible,attr" bson:"NombreCombustible"`
	FolioOperacion    string                               `xml:"folioOperacion,attr" bson:"FolioOperacion"`
	ValorUnitario     float64                              `xml:"valorUnitario,attr" bson:"ValorUnitario"`
	Importe           float64                              `xml:"importe,attr" bson:"Importe"`
	Traslados         TrasladosEstadoDeCuentaCombustible10 `xml:"Traslados" bson:"Traslados"`
}

func (c *ConceptoEstadoDeCuentaCombustible10) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias ConceptoEstadoDeCuentaCombustible10
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	fechaAutorizacion, err := helpers.ParseDatetime(aux.FechaString)
	if err != nil {
		return err
	}

	*c = ConceptoEstadoDeCuentaCombustible10(aux)
	c.Fecha = fechaAutorizacion

	return nil
}

type TrasladosEstadoDeCuentaCombustible10 struct {
	Traslado []TrasladoEstadoDeCuentaCombustible10 `xml:"Traslado" bson:"Traslado"`
}

type TrasladoEstadoDeCuentaCombustible10 struct {
	Impuesto   string  `xml:"impuesto,attr" bson:"Impuesto"`
	TasaOCuota float64 `xml:"tasa,attr" bson:"TasaOCuota"`
	Importe    float64 `xml:"importe,attr" bson:"Importe"`
}
