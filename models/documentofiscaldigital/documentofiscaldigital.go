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
	KoreModelsVersion  *string             `bson:"KoreModelsVersion" json:"KoreModelsVersion"`
	KoreVersion        *string             `bson:"KoreVersion" json:"KoreVersion"`
	LastUpdate         *time.Time          `bson:"LastUpdate" json:"LastUpdate"`
	CreationDate       *time.Time          `bson:"CreationDate,omitempty" json:"CreationDate,omitempty"`
	ZipProcessor       *ZipProcessor       `bson:"ZipProcessor,omitempty" json:"ZipProcessor,omitempty"`
	XmlPathProcessor   *XmlPathProcessor   `bson:"XmlPathProcessor,omitempty" json:"XmlPathProcessor,omitempty"`
	XmlFileProcessor   *XmlFileProcessor   `bson:"XmlFileProcessor,omitempty" json:"XmlFileProcessor,omitempty"`
	FilePathsProcessor *FilePathsProcessor `bson:"FilePathsProcessor,omitempty" json:"FilePathsProcessor,omitempty"`
	MetaProcessor      *MetaProcessor      `bson:"MetaProcessor,omitempty" json:"MetaProcessor,omitempty"`
	// ZipPath           *string    `bson:"ZipPath,omitempty" json:"ZipPath,omitempty"`
	// XmlPath           *string    `bson:"XmlPath,omitempty" json:"XmlPath,omitempty"`
	// FileName          *string    `bson:"FileName,omitempty" json:"FileName,omitempty"`
	// FilePaths         *string    `bson:"FilePaths,omitempty" json:"FilePaths,omitempty"`
}
type Processor struct {
	Attempt    int       `bson:"Attempt,omitempty" json:"Attempt,omitempty"`
	LastUpdate time.Time `bson:"LastUpdate" json:"LastUpdate"`
}

type ZipProcessor struct {
	ZipPath   string `bson:"ZipPath" json:"ZipPath"`
	Processor `bson:",inline"`
}

type XmlPathProcessor struct {
	XmlPath   string `bson:"XmlPath" json:"XmlPath"`
	Processor `bson:",inline"`
}

type XmlFileProcessor struct {
	FileName  string `bson:"FileName" json:"FileName"`
	Processor `bson:",inline"`
}

type FilePathsProcessor struct {
	FilePaths string `bson:"FilePaths" json:"FilePaths"`
	Processor `bson:",inline"`
}

type MetaProcessor struct {
	MetadataPath string `bson:"MetadataPath" json:"MetadataPath"`
	Processor    `bson:",inline"`
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
