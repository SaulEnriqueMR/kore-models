package nomina

import (
	"encoding/xml"
	"testing"

	nomina11 "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/nomina"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetNomina11ForTest(filename string, t *testing.T) (nomina11.Nomina11, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed nomina11.Nomina11
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	testing2.GenerateJSONFromStructure("nomina11.json", parsed)
	return parsed, errUnmashal
}

func TestFullNomina11(t *testing.T) {
	nomina, _ := GetNomina11ForTest("./nomina11.xml", t)
	InternalTestBaseAttributes(t, nomina)
	InternalTestPercepciones(t, nomina.Percepciones)
	InternalTestDeducciones(t, nomina.Deducciones)
	InternalTestIncapacidades(t, nomina.Incapacidades)
	InternalTestHorasExtras(t, nomina.HorasExtras)
}

func InternalTestBaseAttributes(t *testing.T, nomina nomina11.Nomina11) {
	assert.Equal(t, "1.1", nomina.Version)
	assert.Equal(t, "12345", nomina.NoEmpleado)
	assert.Equal(t, "GODE561231GR8", nomina.Curp)
	assert.Equal(t, 2, nomina.TipoRegimen)
	assert.Equal(t, "2024-09-18", nomina.FechaPagoString)
	assert.Equal(t, "2023-09-01", nomina.FechaInicialPagoString)
	assert.Equal(t, "2023-09-15", nomina.FechaFinalPagoString)
	assert.Equal(t, 14.000000, nomina.NumerosDiasPagados)
	assert.Equal(t, "Quincenal", nomina.PeriodicidadPago)
	assert.Equal(t, "12345678901234567890", *nomina.RegistroPatronal)
	assert.Equal(t, "123456789012345", *nomina.NumSeguridadSocial)
	assert.Equal(t, "Contabilidad", *nomina.Departamento)
	assert.Equal(t, "012345678901234567", *nomina.Clabe)
	assert.Equal(t, 12, *nomina.Banco)
	assert.Equal(t, "2010-05-15", *nomina.FechaInicioRelLaboral)
	assert.Equal(t, 750, *nomina.Antiguedad)
	assert.Equal(t, "Analista Financiero", *nomina.Puesto)
	assert.Equal(t, "Sindicalizado", *nomina.TipoContrato)
	assert.Equal(t, "Diurna", *nomina.TipoJornada)
	assert.Equal(t, 10000.00, *nomina.SalarioBaseCotizacion)
	assert.Equal(t, 1, *nomina.RiesgoPuesto)
	assert.Equal(t, 500.0, *nomina.SalarioDiarioIntegrado)
}

func InternalTestPercepciones(t *testing.T, percepciones *nomina11.PercepcionesNomina11) {
	assert.NotNil(t, percepciones)
	assert.Equal(t, 1000.00, percepciones.TotalExento)
	assert.Equal(t, 15000.00, percepciones.TotalGravado)
	InternalTestPercepcion(t, percepciones.Percepcion)
}

func InternalTestPercepcion(t *testing.T, percepciones []nomina11.PercepcionNomina11) {
	assert.NotNil(t, percepciones)
	assert.Equal(t, len(percepciones), 2)
	percepcion := (percepciones)[0]
	assert.Equal(t, "001", percepcion.Clave)
	assert.Equal(t, "Sueldo Base", percepcion.Concepto)
	assert.Equal(t, 1000.00, percepcion.ImporteExento)
	assert.Equal(t, 15000.00, percepcion.ImporteGravado)
	assert.Equal(t, 001, percepcion.Tipo)

}

func InternalTestDeducciones(t *testing.T, deducciones *nomina11.DeduccionesNomina11) {
	assert.NotNil(t, deducciones)
	assert.Equal(t, 0.00, deducciones.TotalExento)
	assert.Equal(t, 500.00, deducciones.TotalGravado)
	InternalTestDeduccion(t, &deducciones.Deduccion)
}

func InternalTestDeduccion(t *testing.T, deduccion *[]nomina11.DeduccionNomina11) {
	assert.NotNil(t, deduccion)
	assert.Equal(t, len(*deduccion), 2)
	deduccion1 := (*deduccion)[0]
	assert.Equal(t, "001", deduccion1.Clave)
	assert.Equal(t, 001, deduccion1.Tipo)
	assert.Equal(t, 0.00, deduccion1.ImporteExento)
	assert.Equal(t, 500.00, deduccion1.ImporteGravado)
	assert.Equal(t, "ISR", deduccion1.Concepto)
}

func InternalTestIncapacidades(t *testing.T, incapacidades *[]nomina11.IncapacidadNomina11) {
	assert.NotNil(t, incapacidades)
	assert.Equal(t, len(*incapacidades), 2)
	incapacidad1 := (*incapacidades)[0]
	InternalTestIncapacidad(t, &incapacidad1)
}

func InternalTestIncapacidad(t *testing.T, incapacidad *nomina11.IncapacidadNomina11) {
	assert.NotNil(t, incapacidad)
	assert.Equal(t, 2.00, incapacidad.Dias)
	assert.Equal(t, 001, incapacidad.Tipo)
	assert.Equal(t, 200.0, incapacidad.Descuento)
}

func InternalTestHorasExtras(t *testing.T, extraNomina11 *[]nomina11.HorasExtraNomina11) {
	assert.NotNil(t, extraNomina11)
	assert.Equal(t, len(*extraNomina11), 2)
	extraNomina1 := (*extraNomina11)[0]
	InternalTestHorasExtra(t, &extraNomina1)
}

func InternalTestHorasExtra(t *testing.T, horaExtra *nomina11.HorasExtraNomina11) {
	assert.NotNil(t, horaExtra)
	assert.Equal(t, 1, horaExtra.Dias)
	assert.Equal(t, "Dobles", horaExtra.TipoHoras)
	assert.Equal(t, 2, horaExtra.HorasExtra)
	assert.Equal(t, 100.0, horaExtra.ImportePagado)
}
