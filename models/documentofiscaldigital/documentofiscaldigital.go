package documentofiscaldigital

import (
	"time"
)

type DocumentoFiscalDigital struct {
	FechaEmision         time.Time             `bson:"FechaEmision"`
	Uuid                 string                `bson:"Uuid"`
	FechaTimbrado        time.Time             `bson:"FechaTimbrado"`
	Vigente              *bool                 `bson:"Vigente,omitempty"`
	CadenaOriginal       string                `bson:"CadenaOriginal"`
	InformacionAdicional *InformacionAdicional `xml:"InformacionAdicional" bson:"InformacionAdicional"`
	Cancelacion          *Cancelacion          `xml:"Cancelacion" bson:"Cancelacion"`
	Comprobante          bool                  `bson:"Comprobante"`
}

type InformacionAdicional struct {
	StampedByKuantik *bool   `bson:"StampedByKuantik,omitempty"` // Add this field to Kore
	Comentario       *string `bson:"Comentario,omitempty"`
	Desgloce         *string `bson:"Desgloce,omitempty"`
	NotaPie          *string `bson:"NotaPie,omitempty"`
	IdProyecto       *string `bson:"IdProyecto,omitempty"`
	Categoria        *string `bson:"Categoria,omitempty"`
}

type Cancelacion struct {
	CanceledByKuantik  *bool      `bson:"CanceledByKuantik,omitempty"` // Add this field to Kore
	MotivoCancelacion  *string    `bson:"MotivoCancelacion,omitempty"`
	FolioSustitucion   *string    `bson:"FolioSustitucion,omitempty"`
	EstatusCancelacion *string    `bson:"EstatusCancelacion,omitempty"`
	EsCancelable       *string    `bson:"EsCancelable,omitempty"`
	FechaCancelacion   *time.Time `bson:"FechaCancelacion,omitempty"`
}
