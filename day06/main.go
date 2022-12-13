package main

import (
	"fmt"
	"math"
	"math/bits"
	"strings"
	"time"

	"github.com/Leyka/Advent-2022/helpers/utils"
)

func main() {
	input := strings.TrimSpace(utils.ReadFile("input.txt"))

	start := time.Now()
	fmt.Println("part1:", part1(input))
	fmt.Println("part2:", part2(input))
	end := time.Since(start)
	fmt.Println("time:", end) // ~ 746.792µs
	fmt.Println()

	// Xor solution
	// https://www.mattkeeter.com/blog/2022-12-10-xor
	start = time.Now()
	fmt.Println("xorPart1:", xorPart1(input))
	fmt.Println("xorPart2:", xorPart2(input))
	end = time.Since(start)
	fmt.Println("time:", end) // ~ 9.834µs
}

func part1(input string) int {
	return solve(input, 4)
}

func part2(input string) int {
	return solve(input, 14)
}

func solve(input string, buffer int) int {
	size := len(input)

	for i := 0; i < size; i++ {
		from := i
		to := math.Min(float64(size), float64(i+buffer))
		chars := input[from:int(to)]

		if utils.HasUniqueChars(chars) {
			return i + len(chars)
		}
	}

	return 0
}

func xorPart1(input string) int {
	return xor(input, 4)
}

func xorPart2(input string) int {
	return xor(input, 14)
}

func xor(input string, size int) int {
	var set uint32 = 0
	for i := range input {
		// Turn on bits as they enter the window
		set ^= 1 << (input[i] - 'a') // ex. a = 0b1, b = 0b10, c = 0b100, d = 0b1000 etc.
		// fmt.Printf("%d: %b\n", set, set)

		// Turn off bits as they leave the window
		if i >= size {
			set ^= 1 << (input[i-size] - 'a')
		}

		// Check the current window and see if we're done
		if bits.OnesCount32(set) == size {
			return i + 1
		}
	}

	return 0
}
