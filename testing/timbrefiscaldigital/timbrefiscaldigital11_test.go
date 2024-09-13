package timbrefiscaldigital

import (
	"encoding/xml"
	"github.com/saulenriquemr/kore-models/models/timbrefiscaldigital"
	testing2 "github.com/saulenriquemr/kore-models/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFullTimbreFiscalDigital11(t *testing.T) {
	data := testing2.GetFileContentForTest("./timbrefiscaldigital11.xml", t)

	var parsed timbrefiscaldigital.TimbreFiscalDigital11
	errUnmashal := xml.Unmarshal(data, &parsed)
	if errUnmashal != nil {
		t.Error("Error unmarshalling xml: ", errUnmashal)
	}

	assert.Equal(t, "1.1", parsed.Version)
	assert.Equal(t, "3ea43e97-71bf-4ac5-a28e-9374ea9f8b45", parsed.Uuid)
	assert.Equal(t, "2024-04-29T10:46:30", parsed.Fecha)
	assert.Equal(t, "SPR190613I52", parsed.RfcProvCertif)
	assert.Equal(t, "AMifipYnPS5FuNWEf3ysVCru4FVG0eGp3", parsed.SelloCfd)
	assert.Equal(t, "30001000000500003456", parsed.NoCertificadoSat)
	assert.Equal(t, "QKKzapHEGa8tPZFtDuQi+450AtNiQENVSoK1uvQu", parsed.SelloSat)

	assert.NotNil(t, parsed.Leyenda)
	assert.Equal(t, "Leyenda SAT", *parsed.Leyenda)
}

func TestRequiredTimbreFiscalDigital11(t *testing.T) {
	data := testing2.GetFileContentForTest("./timbrefiscaldigital11_requeridos.xml", t)

	var parsed timbrefiscaldigital.TimbreFiscalDigital11
	errUnmashal := xml.Unmarshal(data, &parsed)
	if errUnmashal != nil {
		t.Error("Error unmarshalling xml: ", errUnmashal)
	}

	// Check that Leyenda is nil when it's not present in the XML
	assert.Nil(t, parsed.Leyenda)
}
