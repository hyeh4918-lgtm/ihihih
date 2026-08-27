package main

import (
	"strings"
)

// FixVowel checks if word is "a" or "A" and changes it to "an" or "An" 
// if nextWord begins with a vowel (a, e, i, o, u) or 'h'.
func FixVowel(word string, nextWord string) string {
	cleanWord := strings.Trim(word, "'")
	cleanNext := strings.Trim(nextWord, "'")

	if (cleanWord != "a" && cleanWord != "A") || len(cleanNext) == 0 {
		return word
	}

	vowels := "aeiouhAEIOUH"
	firstChar := string(cleanNext[0])

	if strings.ContainsAny(firstChar, vowels) {
		if cleanWord == "a" {
			return strings.Replace(word, "a", "an", 1)
		}
		return strings.Replace(word, "A", "An", 1)
	}

	return word
}