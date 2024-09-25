package comprobante

import (
	"fmt"
	"time"
)

type Comprobante33 struct {
	Version          string              `xml:"Version,attr" bson:"Version"`
	Serie            *string             `xml:"Serie,attr" bson:"Serie,omitempty"`
	Folio            *string             `xml:"Folio,attr" bson:"Folio,omitempty"`
	Fecha            string              `xml:"Fecha,attr"`
	Sello            string              `xml:"Sello,attr" bson:"Sello"`
	FormaPago        *string             `xml:"FormaPago,attr" bson:"FormaPago,omitempty"`
	NoCertificado    string              `xml:"NoCertificado,attr" bson:"NoCertificado"`
	Certificado      string              `xml:"Certificado,attr" bson:"Certificado"`
	CondicionesPago  *string             `xml:"CondicionesDePago,attr" bson:"CondicionesPago,omitempty"`
	SubTotal         float64             `xml:"SubTotal,attr" bson:"Subtotal"`
	Descuento        *float64            `xml:"Descuento,attr" bson:"Descuento,omitempty"`
	Moneda           string              `xml:"Moneda,attr" bson:"Moneda"`
	TipoCambio       *float64            `xml:"TipoCambio,attr" bson:"TipoCambio,omitempty"`
	Total            float64             `xml:"Total,attr" bson:"Total"`
	TipoComprobante  string              `xml:"TipoDeComprobante,attr" bson:"TipoComprobante"`
	MetodoPago       *string             `xml:"MetodoPago,attr" bson:"MetodoPago,omitempty"`
	LugarExpedicion  string              `xml:"LugarExpedicion,attr" bson:"LugarExpedicion"` // Cifrado
	Confirmacion     *string             `xml:"Confirmacion,attr" bson:"Confirmacion,omitempty"`
	CfdiRelacionados *CfdiRelacionados33 `xml:"CfdiRelacionados" bson:"CfdiRelacionados,omitempty"`
	Emisor           Emisor33            `xml:"Emisor" bson:"Emisor"`
	Receptor         Receptor33          `xml:"Receptor" bson:"Receptor"`
	Conceptos        []Concepto33        `xml:"Conceptos>Concepto" bson:"Conceptos"`
	Impuestos        *Impuestos33        `xml:"Impuestos" bson:"Impuestos,omitempty"`
	Complemento      *Complemento        `xml:"Complemento" bson:"Complemento,omitempty"`
	/* Atributo convertido */
	FechaEmision time.Time `bson:"FechaEmision"`
	/* Atributos extraidos desde tfd */
	Uuid          string    `bson:"Uuid"`
	FechaTimbrado time.Time `bson:"FechaTimbrado"`
	/* Atributos adicionales, generalmente actualizados por fuentes externas */
	InformacionAdicional InformacionAdicional `xml:"InformacionAdicional" bson:"InformacionAdicional"`
	Cancelacion          Cancelacion          `xml:"Cancelacion" bson:"Cancelacion"`
	Vigente              bool                 `bson:"Vigente"`
	/* Cadena original */
	CadenaOriginal string `bson:"CadenaOriginal"`
}

type CfdiRelacionados33 struct {
	TipoRelacion    string              `xml:"TipoRelacion,attr" bson:"TipoRelacion"`
	CfdiRelacionado []CfdiRelacionado33 `xml:"CfdiRelacionado" bson:"CfdiRelacionado"`
}

type CfdiRelacionado33 struct {
	UUID string `xml:"UUID,attr" bson:"Uuid"`
}

type Emisor33 struct {
	Rfc           string  `xml:"Rfc,attr" bson:"Rfc"`                 // Cifrado
	Nombre        *string `xml:"Nombre,attr" bson:"Nombre,omitempty"` // Cifrado
	RegimenFiscal string  `xml:"RegimenFiscal,attr" bson:"RegimenFiscal"`
}

type Receptor33 struct {
	Rfc              string  `xml:"Rfc,attr" bson:"Rfc"`                                     // Cifrado
	Nombre           *string `xml:"Nombre,attr" bson:"Nombre,omitempty"`                     // Cifrado
	ResidenciaFiscal *string `xml:"ResidenciaFiscal,attr" bson:"ResidenciaFiscal,omitempty"` // Cifrado
	NumRegIdTrib     *string `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty"`         // Cifrado
	UsoCFDI          string  `xml:"UsoCFDI,attr" bson:"UsoCFDI"`
}

