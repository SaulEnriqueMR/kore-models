package nomina

import (
	"encoding/xml"
	"time"

	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

// Nomina12 Presente en Comprobante 3.3 y 4.0
type Nomina12 struct {
	Version                string                 `xml:"Version,attr" bson:"Version" json:"Version"`
	TipoNomina             string                 `xml:"TipoNomina,attr" bson:"TipoNomina" json:"TipoNomina"`
	FechaPagoString        string                 `xml:"FechaPago,attr" bson:"FechaPagoString" json:"FechaPagoString"`
	FechaPago              time.Time              `bson:"FechaPago" json:"FechaPago"`
	FechaInicialPagoString string                 `xml:"FechaInicialPago,attr" bson:"FechaInicialPagoString" json:"FechaInicialPagoString"`
	FechaInicialPago       time.Time              `bson:"FechaInicialPago" json:"FechaInicialPago"`
	FechaFinalPagoString   string                 `xml:"FechaFinalPago,attr" bson:"FechaFinalPagoString" json:"FechaFinalPagoString"`
	FechaFinalPago         time.Time              `bson:"FechaFinalPago" json:"FechaFinalPago"`
	NumeroDiasPagados      float64                `xml:"NumDiasPagados,attr" bson:"NumeroDiasPagados" json:"NumeroDiasPagados"`
	TotalPercepciones      *float64               `xml:"TotalPercepciones,attr" bson:"TotalPercepciones,omitempty" json:"TotalPercepciones,omitempty"`
	TotalDeducciones       *float64               `xml:"TotalDeducciones,attr" bson:"TotalDeducciones,omitempty" json:"TotalDeducciones,omitempty"`
	TotalOtrosPagos        *float64               `xml:"TotalOtrosPagos,attr" bson:"TotalOtrosPagos,omitempty" json:"TotalOtrosPagos,omitempty"`
	Emisor                 *EmisorNomina12        `xml:"Emisor" bson:"Emisor,omitempty" json:"Emisor,omitempty"`
	Receptor               ReceptorNomina12       `xml:"Receptor" bson:"Receptor" json:"Receptor"`
	Percepciones           *PercepcionesNomina12  `xml:"Percepciones" bson:"Percepciones,omitempty" json:"Percepciones,omitempty"`
	Deducciones            *DeduccionesNomina12   `xml:"Deducciones" bson:"Deducciones,omitempty" json:"Deducciones,omitempty"`
	OtrosPagos             *[]OtroPagoNomina12    `xml:"OtrosPagos>OtroPago" bson:"OtrosPagos,omitempty" json:"OtrosPagos,omitempty"`
	Incapacidades          *[]IncapacidadNomina12 `xml:"Incapacidades>Incapacidad" bson:"Incapacidades,omitempty" json:"Incapacidades,omitempty"`
}

type EmisorNomina12 struct {
	Curp             *string              `xml:"Curp,attr" bson:"Curp,omitempty" json:"Curp,omitempty"`                                     // Cifrado
	RegistroPatronal *string              `xml:"RegistroPatronal,attr" bson:"RegistroPatronal,omitempty" json:"RegistroPatronal,omitempty"` // Cifrado
	RfcPatronOrigen  *string              `xml:"RfcPatronOrigen,attr" bson:"RfcPatronOrigen,omitempty" json:"RfcPatronOrigen,omitempty"`    // Cifrado
	EntidadSncf      *EntidadSNCFNomina12 `xml:"EntidadSNCF" bson:"EntidadSncf,omitempty" json:"EntidadSncf,omitempty"`
}

type EntidadSNCFNomina12 struct {
	OrigenRecurso      string   `xml:"OrigenRecurso,attr" bson:"OrigenRecurso" json:"OrigenRecurso"`
	MontoRecursoPropio *float64 `xml:"MontoRecursoPropio,attr" bson:"MontoRecursoPropio,omitempty" json:"MontoRecursoPropio,omitempty"`
}

type ReceptorNomina12 struct {
	Curp                       string                     `xml:"Curp,attr" bson:"Curp" json:"Curp"`                                                             // Cifrado
	NoSeguridadSocial          *string                    `xml:"NumSeguridadSocial,attr" bson:"NoSeguridadSocial,omitempty" json:"NoSeguridadSocial,omitempty"` // Cifrado
	FechaInicioRelLaboral      *string                    `xml:"FechaInicioRelLaboral,attr" bson:"FechaInicioRelLaboral,omitempty" json:"FechaInicioRelLaboral,omitempty"`
	FechaInicioRelacionLaboral *time.Time                 `bson:"FechaInicioRelacionLaboral,omitempty" json:"FechaInicioRelacionLaboral,omitempty"`
	Antiguedad                 *string                    `xml:"Antigüedad,attr" bson:"Antiguedad,omitempty" json:"Antiguedad,omitempty"`
	TipoContrato               string                     `xml:"TipoContrato,attr" bson:"TipoContrato" json:"TipoContrato"`
	Sindicalizado              *string                    `xml:"Sindicalizado,attr" bson:"Sindicalizado,omitempty" json:"Sindicalizado,omitempty"`
	TipoJornada                *string                    `xml:"TipoJornada,attr" bson:"TipoJornada,omitempty" json:"TipoJornada,omitempty"`
	TipoRegimen                string                     `xml:"TipoRegimen,attr" bson:"TipoRegimen" json:"TipoRegimen"`
	NoEmpleado                 string                     `xml:"NumEmpleado,attr" bson:"NoEmpleado" json:"NoEmpleado"`
	Departamento               *string                    `xml:"Departamento,attr" bson:"Departamento,omitempty" json:"Departamento,omitempty"`
	Puesto                     *string                    `xml:"Puesto,attr" bson:"Puesto,omitempty" json:"Puesto,omitempty"`
	RiesgoPuesto               *string                    `xml:"RiesgoPuesto,attr" bson:"RiesgoPuesto,omitempty" json:"RiesgoPuesto,omitempty"`
	PeriodicidadPago           string                     `xml:"PeriodicidadPago,attr" bson:"PeriodicidadPago" json:"PeriodicidadPago"`
	Banco                      *string                    `xml:"Banco,attr" bson:"Banco,omitempty" json:"Banco,omitempty"`
	CuentaBancaria             *string                    `xml:"CuentaBancaria,attr" bson:"CuentaBancaria,omitempty" json:"CuentaBancaria,omitempty"`
	SalarioBaseCotizacion      *float64                   `xml:"SalarioBaseCotApor,attr" bson:"SalarioBaseCotizacion,omitempty" json:"SalarioBaseCotizacion,omitempty"`
	SalarioDiarioIntegrado     *float64                   `xml:"SalarioDiarioIntegrado,attr" bson:"SalarioDiarioIntegrado,omitempty" json:"SalarioDiarioIntegrado,omitempty"`
	ClaveEntidadFederativa     string                     `xml:"ClaveEntFed,attr" bson:"ClaveEntidadFederativa" json:"ClaveEntidadFederativa"`
	Subcontratacion            *[]SubContratacionNomina12 `xml:"SubContratacion" bson:"Subcontratacion,omitempty" json:"Subcontratacion,omitempty"`
}

func (r *ReceptorNomina12) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias ReceptorNomina12
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*r = ReceptorNomina12(aux)

	if r.FechaInicioRelLaboral != nil {
		fechaIniRelLab, _ := helpers.ParseDatetime(*aux.FechaInicioRelLaboral)
		r.FechaInicioRelacionLaboral = &fechaIniRelLab
	}

	return nil
}

