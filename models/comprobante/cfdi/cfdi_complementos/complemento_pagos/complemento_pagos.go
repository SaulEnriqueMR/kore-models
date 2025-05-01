package complementopagos

import "time"

type ComplementoPagos struct {
	Version string `json:"version"`
	Pagos   []Pago `json:"pagos"`
}

type Pago struct {
	FechaPago              time.Time              `json:"fechaPago"`
	FormaPago              string                 `json:"formaPago"`
	Moneda                 string                 `json:"moneda"`
	TipoCambio             *float64               `json:"tipoCambio"`
	Monto                  float64                `json:"monto"`
	NoOperacion            *string                `json:"noOperacion"`
	CuentaOrdenante        *CuentaOrdenante       `json:"cuentaOrdenante"`
	CuentaBeneficiario     *CuentaBeneficiario    `json:"cuentaBeneficiario"`
	TipoCadenaPago         *string                `json:"tipoCadenaPago"`
	CertificadoPago        *string                `json:"certificadoPago"`
	CadenaPago             *string                `json:"cadenaPago"`
	SelloPago              *string                `json:"selloPago"`
	DocumentosRelacionados []DocumentoRelacionado `json:"documentosRelacionados"`
	DetalleImpuestos       *DetalleImpuestos      `json:"detalleImpuestos"`
}

type CuentaOrdenante struct {
	RfcEmisorCuentaOrdenante       *string `json:"rfcEmisorCuentaOrdenante,omitempty"`
	NombreBancoOrdenanteExtranjero *string `json:"nombreBancoOrdenanteExtranjero,omitempty"`
	CuentaOrdenante                *string `json:"cuentaOrdenante,omitempty"`
}

type CuentaBeneficiario struct {
	RfcEmisorCuentaBeneficiario *string `json:"rfcEmisorCuentaBeneficiario,omitempty"`
	CuentaBeneficiario          *string `json:"cuentaBeneficiario,omitempty"`
}

type DocumentoRelacionado struct {
	IdDocumento          string     `json:"idDocumento"`
	IdDocumentoString    string     `json:"idDocumentoString"`
	Serie                *string    `json:"serie,omitempty"`
	Folio                *string    `json:"folio,omitempty"`
	Moneda               string     `json:"moneda"`
	Equivalencia         *float64   `json:"equivalencia,omitempty"`
	TipoCambio           *float64   `json:"tipoCambio,omitempty"`
	NoParcialidad        float64    `json:"noParcialidad"`
	ImporteSaldoAnterior float64    `json:"importeSaldoAnterior"`
	ImportePagado        float64    `json:"importePagado"`
	ImporteSaldoInsoluto float64    `json:"importeSaldoInsoluto"`
	MetodoPago           string     `json:"metodoPago"`
	ObjetoImpuesto       string     `json:"objetoImpuesto"`
	Impuestos            []Impuesto `json:"retenciones,omitempty"`
}

type Impuesto struct {
	Tipo       string  `json:"tipo"`
	Base       float64 `json:"base"`
	Impuesto   string  `json:"impuesto"`
	TipoFactor string  `json:"tipoFactor"`
	TasaOCuota float64 `json:"tasaOCuota"`
	Importe    float64 `json:"importe"`
}

type DetalleImpuestos struct {
	TotalImpuestosRetenidos   *float64   `json:"totalImpuestosRetenidos,omitempty"`
	TotalImpuestosTrasladados *float64   `json:"totalImpuestosTrasladados,omitempty"`
	Impuestos                 []Impuesto `json:"impuestos,omitempty"`
}
