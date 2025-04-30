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

type Retenciones10 struct {
	documentofiscaldigital.DocumentoFiscalDigital
	Version              string       `xml:"Version,attr" bson:"Version" json:"Version"`
	Folio                *string      `xml:"FolioInt,attr" bson:"Folio,omitempty" json:"Folio,omitempty"`
	Sello                string       `xml:"Sello,attr" bson:"Sello" json:"Sello"`
	NoCertificado        string       `xml:"NumCert,attr" bson:"NoCertificado" json:"NoCertificado"`
	Certificado          string       `xml:"Cert,attr" bson:"Certificado" json:"Certificado"`
	Fecha                string       `xml:"FechaExp,attr" bson:"Fecha" json:"Fecha"`
	ClaveRetencion       string       `xml:"CveRetenc,attr" bson:"ClaveRetencion" json:"ClaveRetencion"`
	DescripcionRetencion *string      `xml:"DescRetenc,attr" bson:"DescripcionRetencion,omitempty" json:"DescripcionRetencion,omitempty"`
	Emisor               Emisor10     `xml:"Emisor" bson:"Emisor" json:"Emisor"`
	Receptor             Receptor10   `xml:"Receptor" bson:"Receptor" json:"Receptor"`
	Periodo              Periodo10    `xml:"Periodo" bson:"Periodo" json:"Periodo"`
	Totales              Totales10    `xml:"Totales" bson:"Totales" json:"Totales"`
	Complemento          *Complemento `xml:"Complemento" bson:"Complemento,omitempty" json:"Complemento,omitempty"`
}

func (c *Retenciones10) DefineTransaccion(rfc string) {
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

type Emisor10 struct {
	Rfc           string `xml:"RFCEmisor,attr" bson:"Rfc" json:"Rfc"`
	Nombre        string `xml:"NomDenRazSocE,attr" bson:"Nombre" json:"Nombre"`
	RegimenFiscal string `xml:"CURPE,attr" bson:"RegimenFiscal" json:"RegimenFiscal"`
}

type Receptor10 struct {
	Nacionalidad string        `xml:"Nacionalidad,attr" bson:"Nacionalidad" json:"Nacionalidad"`
	Nacional     *Nacional10   `xml:"Nacional" bson:"Nacional" json:"Nacional"`
	Extranjero   *Extranjero10 `xml:"Extranjero" bson:"Extranjero" json:"Extranjero"`
}

type Nacional10 struct {
	Rfc    string  `xml:"RFCRecep,attr" bson:"Rfc" json:"Rfc"`
	Nombre string  `xml:"NomDenRazSocR,attr" bson:"Nombre" json:"Nombre"`
	Curp   *string `xml:"CURPR,attr" bson:"Curp,omitempty" json:"Curp,omitempty"`
}

type Extranjero10 struct {
	NumRegIdTrib *string `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty" json:"NumRegIdTrib,omitempty"`
	Nombre       string  `xml:"NomDenRazSocR,attr" bson:"Nombre" json:"Nombre"`
}

type Periodo10 struct {
	MesInicio string `xml:"MesIni,attr" bson:"MesInicio" json:"MesInicio"`
	MesFin    string `xml:"MesFin,attr" bson:"MesFin" json:"MesFin"`
	Ejercicio string `xml:"Ejerc,attr" bson:"Ejercicio" json:"Ejercicio"`
}

type Totales10 struct {
	MontoTotalOperacion float64               `xml:"montoTotOperacion,attr" bson:"MontoTotalOperacion" json:"MontoTotalOperacion"`
	MontoTotalGravado   float64               `xml:"montoTotGrav,attr" bson:"MontoTotalGravado" json:"MontoTotalGravado"`
	MontoTotalExento    float64               `xml:"montoTotExent,attr" bson:"MontoTotalExento" json:"MontoTotalExento"`
	MontoTotalRetenido  float64               `xml:"montoTotRet,attr" bson:"MontoTotalRetenido" json:"MontoTotalRetenido"`
	ImpuestosRetenidos  *[]ImpuestoRetenido10 `xml:"ImpRetenidos" bson:"Impuestos,omitempty" json:"Impuestos,omitempty"`
}

type ImpuestoRetenido10 struct {
	Base     *float64 `xml:"BaseRet,attr" bson:"Base,omitempty" json:"Base,omitempty"`
	Impuesto *string  `xml:"Impuesto,attr" bson:"Impuesto,omitempty" json:"Impuesto,omitempty"`
	Monto    float64  `xml:"montoRet,attr" bson:"Monto" json:"Monto"`
	TipoPago string   `xml:"TipoPagoRet,attr" bson:"TipoPago" json:"TipoPago"`
}

func (r *Retenciones10) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias Retenciones10
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	fechaEmision, err := date.ParseDatetime(aux.Fecha)
	if err != nil {
		return err
	}

	*r = Retenciones10(aux)
	r.FechaEmision = fechaEmision

	if r.Complemento != nil {
		if r.Complemento.TimbreFiscalDigital != nil {
			tfd10 := r.Complemento.TimbreFiscalDigital.TimbreFiscalDigital10
			if tfd10 != nil {
				r.FechaTimbrado = tfd10.FechaTimbrado
				r.Uuid = strings.ToUpper(tfd10.Uuid)
			}

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

func (c Retenciones10) GetBasePath() string {
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

/* func (c *Retenciones10) SetFilePaths() {
	basePath := c.GetBasePath()
	c.XmlPath = strings.Join([]string{basePath, "xml"}, ".")
	c.PdfPath = strings.Join([]string{basePath, "pdf"}, ".")
} */