type SubContratacionNomina12 struct {
	RfcLabora        string  `xml:"RfcLabora,attr" bson:"RfcLabora" json:"RfcLabora"` // Cifrado
	PorcentajeTiempo float64 `xml:"PorcentajeTiempo,attr" bson:"PorcentajeTiempo" json:"PorcentajeTiempo"`
}

type PercepcionesNomina12 struct {
	TotalSueldos                 *float64                         `xml:"TotalSueldos,attr" bson:"TotalSueldos,omitempty" json:"TotalSueldos,omitempty"`
	TotalSeparacionIndemnizacion *float64                         `xml:"TotalSeparacionIndemnizacion,attr" bson:"TotalSeparacionIndemnizacion,omitempty" json:"TotalSeparacionIndemnizacion,omitempty"`
	TotalJubilacionPensionRetiro *float64                         `xml:"TotalJubilacionPensionRetiro,attr" bson:"TotalJubilacionPensionRetiro,omitempty" json:"TotalJubilacionPensionRetiro,omitempty"`
	TotalGravado                 float64                          `xml:"TotalGravado,attr" bson:"TotalGravado" json:"TotalGravado"`
	TotalExento                  float64                          `xml:"TotalExento,attr" bson:"TotalExento" json:"TotalExento"`
	Percepcion                   []PercepcionNomina12             `xml:"Percepcion" bson:"Percepcion" json:"Percepcion"`
	JubilacionPensionRetiro      *JubilacionPensionRetiroNomina12 `xml:"JubilacionPensionRetiro" bson:"JubilacionPensionRetiro,omitempty" json:"JubilacionPensionRetiro,omitempty"`
	SeparacionIndemnizacion      *SeparacionIndemnizacionNomina12 `xml:"SeparacionIndemnizacion" bson:"SeparacionIndemnizacion,omitempty" json:"SeparacionIndemnizacion,omitempty"`
}

