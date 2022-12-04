#!/usr/bin/env bash

last_day=$(ls | grep ^day | grep -Eo '[0-9]*' | tail -1)
last_day_increment=$((last_day + 1))

new_day_padded=$(printf "%02d\n" $last_day_increment)
new_day_dirname="day$new_day_padded"

mkdir "$new_day_dirname"

cat > "$new_day_dirname/main.go" <<EOF
package main

import (
	"fmt"

	"github.com/Leyka/Advent-2022/utils"
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

EOF

cat > "$new_day_dirname/main_test.go" <<EOF
package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const INPUT = \`\`

func TestPart1(t *testing.T) {
	lines := strings.Split(INPUT, "\n")

	result := part1(lines)
	expected := 0

	assert.Equal(t, expected, result)
}

func TestPart2(t *testing.T) {
	lines := strings.Split(INPUT, "\n")

	result := part2(lines)
	expected := 0

	assert.Equal(t, expected, result)
}

EOF

touch "$new_day_dirname/input.txt"
