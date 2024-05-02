package trebuchet_test

import (
	"testing"

	"github.com/ericmariot/goAdventTrebuchet/trebuchet"
	"github.com/stretchr/testify/require"
)

func TestCalibrationValueFromFile_CorrectSum_PartTwo(t *testing.T) {
	filePath := "./input_test_two.txt"

	result, err := trebuchet.CalibrationValueFromFilePartTwo(filePath)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	require.Equal(t, 281, result)
}
