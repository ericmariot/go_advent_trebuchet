package trebuchet_test

import (
	"testing"

	"github.com/ericmariot/goAdventTrebuchet/trebuchet"
)

func TestCalibrationValueFromFile_CorrectSum_PartTwo(t *testing.T) {
	filePath := "./input_test_two.txt"
	want := 281

	result, err := trebuchet.CalibrationValueFromFilePartTwo(filePath)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	if result != want {
		t.Errorf("Expected: %v, got: %v", want, result)
	}
}
