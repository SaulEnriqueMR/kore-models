package documentofiscaldigital

import (
	"time"
)

type DocumentoFiscalDigital struct {
	FechaEmision         time.Time             `bson:"FechaEmision"`
	Uuid                 string                `bson:"Uuid"`
	FechaTimbrado        time.Time             `bson:"FechaTimbrado"`
	Vigente              *bool                 `bson:"Vigente,omitempty" json:"Vigente,omitempty"`
	CadenaOriginal       string                `bson:"CadenaOriginal"`
	InformacionAdicional *InformacionAdicional `bson:"InformacionAdicional,omitempty" json:"InformacionAdicional,omitempty"`
	Cancelacion          *Cancelacion          `bson:"Cancelacion,omitempty" json:"Cancelacion,omitempty"`
	Comprobante          bool                  `bson:"Comprobante"`
	Transaccion          string                `bson:"Transaccion"`
	XmlPath              string                `bson:"XmlPath"`
	PdfPath              string                `bson:"PdfPath"`
	TotalesMonedaLocal   TotalesMonedaLocal    `bson:"TotalesMonedaLocal"`
	ProcessorMetadata    ProcessorMetadata     `bson:"ProcessorMetadata"`
}

type ProcessorMetadata struct {
	KoreModelsVersion *string    `bson:"KoreModelsVersion"`
	KoreVersion       *string    `bson:"KoreVersion"`
	LastUpdate        *time.Time `bson:"LastUpdate"`
	CreationDate      *time.Time `bson:"CreationDate"`
	ZipPath           *string    `bson:"ZipPath"`
	XmlPath           *string    `bson:"XmlPath"`
	FileName          *string    `bson:"FileName"`
	FilePaths         *string    `bson:"FilePaths"`
}

type InformacionAdicional struct {
	StampedByKuantik *bool   `bson:"StampedByKuantik,omitempty" json:"StampedByKuantik,omitempty"`
	Comentario       *string `bson:"Comentario,omitempty"`
	Desgloce         *string `bson:"Desgloce,omitempty"`
	NotaPie          *string `bson:"NotaPie,omitempty"`
	IdProyecto       *string `bson:"IdProyecto,omitempty"`
	Categoria        *string `bson:"Categoria,omitempty"`
}

type Cancelacion struct {
	CanceledByKuantik  *bool      `bson:"CanceledByKuantik,omitempty" json:"CanceledByKuantik,omitempty"`
	MotivoCancelacion  *string    `bson:"MotivoCancelacion,omitempty"`
	FolioSustitucion   *string    `bson:"FolioSustitucion,omitempty"`
	EstatusCancelacion *string    `bson:"EstatusCancelacion,omitempty"`
	EsCancelable       *string    `bson:"EsCancelable,omitempty"`
	FechaCancelacion   *time.Time `bson:"FechaCancelacion,omitempty"`
}

type TotalesMonedaLocal struct {
	Total                     float64  `bson:"Total"`
	Subtotal                  float64  `bson:"Subtotal"`
	Descuento                 *float64 `bson:"Descuento,omitempty"`
	TotalImpuestosRetenidos   *float64 `bson:"TotalImpuestosRetenidos,omitempty"`
	TotalImpuestosTrasladados *float64 `bson:"TotalImpuestosTrasladados,omitempty"`
}
