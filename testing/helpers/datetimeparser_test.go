package helpers

import (
	"testing"
	"time"

	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"github.com/stretchr/testify/assert"
)

func TestParseDatetime(t *testing.T) {
	rfcInput := "2024-01-01"
	expectedRfc, _ := time.ParseInLocation(helpers.Rfc3339DatetimeLayout, rfcInput, helpers.UTC)
	// expectedRfc, _ := time.ParseInLocation(helpers.Rfc3339DatetimeLayout, rfcInput, helpers.MexicoLocation)
	rfcTest, _ := helpers.ParseDatetime(rfcInput)
	assert.Equal(t, expectedRfc, rfcTest)

	isoInput := "2024-01-01T13:20:00"
	expectedIso, _ := time.ParseInLocation(helpers.IsoDatetimeLayout, isoInput, helpers.UTC)
	isoTest, _ := helpers.ParseDatetime(isoInput)
	assert.Equal(t, expectedIso, isoTest)

	xsDateInput := "2024-01-01Z"
	expectedXsDate, _ := time.Parse(time.RFC3339, "2024-01-01T00:00:00Z")
	xsDateTest, _ := helpers.ParseDatetime(xsDateInput)
	assert.Equal(t, expectedXsDate, xsDateTest)

	xsDateTimeInput := "2024-01-01+05:00"
	expectedXsDateTime, _ := time.Parse(time.RFC3339, "2024-01-01T00:00:00+05:00")
	xsDateTimeTest, _ := helpers.ParseDatetime(xsDateTimeInput)
	assert.Equal(t, expectedXsDateTime, xsDateTimeTest)

	xsDateTimeInput2 := "2024-01-01-05:00"
	expectedXsDateTime2, _ := time.Parse(time.RFC3339, "2024-01-01T00:00:00-05:00")
	xsDateTimeTest2, _ := helpers.ParseDatetime(xsDateTimeInput2)
	assert.Equal(t, expectedXsDateTime2, xsDateTimeTest2)

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

func TestISO8601(t *testing.T) {
	isoInput := "2025-03-07T17:18:05.55"
	expectedIso, _ := time.ParseInLocation(helpers.CustomLayout8601, isoInput, helpers.UTC)
	isoTest, _ := helpers.ParseDatetime(isoInput)
	assert.Equal(t, expectedIso, isoTest)
}
