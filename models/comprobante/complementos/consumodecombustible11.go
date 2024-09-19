package complementos

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models"
	"time"
)

type ConsumoDeCombustible11 struct {
	Version        string                     `xml:"version,attr" bson:"Version"`
	TipoOperacion  string                     `xml:"tipoOperacion,attr" bson:"TipoOperacion"`
	NumeroDeCuenta string                     `xml:"numeroDeCuenta,attr" bson:"NumeroDeCuenta"`
	SubTotal       *float64                   `xml:"subTotal,attr" bson:"SubTotal,omitempty"`
	Total          float64                    `xml:"total,attr" bson:"Total"`
	Conceptos      ConceptosConsumoDeCombus11 `xml:"Conceptos" bson:"Conceptos,omitempty"`
}

type ConceptosConsumoDeCombus11 struct {
	ConceptoConsumoDeCombustibles []ConceptoConsumoDeCombustibles11 `xml:"ConceptoConsumoDeCombustibles" bson:"ConceptoConsumoDeCombustibles"`
}

type ConceptoConsumoDeCombustibles11 struct {
	Identificador     string                        `xml:"identificador,attr" bson:"identificador"`
	FechaString       string                        `xml:"fecha,attr"`
	Fecha             time.Time                     `bson:"Fecha"`
	Rfc               string                        `xml:"rfc,attr" bson:"Rfc"`
	ClaveEstacion     string                        `xml:"claveEstacion,attr" bson:"ClaveEstacion"`
	TipoCombustible   string                        `xml:"tipoCombustible,attr" bson:"TipoCombustible"`
	Cantidad          float64                       `xml:"cantidad,attr" bson:"Cantidad"`
	NombreCombustible string                        `xml:"nombreCombustible,attr" bson:"NombreCombustible"`
	FolioOperacion    string                        `xml:"folioOperacion,attr" bson:"FolioOperacion"`
	ValorUnitario     float64                       `xml:"valorUnitario,attr" bson:"ValorUnitario,"`
	Importe           float64                       `xml:"importe,attr" bson:"Importe"`
	Determinados      DeterminadosConsumoDeCombus11 `xml:"Determinados" bson:"Determinados"`
}

type DeterminadosConsumoDeCombus11 struct {
	Determinado []DeterminadoConsumoDeCombus11 `xml:"Determinado" bson:"Determinado"`
}

type DeterminadoConsumoDeCombus11 struct {
	Impuesto   string  `xml:"impuesto,attr" bson:"Impuesto"`
	TasaOCuota float64 `xml:"tasa,attr" bson:"Tasa"`
	Importe    float64 `xml:"importe,attr" bson:"Importe"`
}

func (c *ConceptoConsumoDeCombustibles11) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias ConceptoConsumoDeCombustibles11
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	fecha, err := models.ParseDatetime(aux.FechaString)
	if err != nil {
		return err
	}

	*c = ConceptoConsumoDeCombustibles11(aux)
	c.Fecha = fecha

	return nil
}
