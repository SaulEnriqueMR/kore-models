package pagos

import (
	"encoding/xml"
	"strings"
	"time"

	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type Pagos10 struct {
	Version string        `xml:"Version,attr" bson:"Version" json:"Version"`
	Pago    []PagoPagos10 `xml:"Pago" bson:"Pago" json:"Pago"`
}

type PagoPagos10 struct {
	FechaPagoString                string                     `xml:"FechaPago,attr" bson:"FechaPagoString" json:"FechaPagoString"`
	FechaPago                      time.Time                  `bson:"FechaPago" json:"FechaPago"`
	FormaPago                      string                     `xml:"FormaDePagoP,attr" bson:"FormaPago" json:"FormaPago"`
	Moneda                         string                     `xml:"MonedaP,attr" bson:"Moneda" json:"Moneda"`
	TipoCambio                     *float64                   `xml:"TipoCambioP,attr" bson:"TipoCambio,omitempty" json:"TipoCambio,omitempty"`
	Monto                          float64                    `xml:"Monto,attr" bson:"Monto" json:"Monto"`
	NoOperacion                    *string                    `xml:"NumOperacion,attr" bson:"NoOperacion,omitempty" json:"NoOperacion,omitempty"`
	RfcEmisorCuentaOrdenante       *string                    `xml:"RfcEmisorCtaOrd,attr" bson:"RfcEmisorCuentaOrdenante,omitempty" json:"RfcEmisorCuentaOrdenante,omitempty"`
	NombreBancoOrdenanteExtranjero *string                    `xml:"NomBancoOrdExt,attr" bson:"NombreBancoOrdenanteExtranjero,omitempty" json:"NombreBancoOrdenanteExtranjero,omitempty"`
	CuentaOrdenante                *string                    `xml:"CtaOrdenante,attr" bson:"CuentaOrdenante,omitempty" json:"CuentaOrdenante,omitempty"`
	RfcEmisorCuentaBeneficiario    *string                    `xml:"RfcEmisorCtaBen,attr" bson:"RfcEmisorCuentaBenBeneficiario,omitempty" json:"RfcEmisorCuentaBenBeneficiario,omitempty"`
	CuentaBeneficiario             *string                    `xml:"CtaBeneficiario,attr" bson:"CuentaBeneficiario,omitempty" json:"CuentaBeneficiario,omitempty"`
	TipoCadenaPago                 *string                    `xml:"TipoCadPago,attr" bson:"TipoCadenaPago,omitempty" json:"TipoCadenaPago,omitempty"`
	CertificadoPago                *string                    `xml:"CertPago,attr" bson:"CertificadoPago,omitempty" json:"CertificadoPago,omitempty"`
	CadenaPago                     *string                    `xml:"CadPago,attr" bson:"CadenaPago,omitempty" json:"CadenaPago,omitempty"`
	SelloPago                      *string                    `xml:"SelloPago,attr" bson:"SelloPago,omitempty" json:"SelloPago,omitempty"`
	DocumentosRelacionados         *[]DoctoRelacionadoPagos10 `xml:"DoctoRelacionado" bson:"DocumentosRelacionados,omitempty" json:"DocumentosRelacionados,omitempty"`
	Impuestos                      *[]ImpuestosPagos10        `xml:"Impuestos" bson:"Impuestos,omitempty" json:"Impuestos,omitempty"`
}

type DoctoRelacionadoPagos10 struct {
	IdDocumentoString    string   `xml:"IdDocumento,attr" bson:"IdDocumentoString" json:"IdDocumentoString"`
	IdDocumento          string   `bson:"IdDocumento" json:"IdDocumento"`
	Serie                *string  `xml:"Serie,attr" bson:"Serie,omitempty" json:"Serie,omitempty"`
	Folio                *string  `xml:"Folio,attr" bson:"Folio,omitempty" json:"Folio,omitempty"`
	Moneda               string   `xml:"MonedaDR,attr" bson:"Moneda" json:"Moneda"`
	TipoCambio           *float64 `xml:"TipoCambioDR,attr" bson:"TipoCambio,omitempty" json:"TipoCambio,omitempty"`
	MetodoPago           string   `xml:"MetodoDePagoDR,attr" bson:"MetodoPago" json:"MetodoPago"`
	NoParcialidad        *float64 `xml:"NumParcialidad,attr" bson:"NoParcialidad,omitempty" json:"NoParcialidad,omitempty"`
	ImporteSaldoAnterior *float64 `xml:"ImpSaldoAnt,attr" bson:"ImporteSaldoAnterior,omitempty" json:"ImporteSaldoAnterior,omitempty"`
	ImportePagado        *float64 `xml:"ImpPagado,attr" bson:"ImportePagado,omitempty" json:"ImportePagado,omitempty"`
	ImporteSaldoInsoluto *float64 `xml:"ImpSaldoInsoluto,attr" bson:"ImporteSaldoInsoluto,omitempty" json:"ImporteSaldoInsoluto,omitempty"`
}

func (dr *DoctoRelacionadoPagos10) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias DoctoRelacionadoPagos10
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	*dr = DoctoRelacionadoPagos10(aux)
	dr.IdDocumento = strings.ToUpper(aux.IdDocumentoString)

	return nil
}

type ImpuestosPagos10 struct {
	TotalImpuestosRetenidos   *float64            `xml:"TotalImpuestosRetenidos,attr" bson:"TotalImpuestosRetenidos,omitempty" json:"TotalImpuestosRetenidos,omitempty"`
	TotalImpuestosTrasladados *float64            `xml:"TotalImpuestosTrasladados,attr" bson:"TotalImpuestosTrasladados,omitempty" json:"TotalImpuestosTrasladados,omitempty"`
	Retenciones               *[]RetencionPagos10 `xml:"retenciones>Retencion" bson:"Retenciones,omitempty" json:"Retenciones,omitempty"`
	Traslados                 *[]TrasladoPagos10  `xml:"Traslados>Traslado" bson:"Traslados,omitempty" json:"Traslados,omitempty"`
}

type RetencionPagos10 struct {
	Impuesto string  `xml:"Impuesto,attr" bson:"Impuesto" json:"Impuesto"`
	Importe  float64 `xml:"Importe,attr" bson:"Importe" json:"Importe"`
}

type TrasladoPagos10 struct {
	Impuesto   string  `xml:"Impuesto,attr" bson:"Impuesto" json:"Impuesto"`
	TipoFactor string  `xml:"TipoFactor,attr" bson:"TipoFactor" json:"TipoFactor"`
	TasaOCuota float64 `xml:"TasaOCuota,attr" bson:"TasaOCuota" json:"TasaOCuota"`
	Importe    float64 `xml:"Importe,attr" bson:"Importe" json:"Importe"`
}

func (p *PagoPagos10) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias PagoPagos10
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*p = PagoPagos10(aux)

	fecha, err := helpers.ParseDatetime(aux.FechaPagoString)
	if err != nil {
		return err
	}
	p.FechaPago = fecha

	return nil
}
