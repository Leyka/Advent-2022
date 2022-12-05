package main

import (
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
	result := part1(INPUT)
	expected := "CMZ"

	assert.Equal(t, expected, result)
}

func TestPart2(t *testing.T) {
	result := part2(INPUT)
	expected := "MCD"

	assert.Equal(t, expected, result)
}
