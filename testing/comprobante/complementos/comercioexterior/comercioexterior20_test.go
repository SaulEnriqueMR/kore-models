package comercioexterior

import (
	"encoding/xml"
	"testing"

	comexterior1 "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/comercioexterior"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetTurisComercioExteriorForTest(filename string, t *testing.T) (comexterior1.ComercioExterior20, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed comexterior1.ComercioExterior20
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	return parsed, errUnmashal
}

func TestFullIntereses(t *testing.T) {
	comexterior, _ := GetTurisComercioExteriorForTest("./comercioexterior20.xml", t)
	assert.NotNil(t, comexterior)
	InternalTestComercioExterior(t, comexterior)
}

func InternalTestComercioExterior(t *testing.T, comercioExterior comexterior1.ComercioExterior20) {
	assert.Equal(t, "2.0", comercioExterior.Version)
	assert.Equal(t, "01", *comercioExterior.MotivoTraslado)
	assert.Equal(t, "A1", comercioExterior.ClavePedimento)
	assert.Equal(t, 1.0, comercioExterior.CertificadoOrigen)
	assert.Equal(t, "1234567890", *comercioExterior.NumCertificadoOrigen)
	assert.Equal(t, "EXCONFIABLE123", *comercioExterior.NumeroExportadorConfiable)
	assert.Equal(t, "FOB", *comercioExterior.Incoterm)
	assert.Equal(t, "Exportación de mercancías desde México.", *comercioExterior.Observaciones)
	assert.Equal(t, 18.50, comercioExterior.TipoCambioUSD)
	assert.Equal(t, 10000.00, comercioExterior.TotalUSD)
	InternalTestEmisorComercioExterior(t, comercioExterior.Emisor)
	InternalTestPropietarioComercioExterior(t, comercioExterior.Propietario)
	InternalTestReceptorComercioExterior(t, comercioExterior.Receptor)
	InternalTestDestinatarioComercioExterior(t, comercioExterior.Destinatario)
	InternalTestMercanciaComercioExterior(t, comercioExterior.Mercancias)
}

func InternalTestEmisorComercioExterior(t *testing.T, emisor *comexterior1.EmisorComercioExterior20) {
	assert.Equal(t, "GOCM890107HDFRRR02", *emisor.Curp)
	InternalTestDomicilioComercioExterior(t, emisor.Domicilio)
}

func InternalTestDomicilioComercioExterior(t *testing.T, domicilio comexterior1.DomicilioComercioExterior20) {
	assert.Equal(t, "Avenida Revolución", domicilio.Calle)
	assert.Equal(t, "123", *domicilio.NumeroExterior)
	assert.Equal(t, "4B", *domicilio.NumeroInterior)
	assert.Equal(t, "0348", *domicilio.Colonia)
	assert.Equal(t, "001", *domicilio.Localidad)
	assert.Equal(t, "Cerca de la torre", *domicilio.Referencia)
	assert.Equal(t, "012", *domicilio.Municipio)
	assert.Equal(t, "CDMX", domicilio.Estado)
	assert.Equal(t, "MEX", domicilio.Pais)
	assert.Equal(t, "03100", domicilio.CodigoPostal)
}

func InternalTestPropietarioComercioExterior(t *testing.T, propietarios *[]comexterior1.PropietarioComercioExterior20) {
	assert.NotNil(t, propietarios)
	assert.Equal(t, len(*propietarios), 2)
	propietario := (*propietarios)[0]
	assert.Equal(t, "ABC123456789", propietario.NumRegIdTrib)
	assert.Equal(t, "USA", propietario.ResidenciaFiscal)
}
func InternalTestReceptorComercioExterior(t *testing.T, receptor *comexterior1.ReceptorComercioExterior20) {
	assert.Equal(t, "XYZ123456789", *receptor.NumRegIdTrib)
	assert.NotNil(t, receptor.Domicilio)
}
func InternalTestDestinatarioComercioExterior(t *testing.T, destinatarios *[]comexterior1.DestinatarioComercioExterior20) {
	assert.NotNil(t, destinatarios)
	assert.Equal(t, len(*destinatarios), 3)
	destinatario0 := (*destinatarios)[0]
	assert.Equal(t, "X12345678", *destinatario0.NumRegIdTrib)
	assert.Equal(t, "Comercializadora Internacional S.A. de C.V.", *destinatario0.Nombre)
	destinatario1 := (*destinatarios)[1]
	assert.Equal(t, len(destinatario1.Domicilio), 3)
}
func InternalTestMercanciaComercioExterior(t *testing.T, mercancias []comexterior1.MercanciaComercioExterior20) {
	assert.NotNil(t, mercancias)
	assert.Equal(t, len(mercancias), 2)
	mercancia0 := mercancias[0]
	assert.Equal(t, "12345ABC", mercancia0.NoIdentificacion)
	assert.Equal(t, "010121", *mercancia0.FraccionArancelaria)
	assert.Equal(t, 100.000, *mercancia0.CantidadAduana)
	assert.Equal(t, "KGM", *mercancia0.UnidadAduana)
	assert.Equal(t, 50.123456, *mercancia0.ValorUnitarioAduana)
	assert.Equal(t, 5000.1234, mercancia0.ValorDolares)
	mercancia1 := mercancias[1]
	assert.Equal(t, len(*mercancia1.DescripcionesEspecificas), 2)
	InternalTestDescripcionesEspecificasComercioExterior(t, mercancia1.DescripcionesEspecificas)
}
func InternalTestDescripcionesEspecificasComercioExterior(t *testing.T, descripciones *[]comexterior1.DescripcionesEspecificasComercioExterior20) {
	descripcion := (*descripciones)[1]
	assert.NotNil(t, descripcion)
	assert.Equal(t, "Ford", descripcion.Marca)
	assert.Equal(t, "Focus", *descripcion.Modelo)
	assert.Equal(t, "SE", *descripcion.SubModelo)
	assert.Equal(t, "DEF456XYZ789", *descripcion.NumeroSerie)
}
