package testing

import (
	"kore-models/models"
	"testing"
	"time"
)

func TestParseDatetime(t *testing.T) {
	rfcInput := "2024-01-01"
	expectedRfc, _ := time.Parse(models.Rfc3339DatetimeLayout, rfcInput)
	rfcTest, _ := models.ParseDatetime(rfcInput)
	if rfcTest != expectedRfc {
		t.Errorf("ParseDatetime failed: expected %s, got %s", expectedRfc, rfcTest)
	}

	isoInput := "2024-01-01T13:20:00"
	expectedIso, _ := time.Parse(models.IsoDatetimeLayout, isoInput)
	isoTest, _ := models.ParseDatetime(isoInput)
	if isoTest != expectedIso {
		t.Errorf("ParseDatetime failed: expected %s, got %s", expectedIso, isoTest)
	}

	emptyInput := ""
	_, errEmpty := models.ParseDatetime(emptyInput)
	if errEmpty == nil {
		t.Errorf("ParseDatetime should have failed due to empty string")
	}

	badInput := "sdasdasdouhe"
	_, errBad := models.ParseDatetime(badInput)
	if errBad == nil {
		t.Errorf("ParseDatetime should have failed due to bad string")
	}

	unsupportedInput := "2024/01/01"
	_, errUnsupported := models.ParseDatetime(unsupportedInput)
	if errUnsupported == nil {
		t.Errorf("ParseDatetime should have failed due to unsupported input")
	}
}
