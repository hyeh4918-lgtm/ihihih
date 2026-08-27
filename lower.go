package main

import "strings"

func makeLowerCase(words []string, count int) []string {
	if count > len(words) {
		count = len(words)
	}
	start := len(words) - count
	for i := start; i < len(words); i++ {
		words[i] = strings.ToLower(words[i])
	}
	return words
}
