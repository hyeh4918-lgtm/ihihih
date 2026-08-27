package main

import (
	"regexp"
)

// ensures a space exists after the closing quote.
func FormatQuotes(text string) string {
	// Remove spaces inside the quote pair: '  word  ' -> 'word'
	re := regexp.MustCompile(`'\s*(.*?)\s*'`)
	output := re.ReplaceAllString(text, "'$1'")

	// Add a space ONLY after closing quotes if attached to the next word
	spaceFixer := regexp.MustCompile(`('.*?')([a-zA-Z0-9])`)
	return spaceFixer.ReplaceAllString(output, "$1 $2")
}
