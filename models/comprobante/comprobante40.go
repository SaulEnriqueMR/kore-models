package comprobante

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"

	"github.com/SaulEnriqueMR/kore-models/models"

	"github.com/SaulEnriqueMR/kore-models/models/documentofiscaldigital"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type Comprobante40 struct {
	documentofiscaldigital.DocumentoFiscalDigital `bson:",inline"`
	Version                                       string                 `xml:"Version,attr" bson:"Version" json:"Version"`
	Serie                                         *string                `xml:"Serie,attr" bson:"Serie,omitempty" json:"Serie"`
	Folio                                         *string                `xml:"Folio,attr" bson:"Folio,omitempty" json:"Folio"`
	Fecha                                         string                 `xml:"Fecha,attr" bson:"Fecha" json:"Fecha"`
	Sello                                         string                 `xml:"Sello,attr" bson:"Sello"`
	FormaPago                                     *string                `xml:"FormaPago,attr" bson:"FormaPago,omitempty" json:"FormaPago"`
	NoCertificado                                 string                 `xml:"NoCertificado,attr" bson:"NoCertificado" json:"NoCertificado"`
	Certificado                                   string                 `xml:"Certificado,attr" bson:"Certificado" json:"Certificado"`
	CondicionesPago                               *string                `xml:"CondicionesDePago,attr" bson:"CondicionesPago,omitempty" json:"CondicionesPago"`
	Subtotal                                      float64                `xml:"SubTotal,attr" bson:"Subtotal" json:"Subtotal"`
	Descuento                                     *float64               `xml:"Descuento,attr" bson:"Descuento,omitempty" json:"Descuento"`
	Moneda                                        string                 `xml:"Moneda,attr" bson:"Moneda" json:"Moneda"`
	TipoCambio                                    *float64               `xml:"TipoCambio,attr" bson:"TipoCambio,omitempty" json:"TipoCambio"`
	Total                                         float64                `xml:"Total,attr" bson:"Total" json:"Total"`
	TipoComprobante                               string                 `xml:"TipoDeComprobante,attr" bson:"TipoComprobante" json:"TipoComprobante"`
	Exportacion                                   string                 `xml:"Exportacion,attr" bson:"Exportacion" json:"Exportacion"`
	MetodoPago                                    *string                `xml:"MetodoPago,attr" bson:"MetodoPago,omitempty" json:"MetodoPago"`
	LugarExpedicion                               string                 `xml:"LugarExpedicion,attr" bson:"LugarExpedicion" json:"LugarExpedicion"`
	Confirmacion                                  *string                `xml:"Confirmacion,attr" bson:"Confirmacion,omitempty" json:"Confirmacion"`
	InformacionGlobal                             *InformacionGlobal40   `xml:"InformacionGlobal" bson:"InformacionGlobal,omitempty" json:"InformacionGlobal"`
	CfdisRelacionados                             *[]CfdisRelacionados40 `xml:"CfdiRelacionados" bson:"CfdisRelacionados,omitempty" json:"CfdisRelacionados"`
	Emisor                                        Emisor40               `xml:"Emisor" bson:"Emisor" json:"Emisor"`
	Receptor                                      Receptor40             `xml:"Receptor" bson:"Receptor" json:"Receptor"`
	RfcProvCertif                                 string                 `bson:"RfcProvCertif"`
	Conceptos                                     []Concepto40           `xml:"Conceptos>Concepto" bson:"Conceptos" json:"Conceptos"`
	Impuestos                                     *Impuestos40           `xml:"Impuestos" bson:"Impuestos,omitempty" json:"Impuestos"`
	Complemento                                   Complemento            `xml:"Complemento" bson:"Complemento" json:"Complemento"`
	Addenda                                       *Addenda               `xml:"Addenda" bson:"Addenda,omitempty"`
}

func (c *Comprobante40) DefineTransaccion(rfc string) {
	if c.Receptor.Rfc == rfc {
		c.Transaccion = "RECIBIDO"
	}
	if c.Emisor.Rfc == rfc {
		c.Transaccion = "EMITIDO"
	}
	if c.Receptor.Rfc == rfc && c.Emisor.Rfc == rfc {
		c.Transaccion = "AUTOFACTURACION"
	}
}

