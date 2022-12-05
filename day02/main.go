package main

import (
	"fmt"
	"strings"

	"github.com/Leyka/Advent-2022/helpers/utils"
)

const LOST = 0
const DRAW = 3
const WIN = 6

var shapePoints = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func main() {
	input := utils.ReadFile("input.txt")

	fmt.Println("part1:", part1(input))
	fmt.Println("part2:", part2(input))
}

func part1(input string) int {
	var rules = map[string]int{
		"A X": DRAW,
		"A Y": WIN,
		"A Z": LOST,
		"B X": LOST,
		"B Y": DRAW,
		"B Z": WIN,
		"C X": WIN,
		"C Y": LOST,
		"C Z": DRAW,
	}

	total := 0
	for _, turn := range strings.Split(input, "\n") {
		if turn == "" {
			continue
		}

		playedShape := turn[len(turn)-1:] // last char
		total += rules[turn] + shapePoints[playedShape]
	}

	return total
}

func part2(input string) int {
	var rules = map[int]map[string]string{
		WIN:  {"A": "Y", "B": "Z", "C": "X"},
		DRAW: {"A": "X", "B": "Y", "C": "Z"},
		LOST: {"A": "Z", "B": "X", "C": "Y"},
	}
	var strategies = map[string]int{
		"X": LOST,
		"Y": DRAW,
		"Z": WIN,
	}

	total := 0
	for _, turn := range strings.Split(input, "\n") {
		if turn == "" {
			continue
		}

		elfPlay := turn[0:1]
		strategy := strategies[turn[len(turn)-1:]]
		shapeToPlay := rules[strategy][elfPlay]

		total += strategy + shapePoints[shapeToPlay]
	}

	return total
}
