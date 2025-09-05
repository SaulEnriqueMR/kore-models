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

type Comprobante33 struct {
	documentofiscaldigital.DocumentoFiscalDigital `bson:",inline"`
	Version                                       string              `xml:"Version,attr" bson:"Version" json:"Version"`
	Serie                                         *string             `xml:"Serie,attr" bson:"Serie,omitempty" json:"Serie,omitempty"`
	Folio                                         *string             `xml:"Folio,attr" bson:"Folio,omitempty" json:"Folio,omitempty"`
	Fecha                                         string              `xml:"Fecha,attr" bson:"Fecha" json:"Fecha"`
	Sello                                         string              `xml:"Sello,attr" bson:"Sello" json:"Sello"`
	FormaPago                                     *string             `xml:"FormaPago,attr" bson:"FormaPago,omitempty" json:"FormaPago,omitempty"`
	NoCertificado                                 string              `xml:"NoCertificado,attr" bson:"NoCertificado" json:"NoCertificado"`
	Certificado                                   string              `xml:"Certificado,attr" bson:"Certificado" json:"Certificado"`
	CondicionesPago                               *string             `xml:"CondicionesDePago,attr" bson:"CondicionesPago,omitempty" json:"CondicionesPago,omitempty"`
	Subtotal                                      float64             `xml:"SubTotal,attr" bson:"Subtotal" json:"Subtotal"`
	Descuento                                     *float64            `xml:"Descuento,attr" bson:"Descuento,omitempty" json:"Descuento,omitempty"`
	Moneda                                        string              `xml:"Moneda,attr" bson:"Moneda" json:"Moneda"`
	TipoCambio                                    *float64            `xml:"TipoCambio,attr" bson:"TipoCambio,omitempty" json:"TipoCambio,omitempty"`
	Total                                         float64             `xml:"Total,attr" bson:"Total" json:"Total"`
	TipoComprobante                               string              `xml:"TipoDeComprobante,attr" bson:"TipoComprobante" json:"TipoComprobante"`
	MetodoPago                                    *string             `xml:"MetodoPago,attr" bson:"MetodoPago,omitempty" json:"MetodoPago,omitempty"`
	LugarExpedicion                               string              `xml:"LugarExpedicion,attr" bson:"LugarExpedicion" json:"LugarExpedicion"`
	Confirmacion                                  *string             `xml:"Confirmacion,attr" bson:"Confirmacion,omitempty" json:"Confirmacion,omitempty"`
	CfdisRelacionados                             *CfdiRelacionados33 `xml:"CfdiRelacionados" bson:"CfdisRelacionados,omitempty" json:"CfdisRelacionados,omitempty"`
	Emisor                                        Emisor33            `xml:"Emisor" bson:"Emisor" json:"Emisor"`
	Receptor                                      Receptor33          `xml:"Receptor" bson:"Receptor" json:"Receptor"`
	RfcProvCertif                                 string              `bson:"RfcProvCertif" json:"RfcProvCertif"`
	Conceptos                                     []Concepto33        `xml:"Conceptos>Concepto" bson:"Conceptos" json:"Conceptos"`
	Impuestos                                     *Impuestos33        `xml:"Impuestos" bson:"Impuestos,omitempty" json:"Impuestos,omitempty"`
	Complemento                                   *Complemento        `xml:"Complemento" bson:"Complemento,omitempty" json:"Complemento,omitempty"`
	Addenda                                       *Addenda            `xml:"Addenda" bson:"Addenda,omitempty" json:"Addenda,omitempty"`
	KuantikMetadata                               KuantikMetadata     `bson:"KuantikMetadata,omitempty" json:"KuantikMetadata,omitempty"`
}

func (c *Comprobante33) DefineTransaccion(rfc string) {
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

type CfdiRelacionados33 struct {
	TipoRelacion      string              `xml:"TipoRelacion,attr" bson:"TipoRelacion" json:"TipoRelacion"`
	UuidsRelacionados []UuidRelacionado33 `xml:"CfdiRelacionado" bson:"UuidsRelacionados" json:"UuidsRelacionados"`
}

type UuidRelacionado33 struct {
	UUID string `xml:"UUID,attr" bson:"UUID" json:"UUID"`
	Uuid string `bson:"Uuid" json:"Uuid"`
}

func (ur *UuidRelacionado33) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias UuidRelacionado33
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	*ur = UuidRelacionado33(aux)
	ur.Uuid = strings.ToUpper(aux.UUID)

	return nil
}

