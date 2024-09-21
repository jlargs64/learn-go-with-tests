package iteration

import "strings"

// Repeats a character or string N amount of times
func Repeat(character string, n int) string {
	var repeated string
	for i := 0; i < n; i++ {
		repeated += character
	}
	return repeated
}

// Generates a string that makes a list comma separated
func GenerateCommaList(words []string) string {
	if len(words) == 0 {
		return ""
	}
	return strings.Join(words, ",")
}
