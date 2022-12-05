package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const INPUT = `    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

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
