package nomina

import (
	"encoding/xml"
	"time"

	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

// Nomina11 Presente en Comprobante 3.2 y 3.3
type Nomina11 struct {
	Version                    string                `xml:"Version,attr" bson:"Version"`
	RegistroPatronal           *string               `xml:"RegistroPatronal,attr" bson:"RegistroPatronal,omitempty"`
	NoEmpleado                 string                `xml:"NumEmpleado,attr" bson:"NoEmpleado"`
	Curp                       string                `xml:"CURP,attr" bson:"Curp"`
	TipoRegimen                int                   `xml:"TipoRegimen,attr" bson:"TipoRegimen"`
	NumSeguridadSocial         *string               `xml:"NumSeguridadSocial,attr" bson:"NoSeguridadSocial,omitempty"`
	FechaPagoString            string                `xml:"FechaPago,attr" bson:"FechaPagoString"`
	FechaPago                  time.Time             `bson:"FechaPago"`
	FechaInicialPagoString     string                `xml:"FechaInicialPago,attr" bson:"FechaInicialPagoString"`
	FechaInicialPago           time.Time             `bson:"FechaInicialPago"`
	FechaFinalPagoString       string                `xml:"FechaFinalPago,attr" bson:"FechaFinalPagoString"`
	FechaFinalPago             time.Time             `bson:"FechaFinalPago"`
	NumerosDiasPagados         float64               `xml:"NumDiasPagados,attr" bson:"NumerosDiasPagados"`
	Departamento               *string               `xml:"Departamento,attr" bson:"Departamento,omitempty"`
	Clabe                      *string               `xml:"CLABE,attr" bson:"Clabe,omitempty"`
	Banco                      *int                  `xml:"Banco,attr" bson:"Banco,omitempty"`
	FechaInicioRelLaboral      *string               `xml:"FechaInicioRelLaboral,attr" bson:"FechaInicioRelLaboral,omitempty"`
	FechaInicioRelacionLaboral *time.Time            `bson:"FechaInicioRelacionLaboral,omitempty"`
	Antiguedad                 *int                  `xml:"Antiguedad,attr" bson:"Antiguedad,omitempty"`
	Puesto                     *string               `xml:"Puesto,attr" bson:"Puesto,omitempty"`
	TipoContrato               *string               `xml:"TipoContrato,attr" bson:"TipoContrato,omitempty"`
	TipoJornada                *string               `xml:"TipoJornada,attr" bson:"TipoJornada,omitempty"`
	PeriodicidadPago           string                `xml:"PeriodicidadPago,attr" bson:"PeriodicidadPago"`
	SalarioBaseCotizacion      *float64              `xml:"SalarioBaseCotApor,attr" bson:"SalarioBaseCotizacion,omitempty"`
	RiesgoPuesto               *int                  `xml:"RiesgoPuesto,attr" bson:"RiesgoPuesto,omitempty"`
	SalarioDiarioIntegrado     *float64              `xml:"SalarioDiarioIntegrado,attr" bson:"SalarioDiarioIntegrado,omitempty"`
	Percepciones               *PercepcionesNomina11 `xml:"Percepciones" bson:"Percepciones,omitempty"`
	Deducciones                *DeduccionesNomina11  `xml:"Deducciones" bson:"Deducciones,omitempty"`
	Incapacidades              *IncapacidadNomina11  `xml:"Incapacidades>Incapacidad" bson:"Incapacidades,omitempty"`
	HorasExtras                *HorasExtraNomina11   `xml:"HorasExtras>HorasExtra" bson:"HorasExtras,omitempty"`
}

type PercepcionesNomina11 struct {
	TotalGravado float64              `xml:"TotalGravado,attr" bson:"TotalGravado"`
	TotalExento  float64              `xml:"TotalExento,attr" bson:"TotalExento"`
	Percepcion   []PercepcionNomina11 `xml:"Percepcion" bson:"Percepcion"`
}

type PercepcionNomina11 struct {
	TipoPercepcion int     `xml:"TipoPercepcion,attr" bson:"Tipo"`
	Clave          string  `xml:"Clave,attr" bson:"Clave"`
	Concepto       string  `xml:"Concepto,attr" bson:"Concepto"`
	ImporteGravado float64 `xml:"ImporteGravado,attr" bson:"ImporteGravado"`
	ImporteExento  float64 `xml:"ImporteExento,attr" bson:"ImporteExento"`
}

type DeduccionesNomina11 struct {
	TotalGravado float64             `xml:"TotalGravado,attr" bson:"TotalGravado"`
	TotalExento  float64             `xml:"TotalExento,attr" bson:"TotalExento"`
	Deduccion    []DeduccionNomina11 `xml:"Deduccion" bson:"Deduccion"`
}

type DeduccionNomina11 struct {
	TipoDeduccion  int     `xml:"TipoDeduccion,attr" bson:"Tipo"`
	Clave          string  `xml:"Clave,attr" bson:"Clave"`
	Concepto       string  `xml:"Concepto,attr" bson:"Concepto"`
	ImporteGravado float64 `xml:"ImporteGravado,attr" bson:"ImporteGravado"`
	ImporteExento  float64 `xml:"ImporteExento,attr" bson:"ImporteExento"`
}

type IncapacidadNomina11 struct {
	Dias      float64 `xml:"DiasIncapacidad,attr" bson:"Dias"`
	Tipo      int     `xml:"TipoIncapacidad,attr" bson:"Tipo"`
	Descuento float64 `xml:"Descuento,attr" bson:"Descuento"`
}

type HorasExtraNomina11 struct {
	Dias          int     `xml:"Dias,attr" bson:"Dias"`
	TipoHoras     string  `xml:"TipoHoras,attr" bson:"TipoHoras"`
	HorasExtra    int     `xml:"HorasExtra,attr" bson:"HorasExtra"`
	ImportePagado float64 `xml:"ImportePagado,attr" bson:"ImportePagado"`
}

func (n *Nomina11) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias Nomina11
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*n = Nomina11(aux)

	fecha, err := helpers.ParseDatetime(aux.FechaPagoString)
	if err != nil {
		return err
	}
	n.FechaPago = fecha

	fecha1, err1 := helpers.ParseDatetime(aux.FechaInicialPagoString)
	if err1 != nil {
		return err1
	}
	n.FechaInicialPago = fecha1

	fecha2, err2 := helpers.ParseDatetime(aux.FechaFinalPagoString)
	if err2 != nil {
		return err2
	}
	n.FechaFinalPago = fecha2

	if aux.FechaInicioRelLaboral != nil {
		fechaIniRelLab, _ := helpers.ParseDatetime(*aux.FechaInicioRelLaboral)
		n.FechaInicioRelacionLaboral = &fechaIniRelLab
	}

	return nil
}
