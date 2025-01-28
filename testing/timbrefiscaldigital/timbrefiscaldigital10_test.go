package timbrefiscaldigital

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"github.com/SaulEnriqueMR/kore-models/models/timbrefiscaldigital"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTimbreFiscalDigital10UnmarshalXML(t *testing.T) {
	data := testing2.GetFileContentForTest("./timbrefiscaldigital10.xml", t)

	var parsed timbrefiscaldigital.TimbreFiscalDigital10
	_ = xml.Unmarshal(data, &parsed)

	assert.Equal(t, "1.0", parsed.Version)
	assert.Equal(t, "3ea43e97-71bf-4ac5-a28e-9374ea9f8b45", parsed.Uuid)
	assert.Equal(t, "2024-04-29T10:46:30", parsed.Fecha)
	assert.Equal(t, "AMifipYnPS5FuNWEf3ysVCru4FVG0eGp3", parsed.SelloCfd)
	assert.Equal(t, "30001000000500003456", parsed.NoCertificadoSat)
	assert.Equal(t, "QKKzapHEGa8tPZFtDuQi+450AtNiQENVSoK1uvQu", parsed.SelloSat)

	InternalTestTimbreFiscalDigital10(parsed, t)
}

func InternalTestTimbreFiscalDigital10(tfd timbrefiscaldigital.TimbreFiscalDigital10, t *testing.T) {
	cadena := helpers.CreateCadenaOriginalFromTfd(tfd)
	assert.Equal(t, "||1.0|3ea43e97-71bf-4ac5-a28e-9374ea9f8b45|2024-04-29T10:46:30|AMifipYnPS5FuNWEf3ysVCru4FVG0eGp3|30001000000500003456|QKKzapHEGa8tPZFtDuQi+450AtNiQENVSoK1uvQu||", cadena)
}
