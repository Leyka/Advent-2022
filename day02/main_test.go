package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const INPUT = `
A Y
B X
C Z
`

func TestPart1(t *testing.T) {
	result := part1(INPUT)
	expected := 15

	assert.Equal(t, expected, result)
}

func TestPart2(t *testing.T) {
	result := part2(INPUT)
	expected := 12

	assert.Equal(t, expected, result)
}
