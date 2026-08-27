package main

import (
	"regexp"
)

// FormatPunctuation attaches punctuation/groups to previous words and spaces next ones
func FormatPunctuation(text string) string {
	// Remove space before punctuation marks or grouped punctuation
	reBefore := regexp.MustCompile(`\s+([,\.!\?:;]+)`)
	text = reBefore.ReplaceAllString(text, "$1")

	// Ensure a single space after punctuation if followed by a letter/digit
	reAfter := regexp.MustCompile(`([,\.!\?:;]+)([a-zA-Z0-9])`)
	text = reAfter.ReplaceAllString(text, "$1 $2")

	return text
}
