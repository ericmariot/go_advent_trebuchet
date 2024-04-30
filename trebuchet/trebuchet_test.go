package trebuchet

import "testing"

func TestFirstItem(t *testing.T) {
	filePath := "./input_test.txt"
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
