package main

import (
	"fmt"
	"log"

	"github.com/ericmariot/goAdventTrebuchet/trebuchet"
)

func main() {
	count, err := trebuchet.CalibrationValueFromFile("trebuchet/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total count:", count)
}
