package pagos

import (
	"encoding/xml"
	"strings"
	"time"

	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type Pagos20 struct {
	Version string         `xml:"Version,attr" bson:"Version"`
	Totales TotalesPagos20 `xml:"Totales" bson:"Totales"`
	Pago    []PagoPagos20  `xml:"Pago" bson:"Pago"`
}

type TotalesPagos20 struct {
	TotalRetencionesIva         *float64 `xml:"TotalRetencionesIVA,attr" bson:"TotalRetencionesIva,omitempty"`
	TotalRetencionesIsr         *float64 `xml:"TotalRetencionesISR,attr" bson:"TotalRetencionesIsr,omitempty"`
	TotalRetencionesIeps        *float64 `xml:"TotalRetencionesIEPS,attr" bson:"TotalRetencionesIeps,omitempty"`
	TotalTrasladosBaseIva16     *float64 `xml:"TotalTrasladosBaseIVA16,attr" bson:"TotalTrasladosBaseIva16,omitempty"`
	TotalTrasladosImpuestoIva16 *float64 `xml:"TotalTrasladosImpuestoIVA16,attr" bson:"TotalTrasladosImpuestoIva16,omitempty"`
	TotalTrasladosBaseIva8      *float64 `xml:"TotalTrasladosBaseIVA8,attr" bson:"TotalTrasladosBaseIva8,omitempty"`
	TotalTrasladosImpuestoIva8  *float64 `xml:"TotalTrasladosImpuestoIVA8,attr" bson:"TotalTrasladosImpuestoIva8,omitempty"`
	TotalTrasladosBaseIva0      *float64 `xml:"TotalTrasladosBaseIVA0,attr" bson:"TotalTrasladosBaseIva0,omitempty"`
	TotalTrasladosImpuestoIva0  *float64 `xml:"TotalTrasladosImpuestoIVA0,attr" bson:"TotalTrasladosImpuestoIva0,omitempty"`
	TotalTrasladosBaseIvaExento *float64 `xml:"TotalTrasladosBaseIVAExento,attr" bson:"TotalTrasladosBaseIvaExento,omitempty"`
	MontoTotalPagos             float64  `xml:"MontoTotalPagos,attr" bson:"MontoTotalPagos"`
}

type PagoPagos20 struct {
	FechaPagoString                string                    `xml:"FechaPago,attr" bson:"FechaPagoString"`
	FechaPago                      time.Time                 `bson:"FechaPago"`
	FormaPago                      string                    `xml:"FormaDePagoP,attr" bson:"FormaPago"`
	Moneda                         string                    `xml:"MonedaP,attr" bson:"Moneda"`
	TipoCambio                     *float64                  `xml:"TipoCambioP,attr" bson:"TipoCambio,omitempty"`
	Monto                          float64                   `xml:"Monto,attr" bson:"Monto"`
	NoOperacion                    *string                   `xml:"NumOperacion,attr" bson:"NoOperacion,omitempty"`
	RfcEmisorCuentaOrdenante       *string                   `xml:"RfcEmisorCtaOrd,attr" bson:"RfcEmisorCuentaOrdenante,omitempty"`
	NombreBancoOrdenanteExtranjero *string                   `xml:"NomBancoOrdExt,attr" bson:"NombreBancoOrdenanteExtranjero,omitempty"` // Cifrado
	CuentaOrdenante                *string                   `xml:"CtaOrdenante,attr" bson:"CuentaOrdenante,omitempty"`                  // Cifrado
	RfcEmisorCuentaBeneficiario    *string                   `xml:"RfcEmisorCtaBen,attr" bson:"RfcEmisorCuentaBeneficiario,omitempty"`   // Cifrado
	CuentaBeneficiario             *string                   `xml:"CtaBeneficiario,attr" bson:"CuentaBeneficiario,omitempty"`            // Cifrado
	TipoCadenaPago                 *string                   `xml:"TipoCadPago,attr" bson:"TipoCadenaPago,omitempty"`
	CertificadoPago                *string                   `xml:"CertPago,attr" bson:"CertificadoPago,omitempty"`
	CadenaPago                     *string                   `xml:"CadPago,attr" bson:"CadenaPago,omitempty"`
	SelloPago                      *string                   `xml:"SelloPago,attr" bson:"SelloPago,omitempty"`
	DocumentosRelacionados         []DoctoRelacionadoPagos20 `xml:"DoctoRelacionado" bson:"DocumentosRelacionados"`
	Impuestos                      *ImpuestosPagos20         `xml:"ImpuestosP" bson:"Impuestos,omitempty"`
}

type DoctoRelacionadoPagos20 struct {
	IdDocumentoString    string              `xml:"IdDocumento,attr" bson:"IdDocumentoString"`
	IdDocumento          string              `bson:"IdDocumento"`
	Serie                *string             `xml:"Serie,attr" bson:"Serie,omitempty"`
	Folio                *string             `xml:"Folio,attr" bson:"Folio,omitempty"`
	Moneda               string              `xml:"MonedaDR,attr" bson:"Moneda"`
	Equivalencia         *float64            `xml:"EquivalenciaDR,attr" bson:"Equivalencia,omitempty"`
	NoParcialidad        float64             `xml:"NumParcialidad,attr" bson:"NoParcialidad"`
	ImporteSaldoAnterior float64             `xml:"ImpSaldoAnt,attr" bson:"ImporteSaldoAnterior"`
	ImportePagado        float64             `xml:"ImpPagado,attr" bson:"ImportePagado"`
	ImporteSaldoInsoluto float64             `xml:"ImpSaldoInsoluto,attr" bson:"ImporteSaldoInsoluto"`
	ObjetoImpuesto       string              `xml:"ObjetoImpDR,attr" bson:"ObjetoImpuesto"`
	Impuestos            *ImpuestosDRPagos20 `xml:"ImpuestosDR" bson:"Impuestos,omitempty"`
}

func (dr *DoctoRelacionadoPagos20) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias DoctoRelacionadoPagos20
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	*dr = DoctoRelacionadoPagos20(aux)
	dr.IdDocumento = strings.ToUpper(aux.IdDocumentoString)

	return nil
}

