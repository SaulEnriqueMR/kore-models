package helpers

import (
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParseDatetime(t *testing.T) {
	rfcInput := "2024-01-01"
	expectedRfc, _ := time.ParseInLocation(helpers.Rfc3339DatetimeLayout, rfcInput, helpers.MexicoLocation)
	rfcTest, _ := helpers.ParseDatetime(rfcInput)
	assert.Equal(t, expectedRfc, rfcTest)

	isoInput := "2024-01-01T13:20:00"
	expectedIso, _ := time.ParseInLocation(helpers.IsoDatetimeLayout, isoInput, helpers.MexicoLocation)
	isoTest, _ := helpers.ParseDatetime(isoInput)
	assert.Equal(t, expectedIso, isoTest)

	emptyInput := ""
	_, errEmpty := helpers.ParseDatetime(emptyInput)
	assert.Error(t, errEmpty)

	badInput := "sdasdasdouhe"
	_, errBad := helpers.ParseDatetime(badInput)
	assert.Error(t, errBad)

	unsupportedInput := "2024/01/01"
	_, errUnsupported := helpers.ParseDatetime(unsupportedInput)
	assert.Error(t, errUnsupported)
}
