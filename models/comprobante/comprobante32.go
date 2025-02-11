package comprobante

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/SaulEnriqueMR/kore-models/models"

	"github.com/SaulEnriqueMR/kore-models/models/documentofiscaldigital"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type Comprobante32 struct {
	documentofiscaldigital.DocumentoFiscalDigital `bson:",inline"`
	Version                                       string       `xml:"version,attr" bson:"Version"`
	Serie                                         *string      `xml:"serie,attr"  bson:"Serie,omitempty"`
	Folio                                         *string      `xml:"folio,attr" bson:"Folio,omitempty"`
	Fecha                                         string       `xml:"fecha,attr" bson:"Fecha"`
	Sello                                         string       `xml:"sello,attr" bson:"Sello"`
	FormaPago                                     string       `xml:"formaDePago,attr" bson:"FormaPago"`
	NoCertificado                                 string       `xml:"noCertificado,attr" bson:"NoCertificado"`
	Certificado                                   string       `xml:"certificado,attr" bson:"Certificado"`
	CondicionesPago                               *string      `xml:"condicionesDePago,attr" bson:"CondicionesPago,omitempty"`
	Subtotal                                      float64      `xml:"subTotal,attr" bson:"Subtotal"`
	Descuento                                     *float64     `xml:"descuento,attr" bson:"Descuento,omitempty"`
	MotivoDescuento                               *string      `xml:"motivoDescuento,attr" bson:"MotivoDescuento,omitempty"`
	TipoDeCambio                                  *string      `xml:"TipoCambio,attr" bson:"TipoDeCambio,omitempty"`
	TipoCambio                                    float64      `bson:"TipoCambio,omitempty"`
	Moneda                                        *string      `xml:"Moneda,attr" bson:"Moneda,omitempty"`
	Total                                         float64      `xml:"total,attr" bson:"Total"`
	TipoDeComprobante                             string       `xml:"tipoDeComprobante,attr" bson:"TipoDeComprobante"`
	TipoComprobante                               string       `bson:"TipoComprobante"`
	MetodoPago                                    string       `xml:"metodoDePago,attr" bson:"MetodoPago"`
	LugarExpedicion                               string       `xml:"LugarExpedicion,attr" bson:"LugarExpedicion"`
	NumeroCuentaPago                              *string      `xml:"NumCtaPago,attr" bson:"NumeroCuentaPago,omitempty"`
	FolioFiscalOriginal                           *string      `bson:"FolioFiscalOriginal,omitempty"`
	FolioFiscalOrig                               *string      `xml:"FolioFiscalOrig,attr" bson:"FolioFiscalOrig,omitempty"`
	SerieFolioFiscalOriginal                      *string      `xml:"SerieFolioFiscalOrig,attr" bson:"SerieFolioFiscalOriginal,omitempty"`
	FechaFolioFiscalOriginal                      *time.Time   `bson:"FechaFolioFiscalOriginal,omitempty"`
	FechaFolioFiscalOrig                          *string      `xml:"FechaFolioFiscalOrig,attr" bson:"FechaFolioFiscalOrig,omitempty"`
	MontoFolioFiscalOriginal                      *float64     `xml:"MontoFolioFiscalOrig,attr" bson:"MontoFolioFiscalOriginal,omitempty"`
	Emisor                                        Emisor32     `xml:"Emisor" bson:"Emisor"`
	Receptor                                      Receptor32   `xml:"Receptor" bson:"Receptor"`
	Conceptos                                     []Concepto32 `xml:"Conceptos>Concepto" bson:"Conceptos"`
	Impuestos                                     Impuestos32  `xml:"Impuestos" bson:"Impuestos"`
	Complemento                                   *Complemento `xml:"Complemento" bson:"Complemento,omitempty"`
	Addenda                                       *Addenda     `xml:"Addenda" bson:"Addenda,omitempty"`
}

