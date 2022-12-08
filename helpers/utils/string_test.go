package utils

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestFindUniqueCommonLetters(t *testing.T) {
	unique := true
	assert.Equal(t, []string{}, FindCommonLetters([]string{"abc"}, unique))
	assert.Equal(t, []string{"c"}, FindCommonLetters([]string{"abc", "decf"}, unique))
	assert.Equal(t, []string{"E", "c"}, FindCommonLetters([]string{"Eabc", "dEcf"}, unique))
	assert.Equal(t, []string{"E"}, FindCommonLetters([]string{"Eabc", "dEEcf", "sEEEEE"}, unique))

	unique = false
	assert.Equal(t, []string{}, FindCommonLetters([]string{"abc"}, unique))
	assert.Equal(t, []string{"c", "c", "c"}, FindCommonLetters([]string{"abccc", "deccfc"}, unique))
	assert.Equal(t, []string{"E", "c"}, FindCommonLetters([]string{"Eabc", "dEcfE"}, unique))
}

func TestRemoveEmptyStringFromArray(t *testing.T) {
	assert.Equal(t, []string{}, RemoveEmptyStringFromArray([]string{}))
	assert.Equal(t, []string{}, RemoveEmptyStringFromArray([]string{"", " ", "  ", "   "}))
	assert.Equal(t, []string{"a"}, RemoveEmptyStringFromArray([]string{"a", " ", "  ", "   "}))
	assert.Equal(t, []string{"a", "b"}, RemoveEmptyStringFromArray([]string{"a", " ", "  ", "b", "   "}))
}

func TestHasUniqueChars(t *testing.T) {
	assert.True(t, HasUniqueChars("abc"))
	assert.False(t, HasUniqueChars("abcc"))
	assert.True(t, HasUniqueChars("ab"))
	assert.False(t, HasUniqueChars("AhelloA"))
	assert.True(t, HasUniqueChars(""))
}

func TestReverse(t *testing.T) {
	assert.Equal(t, "", StringReverse(""))
	assert.Equal(t, "Abc", StringReverse("cbA"))
	assert.Equal(t, "Hello~World 123", StringReverse(StringReverse("Hello~World 123")))
}
