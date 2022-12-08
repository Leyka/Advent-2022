package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const INPUT = `30373
25512
65332
33549
35390`

func TestPart1(t *testing.T) {
	lines := strings.Split(INPUT, "\n")

	result := part1(lines)
	expected := 21

	assert.Equal(t, expected, result)
}

func TestPart2(t *testing.T) {
	lines := strings.Split(INPUT, "\n")

	result := part2(lines)
	expected := 8

	assert.Equal(t, expected, result)
}
