package main

import (
	"fmt"
	"log"

	"github.com/ericmariot/goAdventTrebuchet/trebuchet"
)

func main() {
	// Getting the count for Part One
	countOne, err := trebuchet.CalibrationValueFromFilePartOne("trebuchet/input.txt")
	if err != nil {
		log.Fatalf("Error in Part One: %v", err)
	}
	fmt.Println("Total count from Part One:", countOne)

	// Getting the count for Part Two
	countTwo, err2 := trebuchet.CalibrationValueFromFilePartTwo("trebuchet/input.txt")
	if err2 != nil {
		log.Fatalf("Error in Part Two: %v", err2)
	}
	fmt.Println("Total count from Part Two:", countTwo)
}
