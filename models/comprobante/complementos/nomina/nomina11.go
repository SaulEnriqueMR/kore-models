package nomina

import (
	"encoding/json"
	"encoding/xml"
	"time"

	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

// Nomina11 Presente en Comprobante 3.2 y 3.3
type Nomina11 struct {
	Version                    string                 `xml:"Version,attr" bson:"Version" json:"Version"`
	RegistroPatronal           *string                `xml:"RegistroPatronal,attr" bson:"RegistroPatronal,omitempty" json:"RegistroPatronal,omitempty"`
	NoEmpleado                 string                 `xml:"NumEmpleado,attr" bson:"NoEmpleado" json:"NoEmpleado"`
	Curp                       string                 `xml:"CURP,attr" bson:"Curp" json:"Curp"`
	TipoRegimen                int                    `xml:"TipoRegimen,attr" bson:"TipoRegimen" json:"TipoRegimen"`
	NumSeguridadSocial         *string                `xml:"NumSeguridadSocial,attr" bson:"NoSeguridadSocial,omitempty" json:"NoSeguridadSocial,omitempty"`
	FechaPagoString            string                 `xml:"FechaPago,attr" bson:"FechaPagoString" json:"FechaPagoString"`
	FechaPago                  time.Time              `bson:"FechaPago" json:"FechaPago"`
	FechaInicialPagoString     string                 `xml:"FechaInicialPago,attr" bson:"FechaInicialPagoString" json:"FechaInicialPagoString"`
	FechaInicialPago           time.Time              `bson:"FechaInicialPago" json:"FechaInicialPago"`
	FechaFinalPagoString       string                 `xml:"FechaFinalPago,attr" bson:"FechaFinalPagoString" json:"FechaFinalPagoString"`
	FechaFinalPago             time.Time              `bson:"FechaFinalPago" json:"FechaFinalPago"`
	NumerosDiasPagados         float64                `xml:"NumDiasPagados,attr" bson:"NumerosDiasPagados" json:"NumerosDiasPagados"`
	Departamento               *string                `xml:"Departamento,attr" bson:"Departamento,omitempty" json:"Departamento,omitempty"`
	Clabe                      *string                `xml:"CLABE,attr" bson:"Clabe,omitempty" json:"Clabe,omitempty"`
	Banco                      *int                   `xml:"Banco,attr" bson:"Banco,omitempty" json:"Banco,omitempty"`
	FechaInicioRelLaboral      *string                `xml:"FechaInicioRelLaboral,attr" bson:"FechaInicioRelLaboral,omitempty" json:"FechaInicioRelLaboral,omitempty"`
	FechaInicioRelacionLaboral *time.Time             `bson:"FechaInicioRelacionLaboral,omitempty" json:"FechaInicioRelacionLaboral,omitempty"`
	Antiguedad                 *int                   `xml:"Antiguedad,attr" bson:"Antiguedad,omitempty" json:"Antiguedad,omitempty"`
	Puesto                     *string                `xml:"Puesto,attr" bson:"Puesto,omitempty" json:"Puesto,omitempty"`
	TipoContrato               *string                `xml:"TipoContrato,attr" bson:"TipoContrato,omitempty" json:"TipoContrato,omitempty"`
	TipoJornada                *string                `xml:"TipoJornada,attr" bson:"TipoJornada,omitempty" json:"TipoJornada,omitempty"`
	PeriodicidadPago           string                 `xml:"PeriodicidadPago,attr" bson:"PeriodicidadPago" json:"PeriodicidadPago"`
	SalarioBaseCotizacion      *float64               `xml:"SalarioBaseCotApor,attr" bson:"SalarioBaseCotizacion,omitempty" json:"SalarioBaseCotizacion,omitempty"`
	RiesgoPuesto               *int                   `xml:"RiesgoPuesto,attr" bson:"RiesgoPuesto,omitempty" json:"RiesgoPuesto,omitempty"`
	SalarioDiarioIntegrado     *float64               `xml:"SalarioDiarioIntegrado,attr" bson:"SalarioDiarioIntegrado,omitempty" json:"SalarioDiarioIntegrado,omitempty"`
	Percepciones               *PercepcionesNomina11  `xml:"Percepciones" bson:"Percepciones,omitempty" json:"Percepciones,omitempty"`
	Deducciones                *DeduccionesNomina11   `xml:"Deducciones" bson:"Deducciones,omitempty" json:"Deducciones,omitempty"`
	Incapacidades              *[]IncapacidadNomina11 `xml:"Incapacidades>Incapacidad" bson:"Incapacidades,omitempty" json:"Incapacidades,omitempty"`
	HorasExtras                *[]HorasExtraNomina11  `xml:"HorasExtras>HorasExtra" bson:"HorasExtras,omitempty" json:"HorasExtras,omitempty"`
}

type PercepcionesNomina11 struct {
	TotalGravado float64              `xml:"TotalGravado,attr" bson:"TotalGravado" json:"TotalGravado"`
	TotalExento  float64              `xml:"TotalExento,attr" bson:"TotalExento" json:"TotalExento"`
	Percepcion   []PercepcionNomina11 `xml:"Percepcion" bson:"Percepcion" json:"Percepcion"`
}

type PercepcionNomina11 struct {
	Tipo           int     `xml:"TipoPercepcion,attr" bson:"Tipo" json:"Tipo"`
	Clave          string  `xml:"Clave,attr" bson:"Clave" json:"Clave"`
	Concepto       string  `xml:"Concepto,attr" bson:"Concepto" json:"Concepto"`
	ImporteGravado float64 `xml:"ImporteGravado,attr" bson:"ImporteGravado" json:"ImporteGravado"`
	ImporteExento  float64 `xml:"ImporteExento,attr" bson:"ImporteExento" json:"ImporteExento"`
}

type DeduccionesNomina11 struct {
	TotalGravado float64             `xml:"TotalGravado,attr" bson:"TotalGravado" json:"TotalGravado"`
	TotalExento  float64             `xml:"TotalExento,attr" bson:"TotalExento" json:"TotalExento"`
	Deduccion    []DeduccionNomina11 `xml:"Deduccion" bson:"Deduccion" json:"Deduccion"`
}

type DeduccionNomina11 struct {
	Tipo           int     `xml:"TipoDeduccion,attr" bson:"Tipo" json:"Tipo"`
	Clave          string  `xml:"Clave,attr" bson:"Clave" json:"Clave"`
	Concepto       string  `xml:"Concepto,attr" bson:"Concepto" json:"Concepto"`
	ImporteGravado float64 `xml:"ImporteGravado,attr" bson:"ImporteGravado" json:"ImporteGravado"`
	ImporteExento  float64 `xml:"ImporteExento,attr" bson:"ImporteExento" json:"ImporteExento"`
}

type IncapacidadNomina11 struct {
	Dias      float64 `xml:"DiasIncapacidad,attr" bson:"Dias" json:"Dias"`
	Tipo      int     `xml:"TipoIncapacidad,attr" bson:"Tipo" json:"Tipo"`
	Descuento float64 `xml:"Descuento,attr" bson:"Descuento" json:"Descuento"`
}

type HorasExtraNomina11 struct {
	Dias          int     `xml:"Dias,attr" bson:"Dias" json:"Dias"`
	TipoHoras     string  `xml:"TipoHoras,attr" bson:"TipoHoras" json:"TipoHoras"`
	HorasExtra    int     `xml:"HorasExtra,attr" bson:"HorasExtra" json:"HorasExtra"`
	ImportePagado float64 `xml:"ImportePagado,attr" bson:"ImportePagado" json:"ImportePagado"`
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

func (n *Nomina11) UnmarshalJSON(data []byte) error {
	// Create an alias to avoid recursion
	type Alias Nomina11
	var aux Alias

	// Unmarshal the XML into the alias
	if err := json.Unmarshal(data, &aux); err != nil {
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
