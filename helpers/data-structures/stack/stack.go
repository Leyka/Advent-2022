package stack

import (
	"errors"
)

type Stack[T any] []T

// Stack implementation (LIFO). Not thread-safe.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) Push(element T) {
	*s = append(*s, element)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var empty T
		return empty, errors.New("empty stack")
	}

	idxLastElement := len(*s) - 1
	lastElement := (*s)[idxLastElement]
	*s = (*s)[:idxLastElement]

	return lastElement, nil
}
