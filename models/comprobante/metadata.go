package comprobante

import (
	"time"
)

type ComprobanteMetadata struct {
	Uuid            string            `bson:"Uuid" json:"Uuid"`
	Fecha           time.Time         `bson:"FechaEmision" json:"FechaEmision"`
	Vigente         bool              `bson:"Vigente" json:"Estatus"`
	FechaTimbrado   time.Time         `bson:"FechaTimbrado" json:"FechaTimbrado"`
	Total           float64           `bson:"Total" json:"Total"`
	TipoComprobante string            `bson:"TipoComprobante" json:"TipoComprobante"`
	Emisor          RfcEmisorReceptor `bson:"Emisor" json:"Emisor"`
	Receptor        RfcEmisorReceptor `bson:"Receptor" json:"Receptor"`
	RfcProvCertif   string            `bson:"RfcProvCertif" json:"RfcProvCertif"`
	Metadata        bool              `bson:"Metadata" json:"Metadata"`
	// FechaCancelacion *time.Time        `bson:"FechaCancelacion,omitempty" json:"FechaCancelacion,omitempty"`
	Cancelacion *Cancelacion `xml:"Cancelacion" bson:"Cancelacion,omitempty" json:"Cancelacion,omitempty"`
}

type Cancelacion struct {
	CanceledByKuantik *bool      `bson:"CanceledByKuantik,omitempty" json:"CanceledByKuantik,omitempty"`
	FechaCancelacion  *time.Time `bson:"FechaCancelacion,omitempty" json:"FechaCancelacion,omitempty"`
}
type RfcEmisorReceptor struct {
	Rfc    string `bson:"Rfc" json:"Rfc"`
	Nombre string `bson:"Nombre" json:"Nombre"`
}

type PagoTercerosMetadata struct {
	Uuid          string `bson:"Uuid" json:"UuidComprobante"`
	RfcTercero    string `bson:"RfcTercero" json:"RfcTercero"`
	NombreTercero string `bson:"NombreTercero" json:"NombreTercero"`
}

type ComprobanteTerceros struct {
	Uuid     string              `bson:"Uuid" json:"UuidComprobante"`
	Terceros []RfcEmisorReceptor `bson:"Terceros" json:"Terceros"`
}
