package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const INPUT = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

func TestPart1(t *testing.T) {
	lines := strings.Split(INPUT, "\n")

	result := part1(lines)
	expected := 2

	assert.Equal(t, expected, result)
}

func TestPart2(t *testing.T) {
	lines := strings.Split(INPUT, "\n")

	result := part2(lines)
	expected := 4

	assert.Equal(t, expected, result)
}
