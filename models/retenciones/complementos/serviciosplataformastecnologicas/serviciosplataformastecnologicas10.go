package serviciosplataformastecnologicas

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"time"
)

type ServiciosPlataformasTecnologicas10 struct {
	Version                                  string                            `xml:"Version,attr" bson:"Version" json:"Version"`
	Periodicidad                             string                            `xml:"Periodicidad,attr" bson:"Periodicidad" json:"Periodicidad"`
	NoServicio                               string                            `xml:"NumServ,attr" bson:"NoServicio" json:"NoServicio"`
	MontoTotalServicioSinIva                 float64                           `xml:"MonTotServSIVA,attr" bson:"MontoTotalServicioSinIva" json:"MontoTotalServicioSinIva"`
	TotalIvaTrasladado                       float64                           `xml:"TotalIVATrasladado,attr" bson:"TotalIvaTrasladado" json:"TotalIvaTrasladado"`
	TotalIvaRetenido                         float64                           `xml:"TotalIVARetenido,attr" bson:"TotalIvaRetenido" json:"TotalIvaRetenido"`
	TotalIsrRetenido                         float64                           `xml:"TotalISRRetenido,attr" bson:"TotalIsrRetenido" json:"TotalIsrRetenido"`
	DiferenciaIvaEntregadoPrestadorServicios float64                           `xml:"DifIVAEntregadoPrestServ,attr" bson:"DiferenciaIvaEntregadoPrestadorServicios" json:"DiferenciaIvaEntregadoPrestadorServicios"`
	MontoTotalUsoPlataforma                  float64                           `xml:"MonTotalporUsoPlataforma,attr" bson:"MontoTotalUsoPlataforma" json:"MontoTotalUsoPlataforma"`
	MontoTotalContribucionGubernamental      *float64                          `xml:"MonTotalContribucionGubernamental,attr" bson:"MontoTotalContribucionGubernamental,omitempty" json:"MontoTotalContribucionGubernamental,omitempty"`
	DetallesDelServicio                      *[]DetallesDelServicioPlatTecno10 `xml:"Servicios>DetallesDelServicio" bson:"Servicios,omitempty" json:"Servicios,omitempty"`
}

type DetallesDelServicioPlatTecno10 struct {
	FormaPagoServicio            string                                   `xml:"FormaPagoServ,attr" bson:"FormaPagoServicio" json:"FormaPagoServicio"`
	TipoServicio                 string                                   `xml:"TipoDeServ,attr" bson:"TipoServicio" json:"TipoServicio"`
	SubtipoServicio              *string                                  `xml:"SubTipServ,attr" bson:"SubtipoServicio" json:"SubtipoServicio"`
	RfcTerceroAutorizado         *string                                  `xml:"RFCTerceroAutorizado,attr" bson:"RfcTerceroAutorizado,omitempty" json:"RfcTerceroAutorizado,omitempty"`
	FechaServ                    string                                   `xml:"FechaServ,attr" bson:"FechaServ" json:"FechaServ"`
	FechaServicio                time.Time                                `bson:"FechaServicio" json:"FechaServicio"`
	PrecioServicioSinIva         float64                                  `xml:"PrecioServSinIVA,attr" bson:"PrecioServicioSinIva" json:"PrecioServicioSinIva"`
	ImpuestosTrasladadosServicio ImpuestosTrasladosDelServicioPlatTecno10 `xml:"ImpuestosTrasladadosdelServicio" bson:"ImpuestosTrasladadosServicio,omitempty" json:"ImpuestosTrasladadosServicio,omitempty"`
	ContribucionGubernamental    *ContribucionGubernamentalPlatTecno10    `xml:"ContribucionGubernamental" bson:"ContribucionGubernamental,omitempty" json:"ContribucionGubernamental,omitempty"`
	ComisionServicio             ComisionDelServicioPlatTecno10           `xml:"ComisionDelServicio" bson:"ComisionServicio,omitempty" json:"ComisionServicio,omitempty"`
}

type ImpuestosTrasladosDelServicioPlatTecno10 struct {
	Base       float64 `xml:"Base,attr" bson:"Base" json:"Base"`
	Impuesto   string  `xml:"Impuesto,attr" bson:"Impuesto" json:"Impuesto"`
	TipoFactor string  `xml:"TipoFactor,attr" bson:"TipoFactor" json:"TipoFactor"`
	TasaOCuota float64 `xml:"TasaCuota,attr" bson:"TasaOCuota" json:"TasaOCuota"`
	Importe    float64 `xml:"Importe,attr" bson:"Importe" json:"Importe"`
}

type ContribucionGubernamentalPlatTecno10 struct {
	ImporteContribucion     float64 `xml:"ImpContrib,attr" bson:"ImporteContribucion" json:"ImporteContribucion"`
	EntidadPagoContribucion string  `xml:"EntidadDondePagaLaContribucion,attr" bson:"EntidadPagoContribucion" json:"EntidadPagoContribucion"`
}

type ComisionDelServicioPlatTecno10 struct {
	Base       *float64 `xml:"Base,attr" bson:"Base,omitempty" json:"Base,omitempty"`
	Porcentaje *float64 `xml:"Porcentaje,attr" bson:"Porcentaje,omitempty" json:"Porcentaje,omitempty"`
	Importe    *float64 `xml:"Importe,attr" bson:"Importe,omitempty" json:"Importe,omitempty"`
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
	fechaServ, err := helpers.ParseDatetime(aux.FechaServ)
	if err != nil {
		return err
	}

	// Assign the parsed date and other fields back to the original struct
	*t = DetallesDelServicioPlatTecno10(aux)
	t.FechaServicio = fechaServ

	return nil
}
