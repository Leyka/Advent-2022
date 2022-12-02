package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/Leyka/Advent-2022/utils"
)

func main() {
	input := utils.ReadFile("input.txt")

	fmt.Println("part1:", part1(input))
	fmt.Println("part2:", part2(input))
}

func sortCaloriesDesc(input string) []int {
	groupCalories := strings.Split(input, "\n\n")
	totalCaloriesByGroup := []int{}

	for _, group := range groupCalories {
		total := 0
		calories := strings.Split(group, "\n")

		for _, calorie := range calories {
			calorieInt, _ := strconv.Atoi(calorie)
			total += calorieInt
		}

		totalCaloriesByGroup = append(totalCaloriesByGroup, total)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(totalCaloriesByGroup)))
	return totalCaloriesByGroup
}

func part1(input string) int {
	totalCaloriesByGroup := sortCaloriesDesc(input)
	return totalCaloriesByGroup[0]
}

func part2(input string) int {
	totalCaloriesByGroup := sortCaloriesDesc(input)
	return totalCaloriesByGroup[0] + totalCaloriesByGroup[1] + totalCaloriesByGroup[2]
}
