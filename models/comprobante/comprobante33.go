package comprobante

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/SaulEnriqueMR/kore-models/models/documentofiscaldigital"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type Comprobante33 struct {
	documentofiscaldigital.DocumentoFiscalDigital `bson:",inline"`
	Version                                       string              `xml:"Version,attr" bson:"Version" json:"Version"`
	Serie                                         *string             `xml:"Serie,attr" bson:"Serie,omitempty" json:"Serie"`
	Folio                                         *string             `xml:"Folio,attr" bson:"Folio,omitempty" json:"Folio"`
	Fecha                                         string              `xml:"Fecha,attr" bson:"Fecha" json:"Fecha"`
	Sello                                         string              `xml:"Sello,attr" bson:"Sello" json:"Sello"`
	FormaPago                                     *string             `xml:"FormaPago,attr" bson:"FormaPago,omitempty" json:"FormaPago"`
	NoCertificado                                 string              `xml:"NoCertificado,attr" bson:"NoCertificado" json:"NoCertificado"`
	Certificado                                   string              `xml:"Certificado,attr" bson:"Certificado" json:"Certificado"`
	CondicionesPago                               *string             `xml:"CondicionesDePago,attr" bson:"CondicionesPago,omitempty" json:"CondicionesPago"`
	Subtotal                                      float64             `xml:"SubTotal,attr" bson:"Subtotal" json:"Subtotal"`
	Descuento                                     *float64            `xml:"Descuento,attr" bson:"Descuento,omitempty" json:"Descuento"`
	Moneda                                        string              `xml:"Moneda,attr" bson:"Moneda" json:"Moneda"`
	TipoCambio                                    *float64            `xml:"TipoCambio,attr" bson:"TipoCambio,omitempty" json:"TipoCambio"`
	Total                                         float64             `xml:"Total,attr" bson:"Total" json:"Total"`
	TipoComprobante                               string              `xml:"TipoDeComprobante,attr" bson:"TipoComprobante" json:"TipoComprobante"`
	MetodoPago                                    *string             `xml:"MetodoPago,attr" bson:"MetodoPago,omitempty" json:"MetodoPago"`
	LugarExpedicion                               string              `xml:"LugarExpedicion,attr" bson:"LugarExpedicion" json:"LugarExpedicion"`
	Confirmacion                                  *string             `xml:"Confirmacion,attr" bson:"Confirmacion,omitempty" json:"Confirmacion"`
	CfdisRelacionados                             *CfdiRelacionados33 `xml:"CfdiRelacionados" bson:"CfdisRelacionados,omitempty" json:"CfdisRelacionados"`
	Emisor                                        Emisor33            `xml:"Emisor" bson:"Emisor" json:"Emisor"`
	Receptor                                      Receptor33          `xml:"Receptor" bson:"Receptor" json:"Receptor"`
	RfcProvCertif                                 string              `bson:"RfcProvCertif" json:"RfcProvCertif"`
	Conceptos                                     []Concepto33        `xml:"Conceptos>Concepto" bson:"Conceptos" json:"Conceptos"`
	Impuestos                                     *Impuestos33        `xml:"Impuestos" bson:"Impuestos,omitempty" json:"Impuestos"`
	Complemento                                   *Complemento        `xml:"Complemento" bson:"Complemento,omitempty" json:"Complemento"`
}

func (c *Comprobante33) DefineTransaccion(rfc string) {
	if c.Receptor.Rfc == rfc {
		c.Transaccion = "RECIBIDO"
	}
	if c.Emisor.Rfc == rfc {
		c.Transaccion = "EMITIDO"
	}
}

type CfdiRelacionados33 struct {
	TipoRelacion     string              `xml:"TipoRelacion,attr" bson:"TipoRelacion" json:"TipoRelacion"`
	UuidRelacionados []UuidRelacionado33 `xml:"CfdiRelacionado" bson:"UuidRelacionados" json:"UuidRelacionados"`
}

type UuidRelacionado33 struct {
	UUID string `xml:"UUID,attr" bson:"UUID" json:"UUID"`
	Uuid string `bson:"Uuid"`
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
	Nombre        *string `xml:"Nombre,attr" bson:"Nombre,omitempty" json:"Nombre"`
	RegimenFiscal string  `xml:"RegimenFiscal,attr" bson:"RegimenFiscal" json:"RegimenFiscal"`
}

type Receptor33 struct {
	Rfc              string  `xml:"Rfc,attr" bson:"Rfc" json:"Rfc"`                                                  // Cifrado
	Nombre           *string `xml:"Nombre,attr" bson:"Nombre,omitempty" json:"Nombre"`                               // Cifrado
	ResidenciaFiscal *string `xml:"ResidenciaFiscal,attr" bson:"ResidenciaFiscal,omitempty" json:"ResidenciaFiscal"` // Cifrado
	NumRegIdTrib     *string `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty" json:"NumRegIdTrib"`             // Cifrado
	UsoCFDI          string  `xml:"UsoCFDI,attr" bson:"UsoCFDI" json:"UsoCFDI"`
}

