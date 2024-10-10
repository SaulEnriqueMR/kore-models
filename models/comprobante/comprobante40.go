package comprobante

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/SaulEnriqueMR/kore-models/models/documentofiscaldigital"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type Comprobante40 struct {
	documentofiscaldigital.DocumentoFiscalDigital `bson:",inline"`
	Version                                       string                 `xml:"Version,attr" bson:"Version"`
	Serie                                         *string                `xml:"Serie,attr" bson:"Serie,omitempty"`
	Folio                                         *string                `xml:"Folio,attr" bson:"Folio,omitempty"`
	Fecha                                         string                 `xml:"Fecha,attr"`
	Sello                                         string                 `xml:"Sello,attr" bson:"Sello"`
	FormaPago                                     *string                `xml:"FormaPago,attr" bson:"FormaPago,omitempty"`
	NoCertificado                                 string                 `xml:"NoCertificado,attr" bson:"NoCertificado"`
	Certificado                                   string                 `xml:"Certificado,attr" bson:"Certificado"`
	CondicionesPago                               *string                `xml:"CondicionesDePago,attr" bson:"CondicionesPago,omitempty"`
	Subtotal                                      float64                `xml:"SubTotal,attr" bson:"Subtotal"`
	Descuento                                     *float64               `xml:"Descuento,attr" bson:"Descuento,omitempty"`
	Moneda                                        string                 `xml:"Moneda,attr" bson:"Moneda"`
	TipoCambio                                    *float64               `xml:"TipoCambio,attr" bson:"TipoCambio,omitempty"`
	Total                                         float64                `xml:"Total,attr" bson:"Total"`
	TipoComprobante                               string                 `xml:"TipoDeComprobante,attr" bson:"TipoComprobante"`
	Exportacion                                   string                 `xml:"Exportacion,attr" bson:"Exportacion"`
	MetodoPago                                    *string                `xml:"MetodoPago,attr" bson:"MetodoPago,omitempty"`
	LugarExpedicion                               string                 `xml:"LugarExpedicion,attr" bson:"LugarExpedicion"`
	Confirmacion                                  *string                `xml:"Confirmacion,attr" bson:"Confirmacion,omitempty"`
	InformacionGlobal                             *InformacionGlobal40   `xml:"InformacionGlobal" bson:"InformacionGlobal,omitempty"`
	CfdisRelacionados                             *[]CfdisRelacionados40 `xml:"CfdiRelacionados" bson:"CfdisRelacionados,omitempty"`
	Emisor                                        Emisor40               `xml:"Emisor" bson:"Emisor"`
	Receptor                                      Receptor40             `xml:"Receptor" bson:"Receptor"`
	RfcProvCertif                                 string                 `bson:"RfcProvCertif"`
	Conceptos                                     []Concepto40           `xml:"Conceptos>Concepto" bson:"Conceptos"`
	Impuestos                                     *Impuestos40           `xml:"Impuestos" bson:"Impuestos,omitempty"`
	Complemento                                   Complemento            `xml:"Complemento" bson:"Complemento"`
}

type InformacionGlobal40 struct {
	Periodicidad string `xml:"Periodicidad,attr" bson:"Periodicidad"`
	Meses        string `xml:"Meses,attr" bson:"Meses"`
	Anio         string `xml:"Año,attr" bson:"Anio"`
}

type CfdisRelacionados40 struct {
	TipoRelacion     string              `xml:"TipoRelacion,attr" bson:"TipoRelacion"`
	UuidRelacionados []UuidRelacionado40 `xml:"CfdiRelacionado" bson:"UuidsRelacionados"`
}

type UuidRelacionado40 struct {
	Uuid string `xml:"UUID,attr" bson:"Uuid"`
}

type Emisor40 struct {
	Rfc              string  `xml:"Rfc,attr" bson:"Rfc"`
	Nombre           string  `xml:"Nombre,attr" bson:"Nombre"`
	RegimenFiscal    string  `xml:"RegimenFiscal,attr" bson:"RegimenFiscal"`
	FactAtrAdquirent *string `xml:"FacAtrAdquirente,attr" bson:"FacAtrAdquirent,omitempty"`
}