type ImpuestosDRPagos20 struct {
	Retenciones *[]RetencionDRPagos20 `xml:"RetencionesDR>RetencionDR" bson:"Retenciones,omitempty"`
	Traslados   *[]TrasladoDRPagos20  `xml:"TrasladosDR>TrasladoDR" bson:"Traslados,omitempty"`
}

type RetencionDRPagos20 struct {
	Base       float64 `xml:"BaseDR,attr" bson:"Base"`
	Impuesto   string  `xml:"ImpuestoDR,attr" bson:"Impuesto"`
	TipoFactor string  `xml:"TipoFactorDR,attr" bson:"TipoFactor"`
	TasaOCuota float64 `xml:"TasaOCuotaDR,attr" bson:"TasaOCuota"`
	Importe    float64 `xml:"ImporteDR,attr" bson:"Importe"`
}

type TrasladoDRPagos20 struct {
	Base       float64  `xml:"BaseDR,attr" bson:"Base"`
	Impuesto   string   `xml:"ImpuestoDR,attr" bson:"Impuesto"`
	TipoFactor string   `xml:"TipoFactorDR,attr" bson:"TipoFactor"`
	TasaOCuota *float64 `xml:"TasaOCuotaDR,attr" bson:"TasaOCuota,omitempty"`
	Importe    *float64 `xml:"ImporteDR,attr" bson:"Importe,omitempty"`
}

type ImpuestosPagos20 struct {
	Retenciones *[]RetencionPagos20 `xml:"RetencionesP>RetencionP" bson:"Retenciones,omitempty"`
	Traslados   *[]TrasladoPagos20  `xml:"TrasladosP>TrasladoP" bson:"Traslados,omitempty"`
}

type RetencionPagos20 struct {
	Impuesto string  `xml:"ImpuestoP,attr" bson:"Impuesto"`
	Importe  float64 `xml:"ImporteP,attr" bson:"Importe"`
}

type TrasladoPagos20 struct {
	Base       float64  `xml:"BaseP,attr" bson:"Base"`
	Impuesto   string   `xml:"ImpuestoP,attr" bson:"Impuesto"`
	TipoFactor string   `xml:"TipoFactorP,attr" bson:"TipoFactor"`
	TasaOCuota *float64 `xml:"TasaOCuotaP,attr" bson:"TasaOCuota,omitempty"`
	Importe    *float64 `xml:"ImporteP,attr" bson:"Importe,omitempty"`
}

func (p *PagoPagos20) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias PagoPagos20
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*p = PagoPagos20(aux)

	fecha, err := helpers.ParseDatetime(aux.FechaPagoString)
	if err != nil {
		return err
	}
	p.FechaPago = fecha

	return nil
}
