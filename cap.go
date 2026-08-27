package main

import (
	"strings"
)

func capitalizeWord(word string) string {
	if len(word) == 0 {
		return word
	}
	lower := strings.ToLower(word)
	return strings.ToUpper(lower[:1]) + lower[1:]
}

func makeCapitalized(words []string, count int) []string {
	if count > len(words) {
		count = len(words)
	}
	start := len(words) - count
	for i := start; i < len(words); i++ {
		words[i] = capitalizeWord(words[i])
	}
	return words
}
