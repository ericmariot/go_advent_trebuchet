package trebuchet

import (
	"os"
	"strings"
)

// ReadLinesFromFile reads a file and returns its contents as a slice of strings.
func ReadLinesFromFile(filePath string) ([]string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	return lines, nil
}
