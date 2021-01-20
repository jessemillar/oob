package oob

import (
	"strings"
)

func isLowerVowel(c rune) bool {
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	for _, value := range vowels {
		if value == c {
			return true
		}
	}
	return false
}

func isUpperVowel(c rune) bool {
	vowels := []rune{'A', 'E', 'I', 'O', 'U'}
	for _, value := range vowels {
		if value == c {
			return true
		}
	}
	return false
}

// Oob takes an input string and replaces all vowels with the string "oob"
func Oob(input string) string {
	var sb strings.Builder
	for _, value := range input {
		if isLowerVowel(value) {
			sb.WriteString("oob")
		} else if isUpperVowel(value) {
			sb.WriteString("OOB")
		} else {
			sb.WriteRune(value)
		}
	}
	return sb.String()
}
