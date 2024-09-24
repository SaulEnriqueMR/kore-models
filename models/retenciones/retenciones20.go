package retenciones

import (
	"encoding/xml"
	"time"

	"github.com/SaulEnriqueMR/kore-models/app/timbrefiscaldigital"
	"github.com/SaulEnriqueMR/kore-models/models"
	"github.com/SaulEnriqueMR/kore-models/models/dividendos"
	"github.com/SaulEnriqueMR/kore-models/models/enajenaciondeacciones"
	"github.com/SaulEnriqueMR/kore-models/models/intereses"
	"github.com/SaulEnriqueMR/kore-models/models/pagosaextranjeros"
)

type Retenciones20 struct {
	Version              string                  `xml:"Version,attr" bson:"Version"`
	Folio                *string                 `xml:"FolioInt,attr" bson:"Folio,omitempty"`
	Sello                string                  `xml:"Sello,attr" bson:"Sello"`
	NoCertificado        string                  `xml:"NoCertificado,attr" bson:"NoCertificado"`
	Certificado          string                  `xml:"Certificado,attr" bson:"Certificado"`
	FechaExp             string                  `xml:"FechaExp,attr"`
	LugarExpedicion      string                  `xml:"LugarExpRetenc,attr" bson:"LugarExpedicion"`
	ClaveRetencion       string                  `xml:"CveRetenc,attr" bson:"ClaveRetencion"`
	DescripcionRetencion *string                 `xml:"DescRetenc,attr" bson:"DescripcionRetencion,omitempty"`
	RetencionRelacionada *RetencionRelacionada20 `xml:"CfdiRetenRelacionados" bson:"RetencionRelacionada,omitempty"`
	Emisor               Emisor20                `xml:"Emisor" bson:"Emisor"`
	Receptor             Receptor20              `xml:"Receptor" bson:"Receptor"`
	Periodo              Periodo20               `xml:"Periodo" bson:"Periodo"`
	Totales              Totales20               `xml:"Totales" bson:"Totales"`
	Complemento          *ComplementoReten20     `xml:"Complemento" bson:"Complemento,omitempty"`
	/* Atributo convertido */
	FechaEmision time.Time `bson:"FechaEmision"`
	/* Atributos extraidos desde tfd */
	Uuid          string    `bson:"Uuid"`
	FechaTimbrado time.Time `bson:"FechaTimbrado"`
}

type RetencionRelacionada20 struct {
	TipoRelacion string `xml:"TipoRelacion,attr" bson:"TipoRelacion"`
	Uuid         string `xml:"UUID,attr" bson:"Uuid"`
}

type Emisor20 struct {
	Rfc           string `xml:"RfcE,attr" bson:"Rfc"`
	Nombre        string `xml:"NomDenRazSocE,attr" bson:"Nombre"`
	RegimenFiscal string `xml:"RegimenFiscalE,attr" bson:"RegimenFiscal"`
}

type Receptor20 struct {
	Nacionalidad string        `xml:"NacionalidadR,attr" bson:"Nacionalidad"`
	Nacional     *Nacional20   `xml:"Nacional" bson:"Nacional"`
	Extranjero   *Extranjero20 `xml:"Extranjero" bson:"Extranjero"`
}

type Nacional20 struct {
	Rfc             string  `xml:"RfcR,attr" bson:"Rfc"`
	Nombre          string  `xml:"NomDenRazSocR,attr" bson:"Nombre"`
	Curp            *string `xml:"CurpR,attr" bson:"Curp,omitempty"`
	DomicilioFiscal string  `xml:"DomicilioFiscalR,attr" bson:"DomicilioFiscal"`
}

type Extranjero20 struct {
	NumRegIdTrib *string `xml:"NumRegIdTribR,attr" bson:"NumRegIdTrib,omitempty"`
	Nombre       string  `xml:"NomDenRazSocR,attr" bson:"Nombre"`
}

type Periodo20 struct {
	MesInicio string `xml:"MesIni,attr" bson:"MesInicio"`
	MesFin    string `xml:"MesFin,attr" bson:"MesFin"`
	Ejercicio string `xml:"Ejercicio,attr" bson:"Ejercicio"`
}

type Totales20 struct {
	MontoTotalOperacion float64               `xml:"MontoTotOperacion,attr" bson:"MontoTotalOperacion"`
	MontoTotalGravado   float64               `xml:"MontoTotGrav,attr" bson:"MontoTotalGravado"`
	MontoTotalExento    float64               `xml:"MontoTotExent,attr" bson:"MontoTotalExento"`
	MontoTotalRetenido  float64               `xml:"MontoTotRet,attr" bson:"MontoTotalRetenido"`
	UtilidadBimestral   *float64              `xml:"UtilidadBimestral,attr" bson:"UtilidadBimestral,omitempty"`
	IsrCorrespondiente  *float64              `xml:"ISRCorrespondiente,attr" bson:"IsrCorrespondiente,omitempty"`
	ImpuestosRetenidos  *[]ImpuestoRetenido20 `xml:"ImpRetenidos" bson:"ImpuestosRetenidos,omitempty"`
}

type ImpuestoRetenido20 struct {
	Base     *float64 `xml:"BaseRet,attr" bson:"BaseRet,omitempty"`
	Impuesto *string  `xml:"ImpuestoRet,attr" bson:"ImpuestoRet,omitempty"`
	Monto    float64  `xml:"MontoRet,attr" bson:"MontoRet"`
	TipoPago string   `xml:"TipoPagoRet,attr" bson:"TipoPagoRet"`
}

type ComplementoReten20 struct {
	Dividendos          *dividendos.Dividendos10                       `xml:"Dividendos" bson:"Dividendos,omitempty"`
	EnajenacionAcciones *enajenaciondeacciones.EnajenacionDeAcciones10 `xml:"EnajenaciondeAcciones" bson:"EnajenacionAcciones,omitempty"`
	Intereses           *intereses.Intereses10                         `xml:"Intereses" bson:"Intereses,omitempty"`
	PagosAExtranjeros   *pagosaextranjeros.Pagosaextranjeros10         `xml:"Pagosaextranjeros" bson:"PagosAExtranjeros,omitempty"`
	TimbreFiscalDigital *timbrefiscaldigital.TimbreFiscalDigital       `xml:"TimbreFiscalDigital" bson:"TimbreFiscalDigital,omitempty"`
}

func (r *Retenciones20) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias Retenciones20
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	fechaEmision, err := models.ParseDatetime(aux.FechaExp)
	if err != nil {
		return err
	}

	*r = Retenciones20(aux)
	r.FechaEmision = fechaEmision

	return nil
}