type Concepto33 struct {
	ClaveProdServ       string                   `xml:"ClaveProdServ,attr" bson:"ClaveProdServ" json:"ClaveProdServ"`
	NoIdentificacion    *string                  `xml:"NoIdentificacion,attr" bson:"NoIdentificacion,omitempty" json:"NoIdentificacion"`
	Cantidad            float64                  `xml:"Cantidad,attr" bson:"Cantidad" json:"Cantidad"`
	ClaveUnidad         string                   `xml:"ClaveUnidad,attr" bson:"ClaveUnidad" json:"ClaveUnidad"`
	Unidad              *string                  `xml:"Unidad,attr" bson:"Unidad,omitempty" json:"Unidad"`
	Descripcion         string                   `xml:"Descripcion,attr" bson:"Descripcion" json:"Descripcion"`
	ValorUnitario       float64                  `xml:"ValorUnitario,attr" bson:"ValorUnitario" json:"ValorUnitario"`
	Importe             float64                  `xml:"Importe,attr" bson:"Importe" json:"Importe"`
	Descuento           *float64                 `xml:"Descuento,attr" bson:"Descuento,omitempty" json:"Descuento"`
	Impuestos           *ImpuestosConcepto33     `xml:"Impuestos" bson:"Impuestos,omitempty" json:"Impuestos"`
	InformacionAduanera *[]InformacionAduanera33 `xml:"InformacionAduanera" bson:"InformacionAduanera,omitempty" json:"InformacionAduanera"`
	CuentaPredial       *CuentaPredial33         `xml:"CuentaPredial" bson:"CuentaPredial,omitempty" json:"CuentaPredial"`
	ComplementoConcepto *ComplementoConcepto     `xml:"ComplementoConcepto" bson:"ComplementoConcepto,omitempty" json:"ComplementoConcepto"`
	Parte               *[]Parte33               `xml:"Parte" bson:"Parte,omitempty" json:"Parte"`
	DetallesConcepto    string                   `xml:"DetallesConcepto" bson:"DetallesConcepto" json:"DetallesConcepto"`
}

type ImpuestosConcepto33 struct {
	Traslados   *[]TrasladoConcepto33  `xml:"Traslados>Traslado" bson:"Traslados,omitempty" json:"Traslados"`
	Retenciones *[]RetencionConcepto33 `xml:"Retenciones>Retencion" bson:"Retenciones,omitempty" json:"Retenciones"`
}

type TrasladoConcepto33 struct {
	Base       float64  `xml:"Base,attr" bson:"Base" json:"Base"`
	Impuesto   string   `xml:"Impuesto,attr" bson:"Impuesto" json:"Impuesto"`
	TipoFactor string   `xml:"TipoFactor,attr" bson:"TipoFactor" json:"TipoFactor"`
	TasaOCuota *float64 `xml:"TasaOCuota,attr" bson:"TasaOCuota,omitempty" json:"TasaOCuota"`
	Importe    *float64 `xml:"Importe,attr" bson:"Importe,omitempty" json:"Importe"`
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
	NoIdentificacion    *string                  `xml:"NoIdentificacion,attr" bson:"NoIdentificacion,omitempty" json:"NoIdentificacion"`
	Cantidad            float64                  `xml:"Cantidad,attr" bson:"Cantidad" json:"Cantidad"`
	Unidad              *string                  `xml:"Unidad,attr" bson:"Unidad,omitempty" json:"Unidad"`
	Descripcion         string                   `xml:"Descripcion,attr" bson:"Descripcion" json:"Descripcion"`
	ValorUnitario       *float64                 `xml:"ValorUnitario,attr" bson:"ValorUnitario,omitempty" json:"ValorUnitario"`
	Importe             *float64                 `xml:"Importe,attr" bson:"Importe,omitempty" json:"Importe"`
	InformacionAduanera *[]InformacionAduanera33 `xml:"InformacionAduanera" bson:"InformacionAduanera,omitempty" json:"InformacionAduanera"`
}

type Impuestos33 struct {
	TotalImpuestosRetenidos   *float64                `xml:"TotalImpuestosRetenidos,attr" bson:"TotalImpuestosRetenidos,omitempty" json:"TotalImpuestosRetenidos"`
	TotalImpuestosTrasladados *float64                `xml:"TotalImpuestosTrasladados,attr" bson:"TotalImpuestosTrasladados,omitempty" json:"TotalImpuestosTrasladados"`
	Retenciones               *[]RetencionImpuestos33 `xml:"Retenciones>Retencion" bson:"Retenciones,omitempty" json:"Retenciones"`
	Traslados                 *[]TrasladoImpuestos33  `xml:"Traslados>Traslado" bson:"Traslados,omitempty" json:"Traslados"`
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
