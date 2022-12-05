package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Leyka/Advent-2022/helpers/utils"
)

type Range struct {
	min int
	max int
}

type ContainsFn func(Range, Range) bool

func main() {
	lines := utils.ReadFileLines("input.txt")
	cleanLines := utils.RemoveEmptyStringFromArray(lines)

	fmt.Println("part1:", part1(cleanLines))
	fmt.Println("part2:", part2(cleanLines))
}

func part1(lines []string) int {
	return solve(lines, containsFully)
}

func part2(lines []string) int {
	return solve(lines, containsPartially)
}

func solve(lines []string, containsFn ContainsFn) int {
	count := 0
	for _, line := range lines {
		ranges := parseRanges(line)
		if containsFn(ranges[0], ranges[1]) {
			count++
		}
	}

	return count
}

func parseRanges(line string) []Range {
	ranges := []Range{}
	for _, rangeStr := range strings.Split(line, ",") {
		minMax := strings.Split(rangeStr, "-")
		min, _ := strconv.Atoi(minMax[0])
		max, _ := strconv.Atoi(minMax[1])

		ranges = append(ranges, Range{min, max})
	}

	return ranges
}

func containsFully(firstRange Range, secondRange Range) bool {
	firstContainsSecond := firstRange.min <= secondRange.min && firstRange.max >= secondRange.max
	secondContainsFirst := secondRange.min <= firstRange.min && secondRange.max >= firstRange.max

	return firstContainsSecond || secondContainsFirst
}

func containsPartially(firstRange Range, secondRange Range) bool {
	partialFirst := firstRange.min <= secondRange.min && firstRange.max >= secondRange.min ||
		firstRange.min <= secondRange.max && firstRange.max >= secondRange.max
	partialSecond := secondRange.min <= firstRange.min && secondRange.max >= firstRange.min ||
		secondRange.min <= firstRange.max && secondRange.max >= firstRange.max

	return partialFirst || partialSecond
}
