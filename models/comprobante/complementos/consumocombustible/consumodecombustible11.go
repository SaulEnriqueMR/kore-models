package consumocombustible

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"time"
)

type ConsumoDeCombustible11 struct {
	Version       string                     `xml:"version,attr" bson:"Version" json:"Version"`
	TipoOperacion string                     `xml:"tipoOperacion,attr" bson:"TipoOperacion" json:"TipoOperacion"`
	NoCuenta      string                     `xml:"numeroDeCuenta,attr" bson:"NoCuenta" json:"NoCuenta"`
	Subtotal      *float64                   `xml:"subTotal,attr" bson:"Subtotal,omitempty" json:"Subtotal,omitempty"`
	Total         float64                    `xml:"total,attr" bson:"Total" json:"Total"`
	Conceptos     ConceptosConsumoDeCombus11 `xml:"Conceptos" bson:"Conceptos,omitempty" json:"Conceptos,omitempty"`
}

type ConceptosConsumoDeCombus11 struct {
	Concepto []ConceptoConsumoDeCombustibles11 `xml:"ConceptoConsumoDeCombustibles" bson:"Concepto" json:"Concepto"`
}

type ConceptoConsumoDeCombustibles11 struct {
	Identificador     string                        `xml:"identificador,attr" bson:"Identificador" json:"Identificador"`
	FechaString       string                        `xml:"fecha,attr" bson:"FechaString" json:"FechaString"`
	Fecha             time.Time                     `bson:"Fecha" json:"Fecha"`
	Rfc               string                        `xml:"rfc,attr" bson:"Rfc" json:"Rfc"`
	ClaveEstacion     string                        `xml:"claveEstacion,attr" bson:"ClaveEstacion" json:"ClaveEstacion"`
	TipoCombustible   string                        `xml:"tipoCombustible,attr" bson:"TipoCombustible" json:"TipoCombustible"`
	Cantidad          float64                       `xml:"cantidad,attr" bson:"Cantidad" json:"Cantidad"`
	NombreCombustible string                        `xml:"nombreCombustible,attr" bson:"NombreCombustible" json:"NombreCombustible"`
	FolioOperacion    string                        `xml:"folioOperacion,attr" bson:"FolioOperacion" json:"FolioOperacion"`
	ValorUnitario     float64                       `xml:"valorUnitario,attr" bson:"ValorUnitario," json:"ValorUnitario"`
	Importe           float64                       `xml:"importe,attr" bson:"Importe" json:"Importe"`
	Determinados      DeterminadosConsumoDeCombus11 `xml:"Determinados" bson:"Determinados" json:"Determinados"`
}

type DeterminadosConsumoDeCombus11 struct {
	Determinado []DeterminadoConsumoDeCombus11 `xml:"Determinado" bson:"Determinado" json:"Determinado"`
}

type DeterminadoConsumoDeCombus11 struct {
	Impuesto   string  `xml:"impuesto,attr" bson:"Impuesto" json:"Impuesto"`
	TasaOCuota float64 `xml:"tasa,attr" bson:"Tasa" json:"Tasa"`
	Importe    float64 `xml:"importe,attr" bson:"Importe" json:"Importe"`
}

func (c *ConceptoConsumoDeCombustibles11) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias ConceptoConsumoDeCombustibles11
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	fecha, err := helpers.ParseDatetime(aux.FechaString)
	if err != nil {
		return err
	}

	*c = ConceptoConsumoDeCombustibles11(aux)
	c.Fecha = fecha

	return nil
}
