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
	input := utils.ReadFile("input.txt")

	fmt.Println("part1:", part1(input))
	fmt.Println("part2:", part2(input))
}

func part1(input string) int {
	return 0
}

func part2(input string) int {
	return 0
}

EOF

cat > "$new_day_dirname/main_test.go" <<EOF
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const INPUT = \`

\`

func TestPart1(t *testing.T) {
	result := part1(INPUT)
	expected := 0

	assert.Equal(t, expected, result)
}

func TestPart2(t *testing.T) {
	result := part2(INPUT)
	expected := 0

	assert.Equal(t, expected, result)
}

EOF

touch "$new_day_dirname/input.txt"
