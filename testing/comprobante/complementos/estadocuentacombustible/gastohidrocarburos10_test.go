package complementos

import (
	"encoding/xml"
	"testing"

	hidrocarburos2 "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/gastohidrocarburos"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetGastoHidrocarburos10ForTest(filename string, t *testing.T) (hidrocarburos2.GastoHidrocarburos10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed hidrocarburos2.GastoHidrocarburos10
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	return parsed, errUnmarshal
}

func TestGastoHidrocarburos10(t *testing.T) {
	gastoHidrocarburos10, _ := GetGastoHidrocarburos10ForTest("./gastohidrocarburos10.xml", t)
	InternalTestFullAtributesGastoHidrocarburos10(t, gastoHidrocarburos10)
	InternalTestFullAtributesErogacionHidrocarburos10(t, gastoHidrocarburos10.Erogacion[0])
	erogacion := gastoHidrocarburos10.Erogacion[0]
	InternalTestFullAtributesDocumentoRelacionadoGastoHidrocar10(t, erogacion.DocumentoRelacionado[0])
	InternalTestFullAtributesActividadesHidrocar10(t, *erogacion.Actividades)
	actividades := erogacion.Actividades
	actividad := (*actividades)[0]
	subactividades := actividad.SubActividades
	InternalTestFullSubActividadesHidrocar10(t, *subactividades)
	subactividad := (*subactividades)[0]
	tareas := subactividad.Tareas
	InternalTestFullAtributesTareasHidrocar10(t, *tareas)
	InternalTestFullAtributesCentroCostosHidrocar10(t, *erogacion.CentroCostos)
	centroCostos := erogacion.CentroCostos
	centroCosto := (*centroCostos)[0]
	yacimientos := centroCosto.Yacimientos
	yacimiento := (*yacimientos)[0]
	InternalTestFullAtributesYacimientosHidrocar10(t, *centroCosto.Yacimientos)
	InternalTestFullAtributesPozosHidrocar10(t, *yacimiento.Pozos)
}

func InternalTestFullAtributesGastoHidrocarburos10(t *testing.T, hidrocarburos10 hidrocarburos2.GastoHidrocarburos10) {
	assert.Equal(t, "1.0", hidrocarburos10.Version)
	assert.Equal(t, "ABC123456", hidrocarburos10.NumeroContrato)
	assert.Equal(t, "CentroCosto01", *hidrocarburos10.AreaContractual)
}

func InternalTestFullAtributesErogacionHidrocarburos10(t *testing.T, hidrocar10 hidrocarburos2.ErogacionGastoHidrocar10) {
	assert.Equal(t, "Gasto", hidrocar10.TipoErogacion)
	assert.Equal(t, "1500.00", hidrocar10.MontoCuErogacion)
	assert.Equal(t, "10.500", hidrocar10.Porcentaje)
}

func InternalTestFullAtributesDocumentoRelacionadoGastoHidrocar10(t *testing.T, hidrocarburos10 hidrocarburos2.DocumentoRelacionadoGastoHidrocar10) {
	assert.Equal(t, "Nacional", hidrocarburos10.OrigenErogacion)
	assert.Equal(t, "123e4567-e89b-12d3-a456-426614174000", *hidrocarburos10.FolioFiscalVinculado)
	assert.Equal(t, "ABC123456789", *hidrocarburos10.RFCProveedor)
	assert.Equal(t, 200.00, *hidrocarburos10.MontoTotalIVA)
	assert.Equal(t, 50.00, *hidrocarburos10.MontoRetencionISR)
	assert.Equal(t, 30.00, *hidrocarburos10.MontoRetencionIVA)
	assert.Equal(t, 10.00, *hidrocarburos10.MontoRetencionOtrosImpuestos)
	assert.Equal(t, "21011234", *hidrocarburos10.NumeroPedimentoVinculado)
	assert.Equal(t, "123456789", *hidrocarburos10.ClavePedimentoVinculado)
	assert.Equal(t, "987654321", *hidrocarburos10.ClavePagoPedimentoVinculado)
	assert.Equal(t, 100.00, *hidrocarburos10.MontoIVAPedimento)
	assert.Equal(t, 20.00, *hidrocarburos10.OtrosImpuestosPagadosPedimento)
	assert.Equal(t, "2024-01-15", *hidrocarburos10.FechaFolioFiscalVinculado)
	assert.Equal(t, "Febrero", hidrocarburos10.Mes)
	assert.Equal(t, 3000.00, hidrocarburos10.MontoTotalErogaciones)
}

func InternalTestFullAtributesActividadesHidrocar10(t *testing.T, actividades []hidrocarburos2.ActividadesGastoHidrocar10) {
	assert.NotNil(t, actividades)
	assert.Equal(t, 1, len(actividades))
	assert.Equal(t, "Actividad1", *actividades[0].ActividadRelacionada)
}

func InternalTestFullSubActividadesHidrocar10(t *testing.T, subActividades []hidrocarburos2.SubActividadesGastoHidrocar10) {
	assert.NotNil(t, subActividades)
	assert.Equal(t, 1, len(subActividades))
	assert.Equal(t, "SubActividadA", *subActividades[0].SubActividadRelacionada)
}

func InternalTestFullAtributesTareasHidrocar10(t *testing.T, tareas []hidrocarburos2.TareasGastoHidrocar10) {
	assert.NotNil(t, tareas)
	assert.Equal(t, 1, len(tareas))
	assert.Equal(t, "Tarea1", *tareas[0].TareaRelacionada)
}

func InternalTestFullAtributesCentroCostosHidrocar10(t *testing.T, centroCostos []hidrocarburos2.CentroCostosGastoHidrocar10) {
	assert.NotNil(t, centroCostos)
	assert.Equal(t, 1, len(centroCostos))
	assert.Equal(t, "Campo01", *centroCostos[0].Campo)
}

func InternalTestFullAtributesYacimientosHidrocar10(t *testing.T, yacimientos []hidrocarburos2.YacimientosGastoHidrocar10) {
	assert.NotNil(t, yacimientos)
	assert.Equal(t, 1, len(yacimientos))
	assert.Equal(t, "YacimientoA", *yacimientos[0].Yacimiento)
}

func InternalTestFullAtributesPozosHidrocar10(t *testing.T, pozos []hidrocarburos2.PozosGastoHidrocar10) {
	assert.NotNil(t, pozos)
	assert.Equal(t, 1, len(pozos))
	assert.Equal(t, "Pozo1", *pozos[0].Pozo)
}
