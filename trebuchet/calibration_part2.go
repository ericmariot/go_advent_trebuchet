package trebuchet

import (
	"fmt"
	"strings"
)

/*
How im going to try and solve this problem
create a map for every number that could be spelled out:
map[string]int{"zero": 0, "one": 1, ... "nine": 9}
iterate through each line from the given input
iterate through the string
collect and convert first match: number, or spelled out
save the first occurrence
find others occurrences until the end of the line
if didn't find any other occurrence first=last else create last occurrence
*/

/*
How I actually solved the problem
Refactored calibration_part_1 so there's a FindFirstAndLastDigit function
that iterate through a string(line) and find the first and last digit occurrence
Created a function to replace every number word to the first and last letter
of the number as well as the digit in the middle, making edge cases like
eightwone still be readable as e8tt2oo1e making it easy to find the digits
Send the formatted lines to the FindFirstAndLastDigit function and sum it.
*/

func CalibrationValueFromFilePartTwo(filePath string) (int, error) {
	lines, err := ReadLinesFromFile(filePath)
	if err != nil {
		return 0, err
	}
	word_to_number := map[string]string{
		"zero": "z0o", "one": "o1e", "two": "t2o", "three": "t3e", "four": "f4r",
		"five": "f5e", "six": "s6x", "seven": "s7n", "eight": "e8t", "nine": "n9e",
	}
	sum := 0

	for _, line := range lines {
		digits_line := replace_words(line, word_to_number)
		sum += FindFirstAndLastDigit(digits_line)
	}

	return sum, nil
}

func replace_words(s string, replacements map[string]string) string {
	for key, value := range replacements {
		s = strings.Replace(s, key, fmt.Sprint(value), -1)
	}
	return s
}