func (c *Comprobante32) DefineTransaccion(rfc string) {
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

type Emisor32 struct {
	Rfc               string            `xml:"rfc,attr" bson:"Rfc"`
	Nombre            *string           `xml:"nombre,attr" bson:"Nombre,omitempty"`
	RegimenesFiscales []RegimenFiscal32 `xml:"RegimenFiscal" bson:"RegimenesFiscales"`
	DomicilioFiscal   *Domicilio32      `xml:"DomicilioFiscal" bson:"DomicilioFiscal,omitempty"`
	ExpedidoEn        *Domicilio32      `xml:"ExpedidoEn" bson:"ExpedidoEn,omitempty"`
}

type RegimenFiscal32 struct {
	RegimenFiscal string `xml:"Regimen,attr" bson:"RegimenFiscal"`
}

type Receptor32 struct {
	Rfc       string       `xml:"rfc,attr" bson:"Rfc"`
	Nombre    *string      `xml:"nombre,attr" bson:"Nombre,omitempty"`
	Domicilio *Domicilio32 `xml:"Domicilio" bson:"Domicilio,omitempty"`
}

type Domicilio32 struct {
	Calle        string  `xml:"calle,attr" bson:"Calle"`
	NoExterior   *string `xml:"noExterior,attr" bson:"NoExterior,omitempty"`
	NoInterior   *string `xml:"noInterior,attr" bson:"NoInterior,omitempty"`
	Colonia      *string `xml:"colonia,attr" bson:"Colonia,omitempty"`
	Localidad    *string `xml:"localidad,attr" bson:"Localidad,omitempty"`
	Referencia   *string `xml:"referencia,attr" bson:"Referencia,omitempty"`
	Municipio    string  `xml:"municipio,attr" bson:"Municipio"`
	Estado       string  `xml:"estado,attr" bson:"Estado"`
	Pais         string  `xml:"pais,attr" bson:"Pais"`
	CodigoPostal string  `xml:"codigoPostal,attr" bson:"CodigoPostal"`
}

type Concepto32 struct {
	Cantidad            float64                  `xml:"cantidad,attr" bson:"Cantidad"`
	Unidad              string                   `xml:"unidad,attr" bson:"Unidad"`
	NoIdentificacion    *string                  `xml:"noIdentificacion,attr" bson:"NoIdentificacion,omitempty"`
	Descripcion         string                   `xml:"descripcion,attr" bson:"Descripcion"`
	ValorUnitario       float64                  `xml:"valorUnitario,attr" bson:"ValorUnitario"`
	Importe             float64                  `xml:"importe,attr" bson:"Importe"`
	InformacionAduanera *[]InformacionAduanera32 `xml:"InformacionAduanera" bson:"InformacionAduanera,omitempty"`
	CuentaPredial       *CuentaPredial32         `xml:"CuentaPredial" bson:"CuentaPredial,omitempty"`
	ComplementoConcepto *ComplementoConcepto     `xml:"ComplementoConcepto" bson:"ComplementoConcepto,omitempty"`
	Parte               *[]Parte32               `xml:"Parte" bson:"Parte,omitempty"`
}

type InformacionAduanera32 struct {
	Numero      string    `xml:"numero,attr" bson:"Numero"`
	FechaString string    `xml:"fecha,attr" bson:"fecha"`
	Fecha       time.Time `bson:"Fecha"`
	Aduana      *string   `xml:"aduana,attr" bson:"Aduana,omitempty"`
}

func (ia32 *InformacionAduanera32) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias InformacionAduanera32
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	*ia32 = InformacionAduanera32(aux)

	if ia32 != nil {
		fecha, err := helpers.ParseDatetime(ia32.FechaString)
		if err != nil {
			return err
		}
		ia32.Fecha = fecha
	}

	return nil
}

type CuentaPredial32 struct {
	Numero string `xml:"numero,attr" bson:"Numero"`
}

type Parte32 struct {
	Cantidad            float64                  `xml:"cantidad,attr" bson:"Cantidad"`
	Unidad              *string                  `xml:"unidad,attr" bson:"Unidad,omitempty"`
	NoIdentificacion    *string                  `xml:"noIdentificacion,attr" bson:"NoIdentificacion,omitempty"`
	Descripcion         string                   `xml:"descripcion,attr" bson:"Descripcion"`
	ValorUnitario       *float64                 `xml:"valorUnitario,attr" bson:"ValorUnitario,omitempty"`
	Importe             *float64                 `xml:"importe,attr" bson:"Importe,omitempty"`
	InformacionAduanera *[]InformacionAduanera32 `xml:"InformacionAduanera" bson:"InformacionAduanera,omitempty"`
}

type Impuestos32 struct {
	Retenciones               *[]Retencion32 `xml:"Retenciones>Retencion" bson:"Retenciones,omitempty"`
	Traslados                 *[]Traslado32  `xml:"Traslados>Traslado" bson:"Traslados,omitempty"`
	TotalImpuestosRetenidos   *float64       `xml:"totalImpuestosRetenidos,attr" bson:"TotalImpuestosRetenidos,omitempty"`
	TotalImpuestosTrasladados *float64       `xml:"totalImpuestosTrasladados,attr" bson:"TotalImpuestosTrasladados,omitempty"`
}

type Retencion32 struct {
	TipoImpuesto string  `xml:"impuesto,attr" bson:"TipoImpuesto"`
	Impuesto     string  `bson:"Impuesto"`
	Importe      float64 `xml:"importe,attr" bson:"Importe"`
}

func (i *Retencion32) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Alias Retencion32
	var aux Alias

	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	*i = Retencion32(aux)

	switch strings.ToLower(i.TipoImpuesto) {
	case "isr":
		i.Impuesto = "001"
		break
	case "iva":
		i.Impuesto = "002"
		break
	case "ieps":
		i.Impuesto = "003"
		break
	}

	return nil
}

