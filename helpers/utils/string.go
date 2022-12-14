package utils

import (
	"strings"
)

func FindCommonLetters(words []string, unique bool) []string {
	if (len(words)) <= 1 {
		return []string{}
	}

	commonLetters := []string{}

	baseWord := words[0]
	for _, letter := range baseWord {
		common := true
		for _, word := range words[1:] {
			if !strings.Contains(word, string(letter)) {
				common = false
				break
			}
		}

		if common {
			if ArrayContains(&commonLetters, string(letter)) && unique {
				continue
			}

			commonLetters = append(commonLetters, string(letter))
		}
	}

	return commonLetters
}

func RemoveEmptyStringFromArray(array []string) []string {
	if len(array) == 0 {
		return []string{}
	}

	result := []string{}
	for _, item := range array {
		if item == "" || strings.TrimSpace(item) == "" {
			continue
		}

		result = append(result, item)
	}

	return result
}

func HasUniqueChars(word string) bool {
	chars := []string{}
	for _, char := range word {
		if ArrayContains(&chars, string(char)) {
			return false
		}

		chars = append(chars, string(char))
	}

	return true
}

func StringReverse(s string) string {
	if len(s) == 0 {
		return ""
	}

	n := len(s)
	runes := []rune(s)
	for _, rune := range s {
		n--
		runes[n] = rune
	}

	return string(runes[n:])
}
