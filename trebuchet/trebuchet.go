package trebuchet

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func CalibrationValueFromFile(filePath string) (int, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return 0, err
	}

	var count int

	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		var first, last rune
		var calibration string

		for _, char := range line {
			if unicode.IsDigit(char) {
				if first == 0 {
					first = char
				}
				last = char
			}
		}
		calibration = string(first) + string(last)
		calibrationNum, err := strconv.Atoi(calibration)

		if err != nil {
			fmt.Printf("Error converting calibration to integer: %v", err)
			continue
		}

		count += calibrationNum
	}

	return count, nil
}
