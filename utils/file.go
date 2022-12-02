package utils

import (
	"fmt"
	"log"
	"os"
)

func ReadFile(filename string) string {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}

	inputContent := string(file)
	if len(inputContent) == 0 {
		log.Fatalf("Empty File %s", filename)
	}

	return inputContent
}
