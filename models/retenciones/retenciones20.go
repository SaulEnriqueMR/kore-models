package retenciones

import (
	"encoding/xml"
	"strconv"
	"strings"
	"time"

	"github.com/SaulEnriqueMR/kore-models/models"

	"github.com/SaulEnriqueMR/kore-models/models/documentofiscaldigital"
	date "github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type Retenciones20 struct {
	documentofiscaldigital.DocumentoFiscalDigital
	Version              string                  `xml:"Version,attr" bson:"Version" json:"Version"`
	Folio                *string                 `xml:"FolioInt,attr" bson:"Folio,omitempty" json:"Folio,omitempty"`
	Sello                string                  `xml:"Sello,attr" bson:"Sello" json:"Sello"`
	NoCertificado        string                  `xml:"NoCertificado,attr" bson:"NoCertificado" json:"NoCertificado"`
	Certificado          string                  `xml:"Certificado,attr" bson:"Certificado" json:"Certificado"`
	Fecha                string                  `xml:"FechaExp,attr" bson:"Fecha" json:"Fecha"`
	LugarExpedicion      string                  `xml:"LugarExpRetenc,attr" bson:"LugarExpedicion" json:"LugarExpedicion"`
	ClaveRetencion       string                  `xml:"CveRetenc,attr" bson:"ClaveRetencion" json:"ClaveRetencion"`
	DescripcionRetencion *string                 `xml:"DescRetenc,attr" bson:"DescripcionRetencion,omitempty" json:"DescripcionRetencion,omitempty"`
	RetencionRelacionada *RetencionRelacionada20 `xml:"CfdiRetenRelacionados" bson:"RetencionRelacionada,omitempty" json:"RetencionRelacionada,omitempty"`
	Emisor               Emisor20                `xml:"Emisor" bson:"Emisor" json:"Emisor"`
	Receptor             Receptor20              `xml:"Receptor" bson:"Receptor" json:"Receptor"`
	Periodo              Periodo20               `xml:"Periodo" bson:"Periodo" json:"Periodo"`
	Totales              Totales20               `xml:"Totales" bson:"Totales" json:"Totales"`
	Complemento          *Complemento            `xml:"Complemento" bson:"Complemento,omitempty" json:"Complemento,omitempty"`
}

func (c *Retenciones20) DefineTransaccion(rfc string) {
	if c.Receptor.Nacional.Rfc == rfc {
		c.Transaccion = "RECIBIDO"
	}
	if c.Emisor.Rfc == rfc {
		c.Transaccion = "EMITIDO"
	}
	if c.Receptor.Nacional.Rfc == rfc && c.Emisor.Rfc == rfc {
		c.Transaccion = "AUTOFACTURACION"
	}
}

type RetencionRelacionada20 struct {
	TipoRelacion string `xml:"TipoRelacion,attr" bson:"TipoRelacion" json:"TipoRelacion"`
	UUID         string `xml:"UUID,attr" bson:"UUID" json:"UUID"`
	Uuid         string `bson:"Uuid" json:"Uuid"`
}

func (rr *RetencionRelacionada20) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias RetencionRelacionada20
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	*rr = RetencionRelacionada20(aux)
	rr.Uuid = strings.ToUpper(aux.UUID)

	return nil
}

type Emisor20 struct {
	Rfc           string `xml:"RfcE,attr" bson:"Rfc" json:"Rfc"`
	Nombre        string `xml:"NomDenRazSocE,attr" bson:"Nombre" json:"Nombre"`
	RegimenFiscal string `xml:"RegimenFiscalE,attr" bson:"RegimenFiscal" json:"RegimenFiscal"`
}

type Receptor20 struct {
	Nacionalidad string        `xml:"NacionalidadR,attr" bson:"Nacionalidad" json:"Nacionalidad"`
	Nacional     *Nacional20   `xml:"Nacional" bson:"Nacional" json:"Nacional"`
	Extranjero   *Extranjero20 `xml:"Extranjero" bson:"Extranjero" json:"Extranjero"`
}

type Nacional20 struct {
	Rfc             string  `xml:"RfcR,attr" bson:"Rfc" json:"Rfc"`
	Nombre          string  `xml:"NomDenRazSocR,attr" bson:"Nombre" json:"Nombre"`
	Curp            *string `xml:"CurpR,attr" bson:"Curp,omitempty" json:"Curp,omitempty"`
	DomicilioFiscal string  `xml:"DomicilioFiscalR,attr" bson:"DomicilioFiscal" json:"DomicilioFiscal"`
}