type Receptor40 struct {
	Rfc              string  `xml:"Rfc,attr" bson:"Rfc"`
	Nombre           string  `xml:"Nombre,attr" bson:"Nombre"`
	DomicilioFiscal  string  `xml:"DomicilioFiscalReceptor,attr" bson:"DomicilioFiscal"`
	ResidenciaFiscal *string `xml:"ResidenciaFiscal,attr" bson:"ResidenciaFiscal,omitempty"`
	NumRegIdTrib     *string `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty"`
	RegimenFiscal    string  `xml:"RegimenFiscalReceptor,attr" bson:"RegimenFiscal"`
	UsoCFDI          string  `xml:"UsoCFDI,attr" bson:"UsoCfdi"`
}

type Concepto40 struct {
	ClaveProdServ       string                   `xml:"ClaveProdServ,attr" bson:"ClaveProdServ"`
	NoIdentificacion    *string                  `xml:"NoIdentificacion,attr" bson:"NoIdentificacion,omitempty"`
	Cantidad            float64                  `xml:"Cantidad,attr" bson:"Cantidad"`
	ClaveUnidad         string                   `xml:"ClaveUnidad,attr" bson:"ClaveUnidad"`
	Unidad              *string                  `xml:"Unidad,attr" bson:"Unidad,omitempty"`
	Descripcion         string                   `xml:"Descripcion,attr" bson:"Descripcion"`
	ValorUnitario       float64                  `xml:"ValorUnitario,attr" bson:"ValorUnitario"`
	Importe             float64                  `xml:"Importe,attr" bson:"Importe"`
	Descuento           *float64                 `xml:"Descuento,attr" bson:"Descuento,omitempty"`
	ObjetoImpuesto      string                   `xml:"ObjetoImp,attr" bson:"ObjetoImpuesto"`
	Impuestos           *ImpuestosConcepto40     `xml:"Impuestos" bson:"Impuestos,omitempty"`
	ACuentaTerceros     *ACuentaTerceros40       `xml:"ACuentaTerceros" bson:"ACuentaTerceros,omitempty"`
	InformacionAduanera *[]InformacionAduanera40 `xml:"InformacionAduanera" bson:"InformacionAduanera,omitempty"`
	CuentaPredial       *[]CuentaPredial40       `xml:"CuentaPredial" bson:"CuentaPredial,omitempty"`
	ComplementoConcepto *ComplementoConcepto     `xml:"ComplementoConcepto" bson:"ComplementoConcepto,omitempty"`
	Parte               *[]Parte40               `xml:"Parte" bson:"Parte,omitempty"`
}

type ImpuestosConcepto40 struct {
	Traslados   *[]TrasladoConcepto40  `xml:"Traslados>Traslado" bson:"Traslados,omitempty"`
	Retenciones *[]RetencionConcepto40 `xml:"Retenciones>Retencion" bson:"Retenciones,omitempty"`
}

type TrasladoConcepto40 struct {
	Base       float64  `xml:"Base,attr" bson:"Base"`
	Impuesto   string   `xml:"Impuesto,attr" bson:"Impuesto"`
	TipoFactor string   `xml:"TipoFactor,attr" bson:"TipoFactor"`
	TasaOCuota *float64 `xml:"TasaOCuota,attr" bson:"TasaOCuota,omitempty"`
	Importe    *float64 `xml:"Importe,attr" bson:"Importe,omitempty"`
}

type RetencionConcepto40 struct {
	Base       float64 `xml:"Base,attr" bson:"Base"`
	Impuesto   string  `xml:"Impuesto,attr" bson:"Impuesto"`
	TipoFactor string  `xml:"TipoFactor,attr" bson:"TipoFactor"`
	TasaOCuota float64 `xml:"TasaOCuota,attr" bson:"TasaOCuota"`
	Importe    float64 `xml:"Importe,attr" bson:"Importe"`
}

type ACuentaTerceros40 struct {
	Rfc             string `xml:"RfcACuentaTerceros,attr" bson:"Rfc"`
	Nombre          string `xml:"NombreACuentaTerceros,attr" bson:"Nombre"`
	RegimenFiscal   string `xml:"RegimenFiscalACuentaTerceros,attr" bson:"RegimenFiscal"`
	DomicilioFiscal string `xml:"DomicilioFiscalACuentaTerceros,attr" bson:"DomicilioFiscal"`
}

