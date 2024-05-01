package trebuchet

import "testing"

func TestCalibrationValueFromFile_CorrectSum_PartOne(t *testing.T) {
	filePath := "./input_test_one.txt"
	want := 142

	result, err := CalibrationValueFromFile(filePath)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	if result != want {
		t.Errorf("Expected: %v, got: %v", want, result)
	}
}

func TestCalibrationValueFromFile_CorrectSum_PartTwo(t *testing.T) {
	filePath := "./input_test_two.txt"
	want := 281

	result, err := CalibrationValueFromFile(filePath)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	if result != want {
		t.Errorf("Expected: %v, got: %v", want, result)
	}
}
