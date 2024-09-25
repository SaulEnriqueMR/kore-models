package retenciones

import (
	"encoding/xml"
	"time"

	"github.com/SaulEnriqueMR/kore-models/app/timbrefiscaldigital"
	"github.com/SaulEnriqueMR/kore-models/models"
	"github.com/SaulEnriqueMR/kore-models/models/retenciones/complementos/dividendos"
	"github.com/SaulEnriqueMR/kore-models/models/retenciones/complementos/enajenaciondeacciones"
	"github.com/SaulEnriqueMR/kore-models/models/retenciones/complementos/intereses"
	"github.com/SaulEnriqueMR/kore-models/models/retenciones/complementos/pagosaextranjeros"
)

type Retenciones10 struct {
	Version              string              `xml:"Version,attr" bson:"Version"`
	Folio                *string             `xml:"FolioInt,attr" bson:"Folio,omitempty"`
	Sello                string              `xml:"Sello,attr" bson:"Sello"`
	NoCertificado        string              `xml:"NumCert,attr" bson:"NoCertificado"`
	Certificado          string              `xml:"Cert,attr" bson:"Certificado"`
	FechaExp             string              `xml:"FechaExp,attr"`
	ClaveRetencion       string              `xml:"CveRetenc,attr" bson:"ClaveRetencion"`
	DescripcionRetencion *string             `xml:"DescRetenc,attr" bson:"DescripcionRetencion,omitempty"`
	Emisor               Emisor10            `xml:"Emisor" bson:"Emisor"`
	Receptor             Receptor10          `xml:"Receptor" bson:"Receptor"`
	Periodo              Periodo10           `xml:"Periodo" bson:"Periodo"`
	Totales              Totales10           `xml:"Totales" bson:"Totales"`
	Complemento          *ComplementoReten10 `xml:"Complemento" bson:"Complemento,omitempty"`

	/* Atributo convertido */
	FechaEmision time.Time `bson:"FechaEmision"`
	/* Atributos extraidos desde tfd */
	Uuid          string    `bson:"Uuid"`
	FechaTimbrado time.Time `bson:"FechaTimbrado"`
}

type Emisor10 struct {
	Rfc           string `xml:"RFCEmisor,attr" bson:"Rfc"`
	Nombre        string `xml:"NomDenRazSocE,attr" bson:"Nombre"`
	RegimenFiscal string `xml:"CURPE,attr" bson:"RegimenFiscal"`
}

type Receptor10 struct {
	Nacionalidad string        `xml:"Nacionalidad,attr" bson:"Nacionalidad"`
	Nacional     *Nacional10   `xml:"Nacional" bson:"Nacional"`
	Extranjero   *Extranjero10 `xml:"Extranjero" bson:"Extranjero"`
}

type Nacional10 struct {
	Rfc    string  `xml:"RFCRecep,attr" bson:"Rfc"`
	Nombre string  `xml:"NomDenRazSocR,attr" bson:"Nombre"`
	Curp   *string `xml:"CURPR,attr" bson:"Curp,omitempty"`
}

type Extranjero10 struct {
	NumRegIdTrib *string `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty"`
	Nombre       string  `xml:"NomDenRazSocR,attr" bson:"Nombre"`
}

type Periodo10 struct {
	MesInicio string `xml:"MesIni,attr" bson:"MesInicio"`
	MesFin    string `xml:"MesFin,attr" bson:"MesFin"`
	Ejercicio string `xml:"Ejerc,attr" bson:"Ejercicio"`
}

type ComplementoReten10 struct {
	Dividendos          *dividendos.Dividendos10                       `xml:"Dividendos" bson:"Dividendos,omitempty"`
	EnajenacionAcciones *enajenaciondeacciones.EnajenacionDeAcciones10 `xml:"EnajenacionDeAcciones" bson:"EnajenacionAcciones,omitempty"`
	Intereses           *intereses.Intereses10                         `xml:"Intereses" bson:"Intereses,omitempty"`
	PagosAExtranjeros   *pagosaextranjeros.PagosAExtranjeros10         `xml:"PagosAExtranjeros" bson:"PagosAExtranjeros,omitempty"`
	TimbreFiscalDigital *timbrefiscaldigital.TimbreFiscalDigital       `xml:"TimbreFiscalDigital" bson:"TimbreFiscalDigital,omitempty"`
}

type Totales10 struct {
	MontoTotalOperacion float64               `xml:"montoTotOperacion,attr" bson:"MontoTotalOperacion"`
	MontoTotalGravado   float64               `xml:"montoTotGrav,attr" bson:"MontoTotalGravado"`
	MontoTotalExento    float64               `xml:"montoTotExent,attr" bson:"MontoTotalExento"`
	MontoTotalRetenido  float64               `xml:"montoTotRet,attr" bson:"MontoTotalRetenido"`
	ImpuestosRetenidos  *[]ImpuestoRetenido10 `xml:"ImpRetenidos" bson:"ImpuestosRetenidos,omitempty"`
}

type ImpuestoRetenido10 struct {
	Base     *float64 `xml:"BaseRet,attr" bson:"BaseRet,omitempty"`
	Impuesto *string  `xml:"Impuesto,attr" bson:"ImpuestoRet,omitempty"`
	Monto    float64  `xml:"montoRet,attr" bson:"MontoRet"`
	TipoPago string   `xml:"TipoPagoRet,attr" bson:"TipoPagoRet"`
}

func (r *Retenciones10) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias Retenciones10
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	fechaEmision, err := models.ParseDatetime(aux.FechaExp)
	if err != nil {
		return err
	}

	*r = Retenciones10(aux)
	r.FechaEmision = fechaEmision

	return nil
}
