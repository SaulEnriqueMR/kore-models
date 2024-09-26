package comprobante

import (
	"encoding/xml"
	comprobantes "github.com/SaulEnriqueMR/kore-models/models/comprobante"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetComprobante32ForTestComplementoTest(filename string, t *testing.T) (comprobantes.Comprobante32, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed comprobantes.Comprobante32
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	return parsed, errUnmarshal
}

func GetComprobante33ForTestComplementoTest(filename string, t *testing.T) (comprobantes.Comprobante33, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed comprobantes.Comprobante33
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	return parsed, errUnmarshal
}

func TestComprobanteComplementos(t *testing.T) {
	InternalTestComplementoComprobante32(t)
	InternalTestComplementoComprobante33(t)
}

func InternalTestComplementoComprobante32(t *testing.T) {
	comprobante32, _ := GetComprobante32ForTestComplementoTest("./comprobante32.xml", t)
	InternalTestFullAtributesComplementoComprobantes(t, *comprobante32.Complemento)
}

func InternalTestComplementoComprobante33(t *testing.T) {
	comprobante33, _ := GetComprobante33ForTestComplementoTest("./comprobante33.xml", t)
	InternalTestFullAtributesComplementoComprobantes(t, *comprobante33.Complemento)
}

func InternalTestFullAtributesComplementoComprobantes(t *testing.T, complemento comprobantes.Complemento) {
	assert.NotNil(t, complemento)
	assert.Equal(t, 3, len(*complemento.Aerolineas.Aerolineas10))
	assert.Equal(t, 5, len(*complemento.ImpuestoLocales.ImpuestoLocales10))
	assert.Equal(t, 3, len(*complemento.CertificadoDestruccion.CertificadoDeDestruccion10))
	assert.Equal(t, 3, len(*complemento.CFDIRegistroFiscal.CfdiRegistroFiscal10))
	assert.Equal(t, 3, len(complemento.ComplementoSPEI.SpeiTercero))
	assert.Equal(t, 3, len(*complemento.ConsumoCombustibles.ConsumoDeCombustible11))
	assert.Equal(t, 3, len(*complemento.Detallista.Detallista10))
	assert.Equal(t, 3, len(*complemento.Divisas.Divisas10))
	assert.Equal(t, 3, len(*complemento.Donatarias.Donatarias11))
	assert.Equal(t, 2, len(*complemento.EstadoCuentaCombustible.EstadoDeCuentaCombustible11))
	assert.Equal(t, 2, len(*complemento.EstadoCuentaCombustible.EstadoDeCuentaCombustible12))
	assert.Equal(t, 3, len(*complemento.GastoHidrocarburos.GastoHidrocarburos10))
	assert.Equal(t, 3, len(*complemento.Ine.INE11))
	assert.Equal(t, 3, len(*complemento.LeyendasFiscales.LeyendasFiscales10))
	assert.Equal(t, 2, len(*complemento.Nomina.Nomina11))
	assert.Equal(t, 2, len(*complemento.Nomina.Nomina12))
	assert.Equal(t, 2, len(*complemento.NotariosPublicos.NotariosPublicos10))
	assert.Equal(t, 2, len(*complemento.ComercioExterior.ComercioExterior10))
	assert.Equal(t, 2, len(*complemento.ComercioExterior.ComercioExterior11))
	assert.Equal(t, 3, len(*complemento.ComercioExterior.ComercioExterior20))
	assert.Equal(t, 2, len(*complemento.CartaPorte.CartaPorte20))
	assert.Equal(t, 2, len(*complemento.CartaPorte.CartaPorte30))
	assert.Equal(t, 2, len(*complemento.CartaPorte.CartaPorte31))
	assert.Equal(t, 2, len(*complemento.ObrasAntiguedades.ObrasAntiguedades10))
	assert.Equal(t, 2, len(*complemento.PagoEnEspecie.PagoEspecie10))
	assert.Equal(t, 2, len(*complemento.Pagos.Pagos10))
	assert.Equal(t, 2, len(*complemento.Pagos.Pagos20))
	assert.Equal(t, 2, len(*complemento.PFIntegrantesCoordinado.PFIntegranteCoordinado10))
	assert.Equal(t, 2, len(*complemento.ParcialesConstruccion.ParcialesConstruccion10))
	assert.Equal(t, 2, len(*complemento.RenovacionSustitucionVehiculos.RenovacionSustitucion10))
	assert.Equal(t, 2, len(*complemento.TuristaPasajeroExtranjero.TuristaPasajeroExtranjero10))
	assert.Equal(t, 2, len(*complemento.ValesDeDespensa.ValesDespensa10))
	assert.Equal(t, 1, len(*complemento.VehiculoUsado.VehiculoUsado10))
}