type Emisor33 struct {
	Rfc           string  `xml:"Rfc,attr" bson:"Rfc" json:"Rfc"`
	Nombre        *string `xml:"Nombre,attr" bson:"Nombre,omitempty" json:"Nombre,omitempty"`
	RegimenFiscal string  `xml:"RegimenFiscal,attr" bson:"RegimenFiscal" json:"RegimenFiscal"`
}

type Receptor33 struct {
	Rfc              string  `xml:"Rfc,attr" bson:"Rfc" json:"Rfc"`                                                            // Cifrado
	Nombre           *string `xml:"Nombre,attr" bson:"Nombre,omitempty" json:"Nombre,omitempty"`                               // Cifrado
	ResidenciaFiscal *string `xml:"ResidenciaFiscal,attr" bson:"ResidenciaFiscal,omitempty" json:"ResidenciaFiscal,omitempty"` // Cifrado
	NumRegIdTrib     *string `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty" json:"NumRegIdTrib,omitempty"`             // Cifrado
	UsoCFDI          string  `xml:"UsoCFDI,attr" bson:"UsoCFDI" json:"UsoCFDI"`
}

type Concepto33 struct {
	ClaveProdServ       string                   `xml:"ClaveProdServ,attr" bson:"ClaveProdServ" json:"ClaveProdServ"`
	NoIdentificacion    *string                  `xml:"NoIdentificacion,attr" bson:"NoIdentificacion,omitempty" json:"NoIdentificacion,omitempty"`
	Cantidad            float64                  `xml:"Cantidad,attr" bson:"Cantidad" json:"Cantidad"`
	ClaveUnidad         string                   `xml:"ClaveUnidad,attr" bson:"ClaveUnidad" json:"ClaveUnidad"`
	Unidad              *string                  `xml:"Unidad,attr" bson:"Unidad,omitempty" json:"Unidad,omitempty"`
	Descripcion         string                   `xml:"Descripcion,attr" bson:"Descripcion" json:"Descripcion"`
	ValorUnitario       float64                  `xml:"ValorUnitario,attr" bson:"ValorUnitario" json:"ValorUnitario"`
	Importe             float64                  `xml:"Importe,attr" bson:"Importe" json:"Importe"`
	Descuento           *float64                 `xml:"Descuento,attr" bson:"Descuento,omitempty" json:"Descuento,omitempty"`
	Impuestos           *ImpuestosConcepto33     `xml:"Impuestos" bson:"Impuestos,omitempty" json:"Impuestos,omitempty"`
	InformacionAduanera *[]InformacionAduanera33 `xml:"InformacionAduanera" bson:"InformacionAduanera,omitempty" json:"InformacionAduanera,omitempty"`
	CuentaPredial       *CuentaPredial33         `xml:"CuentaPredial" bson:"CuentaPredial,omitempty" json:"CuentaPredial,omitempty"`
	ComplementoConcepto *ComplementoConcepto     `xml:"ComplementoConcepto" bson:"ComplementoConcepto,omitempty" json:"ComplementoConcepto,omitempty"`
	Parte               *[]Parte33               `xml:"Parte" bson:"Parte,omitempty" json:"Parte,omitempty"`
	DetallesConcepto    string                   `xml:"DetallesConcepto" bson:"DetallesConcepto" json:"DetallesConcepto"`
}

type ImpuestosConcepto33 struct {
	Traslados   *[]TrasladoConcepto33  `xml:"Traslados>Traslado" bson:"Traslados,omitempty" json:"Traslados,omitempty"`
	Retenciones *[]RetencionConcepto33 `xml:"Retenciones>Retencion" bson:"Retenciones,omitempty" json:"Retenciones,omitempty"`
}

type TrasladoConcepto33 struct {
	Base       float64  `xml:"Base,attr" bson:"Base" json:"Base"`
	Impuesto   string   `xml:"Impuesto,attr" bson:"Impuesto" json:"Impuesto"`
	TipoFactor string   `xml:"TipoFactor,attr" bson:"TipoFactor" json:"TipoFactor"`
	TasaOCuota *float64 `xml:"TasaOCuota,attr" bson:"TasaOCuota,omitempty" json:"TasaOCuota,omitempty"`
	Importe    *float64 `xml:"Importe,attr" bson:"Importe,omitempty" json:"Importe,omitempty"`
}

type RetencionConcepto33 struct {
	Base       float64 `xml:"Base,attr" bson:"Base" json:"Base"`
	Impuesto   string  `xml:"Impuesto,attr" bson:"Impuesto" json:"Impuesto"`
	TipoFactor string  `xml:"TipoFactor,attr" bson:"TipoFactor" json:"TipoFactor"`
	TasaOCuota float64 `xml:"TasaOCuota,attr" bson:"TasaOCuota" json:"TasaOCuota"`
	Importe    float64 `xml:"Importe,attr" bson:"Importe" json:"Importe"`
}

