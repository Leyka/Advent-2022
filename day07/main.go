package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Leyka/Advent-2022/helpers/data-structures/node"
	"github.com/Leyka/Advent-2022/helpers/utils"
)

type ComputeFn func(*node.Node, int) int

func main() {
	lines := utils.ReadFileLines("input.txt")
	cleanLines := utils.RemoveEmptyStringFromArray(lines)

	fmt.Println("part1:", part1(cleanLines))
	fmt.Println("part2:", part2(cleanLines))
}

func part1(lines []string) int {
	root := buildFolderStructure(lines)

	compute := func(node *node.Node, sum int) int {
		value := node.ComputeValue()
		if value <= 100_000 {
			sum += value
		}

		return sum
	}

	return solve(root, 0, compute)
}

func part2(lines []string) int {
	root := buildFolderStructure(lines)
	totalUsedSpace := root.ComputeValue()
	target := 30_000_000 - (70_000_000 - totalUsedSpace)

	compute := func(node *node.Node, min int) int {
		value := node.ComputeValue()
		if value >= target && value < min {
			min = value
		}

		return min
	}

	return solve(root, math.MaxInt32, compute)
}

func solve(node *node.Node, acc int, compute ComputeFn) int {
	if node.HasChildren() {
		acc = compute(node, acc)
		for _, child := range node.Children {
			acc = solve(child, acc, compute)
		}
	}

	return acc
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