type Extranjero20 struct {
	NumRegIdTrib *string `xml:"NumRegIdTribR,attr" bson:"NumRegIdTrib,omitempty" json:"NumRegIdTrib,omitempty"`
	Nombre       string  `xml:"NomDenRazSocR,attr" bson:"Nombre" json:"Nombre"`
}

type Periodo20 struct {
	MesInicio string `xml:"MesIni,attr" bson:"MesInicio" json:"MesInicio"`
	MesFin    string `xml:"MesFin,attr" bson:"MesFin" json:"MesFin"`
	Ejercicio string `xml:"Ejercicio,attr" bson:"Ejercicio" json:"Ejercicio"`
}

type Totales20 struct {
	MontoTotalOperacion float64               `xml:"MontoTotOperacion,attr" bson:"MontoTotalOperacion" json:"MontoTotalOperacion"`
	MontoTotalGravado   float64               `xml:"MontoTotGrav,attr" bson:"MontoTotalGravado" json:"MontoTotalGravado"`
	MontoTotalExento    float64               `xml:"MontoTotExent,attr" bson:"MontoTotalExento" json:"MontoTotalExento"`
	MontoTotalRetenido  float64               `xml:"MontoTotRet,attr" bson:"MontoTotalRetenido" json:"MontoTotalRetenido"`
	UtilidadBimestral   *float64              `xml:"UtilidadBimestral,attr" bson:"UtilidadBimestral,omitempty" json:"UtilidadBimestral,omitempty"`
	IsrCorrespondiente  *float64              `xml:"ISRCorrespondiente,attr" bson:"IsrCorrespondiente,omitempty" json:"IsrCorrespondiente,omitempty"`
	ImpuestosRetenidos  *[]ImpuestoRetenido20 `xml:"ImpRetenidos" bson:"Impuestos,omitempty" json:"Impuestos,omitempty"`
}

type ImpuestoRetenido20 struct {
	Base     *float64 `xml:"BaseRet,attr" bson:"Base,omitempty" json:"Base,omitempty"`
	Impuesto *string  `xml:"ImpuestoRet,attr" bson:"Impuesto,omitempty" json:"Impuesto,omitempty"`
	Monto    float64  `xml:"MontoRet,attr" bson:"Monto" json:"Monto"`
	TipoPago string   `xml:"TipoPagoRet,attr" bson:"TipoPago" json:"TipoPago"`
}

func (r *Retenciones20) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias Retenciones20
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	fechaEmision, err := date.ParseDatetime(aux.Fecha)
	if err != nil {
		return err
	}

	*r = Retenciones20(aux)
	r.FechaEmision = fechaEmision
	if r.Complemento != nil {
		if r.Complemento.TimbreFiscalDigital != nil {
			tfd11 := r.Complemento.TimbreFiscalDigital.TimbreFiscalDigital11
			if tfd11 != nil {
				r.FechaTimbrado = tfd11.FechaTimbrado
				r.Uuid = strings.ToUpper(tfd11.Uuid)
			}
		}
	}

	// processDate := time.Now().UTC()
	now := time.Now()
	processDate := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, time.UTC)
	r.ProcessorMetadata.LastUpdate = &processDate
	r.ProcessorMetadata.KoreModelsVersion = &models.KoreModelVersion

	return nil
}

func (c Retenciones20) GetBasePath() string {
	year := c.FechaEmision.Year()
	month := c.FechaEmision.Month()
	sb := strings.Builder{}
	sb.WriteString(c.Emisor.Rfc)
	sb.WriteString("/")
	if c.Receptor.Nacional.Rfc != "" {
		sb.WriteString(c.Receptor.Nacional.Rfc)
		sb.WriteString("/")
	}
	if *c.Receptor.Extranjero.NumRegIdTrib != "" {
		sb.WriteString(*c.Receptor.Extranjero.NumRegIdTrib)
		sb.WriteString("/")
	}
	sb.WriteString(strconv.Itoa(year))
	sb.WriteString("/")
	sb.WriteString(strconv.Itoa(int(month)))
	sb.WriteString("/")
	sb.WriteString(c.Uuid)
	return sb.String()
}

/* func (c *Retenciones20) SetFilePaths() {
	basePath := c.GetBasePath()
	c.XmlPath = strings.Join([]string{basePath, "xml"}, ".")
	c.PdfPath = strings.Join([]string{basePath, "pdf"}, ".")
} */