type InformacionGlobal40 struct {
	Periodicidad string `xml:"Periodicidad,attr" bson:"Periodicidad" json:"Periodicidad"`
	Meses        string `xml:"Meses,attr" bson:"Meses" json:"Meses"`
	Anio         string `xml:"Año,attr" bson:"Anio" json:"Anio"`
}

type CfdisRelacionados40 struct {
	TipoRelacion      string              `xml:"TipoRelacion,attr" bson:"TipoRelacion" json:"TipoRelacion"`
	UuidsRelacionados []UuidRelacionado40 `xml:"CfdiRelacionado" bson:"UuidsRelacionados" json:"UuidsRelacionados"`
}

type UuidRelacionado40 struct {
	Uuid string `bson:"Uuid" json:"Uuid"`
	UUID string `xml:"UUID,attr" bson:"UUID" json:"UUID"`
}

func (ur *UuidRelacionado40) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias UuidRelacionado40
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	*ur = UuidRelacionado40(aux)
	ur.Uuid = strings.ToUpper(aux.UUID)

	return nil
}

type Emisor40 struct {
	Rfc              string  `xml:"Rfc,attr" bson:"Rfc" json:"Rfc"`
	Nombre           string  `xml:"Nombre,attr" bson:"Nombre" json:"Nombre"`
	RegimenFiscal    string  `xml:"RegimenFiscal,attr" bson:"RegimenFiscal" json:"RegimenFiscal"`
	FactAtrAdquirent *string `xml:"FacAtrAdquirente,attr" bson:"FacAtrAdquirent,omitempty" json:"FactAtrAdquirent"`
}

type Receptor40 struct {
	Rfc              string  `xml:"Rfc,attr" bson:"Rfc" json:"Rfc"`
	Nombre           string  `xml:"Nombre,attr" bson:"Nombre" json:"Nombre"`
	DomicilioFiscal  string  `xml:"DomicilioFiscalReceptor,attr" bson:"DomicilioFiscal" json:"DomicilioFiscal"`
	ResidenciaFiscal *string `xml:"ResidenciaFiscal,attr" bson:"ResidenciaFiscal,omitempty" json:"ResidenciaFiscal"`
	NumRegIdTrib     *string `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty" json:"NumRegIdTrib"`
	RegimenFiscal    string  `xml:"RegimenFiscalReceptor,attr" bson:"RegimenFiscal" json:"RegimenFiscal"`
	UsoCFDI          string  `xml:"UsoCFDI,attr" bson:"UsoCfdi" json:"UsoCFDI"`
}

type Concepto40 struct {
	ClaveProdServ       string                   `xml:"ClaveProdServ,attr" bson:"ClaveProdServ" json:"ClaveProductoServicio"`
	NoIdentificacion    *string                  `xml:"NoIdentificacion,attr" bson:"NoIdentificacion,omitempty" json:"NoIdentificacion"`
	Cantidad            float64                  `xml:"Cantidad,attr" bson:"Cantidad" json:"Cantidad" json:"Cantidad"`
	ClaveUnidad         string                   `xml:"ClaveUnidad,attr" bson:"ClaveUnidad" json:"ClaveUnidad"`
	Unidad              *string                  `xml:"Unidad,attr" bson:"Unidad,omitempty" json:"Unidad"`
	Descripcion         string                   `xml:"Descripcion,attr" bson:"Descripcion" json:"Descripcion"`
	ValorUnitario       float64                  `xml:"ValorUnitario,attr" bson:"ValorUnitario" json:"ValorUnitario"`
	Importe             float64                  `xml:"Importe,attr" bson:"Importe" json:"Importe"`
	Descuento           *float64                 `xml:"Descuento,attr" bson:"Descuento,omitempty" json:"Descuento"`
	ObjetoImpuesto      string                   `xml:"ObjetoImp,attr" bson:"ObjetoImpuesto" json:"ObjetoImpuesto"`
	Impuestos           *ImpuestosConcepto40     `xml:"Impuestos" bson:"Impuestos,omitempty" json:"Impuestos"`
	ACuentaTerceros     *ACuentaTerceros40       `xml:"ACuentaTerceros" bson:"ACuentaTerceros,omitempty" json:"ACuentaTerceros"`
	InformacionAduanera *[]InformacionAduanera40 `xml:"InformacionAduanera" bson:"InformacionAduanera,omitempty" json:"InformacionAduanera"`
	CuentaPredial       *[]CuentaPredial40       `xml:"CuentaPredial" bson:"CuentaPredial,omitempty" json:"CuentaPredial"`
	ComplementoConcepto *ComplementoConcepto     `xml:"ComplementoConcepto" bson:"ComplementoConcepto,omitempty" json:"ComplementoConcepto"`
	Parte               *[]Parte40               `xml:"Parte" bson:"Parte,omitempty" json:"Parte"`
}

