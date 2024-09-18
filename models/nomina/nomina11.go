package nomina

type Nomina11 struct {
	Version                string                  `xml:"Version,attr" bson:"Version"`
	RegistroPatronal       *string                 `xml:"RegistroPatronal,attr" bson:"RegistroPatronal,omitempty"` // Cifrado
	NumEmpleado            string                  `xml:"NumEmpleado,attr" bson:"NumEmpleado"`
	Curp                   string                  `xml:"CURP,attr" bson:"Curp"` // Cifrado
	TipoRegimen            int                     `xml:"TipoRegimen,attr" bson:"TipoRegimen"`
	NumSeguridadSocial     *string                 `xml:"NumSeguridadSocial,attr" bson:"NumSeguridadSocial,omitempty"` // Cifrado
	FechaPago              string                  `xml:"FechaPago,attr" bson:"FechaPago"`
	FechaInicialPago       string                  `xml:"FechaInicialPago,attr" bson:"FechaInicialPago"`
	FechaFinalPago         string                  `xml:"FechaFinalPago,attr" bson:"FechaFinalPago"`
	NumDiasPagados         float64                 `xml:"NumDiasPagados,attr" bson:"NumDiasPagados"`
	Departamento           *string                 `xml:"Departamento,attr" bson:"Departamento,omitempty"`
	Clabe                  *string                 `xml:"CLABE,attr" bson:"Clabe,omitempty"` // Cifrado
	Banco                  *int                    `xml:"Banco,attr" bson:"Banco,omitempty"` // Cifrado
	FechaInicioRelLaboral  *string                 `xml:"FechaInicioRelLaboral,attr" bson:"FechaInicioRelLaboral,omitempty"`
	Antiguedad             *int                    `xml:"Antiguedad,attr" bson:"Antiguedad,omitempty"`
	Puesto                 *string                 `xml:"Puesto,attr" bson:"Puesto,omitempty"`
	TipoContrato           *string                 `xml:"TipoContrato,attr" bson:"TipoContrato,omitempty"`
	TipoJornada            *string                 `xml:"TipoJornada,attr" bson:"TipoJornada,omitempty"`
	PeriodicidadPago       string                  `xml:"PeriodicidadPago,attr" bson:"PeriodicidadPago"`
	SalarioBaseCotApor     *float64                `xml:"SalarioBaseCotApor,attr" bson:"SalarioBaseCotApor,omitempty"`
	RiesgoPuesto           *int                    `xml:"RiesgoPuesto,attr" bson:"RiesgoPuesto,omitempty"`
	SalarioDiarioIntegrado *float64                `xml:"SalarioDiarioIntegrado,attr" bson:"SalarioDiarioIntegrado,omitempty"`
	Percepciones           *[]PercepcionesNomina11 `xml:"Percepciones" bson:"Percepciones,omitempty"`
	Deducciones            *[]DeduccionesNomina11  `xml:"Deducciones" bson:"Deducciones,omitempty"`
	Incapacidades          *[]IncapacidadNomina11  `xml:"Incapacidades>Incapacidad" bson:"Incapacidades,omitempty"`
	HorasExtras            *[]HorasExtraNomina11   `xml:"HorasExtras>HorasExtra" bson:"HorasExtras,omitempty"`
}

type PercepcionesNomina11 struct {
	TotalGravado float64              `xml:"TotalGravado,attr" bson:"TotalGravado"`
	TotalExento  float64              `xml:"TotalExento,attr" bson:"TotalExento"`
	Percepcion   []PercepcionNomina11 `xml:"Percepcion" bson:"Percepcion"`
}

type PercepcionNomina11 struct {
	TipoPercepcion int     `xml:"TipoPercepcion,attr" bson:"TipoPercepcion"`
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
	TipoDeduccion  int     `xml:"TipoDeduccion,attr" bson:"TipoDeduccion"`
	Clave          string  `xml:"Clave,attr" bson:"Clave"`
	Concepto       string  `xml:"Concepto,attr" bson:"Concepto"`
	ImporteGravado float64 `xml:"ImporteGravado,attr" bson:"ImporteGravado"`
	ImporteExento  float64 `xml:"ImporteExento,attr" bson:"ImporteExento"`
}

type IncapacidadNomina11 struct {
	DiasIncapacidad float64 `xml:"DiasIncapacidad,attr" bson:"DiasIncapacidad"`
	TipoIncapacidad int     `xml:"TipoIncapacidad,attr" bson:"TipoIncapacidad"`
	Descuento       float64 `xml:"Descuento,attr" bson:"Descuento"`
}

type HorasExtraNomina11 struct {
	Dias          int     `xml:"Dias,attr" bson:"Dias"`
	TipoHoras     string  `xml:"TipoHoras,attr" bson:"TipoHoras"`
	HorasExtra    int     `xml:"HorasExtra,attr" bson:"HorasExtra"`
	ImportePagado float64 `xml:"ImportePagado,attr" bson:"ImportePagado"`
}
