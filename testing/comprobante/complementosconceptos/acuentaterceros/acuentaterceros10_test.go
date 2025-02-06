package acuentaterceros

import (
	"encoding/xml"
	"testing"

	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementosconcepto/acuentaterceros"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetACuentaTercerosForTest(filename string, t *testing.T) (acuentaterceros.ACuentaTerceros11, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed acuentaterceros.ACuentaTerceros11
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	testing2.GenerateJSONFromStructure("acuentaterceros10.json", parsed)
	return parsed, errUnmashal
}

func TestFullACuentaTerceros(t *testing.T) {
	cuenta, _ := GetACuentaTercerosForTest("./acuentaterceros10.xml", t)
	InternalBaseAttributes(t, cuenta)
	InternalTestInformacionFiscalTercero(t, cuenta.InformacionFiscalTercero)
	InternalTestInformacionAduaneraBaseTerceros11(t, cuenta.InformacionAduanera)
	InternalTestParteTerceros11(t, cuenta.Parte)
	InternalTestCuentaPredialTerceros11(t, cuenta.CuentaPredial)
	impuestos := cuenta.Impuestos
	InternalTestRetencionTerceros11(t, impuestos.Retenciones)
	InternalTestTrasladoTerceros11(t, impuestos.Traslados)
}

func InternalTestInformacionAduaneraBaseTerceros11(t *testing.T, info *acuentaterceros.InformacionAduaneraTerceros11) {
	assert.Equal(t, "123456", info.Numero)
	assert.Equal(t, "2024-09-23", info.FechaString)
	assert.Equal(t, "Aduana Ejemplo", *info.Aduana)
}

func InternalBaseAttributes(t *testing.T, cuenta acuentaterceros.ACuentaTerceros11) {
	assert.Equal(t, "1.1", cuenta.Version)
	assert.Equal(t, "AAA010101AAA", cuenta.Rfc)
	assert.Equal(t, "Juan Pérez", *cuenta.Nombre)
}

func InternalTestInformacionFiscalTercero(t *testing.T, infoFiscal *acuentaterceros.InformacionFiscalTerceroTerceros11) {
	assert.Equal(t, "Av. Siempre Viva", infoFiscal.Calle)
	assert.Equal(t, "123", *infoFiscal.NoExterior)
	assert.Equal(t, "Centro", *infoFiscal.Colonia)
	assert.Equal(t, "Ciudad", infoFiscal.Municipio)
	assert.Equal(t, "Estado", infoFiscal.Estado)
	assert.Equal(t, "México", infoFiscal.Pais)
	assert.Equal(t, "01234", infoFiscal.CodigoPostal)
}

func InternalTestParteTerceros11(t *testing.T, partes *[]acuentaterceros.ParteTerceros11) {
	assert.NotNil(t, partes)
	assert.Equal(t, len(*partes), 1)
	parte := (*partes)[0]
	assert.Equal(t, 10.00, parte.Cantidad)
	assert.Equal(t, "Componente X", parte.Descripcion)
	assert.Equal(t, 100.00, *parte.ValorUnitario)
	assert.Equal(t, 1000.00, *parte.Importe)
	InternalTestInformacionAduaneraTerceros11(t, parte.InformacionAduanera)
}

func InternalTestInformacionAduaneraTerceros11(t *testing.T, infos *[]acuentaterceros.InformacionAduaneraTerceros11) {
	assert.NotNil(t, infos)
	assert.Equal(t, len(*infos), 1)
	info := (*infos)[0]
	assert.Equal(t, "123456", info.Numero)
	assert.Equal(t, "2024-09-23", info.FechaString)
	assert.Equal(t, "Aduana Ejemplo", *info.Aduana)
}

func InternalTestCuentaPredialTerceros11(t *testing.T, cuenta *acuentaterceros.CuentaPredialTerceros11) {
	assert.Equal(t, "00000123", cuenta.Numero)
}

func InternalTestRetencionTerceros11(t *testing.T, retenciones *[]acuentaterceros.RetencionTerceros11) {
	assert.NotNil(t, retenciones)
	assert.Equal(t, len(*retenciones), 1)
	retencion := (*retenciones)[0]
	assert.Equal(t, "ISR", retencion.TipoImpuesto)
	assert.Equal(t, "001", retencion.Impuesto)
	assert.Equal(t, 150.00, retencion.Importe)
}

func InternalTestTrasladoTerceros11(t *testing.T, traslados *[]acuentaterceros.TrasladoTerceros11) {
	assert.NotNil(t, traslados)
	assert.Equal(t, len(*traslados), 1)
	traslado := (*traslados)[0]
	assert.Equal(t, "IVA", traslado.TipoImpuesto)
	assert.Equal(t, "002", traslado.Impuesto)
	assert.Equal(t, 16.00, traslado.Tasa)
	assert.Equal(t, 160.00, traslado.Importe)
}
