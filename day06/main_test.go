package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 7, part1("mjqjpqmgbljsphdztnvjfqwrcgsmlb"))
	assert.Equal(t, 5, part1("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	assert.Equal(t, 6, part1("nppdvjthqldpwncqszvftbrmjlhg"))
	assert.Equal(t, 10, part1("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
	assert.Equal(t, 11, part1("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
}

func TestPart1Xor(t *testing.T) {
	assert.Equal(t, 7, xorPart1("mjqjpqmgbljsphdztnvjfqwrcgsmlb"))
	assert.Equal(t, 5, xorPart1("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	assert.Equal(t, 6, xorPart1("nppdvjthqldpwncqszvftbrmjlhg"))
	assert.Equal(t, 10, xorPart1("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
	assert.Equal(t, 11, xorPart1("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 19, part2("mjqjpqmgbljsphdztnvjfqwrcgsmlb"))
	assert.Equal(t, 23, part2("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	assert.Equal(t, 23, part2("nppdvjthqldpwncqszvftbrmjlhg"))
	assert.Equal(t, 29, part2("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
	assert.Equal(t, 26, part2("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
}

func TestPart2Xor(t *testing.T) {
	assert.Equal(t, 19, xorPart2("mjqjpqmgbljsphdztnvjfqwrcgsmlb"))
	assert.Equal(t, 23, xorPart2("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	assert.Equal(t, 23, xorPart2("nppdvjthqldpwncqszvftbrmjlhg"))
	assert.Equal(t, 29, xorPart2("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
	assert.Equal(t, 26, xorPart2("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
}
