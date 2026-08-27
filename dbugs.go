package main

import (
	"regexp"
)

var tagRe = regexp.MustCompile(`\(\s*(up|low|cap|hex|bin)\s*(?:,\s*(\d+)\s*)?\)`)

// Rebuilds each tag into a canonical tight form (e.g. "( cap , 3 )" ->
// "(cap, 3)") so strings.Fields always splits it the same way downstream,
// no matter how it was originally spaced.
func TokenizeInput(text string) string {
	return tagRe.ReplaceAllStringFunc(text, func(match string) string {
		sub := tagRe.FindStringSubmatch(match)
		tag, num := sub[1], sub[2]
		if num != "" {
			return " (" + tag + ", " + num + ") "
		}
		return " (" + tag + ") "
	})
}

var validTagRe = regexp.MustCompile(`^\(\s*(up|low|cap)\s*(,\s*\d+\s*)?\)$|^\(\s*(hex|bin)\s*\)$`)

// Broader than validTagRe on purpose, so near-misses like (bex), (UP),
// (up, 2 6), or spaced-out variants like ( CAP ) get caught as malformed
// instead of passing through as text.
var candidateTagRe = regexp.MustCompile(`\(\s*[a-zA-Z]{2,3}\s*(?:,[^()]*)?\)`)

func FindMalformedTag(line string) string {
	for _, candidate := range candidateTagRe.FindAllString(line, -1) {
		if !validTagRe.MatchString(candidate) {
			return candidate
		}
	}
	return ""
}