type Concepto33 struct {
	ClaveProdServ       string                   `xml:"ClaveProdServ,attr" bson:"ClaveProdServ"`
	NoIdentificacion    *string                  `xml:"NoIdentificacion,attr" bson:"NoIdentificacion,omitempty"`
	Cantidad            float64                  `xml:"Cantidad,attr" bson:"Cantidad"`
	ClaveUnidad         string                   `xml:"ClaveUnidad,attr" bson:"ClaveUnidad"`
	Unidad              *string                  `xml:"Unidad,attr" bson:"Unidad,omitempty"`
	Descripcion         string                   `xml:"Descripcion,attr" bson:"Descripcion"`
	ValorUnitario       float64                  `xml:"ValorUnitario,attr" bson:"ValorUnitario"`
	Importe             float64                  `xml:"Importe,attr" bson:"Importe"`
	Descuento           *float64                 `xml:"Descuento,attr" bson:"Descuento,omitempty"`
	Impuestos           *ImpuestosConcepto33     `xml:"Impuestos" bson:"Impuestos,omitempty"`
	InformacionAduanera *[]InformacionAduanera33 `xml:"InformacionAduanera" bson:"InformacionAduanera,omitempty"`
	CuentaPredial       *CuentaPredial33         `xml:"CuentaPredial" bson:"CuentaPredial,omitempty"`
	ComplementoConcepto *ComplementoConcepto     `xml:"ComplementoConcepto" bson:"ComplementoConcepto,omitempty"`
	Parte               *[]Parte33               `xml:"Parte" bson:"Parte,omitempty"`
	DetallesConcepto    string                   `xml:"DetallesConcepto" bson:"DetallesConcepto"`
}

type ImpuestosConcepto33 struct {
	Traslados   *[]TrasladoConcepto33  `xml:"Traslados>Traslado" bson:"Traslados,omitempty"`
	Retenciones *[]RetencionConcepto33 `xml:"Retenciones>Retencion" bson:"Retenciones,omitempty"`
}

type TrasladoConcepto33 struct {
	Base       float64  `xml:"Base,attr" bson:"Base"`
	Impuesto   string   `xml:"Impuesto,attr" bson:"Impuesto"`
	TipoFactor string   `xml:"TipoFactor,attr" bson:"TipoFactor"`
	TasaOCuota *float64 `xml:"TasaOCuota,attr" bson:"TasaOCuota,omitempty"`
	Importe    *float64 `xml:"Importe,attr" bson:"Importe,omitempty"`
}

type RetencionConcepto33 struct {
	Base       float64 `xml:"Base,attr" bson:"Base"`
	Impuesto   string  `xml:"Impuesto,attr" bson:"Impuesto"`
	TipoFactor string  `xml:"TipoFactor,attr" bson:"TipoFactor"`
	TasaOCuota float64 `xml:"TasaOCuota,attr" bson:"TasaOCuota"`
	Importe    float64 `xml:"Importe,attr" bson:"Importe"`
}

type InformacionAduanera33 struct {
	NumeroPedimento string `xml:"NumeroPedimento,attr" bson:"NumeroPedimento"`
}

type CuentaPredial33 struct {
	Numero string `xml:"Numero,attr" bson:"Numero"`
}

type Parte33 struct {
	ClaveProdServ       string                   `xml:"ClaveProdServ,attr" bson:"ClaveProdServ"`
	NoIdentificacion    *string                  `xml:"NoIdentificacion,attr" bson:"NoIdentificacion,omitempty"`
	Cantidad            float64                  `xml:"Cantidad,attr" bson:"Cantidad"`
	Unidad              *string                  `xml:"Unidad,attr" bson:"Unidad,omitempty"`
	Descripcion         string                   `xml:"Descripcion,attr" bson:"Descripcion"`
	ValorUnitario       *float64                 `xml:"ValorUnitario,attr" bson:"ValorUnitario,omitempty"`
	Importe             *float64                 `xml:"Importe,attr" bson:"Importe,omitempty"`
	InformacionAduanera *[]InformacionAduanera33 `xml:"InformacionAduanera" bson:"InformacionAduanera,omitempty"`
}

type Impuestos33 struct {
	TotalImpuestosRetenidos   *float64                `xml:"TotalImpuestosRetenidos,attr" bson:"TotalImpuestosRetenidos,omitempty"`
	TotalImpuestosTrasladados *float64                `xml:"TotalImpuestosTrasladados,attr" bson:"TotalImpuestosTrasladados,omitempty"`
	Retenciones               *[]RetencionImpuestos33 `xml:"Retenciones>Retencion" bson:"Retenciones,omitempty"`
	Traslados                 *[]TrasladoImpuestos33  `xml:"Traslados>Traslado" bson:"Traslados,omitempty"`
}

type RetencionImpuestos33 struct {
	Impuesto string  `xml:"Impuesto,attr" bson:"Impuesto"`
	Importe  float64 `xml:"Importe,attr" bson:"Importe"`
}

type TrasladoImpuestos33 struct {
	Impuesto   string  `xml:"Impuesto,attr" bson:"Impuesto"`
	TipoFactor string  `xml:"TipoFactor,attr" bson:"TipoFactor"`
	TasaOCuota float64 `xml:"TasaOCuota,attr" bson:"TasaOCuota"`
	Importe    float64 `xml:"Importe,attr" bson:"Importe"`
}

func (c *Comprobante33) GetFileName() string {
	year := fmt.Sprint(c.FechaEmision.Year())
	month := fmt.Sprint(int(c.FechaEmision.Month()))
	return c.Emisor.Rfc + "/" + c.Receptor.Rfc + "/" + year + "/" + month + "/" + c.Uuid + ".xml"
}
