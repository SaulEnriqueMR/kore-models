package complementoconcepto

import (
	"encoding/xml"
	"testing"

	ine112 "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/ine"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetIne11ForTest(filename string, t *testing.T) (ine112.INE11, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed ine112.INE11
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	testing2.GenerateJSONFromStructure("ine11.json", parsed)
	return parsed, errUnmarshal
}

func TestIne11(t *testing.T) {
	ine11, _ := GetIne11ForTest("./ine11.xml", t)
	InternalTestFullAtributesIne11(t, ine11)
	InternalTestFullAtributesEntidadIne11(t, *ine11.Entidad)
	entidades := ine11.Entidad
	contabilidad1 := (*entidades)[0].Contabilidad
	contabilidad2 := (*entidades)[1].Contabilidad
	contabilidades := []ine112.ContabilidadINE11{(*contabilidad1)[0], (*contabilidad2)[0]}
	InternalTestFullAtributesContabilidadIne11(t, contabilidades)
}

func InternalTestFullAtributesIne11(t *testing.T, ine11 ine112.INE11) {
	assert.Equal(t, "1.1", ine11.Version)
	assert.NotNil(t, *ine11.TipoComite)
	assert.Equal(t, "Ordinario", *ine11.TipoComite)
	assert.Equal(t, "Campaña", ine11.TipoProceso)
	assert.Equal(t, int32(123), ine11.IdContabilidad)
}

func InternalTestFullAtributesEntidadIne11(t *testing.T, entidad []ine112.EntidadINE11) {
	assert.NotNil(t, entidad)
	assert.Equal(t, 2, len(entidad))
	assert.Equal(t, "OAX", entidad[0].ClaveEntidad)
	assert.Equal(t, "Local", *entidad[0].Ambito)
	assert.Equal(t, "GUA", entidad[1].ClaveEntidad)
	assert.Equal(t, "Local", *entidad[1].Ambito)
}

func InternalTestFullAtributesContabilidadIne11(t *testing.T, contabilidad []ine112.ContabilidadINE11) {
	assert.NotNil(t, contabilidad)
	assert.Equal(t, 2, len(contabilidad))
	assert.Equal(t, int32(654321), contabilidad[0].IdContabilidad)
	assert.Equal(t, int32(123546), contabilidad[1].IdContabilidad)
}
