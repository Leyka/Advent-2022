package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const INPUT = `
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func TestPart1(t *testing.T) {
	result := part1(INPUT)
	assert.Equal(t, 24_000, result)
}

func TestPart2(t *testing.T) {
	result := part2(INPUT)
	assert.Equal(t, 45_000, result)
}
