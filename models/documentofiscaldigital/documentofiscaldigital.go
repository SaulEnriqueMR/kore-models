package documentofiscaldigital

import (
	"encoding/json"
	"time"
)

type DocumentoFiscalDigital struct {
	FechaEmision  time.Time `bson:"FechaEmision"`
	Uuid          string    `bson:"Uuid"`
	FechaTimbrado time.Time `bson:"FechaTimbrado"`
	// Vigente              bool                 `bson:"Vigente"`
	CadenaOriginal       string               `bson:"CadenaOriginal"`
	InformacionAdicional InformacionAdicional `xml:"InformacionAdicional" bson:"InformacionAdicional"`
	Cancelacion          Cancelacion          `xml:"Cancelacion" bson:"Cancelacion"`
	Comprobante          bool                 `bson:"Comprobante"`
}

type InformacionAdicional struct {
	StampedByKuantik bool    `bson:"StampedByKuantik"`
	Comentario       *string `bson:"Comentario,omitempty"`
	Desgloce         *string `bson:"Desgloce,omitempty"`
	NotaPie          *string `bson:"NotaPie,omitempty"`
	IdProyecto       *string `bson:"IdProyecto,omitempty"`
	Categoria        *string `bson:"Categoria,omitempty"`
}

type Cancelacion struct {
	CanceledByKuantik  bool       `bson:"CanceledByKuantik"`
	MotivoCancelacion  *string    `bson:"MotivoCancelacion,omitempty"`
	FolioSustitucion   *string    `bson:"FolioSustitucion,omitempty"`
	EstatusCancelacion *string    `bson:"EstatusCancelacion,omitempty"`
	EsCancelable       *string    `bson:"EsCancelable,omitempty"`
	FechaCancelacion   *time.Time `bson:"FechaCancelacion,omitempty"`
}

func (d *DocumentoFiscalDigital) UnmarshalJSON(data []byte) error {
	type Alias DocumentoFiscalDigital
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(d),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	d.Comprobante = true
	return nil
}