type PercepcionNomina12 struct {
	Tipo             string                    `xml:"TipoPercepcion,attr" bson:"Tipo" json:"Tipo"`
	Clave            string                    `xml:"Clave,attr" bson:"Clave" json:"Clave"`
	Concepto         string                    `xml:"Concepto,attr" bson:"Concepto" json:"Concepto"`
	ImporteGravado   float64                   `xml:"ImporteGravado,attr" bson:"ImporteGravado" json:"ImporteGravado"`
	ImporteExento    float64                   `xml:"ImporteExento,attr" bson:"ImporteExento" json:"ImporteExento"`
	AccionesOTitulos *AccionesOTitulosNomina12 `xml:"AccionesOTitulos" bson:"AccionesOTitulos,omitempty" json:"AccionesOTitulos,omitempty"`
	HorasExtra       *[]HorasExtraNomina12     `xml:"HorasExtra" bson:"HorasExtra,omitempty" json:"HorasExtra,omitempty"`
}

type AccionesOTitulosNomina12 struct {
	ValorMercado      float64 `xml:"ValorMercado,attr" bson:"ValorMercado" json:"ValorMercado"`
	PrecioAlOtorgarse float64 `xml:"PrecioAlOtorgarse,attr" bson:"PrecioAlOtorgarse" json:"PrecioAlOtorgarse"`
}

type HorasExtraNomina12 struct {
	Dias          int     `xml:"Dias,attr" bson:"Dias" json:"Dias"`
	TipoHoras     string  `xml:"TipoHoras,attr" bson:"TipoHoras" json:"TipoHoras"`
	HorasExtra    int     `xml:"HorasExtra,attr" bson:"HorasExtra" json:"HorasExtra"`
	ImportePagado float64 `xml:"ImportePagado,attr" bson:"ImportePagado" json:"ImportePagado"`
}

type JubilacionPensionRetiroNomina12 struct {
	TotalUnaExhibicion  *float64 `xml:"TotalUnaExhibicion,attr" bson:"TotalUnaExhibicion,omitempty" json:"TotalUnaExhibicion,omitempty"`
	TotalParcialidad    *float64 `xml:"TotalParcialidad,attr" bson:"TotalParcialidad,omitempty" json:"TotalParcialidad,omitempty"`
	MontoDiario         *float64 `xml:"MontoDiario,attr" bson:"MontoDiario,omitempty" json:"MontoDiario,omitempty"`
	IngresoAcumulable   float64  `xml:"IngresoAcumulable,attr" bson:"IngresoAcumulable" json:"IngresoAcumulable"`
	IngresoNoAcumulable float64  `xml:"IngresoNoAcumulable,attr" bson:"IngresoNoAcumulable" json:"IngresoNoAcumulable"`
}

