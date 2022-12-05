package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	s := New[int]()
	s.Push(10)
	s.Push(20)
	s.Push(30)

	assert.Equal(t, 3, len(*s))
	assert.Equal(t, 30, (*s)[2])
}

func TestPop(t *testing.T) {
	s := New[string]()
	s.Push("hello")
	s.Push("world")

	element, err := s.Pop()
	assert.Nil(t, err)
	assert.Equal(t, "world", element)

	element, err = s.Pop()
	assert.Nil(t, err)
	assert.Equal(t, "hello", element)

	_, err = s.Pop()
	assert.NotNil(t, err)
}

func TestIsEmpty(t *testing.T) {
	s := New[int]()
	assert.True(t, s.IsEmpty())

	s.Push(1)
	assert.False(t, s.IsEmpty())

	s.Pop()
	assert.True(t, s.IsEmpty())
}