type Traslado32 struct {
	TipoImpuesto string  `xml:"impuesto,attr" bson:"TipoImpuesto"`
	Impuesto     string  `bson:"Impuesto"`
	Tasa         float64 `xml:"tasa,attr" bson:"Tasa"`
	Importe      float64 `xml:"importe,attr" bson:"Importe"`
}

func (i *Traslado32) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Alias Traslado32
	var aux Alias

	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	*i = Traslado32(aux)

	switch strings.ToLower(i.TipoImpuesto) {
	case "isr":
		i.Impuesto = "001"
		break
	case "iva":
		i.Impuesto = "002"
		break
	case "ieps":
		i.Impuesto = "003"
		break
	}

	return nil
}

func (c *Comprobante32) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias Comprobante32
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	fechaEmision, err := helpers.ParseDatetime(aux.Fecha)
	if err != nil {
		return err
	}

	*c = Comprobante32(aux)

	if aux.FolioFiscalOrig != nil {
		folioOriginal := strings.ToUpper(*aux.FolioFiscalOrig)
		c.FolioFiscalOriginal = &folioOriginal
	}

	c.FechaEmision = fechaEmision
	c.Comprobante = true
	c.Vigente = nil
	c.ProcessorMetadata.CreationDate = nil

	switch strings.ToLower(c.TipoDeComprobante) {
	case "ingreso":
		c.TipoComprobante = "I"
		break
	case "egreso":
		c.TipoComprobante = "E"
		break
	case "traslado":
		c.TipoComprobante = "T"
		break
	}

	tipoCambio := 1.0
	if c.TipoDeCambio != nil {
		tipoDeCambio := *c.TipoDeCambio
		parsedFloat, _ := strconv.ParseFloat(tipoDeCambio, 64)
		if parsedFloat > 0 {
			tipoCambio = parsedFloat
		}
	}

	c.TipoCambio = tipoCambio
	totalesMonedaLocal := documentofiscaldigital.TotalesMonedaLocal{
		Total:    c.Total * tipoCambio,
		Subtotal: c.Subtotal * tipoCambio,
	}

	if c.Descuento != nil {
		descuento := *c.Descuento * tipoCambio
		totalesMonedaLocal.Descuento = &descuento
	}

	if c.Impuestos.TotalImpuestosRetenidos != nil {
		tir := *c.Impuestos.TotalImpuestosRetenidos * tipoCambio
		totalesMonedaLocal.TotalImpuestosRetenidos = &tir
	}

	if c.Impuestos.TotalImpuestosTrasladados != nil {
		tit := *c.Impuestos.TotalImpuestosTrasladados * tipoCambio
		totalesMonedaLocal.TotalImpuestosTrasladados = &tit
	}

	c.TotalesMonedaLocal = totalesMonedaLocal

	if c.InformacionAdicional != nil {
		c.InformacionAdicional.StampedByKuantik = nil
	}
	if c.Cancelacion != nil {
		c.Cancelacion.CanceledByKuantik = nil
	}

	if c.Complemento.TimbreFiscalDigital != nil {
		tfd := c.Complemento.TimbreFiscalDigital.TimbreFiscalDigital10
		if tfd != nil {
			c.FechaTimbrado = tfd.FechaTimbrado
			c.Uuid = strings.ToUpper(tfd.Uuid)
		}
	}
	c.CadenaOriginal = helpers.CreateCadenaOriginal(*c)

	// Calculo de totales
	// processDate := time.Now().UTC()
	now := time.Now()
	processDate := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, time.UTC)
	c.ProcessorMetadata.LastUpdate = &processDate
	c.ProcessorMetadata.KoreModelsVersion = &models.KoreModelVersion
	return nil
}

func (c *Comprobante32) GetFileName() string {
	year := fmt.Sprint(c.FechaEmision.Year())
	month := fmt.Sprint(int(c.FechaEmision.Month()))
	return c.Emisor.Rfc + "/" + c.Receptor.Rfc + "/" + year + "/" + month + "/" + c.Uuid + ".xml"
}

func (c Comprobante32) GetBasePath() string {
	year := c.FechaEmision.Year()
	month := c.FechaEmision.Month()
	sb := strings.Builder{}
	sb.WriteString(c.Emisor.Rfc)
	sb.WriteString("/")
	sb.WriteString(c.Receptor.Rfc)
	sb.WriteString("/")
	sb.WriteString(strconv.Itoa(year))
	sb.WriteString("/")
	sb.WriteString(strconv.Itoa(int(month)))
	sb.WriteString("/")
	sb.WriteString(c.Uuid)
	return sb.String()
}

/* func (c *Comprobante32) SetFilePaths() {
	basePath := c.GetBasePath()
	c.XmlPath = strings.Join([]string{basePath, "xml"}, ".")
	c.PdfPath = strings.Join([]string{basePath, "pdf"}, ".")
} */
