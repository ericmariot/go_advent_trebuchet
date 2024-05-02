package trebuchet_test

import (
	"testing"

	"github.com/ericmariot/goAdventTrebuchet/trebuchet"
)

func TestCalibrationValueFromFile_CorrectSum_PartOne(t *testing.T) {
	filePath := "./input_test_one.txt"
	want := 142

	result, err := trebuchet.CalibrationValueFromFilePartOne(filePath)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	if result != want {
		t.Errorf("Expected: %v, got: %v", want, result)
	}
}