type SeparacionIndemnizacionNomina12 struct {
	TotalPagado                  float64 `xml:"TotalPagado,attr" bson:"TotalPagado" json:"TotalPagado"`
	NumeroAniosServicio          int     `xml:"NumAñosServicio,attr" bson:"NumeroAniosServicio" json:"NumeroAniosServicio"`
	UltimoSueldoMensualOrdinario float64 `xml:"UltimoSueldoMensOrd,attr" bson:"UltimoSueldoMensualOrdinario" json:"UltimoSueldoMensualOrdinario"`
	IngresoAcumulable            float64 `xml:"IngresoAcumulable,attr" bson:"IngresoAcumulable" json:"IngresoAcumulable"`
	IngresoNoAcumulable          float64 `xml:"IngresoNoAcumulable,attr" bson:"IngresoNoAcumulable" json:"IngresoNoAcumulable"`
}

type DeduccionesNomina12 struct {
	TotalOtrasDeducciones   *float64            `xml:"TotalOtrasDeducciones,attr" bson:"TotalOtrasDeducciones,omitempty" json:"TotalOtrasDeducciones,omitempty"`
	TotalImpuestosRetenidos *float64            `xml:"TotalImpuestosRetenidos,attr" bson:"TotalImpuestosRetenidos,omitempty" json:"TotalImpuestosRetenidos,omitempty"`
	Deduccion               []DeduccionNomina12 `xml:"Deduccion" bson:"Deduccion" json:"Deduccion"`
}

type DeduccionNomina12 struct {
	Tipo     string  `xml:"TipoDeduccion,attr" bson:"Tipo" json:"Tipo"`
	Clave    string  `xml:"Clave,attr" bson:"Clave" json:"Clave"`
	Concepto string  `xml:"Concepto,attr" bson:"Concepto" json:"Concepto"`
	Importe  float64 `xml:"Importe,attr" bson:"Importe" json:"Importe"`
}

type OtroPagoNomina12 struct {
	Tipo                     string                            `xml:"TipoOtroPago,attr" bson:"Tipo" json:"Tipo"`
	Clave                    string                            `xml:"Clave,attr" bson:"Clave" json:"Clave"`
	Concepto                 string                            `xml:"Concepto,attr" bson:"Concepto" json:"Concepto"`
	Importe                  float64                           `xml:"Importe,attr" bson:"Importe" json:"Importe"`
	SubsidioAlEmpleo         *SubsidioAlEmpleoNomina12         `xml:"SubsidioAlEmpleo" bson:"SubsidioAlEmpleo,omitempty" json:"SubsidioAlEmpleo,omitempty"`
	CompensacionSaldosAFavor *CompensacionSaldosAFavorNomina12 `xml:"CompensacionSaldosAFavor" bson:"CompensacionSaldosAFavor,omitempty" json:"CompensacionSaldosAFavor,omitempty"`
}

type SubsidioAlEmpleoNomina12 struct {
	SubsidioCausado float64 `xml:"SubsidioCausado,attr" bson:"SubsidioCausado" json:"SubsidioCausado"`
}

type CompensacionSaldosAFavorNomina12 struct {
	SaldoAFavor          float64 `xml:"SaldoAFavor,attr" bson:"SaldoAFavor" json:"SaldoAFavor"`
	Anio                 string  `xml:"Año,attr" bson:"Anio" json:"Anio"`
	RemanenteSaldoAFavor float64 `xml:"RemanenteSalFav,attr" bson:"RemanenteSaldoAFavor" json:"RemanenteSaldoAFavor"`
}

type IncapacidadNomina12 struct {
	Dias    int      `xml:"DiasIncapacidad,attr" bson:"Dias" json:"Dias"`
	Tipo    string   `xml:"TipoIncapacidad,attr" bson:"Tipo" json:"Tipo"`
	Importe *float64 `xml:"ImporteMonetario,attr" bson:"Importe,omitempty" json:"Importe,omitempty"`
}

func (c *Nomina12) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias Nomina12
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*c = Nomina12(aux)
	fecha1, err := helpers.ParseDatetime(aux.FechaPagoString)
	if err != nil {
		return err
	}

	c.FechaPago = fecha1

	fecha2, err := helpers.ParseDatetime(aux.FechaInicialPagoString)
	if err != nil {
		return err
	}

	c.FechaInicialPago = fecha2

	fecha3, err := helpers.ParseDatetime(aux.FechaFinalPagoString)
	if err != nil {
		return err
	}

	c.FechaFinalPago = fecha3

	return nil
}
