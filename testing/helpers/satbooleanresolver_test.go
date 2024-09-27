package helpers

import (
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNormalizeSatBoolean(t *testing.T) {
	expectedSi := "SI"

	assert.Equal(t, expectedSi, helpers.NormalizeSatBoolean("sí"))
	assert.Equal(t, expectedSi, helpers.NormalizeSatBoolean(" sí"))
	assert.Equal(t, expectedSi, helpers.NormalizeSatBoolean(" sí "))
	assert.Equal(t, expectedSi, helpers.NormalizeSatBoolean("sí "))
	assert.Equal(t, expectedSi, helpers.NormalizeSatBoolean("Sí"))
	assert.Equal(t, expectedSi, helpers.NormalizeSatBoolean("SÍ"))
	assert.Equal(t, expectedSi, helpers.NormalizeSatBoolean("sÍ"))

	expectedNo := "NO"
	assert.Equal(t, expectedNo, helpers.NormalizeSatBoolean("nó"))
	assert.Equal(t, expectedNo, helpers.NormalizeSatBoolean(" nó"))
	assert.Equal(t, expectedNo, helpers.NormalizeSatBoolean(" nó "))
	assert.Equal(t, expectedNo, helpers.NormalizeSatBoolean("nó "))
	assert.Equal(t, expectedNo, helpers.NormalizeSatBoolean("Nó"))
	assert.Equal(t, expectedNo, helpers.NormalizeSatBoolean("NÓ"))
	assert.Equal(t, expectedNo, helpers.NormalizeSatBoolean("nÓ"))

	assert.NotEqual(t, expectedSi, helpers.NormalizeSatBoolean("pr4u3bwd"))
}

func TestResolveSatBoolean(t *testing.T) {
	assert.Equal(t, true, helpers.ResolveSatBoolean("SI"))
	assert.Equal(t, true, helpers.ResolveSatBoolean("Sí"))

	assert.Equal(t, false, helpers.ResolveSatBoolean("NO"))
	assert.Equal(t, false, helpers.ResolveSatBoolean("No"))
}
