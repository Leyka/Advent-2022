package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Leyka/Advent-2022/helpers/data-structures/stack"
	"github.com/Leyka/Advent-2022/helpers/utils"
)

type Move struct {
	quantity int
	from     int
	to       int
}

type MoveFn func(Move, *[]*stack.Stack[string])

func main() {
	raw := utils.ReadFile("input.txt")

	fmt.Println("part1:", part1(raw))
	fmt.Println("part2:", part2(raw))
}

func part1(lines string) string {
	moveFn := func(move Move, cratesArray *[]*stack.Stack[string]) {
		from := move.from - 1
		to := move.to - 1

		for i := 0; i < move.quantity; i++ {
			letter, _ := (*cratesArray)[from].Pop()
			(*cratesArray)[to].Push(letter)
		}
	}

	return solve(lines, moveFn)
}

func part2(lines string) string {
	moveFn := func(move Move, cratesArray *[]*stack.Stack[string]) {
		from := move.from - 1
		to := move.to - 1
		var lettersCombo []string

		for i := 0; i < move.quantity; i++ {
			letter, _ := (*cratesArray)[from].Pop()
			lettersCombo = append(lettersCombo, letter)
		}

		for i := len(lettersCombo) - 1; i >= 0; i-- {
			letter := lettersCombo[i]
			(*cratesArray)[to].Push(letter)
		}
	}

	return solve(lines, moveFn)
}

func solve(lines string, fn MoveFn) string {
	rawSplit := strings.Split(lines, "\n\n")
	crates := strings.Split(rawSplit[0], "\n")
	cratesArray := mapCrates(crates)
	moves := strings.Split(rawSplit[1], "\n")

	for _, move := range moves {
		if move == "" {
			continue
		}

		move := parseMove(move)
		fn(move, &cratesArray)
	}

	answer := ""
	for _, crate := range cratesArray {
		letter, _ := crate.Pop()
		answer += letter
	}
	return answer
}

func mapCrates(crateLines []string) []*stack.Stack[string] {
	lastLine := crateLines[len(crateLines)-1]
	nbCrates, _ := strconv.Atoi(lastLine[len(lastLine)-1:])
	cratesArray := make([]*stack.Stack[string], nbCrates)

	// Start from the bottom
	for row := len(crateLines) - 2; row >= 0; row-- {
		line := crateLines[row]

		for crateIdx := 0; crateIdx < nbCrates; crateIdx++ {
			if cratesArray[crateIdx] == nil {
				cratesArray[crateIdx] = stack.New[string]()
			}

			letterIndex := crateIdx*4 + 1
			if letterIndex < len(line) && line[letterIndex] != ' ' {
				letter := string(line[letterIndex])
				cratesArray[crateIdx].Push(letter)
			}
		}
	}

	return cratesArray
}

func parseMove(movesLines string) Move {
	r := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	matches := r.FindStringSubmatch(movesLines)

	quantity, _ := strconv.Atoi(matches[1])
	from, _ := strconv.Atoi(matches[2])
	to, _ := strconv.Atoi(matches[3])

	return Move{quantity, from, to}
}
