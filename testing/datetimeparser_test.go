package testing

import (
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"testing"
	"time"
)

func TestParseDatetime(t *testing.T) {
	rfcInput := "2024-01-01"
	expectedRfc, _ := time.Parse(helpers.Rfc3339DatetimeLayout, rfcInput)
	rfcTest, _ := helpers.ParseDatetime(rfcInput)
	if rfcTest != expectedRfc {
		t.Errorf("ParseDatetime failed: expected %s, got %s", expectedRfc, rfcTest)
	}

	isoInput := "2024-01-01T13:20:00"
	expectedIso, _ := time.Parse(helpers.IsoDatetimeLayout, isoInput)
	isoTest, _ := helpers.ParseDatetime(isoInput)
	if isoTest != expectedIso {
		t.Errorf("ParseDatetime failed: expected %s, got %s", expectedIso, isoTest)
	}

	emptyInput := ""
	_, errEmpty := helpers.ParseDatetime(emptyInput)
	if errEmpty == nil {
		t.Errorf("ParseDatetime should have failed due to empty string")
	}

	badInput := "sdasdasdouhe"
	_, errBad := helpers.ParseDatetime(badInput)
	if errBad == nil {
		t.Errorf("ParseDatetime should have failed due to bad string")
	}

	unsupportedInput := "2024/01/01"
	_, errUnsupported := helpers.ParseDatetime(unsupportedInput)
	if errUnsupported == nil {
		t.Errorf("ParseDatetime should have failed due to unsupported input")
	}
}
