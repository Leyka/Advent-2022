package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/Leyka/Advent-2022/helpers/utils"
)

func main() {
	lines := utils.ReadFileLines("input.txt")
	linesCleaned := utils.RemoveEmptyStringFromArray(lines)

	fmt.Println("part1:", part1(linesCleaned))
	fmt.Println("part2:", part2(linesCleaned))
}

// a through z have priorities 1 through 26.
// A through Z have priorities 27 through 52.
func getLetterValue(letter string) int {
	asciiValueLowered := int(strings.ToLower(letter)[0])
	relativeValue := asciiValueLowered - 96

	if unicode.IsUpper(rune(letter[0])) {
		relativeValue += 26
	}

	return relativeValue
}

func part1(lines []string) int {
	total := 0
	for _, line := range lines {
		firstHalf := line[:len(line)/2]
		secondHalf := line[len(line)/2:]

		uniqueCommonLetters := utils.FindCommonLetters([]string{firstHalf, secondHalf}, true)
		for _, letter := range uniqueCommonLetters {
			total += getLetterValue(letter)
		}
	}

	return total
}

func part2(lines []string) int {
	total := 0
	for i := 0; i < len(lines); i += 3 {
		groups := lines[i : i+3]

		uniqueCommonLetters := utils.FindCommonLetters(groups, true)
		for _, letter := range uniqueCommonLetters {
			total += getLetterValue(letter)
		}
	}

	return total
}
