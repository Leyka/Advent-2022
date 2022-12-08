package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/Leyka/Advent-2022/helpers/utils"
)

func main() {
	lines := utils.ReadFileLines("input.txt")
	cleanLines := utils.RemoveEmptyStringFromArray(lines)

	fmt.Println("part1:", part1(cleanLines))
	fmt.Println("part2:", part2(cleanLines))
}

func part1(lines []string) int {
	countVisible := 0
	for i, row := range lines {
		for j, col := range row {
			isEdge := i == 0 || i == len(lines)-1 || j == 0 || j == len(row)-1
			if isEdge {
				countVisible++
				continue
			}

			currentHeight, _ := strconv.Atoi(string(col))
			allLines := get4Lines(lines, row, i, j)
			if isVisible(currentHeight, allLines) {
				countVisible++
			}
		}
	}

	return countVisible
}

func part2(lines []string) int {
	max := math.MinInt

	for i, row := range lines {
		for j, col := range row {
			currentHeight, _ := strconv.Atoi(string(col))
			allLines := get4Lines(lines, row, i, j)
			score := computeScore(currentHeight, allLines)
			if score > max {
				max = score
			}
		}
	}

	return max
}

// top, left, bottom, right
func get4Lines(lines []string, row string, i int, j int) []string {
	top := utils.StringReverse(getVerticalLine(lines[:i], j)) // relative to current position
	left := utils.StringReverse(row[:j])                      // relative to current position
	bottom := getVerticalLine(lines[i+1:], j)
	right := row[j+1:]

	return []string{top, left, bottom, right}
}

func getVerticalLine(lines []string, col int) string {
	acc := ""
	for _, line := range lines {
		acc += string(line[col])
	}
	return acc
}

func isVisible(currentHeight int, treeLines []string) bool {
	for _, treeLine := range treeLines {
		// If all tree in this line are strictly smaller than current height, then it's visible
		visible := true
		for _, tree := range treeLine {
			treeHeight, _ := strconv.Atoi(string(tree))
			if treeHeight >= currentHeight {
				visible = false
				break
			}
		}
		if visible {
			return true
		}
	}

	return false
}

func computeScore(currentHeight int, treeLines []string) int {
	score := 1

	for _, treeLine := range treeLines {
		countTree := 0
		for _, tree := range treeLine {
			countTree++
			treeHeight, _ := strconv.Atoi(string(tree))
			if treeHeight >= currentHeight {
				// Block the view
				break
			}
		}

		score *= countTree
	}

	return score
}
