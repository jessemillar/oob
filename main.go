package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(oob("Jesse Millar"))
}

func oob(input string) string {
	input = strings.ReplaceAll(input, "o", "oob") // Replace "o" first since "oob" contains "o"
	input = strings.ReplaceAll(input, "a", "oob")
	input = strings.ReplaceAll(input, "e", "oob")
	input = strings.ReplaceAll(input, "i", "oob")
	input = strings.ReplaceAll(input, "u", "oob")
	input = strings.ReplaceAll(input, "y", "oob")

	return input
}
