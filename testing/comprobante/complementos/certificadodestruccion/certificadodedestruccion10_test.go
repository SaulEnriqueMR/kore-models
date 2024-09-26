package complementoconcepto

import (
	"encoding/xml"
	"testing"

	certificadodestruccion2 "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/certificadodestruccion"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetCertificadoDeDestruccion10ForTest(filename string, t *testing.T) (certificadodestruccion2.CertificadoDeDestruccion10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed certificadodestruccion2.CertificadoDeDestruccion10
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	return parsed, errUnmashal
}

func TestCertificadoDeDestruccion10(t *testing.T) {
	certificadoDestruccion10, _ := GetCertificadoDeDestruccion10ForTest("./certificadodedestruccion10.xml", t)
	InternalTestFullAtributesCertDes10(t, certificadoDestruccion10)
	InternalTestFullAtributesVehiculoDestruido10(t, certificadoDestruccion10.VehiculoDestruido)
	InternalTestFullAtributesInformacinoAduaneraCertDest10(t, certificadoDestruccion10.InformacinoAduanera)
}

func InternalTestFullAtributesCertDes10(t *testing.T, certificadoDestruccion certificadodestruccion2.CertificadoDeDestruccion10) {
	assert.Equal(t, "1.0", certificadoDestruccion.Version)
	assert.Equal(t, "SERIE A", certificadoDestruccion.Serie)
	assert.Equal(t, "000323", certificadoDestruccion.NumFolDesVeh)
}

func InternalTestFullAtributesVehiculoDestruido10(t *testing.T, vehiculoDestruido *[]certificadodestruccion2.VehiculoDestruidoCertDest10) {
	assert.NotNil(t, vehiculoDestruido)
	assert.Equal(t, 1, len(*vehiculoDestruido))
	first := (*vehiculoDestruido)[0]

	assert.Equal(t, "Nissan", first.Marca)
	assert.Equal(t, "Camioneta", first.TipoOClase)
	assert.Equal(t, "2008", first.Anio)
	assert.NotNil(t, first.Modelo)
	assert.Equal(t, "Pick-Up", *first.Modelo)
	assert.NotNil(t, first.Niv)
	assert.Equal(t, "39883", *first.Niv)
	assert.NotNil(t, first.NumSerie)
	assert.Equal(t, "IGIYY2189E5138227", *first.NumSerie)
	assert.Equal(t, "TJN-25-29", first.NumPlacas)
	assert.NotNil(t, first.NumMotor)
	assert.Equal(t, "0445542897692", *first.NumMotor)
	assert.Equal(t, "05413312114", first.NumFolTarjCir)
}

func InternalTestFullAtributesInformacinoAduaneraCertDest10(t *testing.T, informacionAduanera *[]certificadodestruccion2.InformacionAduaneraCertDest10) {
	assert.NotNil(t, informacionAduanera)
	assert.Equal(t, 1, len(*informacionAduanera))
	first := (*informacionAduanera)[0]

	assert.Equal(t, "2390321", first.NumPedImp)
	assert.Equal(t, "2014-01-23", first.Fecha)
	assert.Equal(t, "Aduana", first.Aduana)
}
