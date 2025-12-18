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

type Comprobante30 struct {
	documentofiscaldigital.DocumentoFiscalDigital `bson:",inline"`
	Version                                       string          `xml:"version,attr" bson:"Version" json:"Version"`
	Serie                                         *string         `xml:"serie,attr" bson:"Serie" json:"Serie,omitempty"`
	Folio                                         *string         `xml:"folio,attr" bson:"Folio" json:"Folio,omitempty"`
	Fecha                                         string          `xml:"fecha,attr" bson:"Fecha" json:"Fecha"`
	Sello                                         string          `xml:"sello,attr" bson:"Sello" json:"Sello"`
	FormaPago                                     string          `xml:"formaDePago,attr" bson:"FormaDePago" json:"FormaDePago"`
	NoCertificado                                 string          `xml:"noCertificado,attr" bson:"NoCertificado" json:"NoCertificado"`
	Certificado                                   string          `xml:"certificado,attr" bson:"Certificado" json:"Certificado"`
	CondicionesPago                               *string         `xml:"condicionesDePago,attr" bson:"CondicionesDePago" json:"CondicionesDePago,omitempty"`
	Subtotal                                      float64         `xml:"subTotal,attr" bson:"SubTotal" json:"SubTotal"`
	Descuento                                     *float64        `xml:"descuento,attr" bson:"Descuento" json:"Descuento,omitempty"`
	MotivoDescuento                               *string         `xml:"motivoDescuento,attr" bson:"MotivoDescuento" json:"MotivoDescuento,omitempty"`
	TipoDeCambio                                  *string         `xml:"TipoCambio,attr" bson:"TipoDeCambio" json:"TipoDeCambio,omitempty"`
	TipoCambio                                    float64         `bson:"TipoCambio" json:"TipoCambio"`
	Moneda                                        *string         `xml:"Moneda,attr" bson:"Moneda" json:"Moneda,omitempty"`
	Total                                         float64         `xml:"total,attr" bson:"Total" json:"Total"`
	MetodoPago                                    *string         `xml:"metodoDePago,attr" bson:"MetodoDePago" json:"MetodoDePago,omitempty"`
	TipoDeComprobante                             string          `xml:"tipoDeComprobante,attr" bson:"TipoDeComprobante" json:"TipoDeComprobante"`
	TipoComprobante                               string          `bson:"TipoComprobante" json:"TipoComprobante"`
	Emisor                                        Emisor30        `xml:"Emisor" bson:"Emisor" json:"Emisor"`
	Receptor                                      Receptor30      `xml:"Receptor" bson:"Receptor" json:"Receptor"`
	Conceptos                                     []Concepto30    `xml:"Conceptos>Concepto" bson:"Conceptos" json:"Conceptos"`
	Impuestos                                     Impuestos30     `xml:"Impuestos" bson:"Impuestos" json:"Impuestos"`
	Complemento                                   *Complemento    `xml:"Complemento" bson:"Complemento,omitempty" json:"Complemento,omitempty"`
	Addenda                                       *Addenda        `xml:"Addenda" bson:"Addenda,omitempty" json:"Addenda,omitempty"`
	KuantikMetadata                               KuantikMetadata `bson:"KuantikMetadata,omitempty" json:"KuantikMetadata,omitempty"`
}

func (c *Comprobante30) DefineTransaccion(rfc string) {
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

type Emisor30 struct {
	Rfc             string       `xml:"rfc,attr" bson:"Rfc" json:"Rfc"`
	Nombre          string       `xml:"nombre,attr" bson:"Nombre" json:"Nombre"`
	DomicilioFiscal *Domicilio30 `xml:"DomicilioFiscal" bson:"DomicilioFiscal,omitempty" json:"DomicilioFiscal,omitempty"`
	ExpedidoEn      *Domicilio30 `xml:"ExpedidoEn" bson:"ExpedidoEn,omitempty" json:"ExpedidoEn,omitempty"`
}

type Receptor30 struct {
	Rfc       string       `xml:"rfc,attr" bson:"Rfc" json:"Rfc"`
	Nombre    *string      `xml:"nombre,attr" bson:"Nombre,omitempty" json:"Nombre,omitempty"`
	Domicilio *Domicilio30 `xml:"Domicilio" bson:"Domicilio,omitempty" json:"Domicilio,omitempty"`
}

type Domicilio30 struct {
	Calle        *string `xml:"calle,attr" bson:"Calle,omitempty" json:"Calle,omitempty"`
	NoExterior   *string `xml:"noExterior,attr" bson:"NoExterior,omitempty" json:"NoExterior,omitempty"`
	NoInterior   *string `xml:"noInterior,attr" bson:"NoInterior,omitempty" json:"NoInterior,omitempty"`
	Colonia      *string `xml:"colonia,attr" bson:"Colonia,omitempty" json:"Colonia,omitempty"`
	Localidad    *string `xml:"localidad,attr" bson:"Localidad,omitempty" json:"Localidad,omitempty"`
	Referencia   *string `xml:"referencia,attr" bson:"Referencia,omitempty" json:"Referencia,omitempty"`
	Municipio    *string `xml:"municipio,attr" bson:"Municipio,omitempty" json:"Municipio,omitempty"`
	Estado       *string `xml:"estado,attr" bson:"Estado,omitempty" json:"Estado,omitempty"`
	Pais         *string `xml:"pais,attr" bson:"Pais,omitempty" json:"Pais,omitempty"`
	CodigoPostal *string `xml:"codigoPostal,attr" bson:"CodigoPostal,omitempty" json:"CodigoPostal,omitempty"`
}

type Concepto30 struct {
	Cantidad            float64                  `xml:"cantidad,attr" bson:"Cantidad" json:"Cantidad"`
	Unidad              *string                  `xml:"unidad,attr" bson:"Unidad" json:"Unidad"`
	NoIdentificacion    *string                  `xml:"noIdentificacion,attr" bson:"NoIdentificacion,omitempty" json:"NoIdentificacion,omitempty"`
	Descripcion         string                   `xml:"descripcion,attr" bson:"Descripcion" json:"Descripcion"`
	ValorUnitario       float64                  `xml:"valorUnitario,attr" bson:"ValorUnitario" json:"ValorUnitario"`
	Importe             float64                  `xml:"importe,attr" bson:"Importe" json:"Importe"`
	InformacionAduanera *[]InformacionAduanera30 `xml:"InformacionAduanera" bson:"InformacionAduanera,omitempty" json:"InformacionAduanera,omitempty"`
	CuentaPredial       *CuentaPredial30         `xml:"CuentaPredial" bson:"CuentaPredial,omitempty" json:"CuentaPredial,omitempty"`
	ComplementoConcepto *ComplementoConcepto     `xml:"ComplementoConcepto" bson:"ComplementoConcepto,omitempty" json:"ComplementoConcepto,omitempty"`
	Parte               *[]Parte30               `xml:"Parte" bson:"Parte,omitempty" json:"Parte,omitempty"`
}

type InformacionAduanera30 struct {
	Numero      string    `xml:"numero,attr" bson:"Numero" json:"Numero"`
	FechaString string    `xml:"fecha,attr" bson:"fecha" json:"fecha"`
	Fecha       time.Time `bson:"Fecha" json:"Fecha"`
	Aduana      *string   `xml:"aduana,attr" bson:"Aduana,omitempty" json:"Aduana,omitempty"`
}

func (ia30 *InformacionAduanera30) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias InformacionAduanera30
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	*ia30 = InformacionAduanera30(aux)

	if ia30 != nil {
		fecha, err := helpers.ParseDatetime(ia30.FechaString)
		if err != nil {
			return err
		}
		ia30.Fecha = fecha
	}

	return nil
}