type ImpuestosConcepto40 struct {
	Traslados   *[]TrasladoConcepto40  `xml:"Traslados>Traslado" bson:"Traslados,omitempty" json:"Traslados"`
	Retenciones *[]RetencionConcepto40 `xml:"Retenciones>Retencion" bson:"Retenciones,omitempty" json:"Retenciones"`
}

type TrasladoConcepto40 struct {
	Base       float64  `xml:"Base,attr" bson:"Base" json:"Base"`
	Impuesto   string   `xml:"Impuesto,attr" bson:"Impuesto" json:"Impuesto"`
	TipoFactor string   `xml:"TipoFactor,attr" bson:"TipoFactor" json:"TipoFactor"`
	TasaOCuota *float64 `xml:"TasaOCuota,attr" bson:"TasaOCuota,omitempty" json:"TasaOCuota"`
	Importe    *float64 `xml:"Importe,attr" bson:"Importe,omitempty" json:"Importe"`
}

type RetencionConcepto40 struct {
	Base       float64 `xml:"Base,attr" bson:"Base" json:"Base"`
	Impuesto   string  `xml:"Impuesto,attr" bson:"Impuesto" json:"Impuesto"`
	TipoFactor string  `xml:"TipoFactor,attr" bson:"TipoFactor" json:"TipoFactor"`
	TasaOCuota float64 `xml:"TasaOCuota,attr" bson:"TasaOCuota" json:"TasaOCuota"`
	Importe    float64 `xml:"Importe,attr" bson:"Importe" json:"Importe"`
}

type ACuentaTerceros40 struct {
	Rfc             string `xml:"RfcACuentaTerceros,attr" bson:"Rfc" json:"Rfc"`
	Nombre          string `xml:"NombreACuentaTerceros,attr" bson:"Nombre" json:"Nombre"`
	RegimenFiscal   string `xml:"RegimenFiscalACuentaTerceros,attr" bson:"RegimenFiscal" json:"RegimenFiscal"`
	DomicilioFiscal string `xml:"DomicilioFiscalACuentaTerceros,attr" bson:"DomicilioFiscal" json:"DomicilioFiscal"`
}

type InformacionAduanera40 struct {
	NumeroPedimento string `xml:"NumeroPedimento,attr" bson:"NumeroPedimento" json:"NumeroPedimento"`
}

type CuentaPredial40 struct {
	Numero string `xml:"Numero,attr" bson:"Numero" json:"Numero"`
}

type Parte40 struct {
	ClaveProdServ       string                   `xml:"ClaveProdServ,attr" bson:"ClaveProdServ" json:"ClaveProdServ"`
	NoIdentificacion    *string                  `xml:"NoIdentificacion,attr" bson:"NoIdentificacion,omitempty" json:"NoIdentificacion"`
	Cantidad            float64                  `xml:"Cantidad,attr" bson:"Cantidad" json:"Cantidad"`
	Unidad              *string                  `xml:"Unidad,attr" bson:"Unidad,omitempty" json:"Unidad"`
	Descripcion         string                   `xml:"Descripcion,attr" bson:"Descripcion" json:"Descripcion"`
	ValorUnitario       *float64                 `xml:"ValorUnitario,attr" bson:"ValorUnitario,omitempty" json:"ValorUnitario"`
	Importe             *float64                 `xml:"Importe,attr" bson:"Importe,omitempty" json:"Importe"`
	InformacionAduanera *[]InformacionAduanera40 `xml:"InformacionAduanera" bson:"InformacionAduanera,omitempty" json:"InformacionAduanera"`
}