type InformacionAduanera33 struct {
	NumeroPedimento string `xml:"NumeroPedimento,attr" bson:"NumeroPedimento" json:"NumeroPedimento"`
}

type CuentaPredial33 struct {
	Numero string `xml:"Numero,attr" bson:"Numero" json:"Numero"`
}

type Parte33 struct {
	ClaveProdServ       string                   `xml:"ClaveProdServ,attr" bson:"ClaveProdServ" json:"ClaveProdServ"`
	NoIdentificacion    *string                  `xml:"NoIdentificacion,attr" bson:"NoIdentificacion,omitempty" json:"NoIdentificacion,omitempty"`
	Cantidad            float64                  `xml:"Cantidad,attr" bson:"Cantidad" json:"Cantidad"`
	Unidad              *string                  `xml:"Unidad,attr" bson:"Unidad,omitempty" json:"Unidad,omitempty"`
	Descripcion         string                   `xml:"Descripcion,attr" bson:"Descripcion" json:"Descripcion"`
	ValorUnitario       *float64                 `xml:"ValorUnitario,attr" bson:"ValorUnitario,omitempty" json:"ValorUnitario,omitempty"`
	Importe             *float64                 `xml:"Importe,attr" bson:"Importe,omitempty" json:"Importe,omitempty"`
	InformacionAduanera *[]InformacionAduanera33 `xml:"InformacionAduanera" bson:"InformacionAduanera,omitempty" json:"InformacionAduanera,omitempty"`
}

type Impuestos33 struct {
	TotalImpuestosRetenidos   *float64                `xml:"TotalImpuestosRetenidos,attr" bson:"TotalImpuestosRetenidos,omitempty" json:"TotalImpuestosRetenidos,omitempty"`
	TotalImpuestosTrasladados *float64                `xml:"TotalImpuestosTrasladados,attr" bson:"TotalImpuestosTrasladados,omitempty" json:"TotalImpuestosTrasladados,omitempty"`
	Retenciones               *[]RetencionImpuestos33 `xml:"Retenciones>Retencion" bson:"Retenciones,omitempty" json:"Retenciones,omitempty"`
	Traslados                 *[]TrasladoImpuestos33  `xml:"Traslados>Traslado" bson:"Traslados,omitempty" json:"Traslados,omitempty"`
}

type RetencionImpuestos33 struct {
	Impuesto string  `xml:"Impuesto,attr" bson:"Impuesto" json:"Impuesto"`
	Importe  float64 `xml:"Importe,attr" bson:"Importe" json:"Importe"`
}

type TrasladoImpuestos33 struct {
	Impuesto   string  `xml:"Impuesto,attr" bson:"Impuesto" json:"Impuesto"`
	TipoFactor string  `xml:"TipoFactor,attr" bson:"TipoFactor" json:"TipoFactor"`
	TasaOCuota float64 `xml:"TasaOCuota,attr" bson:"TasaOCuota" json:"TasaOCuota"`
	Importe    float64 `xml:"Importe,attr" bson:"Importe" json:"Importe"`
}

func (c *Comprobante33) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias Comprobante33
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	fechaEmision, err := helpers.ParseDatetime(aux.Fecha)
	if err != nil {
		return err
	}

	*c = Comprobante33(aux)
	c.FechaEmision = fechaEmision
	c.Comprobante = true
	c.Vigente = nil
	// c.ProcessorMetadata.CreationDate = nil

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

	sb := strings.Builder{}
	if c.Serie != nil {
		sb.WriteString(*c.Serie)
	}

	sb.WriteString("-")

	if c.Folio != nil {
		sb.WriteString(*c.Folio)
	}

	c.KuantikMetadata.SerieFolio = sb.String()

	return nil
}

func (c *Comprobante33) GetFileName() string {
	year := fmt.Sprint(c.FechaEmision.Year())
	month := fmt.Sprint(int(c.FechaEmision.Month()))
	return c.Emisor.Rfc + "/" + c.Receptor.Rfc + "/" + year + "/" + month + "/" + c.Uuid + ".xml"
}

func (c Comprobante33) GetBasePath() string {
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

/* func (c *Comprobante33) SetFilePaths() {
	basePath := c.GetBasePath()
	c.XmlPath = strings.Join([]string{basePath, "xml"}, ".")
	c.PdfPath = strings.Join([]string{basePath, "pdf"}, ".")
} */
