package comercioexterior

import (
	"encoding/xml"
	"testing"

	comexterior1 "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/comercioexterior"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetComercioExterior11ForTest(filename string, t *testing.T) (comexterior1.ComercioExterior11, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed comexterior1.ComercioExterior11
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	testing2.GenerateJSONFromStructure("comercioexterior11.json", parsed)
	return parsed, errUnmashal
}

func TestFullComercioExterior11(t *testing.T) {
	comexterior, _ := GetComercioExterior11ForTest("./comercioexterior11.xml", t)
	assert.NotNil(t, comexterior)
	InternalTestComercioExterior11(t, comexterior)
}

func InternalTestComercioExterior11(t *testing.T, comercioExterior comexterior1.ComercioExterior11) {
	assert.Equal(t, "1.1", comercioExterior.Version)
	assert.Equal(t, "EXPORTACION", comercioExterior.TipoOperacion)
	assert.Equal(t, "A1", *comercioExterior.ClavePedimento)
	assert.Equal(t, "01", *comercioExterior.MotivoTraslado)
	assert.Equal(t, 1.0, *comercioExterior.CertificadoOrigen)
	assert.Equal(t, "123456789012", *comercioExterior.NoCertificadoOrigen)
	assert.Equal(t, "ABC123456", *comercioExterior.NoExportadorConfiable)
	assert.Equal(t, "CIF", *comercioExterior.Incoterm)
	assert.Equal(t, 0.0, *comercioExterior.Subdivision)
	assert.Equal(t, "Este es un ejemplo de comercio exterior.", *comercioExterior.Observaciones)
	assert.Equal(t, "20.50", *comercioExterior.TipoCambioUsd)
	assert.Equal(t, 10000.00, *comercioExterior.TotalUsd)
	InternalTestEmisorComercioExterior11(t, comercioExterior.Emisor)
	InternalTestPropietarioComercioExterior11(t, comercioExterior.Propietario)
	InternalTestReceptorComercioExterior11(t, comercioExterior.Receptor)
	InternalTestDestinatarioComercioExterior11(t, comercioExterior.Destinatario)
	InternalTestMercanciaComercioExterior11(t, comercioExterior.Mercancias)
}

func InternalTestEmisorComercioExterior11(t *testing.T, emisor *comexterior1.EmisorComercioExterior11) {
	assert.Equal(t, "ABC123456789", *emisor.Curp)

	assert.NotNil(t, emisor.Domicilio)
	domicilio := emisor.Domicilio

	assert.Equal(t, "Av. Reforma", domicilio.Calle)
	assert.Equal(t, "123", *domicilio.NoExterior)
	assert.Equal(t, "456", *domicilio.NoInterior)
	assert.Equal(t, "Centro", *domicilio.Colonia)
	assert.Equal(t, "Cuauhtémoc", *domicilio.Localidad)
	assert.Equal(t, "Cerca del ángel de la independencia", *domicilio.Referencia)
	assert.Equal(t, "Cuauhtémoc", *domicilio.Municipio)
	assert.Equal(t, "DF", domicilio.Estado)
	assert.Equal(t, "MEX", domicilio.Pais)
	assert.Equal(t, "06000", domicilio.CodigoPostal)
}

func InternalTestPropietarioComercioExterior11(t *testing.T, propietarios *[]comexterior1.PropietarioComercioExterior11) {
	assert.NotNil(t, propietarios)
	assert.Equal(t, len(*propietarios), 3)
	propietario := (*propietarios)[0]
	assert.Equal(t, "RFC123456789", propietario.NumRegIdTrib)
	assert.Equal(t, "MEX", propietario.ResidenciaFiscal)
}

func InternalTestReceptorComercioExterior11(t *testing.T, receptor *comexterior1.ReceptorComercioExterior11) {
	assert.Equal(t, "RFC123456789", *receptor.NumRegIdTrib)

	if receptor.Domicilio != nil {
		assert.Equal(t, "Av. Insurgentes Sur", receptor.Domicilio.Calle)
		assert.Equal(t, "123", *receptor.Domicilio.NoExterior)
		assert.Equal(t, "A", *receptor.Domicilio.NoInterior)
		assert.Equal(t, "Roma", *receptor.Domicilio.Colonia)
		assert.Equal(t, "Ciudad de México", *receptor.Domicilio.Localidad)
		assert.Equal(t, "Cerca del parque", *receptor.Domicilio.Referencia)
		assert.Equal(t, "Benito Juárez", *receptor.Domicilio.Municipio)
		assert.Equal(t, "Ciudad de México", receptor.Domicilio.Estado)
		assert.Equal(t, "MEX", receptor.Domicilio.Pais)
		assert.Equal(t, "06760", receptor.Domicilio.CodigoPostal)
	}
}

func InternalTestDestinatarioComercioExterior11(t *testing.T, destinatarios *[]comexterior1.DestinatarioComercioExterior11) {
	assert.NotNil(t, destinatarios)
	assert.Equal(t, len(*destinatarios), 1)
	destinatario := (*destinatarios)[0]
	assert.Equal(t, "RFC987654321", *destinatario.NumRegIdTrib)
	assert.Equal(t, "Juan Pérez", *destinatario.Nombre)

	if len(destinatario.Domicilio) > 0 {
		domicilio := destinatario.Domicilio[0]
		assert.Equal(t, "Calle de la Paz", domicilio.Calle)
		assert.Equal(t, "45", *domicilio.NoExterior)
		assert.Equal(t, "B", *domicilio.NoInterior)
		assert.Equal(t, "Centro", *domicilio.Colonia)
		assert.Equal(t, "Ciudad de México", *domicilio.Localidad)
		assert.Equal(t, "Cerca de la Plaza", *domicilio.Referencia)
		assert.Equal(t, "Cuauhtémoc", *domicilio.Municipio)
		assert.Equal(t, "Ciudad de México", domicilio.Estado)
		assert.Equal(t, "MEX", domicilio.Pais)
		assert.Equal(t, "06000", domicilio.CodigoPostal)
	}
}

func InternalTestMercanciaComercioExterior11(t *testing.T, mercancias *[]comexterior1.MercanciaComercioExterior11) {
	assert.NotNil(t, mercancias)
	assert.Equal(t, len(*mercancias), 2)
	mercancia := (*mercancias)[0]
	assert.Equal(t, "ABC123456", mercancia.NoIdentificacion)
	assert.Equal(t, 1500.00, mercancia.ValorDolares)

	if mercancia.DescripcionesEspecificas != nil && len(*mercancia.DescripcionesEspecificas) > 0 {
		desc1 := (*mercancia.DescripcionesEspecificas)[0]
		assert.Equal(t, "MarcaX", desc1.Marca)
		assert.Equal(t, "ModeloY", *desc1.Modelo)
		assert.Equal(t, "SubModeloZ", *desc1.Submodelo)
		assert.Equal(t, "123456789", *desc1.NoSerie)

		desc2 := (*mercancia.DescripcionesEspecificas)[1]
		assert.Equal(t, "MarcaX", desc2.Marca)
		assert.Equal(t, "ModeloY", *desc2.Modelo)
		assert.Equal(t, "SubModeloZ", *desc2.Submodelo)
		assert.Equal(t, "123456789", *desc2.NoSerie)

		desc3 := (*mercancia.DescripcionesEspecificas)[2]
		assert.Equal(t, "MarcaA", desc3.Marca)
	}
}
