package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Leyka/Advent-2022/helpers/data-structures/node"
	"github.com/Leyka/Advent-2022/helpers/utils"
)

func main() {
	lines := utils.ReadFileLines("input.txt")
	cleanLines := utils.RemoveEmptyStringFromArray(lines)

	fmt.Println("part1:", part1(cleanLines))
	fmt.Println("part2:", part2(cleanLines))
}

func part1(lines []string) int {
	root := buildFolderStructure(lines)
	return solvePart1(root, 0)
}

func solvePart1(node *node.Node, sum int) int {
	threshold := 100_000

	if node.HasChildren() {
		value := node.ComputeValue()
		if value <= threshold {
			sum += value
		}

		for _, child := range node.Children {
			sum = solvePart1(child, sum)
		}
	}

	return sum
}

func part2(lines []string) int {
	root := buildFolderStructure(lines)
	totalUsedSpace := root.ComputeValue()
	target := 30_000_000 - (70_000_000 - totalUsedSpace)

	return solvePart2(root, target, math.MaxInt)
}

func solvePart2(node *node.Node, target int, minCandidate int) int {
	if node.HasChildren() {
		value := node.ComputeValue()
		if value >= target && value < minCandidate {
			minCandidate = value
		}

		for _, child := range node.Children {
			minCandidate = solvePart2(child, target, minCandidate)
		}
	}

	return minCandidate
}

func buildFolderStructure(lines []string) *node.Node {
	var rootNode = node.NewNode("/", 0, nil)
	var currentNode *node.Node

	for _, line := range lines {
		split := strings.Split(line, " ")

		if strings.HasPrefix(line, "$ cd") {
			arg := split[2]
			if arg == "/" {
				currentNode = rootNode
			} else if arg == ".." {
				currentNode = currentNode.Parent
			} else {
				childNode := currentNode.FindChildById(arg)
				if childNode == nil {
					childNode = node.NewNode(arg, 0, currentNode)
				}
				currentNode = childNode
			}
		} else if strings.HasPrefix(line, "$ ls") {
			continue
		} else {
			// file or directory
			size := 0
			if split[0] != "dir" {
				size, _ = strconv.Atoi(split[0])
			}
			currentNode.AddChild(node.NewNode(split[1], size, currentNode))
		}
	}

	return rootNode
}