type CuentaPredial30 struct {
	Numero string `xml:"numero,attr" bson:"Numero" json:"Numero"`
}

type Parte30 struct {
	Cantidad            float64                  `xml:"cantidad,attr" bson:"Cantidad" json:"Cantidad"`
	Unidad              *string                  `xml:"unidad,attr" bson:"Unidad,omitempty" json:"Unidad,omitempty"`
	NoIdentificacion    *string                  `xml:"noIdentificacion,attr" bson:"NoIdentificacion,omitempty" json:"NoIdentificacion,omitempty"`
	Descripcion         string                   `xml:"descripcion,attr" bson:"Descripcion" json:"Descripcion"`
	ValorUnitario       *float64                 `xml:"valorUnitario,attr" bson:"ValorUnitario,omitempty" json:"ValorUnitario,omitempty"`
	Importe             *float64                 `xml:"importe,attr" bson:"Importe,omitempty" json:"Importe,omitempty"`
	InformacionAduanera *[]InformacionAduanera30 `xml:"InformacionAduanera" bson:"InformacionAduanera,omitempty" json:"InformacionAduanera,omitempty"`
}

type Impuestos30 struct {
	Retenciones               *[]Retencion30 `xml:"Retenciones>Retencion" bson:"Retenciones,omitempty" json:"Retenciones,omitempty"`
	Traslados                 *[]Traslado30  `xml:"Traslados>Traslado" bson:"Traslados,omitempty" json:"Traslados,omitempty"`
	TotalImpuestosRetenidos   *float64       `xml:"totalImpuestosRetenidos,attr" bson:"TotalImpuestosRetenidos,omitempty" json:"TotalImpuestosRetenidos,omitempty"`
	TotalImpuestosTrasladados *float64       `xml:"totalImpuestosTrasladados,attr" bson:"TotalImpuestosTrasladados,omitempty" json:"TotalImpuestosTrasladados,omitempty"`
}

type Retencion30 struct {
	TipoImpuesto string  `xml:"impuesto,attr" bson:"TipoImpuesto" json:"TipoImpuesto"`
	Impuesto     string  `bson:"Impuesto" json:"Impuesto"`
	Importe      float64 `xml:"importe,attr" bson:"Importe" json:"Importe"`
}

func (i *Retencion30) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Alias Retencion30
	var aux Alias

	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	*i = Retencion30(aux)

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

type Traslado30 struct {
	TipoImpuesto string  `xml:"impuesto,attr" bson:"TipoImpuesto" json:"TipoImpuesto"`
	Impuesto     string  `bson:"Impuesto" json:"Impuesto"`
	Tasa         float64 `xml:"tasa,attr" bson:"Tasa" json:"Tasa"`
	Importe      float64 `xml:"importe,attr" bson:"Importe" json:"Importe"`
}

func (i *Traslado30) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Alias Traslado30
	var aux Alias

	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	*i = Traslado30(aux)

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

func (c *Comprobante30) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias Comprobante30
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	fechaEmision, err := helpers.ParseDatetime(aux.Fecha)
	if err != nil {
		return err
	}

	*c = Comprobante30(aux)

	c.FechaEmision = fechaEmision
	c.Comprobante = true
	c.Vigente = nil
	// c.ProcessorMetadata.CreationDate = nil

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

	fileName := c.GetFileName()
	c.DocumentoFiscalDigital.S3FilePath = &fileName

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

func (c *Comprobante30) GetFileName() string {
	year := fmt.Sprint(c.FechaEmision.Year())
	month := fmt.Sprint(int(c.FechaEmision.Month()))
	return c.Emisor.Rfc + "/" + c.Receptor.Rfc + "/" + year + "/" + month + "/" + c.Uuid + ".xml"
}

func (c Comprobante30) GetBasePath() string {
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
