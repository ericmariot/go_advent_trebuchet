package trebuchet

import (
	"fmt"
	"strconv"
	"unicode"
)

func CalibrationValueFromFilePartOne(filePath string) (int, error) {
	lines, err := ReadLinesFromFile(filePath)
	if err != nil {
		return 0, err
	}

	var sum int

	for _, line := range lines {
		var first, last rune
		foundFirst := false

		for _, char := range line {
			if unicode.IsDigit(char) {
				if !foundFirst {
					first = char
					foundFirst = true
				}
				last = char
			}
		}

		if foundFirst {
			calibration := strconv.Itoa(int(first-'0')) + strconv.Itoa(int(last-'0'))
			calibrationNum, err := strconv.Atoi(calibration)

			if err != nil {
				fmt.Printf("Error converting calibration to integer: %v", err)
				continue
			}
			sum += calibrationNum
		}
	}

	return sum, nil
}
