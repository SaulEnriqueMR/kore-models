package turistapasajeroextranjero

import (
	"encoding/xml"
	turispasajextra1 "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/turistapasajeroextranjero"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetTurisPasajExtraForTest(filename string, t *testing.T) (turispasajextra1.TuristaPasajeroExtranjero10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed turispasajextra1.TuristaPasajeroExtranjero10
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	testing2.GenerateJSONFromStructure("turista10.json", parsed)
	return parsed, errUnmashal
}

func TestFullTurisPasajExtra(t *testing.T) {
	turispasajextra, _ := GetTurisPasajExtraForTest("./turistapasajeroextranjero.xml", t)
	assert.NotNil(t, turispasajextra)
	InternalTestBase(t, turispasajextra)
}

func InternalTestBase(t *testing.T, turista turispasajextra1.TuristaPasajeroExtranjero10) {
	assert.Equal(t, "1.0", turista.Version)
	//Testear el parseo de fecha ISO
	//Crear Variable ISO para comparar
	// var IsoDatetimeLayout = "2006-01-02T15:04:05"
	//Crear Fecha ISO apartir de la fecha string del cml
	// isoDate, _ := time.Parse(IsoDatetimeLayout, "2024-09-19T14:30:00")
	//Comparar fecha ISO creada con la fecha ISO creada con el Unmarshal
	// assert.Equal(t, isoDate, turista.FechaTransito)
	assert.Equal(t, "2024-09-19T14:30:00", turista.FechaTransitoString)
	assert.Equal(t, "Arribo", turista.TipoTransito)
	InternalTestTuristaPasajeroExtranjero(t, turista.DatosTransito)
}

func InternalTestTuristaPasajeroExtranjero(t *testing.T, datostransito turispasajextra1.DatosTransito) {
	assert.Equal(t, "Aérea", datostransito.Via)
	assert.Equal(t, "Pasaporte", datostransito.TipoId)
	assert.Equal(t, "X123456789", datostransito.NumeroId)
	assert.Equal(t, "Canadiense", datostransito.Nacionalidad)
	assert.Equal(t, "Air Canada", datostransito.EmpresaTransporte)
	assert.Equal(t, "AC102", *datostransito.IdTransporte)
}
