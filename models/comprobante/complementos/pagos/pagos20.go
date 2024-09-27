package pagos

import (
	"encoding/xml"
	"time"

	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type Pagos20 struct {
	Version string         `xml:"Version,attr" bson:"Version"`
	Totales TotalesPagos20 `xml:"Totales" bson:"Totales"`
	Pago    []PagoPagos20  `xml:"Pago" bson:"Pago"`
}

type TotalesPagos20 struct {
	TotalRetencionesIVA         *float64 `xml:"TotalRetencionesIVA,attr" bson:"TotalRetencionesIVA,omitempty"`
	TotalRetencionesISR         *float64 `xml:"TotalRetencionesISR,attr" bson:"TotalRetencionesISR,omitempty"`
	TotalRetencionesIEPS        *float64 `xml:"TotalRetencionesIEPS,attr" bson:"TotalRetencionesIEPS,omitempty"`
	TotalTrasladosBaseIVA16     *float64 `xml:"TotalTrasladosBaseIVA16,attr" bson:"TotalTrasladosBaseIVA16,omitempty"`
	TotalTrasladosImpuestoIVA16 *float64 `xml:"TotalTrasladosImpuestoIVA16,attr" bson:"TotalTrasladosImpuestoIVA16,omitempty"`
	TotalTrasladosBaseIVA8      *float64 `xml:"TotalTrasladosBaseIVA8,attr" bson:"TotalTrasladosBaseIVA8,omitempty"`
	TotalTrasladosImpuestoIVA8  *float64 `xml:"TotalTrasladosImpuestoIVA8,attr" bson:"TotalTrasladosImpuestoIVA8,omitempty"`
	TotalTrasladosBaseIVA0      *float64 `xml:"TotalTrasladosBaseIVA0,attr" bson:"TotalTrasladosBaseIVA0,omitempty"`
	TotalTrasladosImpuestoIVA0  *float64 `xml:"TotalTrasladosImpuestoIVA0,attr" bson:"TotalTrasladosImpuestoIVA0,omitempty"`
	TotalTrasladosBaseIVAExento *float64 `xml:"TotalTrasladosBaseIVAExento,attr" bson:"TotalTrasladosBaseIVAExento,omitempty"`
	MontoTotalPagos             float64  `xml:"MontoTotalPagos,attr"`
}

type PagoPagos20 struct {
	FechaPago        string                    `xml:"FechaPago,attr"`
	FechaPagoDate    time.Time                 `bson:"FechaPago"`
	FormaPagoP       string                    `xml:"FormaDePagoP,attr" bson:"FormaPagoP"`
	MonedaP          string                    `xml:"MonedaP,attr" bson:"MonedaP"`
	TipoCambioP      *float64                  `xml:"TipoCambioP,attr" bson:"TipoCambioP,omitempty"`
	Monto            float64                   `xml:"Monto,attr" bson:"Monto"`
	NumOperacion     *string                   `xml:"NumOperacion,attr" bson:"NumOperacion,omitempty"`       // Cifrado
	RfcEmisorCtaOrd  *string                   `xml:"RfcEmisorCtaOrd,attr" bson:"RfcEmisorCtaOrd,omitempty"` // Cifrado
	NomBancoOrdExt   *string                   `xml:"NomBancoOrdExt,attr" bson:"NomBancoOrdExt,omitempty"`   // Cifrado
	CtaOrdenante     *string                   `xml:"CtaOrdenante,attr" bson:"CtaOrdenante,omitempty"`       // Cifrado
	RfcEmisorCtaBen  *string                   `xml:"RfcEmisorCtaBen,attr" bson:"RfcEmisorCtaBen,omitempty"` // Cifrado
	CtaBeneficiario  *string                   `xml:"CtaBeneficiario,attr" bson:"CtaBeneficiario,omitempty"` // Cifrado
	TipoCadPago      *string                   `xml:"TipoCadPago,attr" bson:"TipoCadPago,omitempty"`
	CertPago         *string                   `xml:"CertPago,attr" bson:"CertPago,omitempty"`
	CadPago          *string                   `xml:"CadPago,attr" bson:"CadPago,omitempty"`
	SelloPago        *string                   `xml:"SelloPago,attr" bson:"SelloPago,omitempty"`
	DoctoRelacionado []DoctoRelacionadoPagos20 `xml:"DoctoRelacionado" bson:"DoctoRelacionado"`
	ImpuestosP       *ImpuestosPPagos20        `xml:"ImpuestosP" bson:"ImpuestosP,omitempty"`
}

type DoctoRelacionadoPagos20 struct {
	IdDocumento      string              `xml:"IdDocumento,attr" bson:"IdDocumento"`
	Serie            *string             `xml:"Serie,attr" bson:"Serie,omitempty"`
	Folio            *string             `xml:"Folio,attr" bson:"Folio,omitempty"`
	MonedaDR         string              `xml:"MonedaDR,attr" bson:"MonedaDR"`
	EquivalenciaDR   *float64            `xml:"EquivalenciaDR,attr" bson:"EquivalenciaDR,omitempty"`
	NumParcialidad   float64             `xml:"NumParcialidad,attr" bson:"NumParcialidad"`
	ImpSaldoAnt      float64             `xml:"ImpSaldoAnt,attr" bson:"ImpSaldoAnt"`
	ImpPagado        float64             `xml:"ImpPagado,attr" bson:"ImpPagado"`
	ImpSaldoInsoluto float64             `xml:"ImpSaldoInsoluto,attr" bson:"ImpSaldoInsoluto"`
	ObjetoImpDR      string              `xml:"ObjetoImpDR,attr" bson:"ObjetoImpDR"`
	ImpuestosDR      *ImpuestosDRPagos20 `xml:"ImpuestosDR" bson:"ImpuestosDR,omitempty"`
}

type ImpuestosDRPagos20 struct {
	RetencionesDR *[]RetencionDRPagos20 `xml:"RetencionesDR>RetencionDR" bson:"RetencionesDR,omitempty"`
	TrasladosDR   *[]TrasladoDRPagos20  `xml:"TrasladosDR>TrasladoDR" bson:"TrasladosDR,omitempty"`
}

type RetencionDRPagos20 struct {
	BaseDR       float64 `xml:"BaseDR,attr" bson:"BaseDR"`
	ImpuestoDR   string  `xml:"ImpuestoDR,attr" bson:"ImpuestoDR"`
	TipoFactorDR string  `xml:"TipoFactorDR,attr" bson:"TipoFactorDR"`
	TasaOCuotaDR float64 `xml:"TasaOCuotaDR,attr" bson:"TasaOCuotaDR"`
	ImporteDR    float64 `xml:"ImporteDR,attr" bson:"ImporteDR"`
}

type TrasladoDRPagos20 struct {
	BaseDR       float64  `xml:"BaseDR,attr" bson:"BaseDR"`
	ImpuestoDR   string   `xml:"ImpuestoDR,attr" bson:"ImpuestoDR"`
	TipoFactorDR string   `xml:"TipoFactorDR,attr" bson:"TipoFactorDR"`
	TasaOCuotaDR *float64 `xml:"TasaOCuotaDR,attr" bson:"TasaOCuotaDR,omitempty"`
	ImporteDR    *float64 `xml:"ImporteDR,attr" bson:"ImporteDR,omitempty"`
}

type ImpuestosPPagos20 struct {
	RetencionesP *[]RetencionPPagos20 `xml:"RetencionesP>RetencionP" bson:"RetencionesP,omitempty"`
	TrasladosP   *[]TrasladoPPagos20  `xml:"TrasladosP>TrasladoP" bson:"TrasladosP,omitempty"`
}

type RetencionPPagos20 struct {
	ImpuestoP string  `xml:"ImpuestoP,attr" bson:"ImpuestoP"`
	ImporteP  float32 `xml:"ImporteP,attr" bson:"ImporteP"`
}

type TrasladoPPagos20 struct {
	BaseP       float64  `xml:"BaseP,attr" bson:"BaseP"`
	ImpuestoP   string   `xml:"ImpuestoP,attr" bson:"ImpuestoP"`
	TipoFactorP string   `xml:"TipoFactorP,attr" bson:"TipoFactorP"`
	TasaOCuotaP *float64 `xml:"TasaOCuotaP,attr" bson:"TasaOCuotaP,omitempty"`
	ImporteP    *float64 `xml:"ImporteP,attr" bson:"ImporteP,omitempty"`
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

	fecha, err := helpers.ParseDatetime(aux.FechaPago)
	if err != nil {
		return err
	}
	p.FechaPagoDate = fecha

	return nil
}