type Impuestos40 struct {
	TotalImpuestosRetenidos   *float64                `xml:"TotalImpuestosRetenidos,attr" bson:"TotalImpuestosRetenidos,omitempty" json:"TotalImpuestosRetenidos"`
	TotalImpuestosTrasladados *float64                `xml:"TotalImpuestosTrasladados,attr" bson:"TotalImpuestosTrasladados,omitempty" json:"TotalImpuestosTrasladados"`
	Retenciones               *[]RetencionImpuestos40 `xml:"Retenciones>Retencion" bson:"Retenciones,omitempty" json:"Retenciones"`
	Traslados                 *[]TrasladoImpuestos40  `xml:"Traslados>Traslado" bson:"Traslados,omitempty" json:"Traslados"`
}

type RetencionImpuestos40 struct {
	Impuesto string  `xml:"Impuesto,attr" bson:"Impuesto" json:"Impuesto"`
	Importe  float64 `xml:"Importe,attr" bson:"Importe" json:"Importe"`
}

type TrasladoImpuestos40 struct {
	Base       float64  `xml:"Base,attr" bson:"Base" json:"Base"`
	Impuesto   string   `xml:"Impuesto,attr" bson:"Impuesto" json:"Impuesto"`
	TipoFactor string   `xml:"TipoFactor,attr" bson:"TipoFactor" json:"TipoFactor"`
	TasaOCuota *float64 `xml:"TasaOCuota,attr" bson:"TasaOCuota,omitempty" json:"TasaOCuota"`
	Importe    *float64 `xml:"Importe,attr" bson:"Importe,omitempty" json:"Importe"`
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
			rfcProvCertif := tfd.RfcProvCertif
			if len(rfcProvCertif) > 0 {
				c.RfcProvCertif = rfcProvCertif
			}
		}
	}

	c.CadenaOriginal = helpers.CreateCadenaOriginal(*c)

	tipoCambio := 1.0
	// Calculo de totales
	if c.TipoCambio != nil {
		tipoCambio = *c.TipoCambio
	}

	totalesMonedaLocal := documentofiscaldigital.TotalesMonedaLocal{
		Total:    c.Total * tipoCambio,
		Subtotal: c.Subtotal * tipoCambio,
	}

	if c.Descuento != nil {
		descuento := *c.Descuento * tipoCambio
		totalesMonedaLocal.Descuento = &descuento
	}

	if c.Impuestos != nil {
		if c.Impuestos.TotalImpuestosRetenidos != nil {
			tir := *c.Impuestos.TotalImpuestosRetenidos * tipoCambio
			totalesMonedaLocal.TotalImpuestosRetenidos = &tir
		}

		if c.Impuestos.TotalImpuestosTrasladados != nil {
			tit := *c.Impuestos.TotalImpuestosTrasladados * tipoCambio
			totalesMonedaLocal.TotalImpuestosTrasladados = &tit
		}
	}

	c.TotalesMonedaLocal = totalesMonedaLocal

	// processDate := time.Now().UTC()
	now := time.Now()
	processDate := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, time.UTC)
	c.ProcessorMetadata.LastUpdate = &processDate
	c.ProcessorMetadata.KoreModelsVersion = &models.KoreModelVersion

	return nil
}

func (c *Comprobante40) GetFileName() string {
	year := fmt.Sprint(c.FechaEmision.Year())
	month := fmt.Sprint(int(c.FechaEmision.Month()))
	return c.Emisor.Rfc + "/" + c.Receptor.Rfc + "/" + year + "/" + month + "/" + c.Uuid + ".xml"
}

func (c Comprobante40) GetBasePath() string {
	year := fmt.Sprint(c.FechaEmision.Year())
	month := fmt.Sprint(int(c.FechaEmision.Month()))
	sb := strings.Builder{}
	sb.WriteString(c.Emisor.Rfc)
	sb.WriteString("/")
	sb.WriteString(c.Receptor.Rfc)
	sb.WriteString("/")
	sb.WriteString(year)
	sb.WriteString("/")
	sb.WriteString(month)
	sb.WriteString("/")
	sb.WriteString(c.Uuid)
	return sb.String()
}

/* func (c *Comprobante40) SetFilePaths() {
	basePath := c.GetBasePath()
	c.XmlPath = strings.Join([]string{basePath, "xml"}, ".")
	c.PdfPath = strings.Join([]string{basePath, "pdf"}, ".")
} */