type InformacionAduanera40 struct {
	NumeroPedimento string `xml:"NumeroPedimento,attr" bson:"NumeroPedimento"`
}

type CuentaPredial40 struct {
	Numero string `xml:"Numero,attr" bson:"Numero"`
}

type Parte40 struct {
	ClaveProdServ       string                   `xml:"ClaveProdServ,attr" bson:"ClaveProdServ"`
	NoIdentificacion    *string                  `xml:"NoIdentificacion,attr" bson:"NoIdentificacion,omitempty"`
	Cantidad            float64                  `xml:"Cantidad,attr" bson:"Cantidad"`
	Unidad              *string                  `xml:"Unidad,attr" bson:"Unidad,omitempty"`
	Descripcion         string                   `xml:"Descripcion,attr" bson:"Descripcion"`
	ValorUnitario       *float64                 `xml:"ValorUnitario,attr" bson:"ValorUnitario,omitempty"`
	Importe             *float64                 `xml:"Importe,attr" bson:"Importe,omitempty"`
	InformacionAduanera *[]InformacionAduanera40 `xml:"InformacionAduanera" bson:"InformacionAduanera,omitempty"`
}

type Impuestos40 struct {
	TotalImpuestosRetenidos   *float64                `xml:"TotalImpuestosRetenidos,attr" bson:"TotalImpuestosRetenidos,omitempty"`
	TotalImpuestosTrasladados *float64                `xml:"TotalImpuestosTrasladados,attr" bson:"TotalImpuestosTrasladados,omitempty"`
	Retenciones               *[]RetencionImpuestos40 `xml:"Retenciones>Retencion" bson:"Retenciones,omitempty"`
	Traslados                 *[]TrasladoImpuestos40  `xml:"Traslados>Traslado" bson:"Traslados,omitempty"`
}

type RetencionImpuestos40 struct {
	Impuesto string  `xml:"Impuesto,attr" bson:"Impuesto"`
	Importe  float64 `xml:"Importe,attr" bson:"Importe"`
}

type TrasladoImpuestos40 struct {
	Base       float64  `xml:"Base,attr" bson:"Base"`
	Impuesto   string   `xml:"Impuesto,attr" bson:"Impuesto"`
	TipoFactor string   `xml:"TipoFactor,attr" bson:"TipoFactor"`
	TasaOCuota *float64 `xml:"TasaOCuota,attr" bson:"TasaOCuota,omitempty"`
	Importe    *float64 `xml:"Importe,attr" bson:"Importe,omitempty"`
}

func (c *Comprobante40) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias Comprobante40
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	fechaEmision, err := helpers.ParseDatetime(aux.Fecha)
	if err != nil {
		return err
	}

	*c = Comprobante40(aux)
	c.FechaEmision = fechaEmision
	c.Comprobante = true
	c.Vigente = nil

	if c.InformacionAdicional != nil {
		c.InformacionAdicional.StampedByKuantik = nil
	}
	if c.Cancelacion != nil {
		c.Cancelacion.CanceledByKuantik = nil
	}

	if c.Complemento.TimbreFiscalDigital != nil {
		tfd := c.Complemento.TimbreFiscalDigital.TimbreFiscalDigital11
		if tfd != nil {
			c.FechaTimbrado = tfd.FechaTimbrado
			c.Uuid = strings.ToUpper(tfd.Uuid)
		}
		rfcProvCertif := tfd.RfcProvCertif
		if len(rfcProvCertif) > 0 {
			c.RfcProvCertif = rfcProvCertif
		}
	}
	c.CadenaOriginal = helpers.CreateCadenaOriginal(c)

	return nil
}

func (c *Comprobante40) GetFileName() string {
	year := fmt.Sprint(c.FechaEmision.Year())
	month := fmt.Sprint(int(c.FechaEmision.Month()))
	return c.Emisor.Rfc + "/" + c.Receptor.Rfc + "/" + year + "/" + month + "/" + c.Uuid + ".xml"
}
