package pagos

type Pagos10 struct {
	Version string        `xml:"Version,attr" bson:"Version"`
	Pago    []PagoPagos10 `xml:"Pago" bson:"Pago"`
}

type PagoPagos10 struct {
	FechaPago        string                     `xml:"FechaPago,attr" bson:"FechaPago"`
	FormaDePagoP     string                     `xml:"FormaDePagoP,attr" bson:"FormaPagoP"`
	MonedaP          string                     `xml:"MonedaP,attr" bson:"MonedaP"`
	TipoCambioP      *float64                   `xml:"TipoCambioP,attr" bson:"TipoCambioP,omitempty"`
	Monto            float64                    `xml:"Monto,attr" bson:"Monto"`
	NumOperacion     *string                    `xml:"NumOperacion,attr" bson:"NumOperacion,omitempty"`       // Cifrado
	RfcEmisorCtaOrd  *string                    `xml:"RfcEmisorCtaOrd,attr" bson:"RfcEmisorCtaOrd,omitempty"` // Cifrado
	NomBancoOrdExt   *string                    `xml:"NomBancoOrdExt,attr" bson:"NomBancoOrdExt,omitempty"`   // Cifrado
	CtaOrdenante     *string                    `xml:"CtaOrdenante,attr" bson:"CtaOrdenante,omitempty"`       // Cifrado
	RfcEmisorCtaBen  *string                    `xml:"RfcEmisorCtaBen,attr" bson:"RfcEmisorCtaBen,omitempty"` // Cifrado
	CtaBeneficiario  *string                    `xml:"CtaBeneficiario,attr" bson:"CtaBeneficiario,omitempty"` // Cifrado
	TipoCadPago      *string                    `xml:"TipoCadPago,attr" bson:"TipoCadPago,omitempty"`
	CertPago         *string                    `xml:"CertPago,attr" bson:"CertPago,omitempty"`
	CadPago          *string                    `xml:"CadPago,attr" bson:"CadPago,omitempty"`
	SelloPago        *string                    `xml:"SelloPago,attr" bson:"SelloPago,omitempty"`
	DoctoRelacionado *[]DoctoRelacionadoPagos10 `xml:"DoctoRelacionado" bson:"DoctoRelacionado,omitempty"`
	Impuestos        *[]ImpuestosPagos10        `xml:"Impuestos" bson:"Impuestos,omitempty"`
}

type DoctoRelacionadoPagos10 struct {
	IdDocumento      string   `xml:"IdDocumento,attr" bson:"IdDocumento"`
	Serie            *string  `xml:"Serie,attr" bson:"Serie,omitempty"`
	Folio            *string  `xml:"Folio,attr" bson:"Folio,omitempty"`
	MonedaDR         string   `xml:"MonedaDR,attr" bson:"MonedaDR"`
	TipoCambioDR     *float64 `xml:"TipoCambioDR,attr" bson:"TipoCambioDR,omitempty"`
	MetodoDePagoDR   string   `xml:"MetodoDePagoDR,attr" bson:"MonedaPagoDR"`
	NumParcialidad   *float64 `xml:"NumParcialidad,attr" bson:"NumParcialidad,omitempty"`
	ImpSaldoAnt      *float64 `xml:"ImpSaldoAnt,attr" bson:"ImpSaldoAnt,omitempty"`
	ImpPagado        *float64 `xml:"ImpPagado,attr" bson:"ImpPagado,omitempty"`
	ImpSaldoInsoluto *float64 `xml:"ImpSaldoInsoluto,attr" bson:"ImpSaldoInsoluto,omitempty"`
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
