package serviciosplataformastecnologicas

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"time"
)

type ServiciosPlataformasTecnologicas10 struct {
	Version                           string                            `xml:"Version,attr" bson:"Version"`
	Periodicidad                      string                            `xml:"Periodicidad,attr" bson:"Periodicidad"`
	NumServ                           string                            `xml:"NumServ,attr" bson:"NumServ"`
	MonTotServSIVA                    float64                           `xml:"MonTotServSIVA,attr" bson:"MonTotServSIVA"`
	TotalIVATrasladado                float64                           `xml:"TotalIVATrasladado,attr" bson:"TotalIVATrasladado"`
	TotalIVARetenido                  float64                           `xml:"TotalIVARetenido,attr" bson:"TotalIVARetenido"`
	TotalISRRetenido                  float64                           `xml:"TotalISRRetenido,attr" bson:"TotalISRRetenido"`
	DifIVAEntregadoPrestServ          float64                           `xml:"DifIVAEntregadoPrestServ,attr" bson:"DifIVAEntregadoPrestServ"`
	MonTotalporUsoPlataforma          float64                           `xml:"MonTotalporUsoPlataforma,attr" bson:"MonTotalporUsoPlataforma"`
	MonTotalContribucionGubernamental *float64                          `xml:"MonTotalContribucionGubernamental,attr" bson:"MonTotalContribucionGubernamental,omitempty"`
	DetallesDelServicio               *[]DetallesDelServicioPlatTecno10 `xml:"Servicios>DetallesDelServicio" bson:"Servicios,omitempty"`
}

type DetallesDelServicioPlatTecno10 struct {
	FormaPagoServ                   string                                   `xml:"FormaPagoServ,attr" bson:"FormaPagoServ"`
	TipoDeServ                      string                                   `xml:"TipoDeServ,attr" bson:"TipoDeServ"`
	SubTipServ                      *string                                  `xml:"SubTipServ,attr" bson:"SubTipServ"`
	RFCTerceroAutorizado            *string                                  `xml:"RFCTerceroAutorizado,attr" bson:"RFCTerceroAutorizado,omitempty"`
	FechaServString                 string                                   `xml:"FechaServ,attr"`
	FechaServ                       time.Time                                `bson:"FechaServ"`
	PrecioServSinIVA                float64                                  `xml:"PrecioServSinIVA,attr" bson:"PrecioServSinIVA"`
	ImpuestosTrasladadosDelServicio ImpuestosTrasladosDelServicioPlatTecno10 `xml:"ImpuestosTrasladadosdelServicio" bson:"ImpuestosTrasladadosDelServicio,omitempty"`
	ContribucionGubernamental       *ContribucionGubernamentalPlatTecno10    `xml:"ContribucionGubernamental" bson:"ContribucionGubernamental,omitempty"`
	ComisionDelServicio             ComisionDelServicioPlatTecno10           `xml:"ComisionDelServicio" bson:"ComisionDelServicio,omitempty"`
}

type ImpuestosTrasladosDelServicioPlatTecno10 struct {
	Base       float64 `xml:"Base,attr" bson:"Base"`
	Impuesto   string  `xml:"Impuesto,attr" bson:"Impuesto"`
	TipoFactor string  `xml:"TipoFactor,attr" bson:"TipoFactor"`
	TasaCuota  float64 `xml:"TasaCuota,attr" bson:"TasaCuota"`
	Importe    float64 `xml:"Importe,attr" bson:"Importe"`
}

type ContribucionGubernamentalPlatTecno10 struct {
	ImpContrib                     float64 `xml:"ImpContrib,attr" bson:"ImpContrib"`
	EntidadDondePagaLaContribucion string  `xml:"EntidadDondePagaLaContribucion,attr" bson:"EntidadDondePagaLaContribucion"`
}

type ComisionDelServicioPlatTecno10 struct {
	Base       *float64 `xml:"Base,attr" bson:"Base,omitempty"`
	Porcentaje *float64 `xml:"Porcentaje,attr" bson:"Porcentaje,omitempty"`
	Importe    *float64 `xml:"Importe,attr" bson:"Importe,omitempty"`
}

func (t *DetallesDelServicioPlatTecno10) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias DetallesDelServicioPlatTecno10
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	// Parse the date from the 'Fecha' field
	fechaServ, err := helpers.ParseDatetime(aux.FechaServString)
	if err != nil {
		return err
	}

	// Assign the parsed date and other fields back to the original struct
	*t = DetallesDelServicioPlatTecno10(aux)
	t.FechaServ = fechaServ

	return nil
}
