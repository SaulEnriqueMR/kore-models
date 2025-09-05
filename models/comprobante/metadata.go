package comprobante

import (
	"time"

	documentofiscaldigital "github.com/SaulEnriqueMR/kore-models/models/documentofiscaldigital"
)

type ComprobanteMetadata struct {
	Uuid               string                                   `bson:"Uuid" json:"Uuid"`
	Fecha              time.Time                                `bson:"FechaEmision" json:"FechaEmision"`
	Vigente            bool                                     `bson:"Vigente" json:"Vigente"`
	FechaTimbrado      time.Time                                `bson:"FechaTimbrado" json:"FechaTimbrado"`
	Total              float64                                  `bson:"Total" json:"Total"`
	TipoComprobante    string                                   `bson:"TipoComprobante" json:"TipoComprobante"`
	Emisor             RfcEmisorReceptor                        `bson:"Emisor" json:"Emisor"`
	Receptor           RfcEmisorReceptor                        `bson:"Receptor" json:"Receptor"`
	RfcProvCertif      string                                   `bson:"RfcProvCertif" json:"RfcProvCertif"`
	Metadata           bool                                     `bson:"Metadata" json:"Metadata"`
	Cancelacion        *CancelacionMetadata                     `xml:"Cancelacion" bson:"Cancelacion,omitempty" json:"Cancelacion,omitempty"`
	Transaccion        string                                   `bson:"Transaccion" json:"Transaccion"`
	CreateDate         *time.Time                               `bson:"CreateDate,omitempty" json:"CreateDate,omitempty"` // --- IGNORE ---
	TotalesMonedaLocal TotalesMonedaLocalMetadata               `bson:"TotalesMonedaLocal" json:"TotalesMonedaLocal"`
	ProcessorMetadata  documentofiscaldigital.ProcessorMetadata `bson:"ProcessorMetadata" json:"ProcessorMetadata"`
	MetaProcessor      documentofiscaldigital.MetaProcessor     `bson:"MetaProcessor" json:"MetaProcessor"`
}

type CancelacionMetadata struct {
	CanceledByKuantik *bool      `bson:"CanceledByKuantik,omitempty" json:"CanceledByKuantik,omitempty"`
	FechaCancelacion  *time.Time `bson:"FechaCancelacion,omitempty" json:"FechaCancelacion,omitempty"`
}
type RfcEmisorReceptor struct {
	Rfc    string `bson:"Rfc" json:"Rfc"`
	Nombre string `bson:"Nombre" json:"Nombre"`
}

type PagoTercerosMetadata struct {
	Uuid          string `bson:"Uuid" json:"Uuid"`
	RfcTercero    string `bson:"RfcTercero" json:"RfcTercero"`
	NombreTercero string `bson:"NombreTercero" json:"NombreTercero"`
}

type ComprobanteTerceros struct {
	Uuid     string              `bson:"Uuid" json:"Uuid"`
	Terceros []RfcEmisorReceptor `bson:"Terceros" json:"Terceros"`
}

type TotalesMonedaLocalMetadata struct {
	Total    float64 `bson:"Total" json:"Total"`
	Subtotal float64 `bson:"Subtotal" json:"Subtotal"`
}
