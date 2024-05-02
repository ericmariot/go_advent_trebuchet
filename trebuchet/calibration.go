package trebuchet

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func ReadLinesFromFile(filePath string) ([]string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	return lines, nil
}

func FindFirstAndLastDigit(line string) int {
	var first, last rune
	sum := 0
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
		}
		sum = calibrationNum
	}

	return sum
}
