package main

import (
	"fmt"

	"github.com/Leyka/Advent-2022/helpers/utils"
)

func main() {

	lines := utils.ReadFileLines("input.txt")
	cleanLines := utils.RemoveEmptyStringFromArray(lines)

	fmt.Println("part1:", part1(cleanLines))
	fmt.Println("part2:", part2(cleanLines))
}

func part1(lines []string) int {
	return 0
}

func part2(lines []string) int {
	return 0
}
