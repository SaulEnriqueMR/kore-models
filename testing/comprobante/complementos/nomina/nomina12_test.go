package nomina

import (
	"encoding/xml"
	"testing"

	nomina12 "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/nomina"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetNomina12ForTest(filename string, t *testing.T) (nomina12.Nomina12, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed nomina12.Nomina12
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	testing2.GenerateJSONFromStructure("nomina12.json", parsed)
	return parsed, errUnmashal
}

func TestFullNomina12(t *testing.T) {
	nomina, _ := GetNomina12ForTest("./nomina12.xml", t)
	InternalTestBaseAttributes12(t, nomina)
	InternalTestEmisor(t, nomina.Emisor)
	InternalTestReceptor(t, nomina.Receptor)
	InternalTestPercepciones12(t, *nomina.Percepciones)
	InternalTestDeducciones12(t, *nomina.Deducciones)
	InternalTestOtroPago12(t, *nomina.OtrosPagos)
	InternalTestIncapacidad12(t, *nomina.Incapacidades)

}

func InternalTestBaseAttributes12(t *testing.T, nomina nomina12.Nomina12) {
	assert.Equal(t, "1.2", nomina.Version)
	assert.Equal(t, "O", nomina.TipoNomina)
	assert.Equal(t, "2024-09-15", nomina.FechaPagoString)
	assert.Equal(t, "2024-09-01", nomina.FechaInicialPagoString)
	assert.Equal(t, "2024-09-14", nomina.FechaFinalPagoString)
	assert.Equal(t, 14.000, nomina.NumeroDiasPagados)
	assert.Equal(t, 15000.00, *nomina.TotalPercepciones)
	assert.Equal(t, 5000.00, *nomina.TotalDeducciones)
	assert.Equal(t, 2000.00, *nomina.TotalOtrosPagos)
}

func InternalTestEmisor(t *testing.T, emisor *nomina12.EmisorNomina12) {
	assert.NotNil(t, emisor)
	assert.Equal(t, "ABC123456HDFGJ90", *emisor.Curp)
	assert.Equal(t, "12345678901234567890", *emisor.RegistroPatronal)
	assert.Equal(t, "RFC123456789", *emisor.RfcPatronOrigen)
	InternalTestEntidadSNCF(t, emisor.EntidadSncf)
}

func InternalTestEntidadSNCF(t *testing.T, entidad *nomina12.EntidadSNCFNomina12) {
	assert.NotNil(t, entidad)
	assert.Equal(t, "Federal", entidad.OrigenRecurso)
	assert.Equal(t, 30000.00, *entidad.MontoRecursoPropio)

}

func InternalTestReceptor(t *testing.T, receptor nomina12.ReceptorNomina12) {
	assert.NotNil(t, receptor)
	assert.Equal(t, "DEF123456HDFGJ01", receptor.Curp)
	assert.Equal(t, "123456789012345", *receptor.NoSeguridadSocial)
	assert.Equal(t, "2022-01-01", *receptor.FechaInicioRelLaboral)
	assert.Equal(t, "P1Y2M10D", *receptor.Antiguedad)
	assert.Equal(t, "01", receptor.TipoContrato)
	assert.Equal(t, "Sí", *receptor.Sindicalizado)
	assert.Equal(t, "Diurna", *receptor.TipoJornada)
	assert.Equal(t, "01", receptor.TipoRegimen)
	assert.Equal(t, "12345", receptor.NoEmpleado)
	assert.Equal(t, "Ventas", *receptor.Departamento)
	assert.Equal(t, "Ejecutivo de Ventas", *receptor.Puesto)
	assert.Equal(t, "B", *receptor.RiesgoPuesto)
	assert.Equal(t, "Quincenal", receptor.PeriodicidadPago)
	assert.Equal(t, "021", *receptor.Banco)
	assert.Equal(t, "12345678901", *receptor.CuentaBancaria)
	assert.Equal(t, 12000.00, *receptor.SalarioBaseCotizacion)
	assert.Equal(t, 800.00, *receptor.SalarioDiarioIntegrado)
	assert.Equal(t, "09", receptor.ClaveEntidadFederativa)
	InternalTestSubContratacion(t, receptor.Subcontratacion)
}

func InternalTestSubContratacion(t *testing.T, subcontrataciones *[]nomina12.SubContratacionNomina12) {
	assert.NotNil(t, subcontrataciones)
	subcontratacion := (*subcontrataciones)[0]
	assert.NotNil(t, subcontratacion)
	assert.Equal(t, "XAXX010101000", subcontratacion.RfcLabora)
	assert.Equal(t, 50.000, subcontratacion.PorcentajeTiempo)
}

func InternalTestPercepciones12(t *testing.T, percepciones nomina12.PercepcionesNomina12) {
	assert.NotNil(t, percepciones)
	assert.Equal(t, 15000.00, *percepciones.TotalSueldos)
	assert.Equal(t, 3000.00, percepciones.TotalExento)
	assert.Equal(t, 5000.00, *percepciones.TotalSeparacionIndemnizacion)
	assert.Equal(t, 2000.00, *percepciones.TotalJubilacionPensionRetiro)
	assert.Equal(t, 22000.00, percepciones.TotalGravado)
	InternalTestPercepcion12(t, percepciones.Percepcion)
	InternalTestJubilacionPensionRetiro12(t, *percepciones.JubilacionPensionRetiro)
	InternalTestSeparacionIndemnizacionNomina12(t, *percepciones.SeparacionIndemnizacion)
}

func InternalTestPercepcion12(t *testing.T, percepciones []nomina12.PercepcionNomina12) {
	assert.NotNil(t, percepciones)
	percepcion := percepciones[0]
	assert.NotNil(t, percepcion)
	assert.Equal(t, "001", percepcion.TipoPercepcion)
	assert.Equal(t, "12345", percepcion.Clave)
	assert.Equal(t, "Sueldo Mensual", percepcion.Concepto)
	assert.Equal(t, 15000.00, percepcion.ImporteGravado)
	assert.Equal(t, 0.00, percepcion.ImporteExento)
	InternalTestAccionesOTitulos12(t, percepcion.AccionesOTitulos)
	InternalTestHorasExtra12(t, *percepcion.HorasExtra)

}

func InternalTestAccionesOTitulos12(t *testing.T, acciones *nomina12.AccionesOTitulosNomina12) {
	assert.NotNil(t, acciones)
	assert.Equal(t, 100.000000, acciones.ValorMercado)
	assert.Equal(t, 50.000000, acciones.PrecioAlOtorgarse)
}

func InternalTestHorasExtra12(t *testing.T, horasExtras []nomina12.HorasExtraNomina12) {
	assert.NotNil(t, horasExtras)
	horasExtra := horasExtras[0]
	assert.NotNil(t, horasExtra)
	assert.Equal(t, 5.0, horasExtra.Dias)
	assert.Equal(t, "1", horasExtra.TipoHoras)
	assert.Equal(t, 10.0, horasExtra.HorasExtra)
	assert.Equal(t, 2000.00, horasExtra.ImportePagado)
}

func InternalTestJubilacionPensionRetiro12(t *testing.T, jubilacion nomina12.JubilacionPensionRetiroNomina12) {
	assert.NotNil(t, jubilacion)
	assert.Equal(t, 1000.00, *jubilacion.TotalUnaExhibicion)
	assert.Equal(t, 500.00, *jubilacion.TotalParcialidad)
	assert.Equal(t, 50.00, *jubilacion.MontoDiario)
	assert.Equal(t, 0.00, jubilacion.IngresoAcumulable)
	assert.Equal(t, 2000.00, jubilacion.IngresoNoAcumulable)
}

func InternalTestSeparacionIndemnizacionNomina12(t *testing.T, separacion nomina12.SeparacionIndemnizacionNomina12) {
	assert.NotNil(t, separacion)
	assert.Equal(t, 5000.00, separacion.TotalPagado)
	assert.Equal(t, 10.0, separacion.NumeroAniosServicio)
	assert.Equal(t, 1500.00, separacion.UltimoSueldoMensualOrdinario)
	assert.Equal(t, 4000.00, separacion.IngresoAcumulable)
	assert.Equal(t, 1000.00, separacion.IngresoNoAcumulable)
}

func InternalTestDeducciones12(t *testing.T, deducciones nomina12.DeduccionesNomina12) {
	assert.NotNil(t, deducciones)
	assert.Equal(t, 3000.00, *deducciones.TotalOtrasDeducciones)
	assert.Equal(t, 2000.00, *deducciones.TotalImpuestosRetenidos)
	InternalTestDeduccion12(t, deducciones.Deduccion)
}

func InternalTestDeduccion12(t *testing.T, deducciones []nomina12.DeduccionNomina12) {
	assert.NotNil(t, deducciones)
	deduccion := deducciones[0]
	assert.NotNil(t, deduccion)
	assert.Equal(t, "003", deduccion.Tipo)
	assert.Equal(t, "D001", deduccion.Clave)
	assert.Equal(t, "Préstamo Personal", deduccion.Concepto)
	assert.Equal(t, 1500.00, deduccion.Importe)
}

func InternalTestOtroPago12(t *testing.T, otrosPagos []nomina12.OtroPagoNomina12) {
	assert.NotNil(t, otrosPagos)
	otroPago := otrosPagos[0]
	assert.NotNil(t, otroPago)
	assert.Equal(t, "001", otroPago.Tipo)
	assert.Equal(t, "OP001", otroPago.Clave)
	assert.Equal(t, "Subsidio de Empleo", otroPago.Concepto)
	assert.Equal(t, 500.00, otroPago.Importe)
	InternalTestSubsidioAlEmpleoNomina12(t, otroPago.SubsidioAlEmpleo)
	otroPago1 := otrosPagos[1]
	assert.NotNil(t, otroPago1)
	InternalTestCompensacionSaldosAFavorNomina12(t, otroPago1.CompensacionSaldosAFavor)
}

func InternalTestSubsidioAlEmpleoNomina12(t *testing.T, subsidio *nomina12.SubsidioAlEmpleoNomina12) {
	assert.NotNil(t, subsidio)
	assert.Equal(t, 500.00, subsidio.SubsidioCausado)
}
func InternalTestCompensacionSaldosAFavorNomina12(t *testing.T, compensacion *nomina12.CompensacionSaldosAFavorNomina12) {
	assert.NotNil(t, compensacion)
	assert.Equal(t, 300.00, compensacion.SaldoAFavor)
	assert.Equal(t, "2023", compensacion.Anio)
	assert.Equal(t, 100.00, compensacion.RemanenteSaldoAFavor)
}
func InternalTestIncapacidad12(t *testing.T, incapacidades []nomina12.IncapacidadNomina12) {
	assert.NotNil(t, incapacidades)
	incapacidad := incapacidades[0]
	assert.NotNil(t, incapacidad)
	assert.Equal(t, 5.0, incapacidad.Dias)
	assert.Equal(t, 01.0, incapacidad.Tipo)
	assert.Equal(t, 1500.00, *incapacidad.Importe)
}
