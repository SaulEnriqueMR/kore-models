package documentofiscaldigital

import (
	"encoding/xml"
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

func (c *DocumentoFiscalDigital) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias DocumentoFiscalDigital
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	*c = DocumentoFiscalDigital(aux)

	c.Comprobante = true

	return nil
}
