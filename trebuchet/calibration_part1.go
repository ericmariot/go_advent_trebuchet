package trebuchet

func CalibrationValueFromFilePartOne(filePath string) (int, error) {
	lines, err := ReadLinesFromFile(filePath)
	if err != nil {
		return 0, err
	}

	sum := 0

	for _, line := range lines {
		calibrate := FindFirstAndLastDigit(line)
		sum += calibrate
	}

	return sum, nil
}
