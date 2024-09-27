package serviciosplataformastecnologicas

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"time"
)

type ServiciosPlataformasTecnologicas10 struct {
	Version                                  string                            `xml:"Version,attr" bson:"Version"`
	Periodicidad                             string                            `xml:"Periodicidad,attr" bson:"Periodicidad"`
	NoServicio                               string                            `xml:"NumServ,attr" bson:"NoServicio"`
	MontoTotalServicioSinIva                 float64                           `xml:"MonTotServSIVA,attr" bson:"MontoTotalServicioSinIva"`
	TotalIvaTrasladado                       float64                           `xml:"TotalIVATrasladado,attr" bson:"TotalIvaTrasladado"`
	TotalIvaRetenido                         float64                           `xml:"TotalIVARetenido,attr" bson:"TotalIvaRetenido"`
	TotalIsrRetenido                         float64                           `xml:"TotalISRRetenido,attr" bson:"TotalIsrRetenido"`
	DiferenciaIvaEntregadoPrestadorServicios float64                           `xml:"DifIVAEntregadoPrestServ,attr" bson:"DiferenciaIvaEntregadoPrestadorServicios"`
	MontoTotalUsoPlataforma                  float64                           `xml:"MonTotalporUsoPlataforma,attr" bson:"MontoTotalUsoPlataforma"`
	MontoTotalContribucionGubernamental      *float64                          `xml:"MonTotalContribucionGubernamental,attr" bson:"MontoTotalContribucionGubernamental,omitempty"`
	DetallesDelServicio                      *[]DetallesDelServicioPlatTecno10 `xml:"Servicios>DetallesDelServicio" bson:"Servicios,omitempty"`
}

type DetallesDelServicioPlatTecno10 struct {
	FormaPagoServicio            string                                   `xml:"FormaPagoServ,attr" bson:"FormaPagoServicio"`
	TipoServicio                 string                                   `xml:"TipoDeServ,attr" bson:"TipoServicio"`
	SubtipoServicio              *string                                  `xml:"SubTipServ,attr" bson:"SubtipoServicio"`
	RfcTerceroAutorizado         *string                                  `xml:"RFCTerceroAutorizado,attr" bson:"RfcTerceroAutorizado,omitempty"`
	FechaServString              string                                   `xml:"FechaServ,attr"`
	FechaServicio                time.Time                                `bson:"FechaServicio"`
	PrecioServicioSinIva         float64                                  `xml:"PrecioServSinIVA,attr" bson:"PrecioServicioSinIva"`
	ImpuestosTrasladadosServicio ImpuestosTrasladosDelServicioPlatTecno10 `xml:"ImpuestosTrasladadosdelServicio" bson:"ImpuestosTrasladadosServicio,omitempty"`
	ContribucionGubernamental    *ContribucionGubernamentalPlatTecno10    `xml:"ContribucionGubernamental" bson:"ContribucionGubernamental,omitempty"`
	ComisionServicio             ComisionDelServicioPlatTecno10           `xml:"ComisionDelServicio" bson:"ComisionServicio,omitempty"`
}

type ImpuestosTrasladosDelServicioPlatTecno10 struct {
	Base       float64 `xml:"Base,attr" bson:"Base"`
	Impuesto   string  `xml:"Impuesto,attr" bson:"Impuesto"`
	TipoFactor string  `xml:"TipoFactor,attr" bson:"TipoFactor"`
	TasaOCuota float64 `xml:"TasaCuota,attr" bson:"TasaOCuota"`
	Importe    float64 `xml:"Importe,attr" bson:"Importe"`
}

type ContribucionGubernamentalPlatTecno10 struct {
	ImporteContribucion     float64 `xml:"ImpContrib,attr" bson:"ImporteContribucion"`
	EntidadPagoContribucion string  `xml:"EntidadDondePagaLaContribucion,attr" bson:"EntidadPagoContribucion"`
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
	t.FechaServicio = fechaServ

	return nil
}
