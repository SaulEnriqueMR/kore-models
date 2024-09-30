package pagos

import (
	"encoding/xml"
	"time"

	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type Pagos10 struct {
	Version string        `xml:"Version,attr" bson:"Version"`
	Pago    []PagoPagos10 `xml:"Pago" bson:"Pago"`
}

type PagoPagos10 struct {
	FechaPagoString                string                     `xml:"FechaPago,attr"`
	FechaPago                      time.Time                  `bson:"FechaPago"`
	FormaPago                      string                     `xml:"FormaDePagoP,attr" bson:"FormaPago"`
	Moneda                         string                     `xml:"MonedaP,attr" bson:"Moneda"`
	TipoCambio                     *float64                   `xml:"TipoCambioP,attr" bson:"TipoCambio,omitempty"`
	Monto                          float64                    `xml:"Monto,attr" bson:"Monto"`
	NoOperacion                    *string                    `xml:"NumOperacion,attr" bson:"NoOperacion,omitempty"`
	RfcEmisorCuentaOrdenante       *string                    `xml:"RfcEmisorCtaOrd,attr" bson:"RfcEmisorCuentaOrdenante,omitempty"`
	NombreBancoOrdenanteExtranjero *string                    `xml:"NomBancoOrdExt,attr" bson:"NombreBancoOrdenanteExtranjero,omitempty"`
	CuentaOrdenante                *string                    `xml:"CtaOrdenante,attr" bson:"CuentaOrdenante,omitempty"`
	RfcEmisorCuentaBeneficiario    *string                    `xml:"RfcEmisorCtaBen,attr" bson:"RfcEmisorCuentaBenBeneficiario,omitempty"`
	CuentaBeneficiario             *string                    `xml:"CtaBeneficiario,attr" bson:"CuentaBeneficiario,omitempty"`
	TipoCadenaPago                 *string                    `xml:"TipoCadPago,attr" bson:"TipoCadenaPago,omitempty"`
	CertificadoPago                *string                    `xml:"CertPago,attr" bson:"CertificadoPago,omitempty"`
	CadenaPago                     *string                    `xml:"CadPago,attr" bson:"CadenaPago,omitempty"`
	SelloPago                      *string                    `xml:"SelloPago,attr" bson:"SelloPago,omitempty"`
	DocumentosRelacionados         *[]DoctoRelacionadoPagos10 `xml:"DoctoRelacionado" bson:"DocumentosRelacionados,omitempty"`
	Impuestos                      *[]ImpuestosPagos10        `xml:"Impuestos" bson:"Impuestos,omitempty"`
}

type DoctoRelacionadoPagos10 struct {
	IdDocumento          string   `xml:"IdDocumento,attr" bson:"IdDocumento"`
	Serie                *string  `xml:"Serie,attr" bson:"Serie,omitempty"`
	Folio                *string  `xml:"Folio,attr" bson:"Folio,omitempty"`
	Moneda               string   `xml:"MonedaDR,attr" bson:"Moneda"`
	TipoCambio           *float64 `xml:"TipoCambioDR,attr" bson:"TipoCambio,omitempty"`
	MetodoPago           string   `xml:"MetodoDePagoDR,attr" bson:"MetodoPago"`
	NoParcialidad        *float64 `xml:"NumParcialidad,attr" bson:"NoParcialidad,omitempty"`
	ImporteSaldoAnterior *float64 `xml:"ImpSaldoAnt,attr" bson:"ImporteSaldoAnterior,omitempty"`
	ImportePagado        *float64 `xml:"ImpPagado,attr" bson:"ImportePagado,omitempty"`
	ImporteSaldoInsoluto *float64 `xml:"ImpSaldoInsoluto,attr" bson:"ImporteSaldoInsoluto,omitempty"`
}

type ImpuestosPagos10 struct {
	TotalImpuestosRetenidos   *float64            `xml:"TotalImpuestosRetenidos,attr" bson:"TotalImpuestosRetenidos,omitempty"`
	TotalImpuestosTrasladados *float64            `xml:"TotalImpuestosTrasladados,attr" bson:"TotalImpuestosTrasladados,omitempty"`
	Retenciones               *[]RetencionPagos10 `xml:"retenciones>Retencion" bson:"Retenciones,omitempty"`
	Traslados                 *[]TrasladoPagos10  `xml:"Traslados>Traslado" bson:"Traslados,omitempty"`
}

type RetencionPagos10 struct {
	Impuesto string  `xml:"Impuesto,attr" bson:"Impuesto"`
	Importe  float64 `xml:"Importe,attr" bson:"Importe"`
}

type TrasladoPagos10 struct {
	Impuesto   string  `xml:"Impuesto,attr" bson:"Impuesto"`
	TipoFactor string  `xml:"TipoFactor,attr" bson:"TipoFactor"`
	TasaOCuota float64 `xml:"TasaOCuota,attr" bson:"TasaOCuota"`
	Importe    float64 `xml:"Importe,attr" bson:"Importe"`
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
