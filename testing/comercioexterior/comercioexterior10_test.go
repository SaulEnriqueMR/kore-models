package comercioexterior

import (
	"encoding/xml"
	"testing"

	comexterior1 "github.com/SaulEnriqueMR/kore-models/models/comercioexterior"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetTurisComercioExterior10ForTest(filename string, t *testing.T) (comexterior1.ComercioExterior10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed comexterior1.ComercioExterior10
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	return parsed, errUnmashal
}

func TestFullComercioExterior10(t *testing.T) {
	comexterior, _ := GetTurisComercioExterior10ForTest("./comercioexterior10.xml", t)
	assert.NotNil(t, comexterior)
	InternalTestComercioExterior10(t, comexterior)
}

func InternalTestComercioExterior10(t *testing.T, comercio comexterior1.ComercioExterior10) {
	assert.Equal(t, "1.0", comercio.Version)
	assert.Equal(t, "1", comercio.TipoOperacion)
	assert.Equal(t, "ABC123", *comercio.ClavePedimento)
	assert.Equal(t, 1.0, *comercio.CertificadoOrigen)
	assert.Equal(t, "123e4567-e89b-12d3-a456-426614174000", *comercio.NumCertificadoOrigen)
	assert.Equal(t, "EX123456", *comercio.NumeroExportadorConfiable)
	assert.Equal(t, "CIF", *comercio.Incoterm)
	assert.Equal(t, 0.0, *comercio.Subdivision)
	assert.Equal(t, "Factura de exportación.", *comercio.Observaciones)
	assert.Equal(t, 20.50, *comercio.TipoCambioUSD)
	assert.Equal(t, 1000.00, *comercio.TotalUSD)
	InternalTestEmisorComercioExterior10(t, comercio.Emisor)
	InternalTestReceptorComercioExterior10(t, comercio.Receptor)
	InternalTestDestinatarioComercioExterior10(t, comercio.Destinatario)
	InternalTestMercanciaComercioExterior10(t, comercio.Mercancias)
}
func InternalTestEmisorComercioExterior10(t *testing.T, emisor *comexterior1.EmisorComercioExterior10) {
	assert.Equal(t, "ABCD123456HDFRLM09", *emisor.Curp) // Asegúrate de que el puntero no sea nil antes de desreferenciar
}
func InternalTestReceptorComercioExterior10(t *testing.T, receptor comexterior1.ReceptorComercioExterior10) {
	assert.Equal(t, "XYZ123456HDFGJ00", *receptor.Curp)
	assert.Equal(t, "RFC123456789", receptor.NumRegIdTrib)
}
func InternalTestDestinatarioComercioExterior10(t *testing.T, destinatario *comexterior1.DestinatarioComercioExterior10) {
	assert.Equal(t, "123456789012", *destinatario.NumRegIdTrib)
	assert.Equal(t, "ABC123456XYZ", *destinatario.Rfc)
	assert.Equal(t, "ABCD123456HDFXYZ00", *destinatario.Curp)
	assert.Equal(t, "Juan Pérez García", *destinatario.Nombre)
	InternalTestDomicilioComercioExterior10(t, destinatario.Domicilio)
}
func InternalTestDomicilioComercioExterior10(t *testing.T, domicilio comexterior1.DomicilioComercioExterior10) {
	assert.Equal(t, "Av. Siempre Viva", domicilio.Calle)
	assert.Equal(t, "Illinois", domicilio.Estado)
	assert.Equal(t, "US", domicilio.Pais)
	assert.Equal(t, "62704", domicilio.CodigoPostal)
	assert.Nil(t, domicilio.NumeroExterior)
	assert.Nil(t, domicilio.NumeroInterior)
	assert.Nil(t, domicilio.Colonia)
	assert.Nil(t, domicilio.Localidad)
	assert.Nil(t, domicilio.Referencia)
	assert.Nil(t, domicilio.Municipio)
}
func InternalTestMercanciaComercioExterior10(t *testing.T, mercancias *[]comexterior1.MercanciaComercioExterior10) {
	assert.NotNil(t, mercancias)
	assert.Equal(t, len(*mercancias), 1)
	mercancia := (*mercancias)[0]
	assert.Equal(t, "ABC123456", mercancia.NoIdentificacion)
	assert.Equal(t, "123456", *mercancia.FraccionArancelaria)
	assert.Equal(t, 10.000, *mercancia.CantidadAduana)
	assert.Equal(t, "PIE", *mercancia.UnidadAduana)
	assert.Equal(t, 150.00, *mercancia.ValorUnitarioAduana)
	assert.Equal(t, 1500.00, mercancia.ValorDolares)
	InternalTestDescripcionesEspecificasComercioExterior10(t, mercancia.DescripcionesEspecificas)
}
func InternalTestDescripcionesEspecificasComercioExterior10(t *testing.T, descripciones *[]comexterior1.DescripcionesEspecificasComercioExterior10) {
	assert.NotNil(t, descripciones)
	assert.Equal(t, len(*descripciones), 1)
	descripcion := (*descripciones)[0]
	assert.Equal(t, "MarcaEjemplo", descripcion.Marca)
	assert.Equal(t, "ModeloEjemplo", *descripcion.Modelo)
	assert.Equal(t, "SubmodeloEjemplo", *descripcion.SubModelo)
	assert.Equal(t, "SN123456789", *descripcion.NumeroSerie)
}
