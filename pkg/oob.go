package oob

import (
	"strings"
)

// Oob takes an input string and replaces all vowels with the string "oob"
func Oob(input string) string {
	input = strings.ReplaceAll(input, "o", "oob") // Replace "o" first since "oob" contains "o"
	input = strings.ReplaceAll(input, "a", "oob")
	input = strings.ReplaceAll(input, "e", "oob")
	input = strings.ReplaceAll(input, "i", "oob")
	input = strings.ReplaceAll(input, "u", "oob")
	input = strings.ReplaceAll(input, "y", "oob")

	input = strings.ReplaceAll(input, "O", "OOB") // Replace "O" first since "OOB" contains "o"
	input = strings.ReplaceAll(input, "A", "OOB")
	input = strings.ReplaceAll(input, "E", "OOB")
	input = strings.ReplaceAll(input, "I", "OOB")
	input = strings.ReplaceAll(input, "U", "OOB")
	input = strings.ReplaceAll(input, "Y", "OOB")

	return input
}
