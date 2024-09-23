package comprobante

import (
	"time"
)

type Comprobante32 struct {
	Version              string       `xml:"version,attr" bson:"Version"`
	Serie                *string      `xml:"serie,attr"  bson:"Serie,omitempty"`
	Folio                *string      `xml:"folio,attr" bson:"Folio,omitempty"`
	Fecha                string       `xml:"fecha,attr" bson:"Fecha"`
	Sello                string       `xml:"sello,attr" bson:"Sello"`
	FormaPago            string       `xml:"formaDePago,attr" bson:"FormaPago"`
	NoCertificado        string       `xml:"noCertificado,attr" bson:"NoCertificado"`
	Certificado          string       `xml:"certificado,attr" bson:"Certificado"`
	CondicionesPago      *string      `xml:"condicionesDePago,attr" bson:"CondicionesPago,omitempty"`
	Subtotal             float64      `xml:"subTotal,attr" bson:"Subtotal"`
	Descuento            *float64     `xml:"descuento,attr" bson:"Descuento,omitempty"`
	MotivoDescuento      *string      `xml:"motivoDescuento,attr" bson:"MotivoDescuento,omitempty"`
	TipoCambio           *string      `xml:"TipoCambio,attr" bson:"TipoCambio,omitempty"`
	Moneda               *string      `xml:"Moneda,attr" bson:"Moneda,omitempty"`
	Total                float64      `xml:"total,attr" bson:"Total"`
	TipoComprobante      string       `xml:"tipoDeComprobante,attr" bson:"TipoComprobante"`
	MetodoPago           string       `xml:"metodoDePago,attr" bson:"MetodoPago"`
	LugarExpedicion      string       `xml:"LugarExpedicion,attr" bson:"LugarExpedicion"`
	NumeroCuentaPago     *string      `xml:"NumCtaPago,attr" bson:"NumeroCuentaPago,omitempty"`
	FolioFiscalOrig      *string      `xml:"FolioFiscalOrig,attr" bson:"FolioFiscalOriginal,omitempty"`
	SerieFolioFiscalOrig *string      `xml:"SerieFolioFiscalOrig,attr" bson:"SerieFolioFiscalOriginal,omitempty"`
	FechaFolioFiscalOrig *string      `xml:"FechaFolioFiscalOrig,attr" bson:"FechaFolioFiscalOriginal,omitempty"`
	MontoFolioFiscalOrig *string      `xml:"MontoFolioFiscalOrig,attr" bson:"MontoFolioFiscalOriginal,omitempty"`
	Emisor               Emisor32     `xml:"Emisor" bson:"Emisor"`
	Receptor             Receptor32   `xml:"Receptor" bson:"Receptor"`
	Conceptos            []Concepto32 `xml:"Conceptos>Concepto" bson:"Conceptos"`
	Impuestos            Impuestos32  `xml:"Impuestos" bson:"Impuestos"`
	Complemento          *Complemento `xml:"Complemento" bson:"Complemento,omitempty"`
	/* Atributo convertido */
	FechaEmision time.Time `bson:"FechaEmision"`
	/* Atributos extraidos desde tfd */
	Uuid          string    `bson:"Uuid"`
	FechaTimbrado time.Time `bson:"FechaTimbrado"`
	/* Atributos adicionales, generalmente actualizados por fuentes externas */
	InformacionAdicional InformacionAdicional `xml:"InformacionAdicional" bson:"InformacionAdicional"`
	Cancelacion          Cancelacion          `xml:"Cancelacion" bson:"Cancelacion"`
	Vigente              bool                 `bson:"Vigente"`
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
	ComplementoConcepto ComplementoConcepto      `xml:"ComplementoConcepto" bson:"ComplementoConcepto,omitempty"`
	Parte               *[]Parte32               `xml:"Parte" bson:"Parte,omitempty"`
}

type InformacionAduanera32 struct {
	Numero string  `xml:"numero,attr" bson:"Numero"`
	Fecha  string  `xml:"fecha,attr" bson:"Fecha"`
	Aduana *string `xml:"aduana,attr" bson:"Aduana,omitempty"`
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
	Impuesto string  `xml:"impuesto,attr" bson:"Impuesto"`
	Importe  float64 `xml:"importe,attr" bson:"Importe"`
}

type Traslado32 struct {
	Impuesto string  `xml:"impuesto,attr" bson:"Impuesto"`
	Tasa     float64 `xml:"tasa,attr" bson:"Tasa"`
	Importe  float64 `xml:"importe,attr" bson:"Importe"`
}
