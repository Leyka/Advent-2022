package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
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

func ReadFileLines(filename string) []string {
	inputContent := ReadFile(filename)
	contentTrimmed := strings.TrimSpace(inputContent)

	return strings.Split(contentTrimmed, "\n")
}
