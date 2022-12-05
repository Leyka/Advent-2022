package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayContains(t *testing.T) {
	assert.False(t, ArrayContains(&[]int{}, 1))
	assert.True(t, ArrayContains(&[]int{1, 2, 3}, 2))
	assert.False(t, ArrayContains(&[]int{1, 2, 3}, 4))

	assert.True(t, ArrayContains(&[]string{"foo", "bar", "baz"}, "foo"))
	assert.False(t, ArrayContains(&[]string{"FOO", "BAR"}, "foo"))
}
