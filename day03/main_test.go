package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const INPUT = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

func TestPart1(t *testing.T) {
	lines := strings.Split(INPUT, "\n")
	result := part1(lines)
	expected := 157

	assert.Equal(t, expected, result)
}

func TestPart2(t *testing.T) {
	lines := strings.Split(INPUT, "\n")
	result := part2(lines)
	expected := 70

	assert.Equal(t, expected, result)
}
