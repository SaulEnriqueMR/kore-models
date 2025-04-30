package estadodecuentacombustible

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"time"
)

type EstadoDeCuentaCombustible11 struct {
	Version       string                          `xml:"Version,attr" bson:"Version" json:"Version"`
	TipoOperacion string                          `xml:"TipoOperacion,attr" bson:"TipoOperacion" json:"TipoOperacion"`
	NoCuenta      string                          `xml:"NumeroDeCuenta,attr" bson:"NoCuenta" json:"NoCuenta"`
	Subtotal      float64                         `xml:"SubTotal,attr" bson:"Subtotal" json:"Subtotal"`
	Total         float64                         `xml:"Total,attr" bson:"Total" json:"Total"`
	Conceptos     ConceptosEstadoDeCuentaCombus11 `xml:"Conceptos" bson:"Conceptos" json:"Conceptos"`
}

type ConceptosEstadoDeCuentaCombus11 struct {
	Concepto []ConceptoEstadoDeCuentaCombustible11 `xml:"ConceptoEstadoDeCuentaCombustible" bson:"Concepto" json:"Concepto"`
}

type ConceptoEstadoDeCuentaCombustible11 struct {
	Identificador     string                          `xml:"Identificador,attr" bson:"Identificador" json:"Identificador"`
	FechaString       string                          `xml:"Fecha,attr" bson:"FechaString" json:"FechaString"`
	Fecha             time.Time                       `bson:"Fecha" json:"Fecha"`
	Rfc               string                          `xml:"Rfc,attr" bson:"Rfc" json:"Rfc"`
	ClaveEstacion     string                          `xml:"ClaveEstacion,attr" bson:"ClaveEstacion" json:"ClaveEstacion"`
	Tar               *string                         `xml:"TAR,attr" bson:"Tar,omitempty" json:"Tar,omitempty"`
	Cantidad          string                          `xml:"Cantidad,attr" bson:"Cantidad" json:"Cantidad"`
	NoIdentificacion  string                          `xml:"NoIdentificacion,attr" bson:"NoIdentificacion" json:"NoIdentificacion"`
	Unidad            *string                         `xml:"Unidad,attr" bson:"Unidad,omitempty" json:"Unidad,omitempty"`
	NombreCombustible string                          `xml:"NombreCombustible,attr" bson:"NombreCombustible" json:"NombreCombustible"`
	FolioOperacion    string                          `xml:"FolioOperacion,attr" bson:"FolioOperacion" json:"FolioOperacion"`
	ValorUnitario     float64                         `xml:"ValorUnitario,attr" bson:"ValorUnitario" json:"ValorUnitario"`
	Importe           float64                         `xml:"Importe,attr" bson:"Importe" json:"Importe"`
	Traslados         TrasladosEstadoDeCuentaCombus11 `xml:"Traslados" bson:"Traslados" json:"Traslados"`
}

type TrasladosEstadoDeCuentaCombus11 struct {
	Traslado []TrasladoEstadoDeCuentaCombus11 `xml:"Traslado" bson:"Traslado" json:"Traslado"`
}

type TrasladoEstadoDeCuentaCombus11 struct {
	Impuesto   string  `xml:"Impuesto,attr" bson:"Impuesto" json:"Impuesto"`
	TasaOCuota float64 `xml:"TasaoCuota,attr" bson:"TasaOCuota" json:"TasaOCuota"`
	Importe    float64 `xml:"Importe,attr" bson:"Importe" json:"Importe"`
}

func (c *ConceptoEstadoDeCuentaCombustible11) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias ConceptoEstadoDeCuentaCombustible11
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	fechaAutorizacion, err := helpers.ParseDatetime(aux.FechaString)
	if err != nil {
		return err
	}

	*c = ConceptoEstadoDeCuentaCombustible11(aux)
	c.Fecha = fechaAutorizacion

	return nil
}
