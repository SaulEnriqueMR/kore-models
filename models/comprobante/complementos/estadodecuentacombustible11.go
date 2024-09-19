package complementos

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models"
	"time"
)

type EstadoDeCuentaCombustible11 struct {
	Version        string                          `xml:"Version,attr" bson:"Version"`
	TipoOperacion  string                          `xml:"TipoOperacion,attr" bson:"TipoOperacion"`
	NumeroDeCuenta string                          `xml:"NumeroDeCuenta,attr" bson:"NumeroDeCuenta"`
	SubTotal       float64                         `xml:"SubTotal,attr" bson:"SubTotal"`
	Total          float64                         `xml:"Total,attr" bson:"Total"`
	Conceptos      ConceptosEstadoDeCuentaCombus11 `xml:"Conceptos" bson:"Conceptos"`
}

type ConceptosEstadoDeCuentaCombus11 struct {
	ConceptoEstadoDeCuentaCombustible []ConceptoEstadoDeCuentaCombustible11 `xml:"ConceptoEstadoDeCuentaCombustible" bson:"ConceptoEstadoDeCuentaCombustible"`
}

type ConceptoEstadoDeCuentaCombustible11 struct {
	Identificador     string                          `xml:"Identificador,attr" bson:"Identificador"`
	FechaString       string                          `xml:"Fecha,attr"`
	Fecha             time.Time                       `bson:"Fecha"`
	Rfc               string                          `xml:"Rfc,attr" bson:"Rfc"`
	ClaveEstacion     string                          `xml:"ClaveEstacion,attr" bson:"ClaveEstacion"`
	Tar               *string                         `xml:"TAR,attr" bson:"Tar,omitempty"`
	Cantidad          string                          `xml:"Cantidad,attr" bson:"Cantidad"`
	NoIdentificacion  string                          `xml:"NoIdentificacion,attr" bson:"NoIdentificacion"`
	Unidad            *string                         `xml:"Unidad,attr" bson:"Unidad,omitempty"`
	NombreCombustible string                          `xml:"NombreCombustible,attr" bson:"NombreCombustible"`
	FolioOperacion    string                          `xml:"FolioOperacion,attr" bson:"FolioOperacion"`
	ValorUnitario     float64                         `xml:"ValorUnitario,attr" bson:"ValorUnitario"`
	Importe           float64                         `xml:"Importe,attr" bson:"Importe"`
	Traslados         TrasladosEstadoDeCuentaCombus11 `xml:"Traslados" bson:"Traslados"`
}

type TrasladosEstadoDeCuentaCombus11 struct {
	Traslado []TrasladoEstadoDeCuentaCombus11 `xml:"Traslado" bson:"Traslado"`
}

type TrasladoEstadoDeCuentaCombus11 struct {
	Impuesto   string  `xml:"Impuesto,attr" bson:"Impuesto"`
	TasaoCuota float64 `xml:"TasaoCuota,attr" bson:"TasaoCuota"`
	Importe    float64 `xml:"Importe,attr" bson:"Importe"`
}

func (c *ConceptoEstadoDeCuentaCombustible11) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias ConceptoEstadoDeCuentaCombustible11
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	fechaAutorizacion, err := models.ParseDatetime(aux.FechaString)
	if err != nil {
		return err
	}

	*c = ConceptoEstadoDeCuentaCombustible11(aux)
	c.Fecha = fechaAutorizacion

	return nil
}
