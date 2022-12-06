package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/Leyka/Advent-2022/helpers/utils"
)

func main() {
	input := strings.TrimSpace(utils.ReadFile("input.txt"))

	fmt.Println("part1:", part1(input))
	fmt.Println("part2:", part2(input))
}

func part1(input string) int {
	return solve(input, 4)
}

func part2(input string) int {
	return solve(input, 14)
}

func solve(input string, buffer int) int {
	size := len(input)

	for i := 0; i < size; i++ {
		from := i
		to := math.Min(float64(size), float64(i+buffer))
		chars := input[from:int(to)]

		if utils.HasUniqueChars(chars) {
			return i + len(chars)
		}
	}

	return 0
}
