package main

import "strings"

const MaxChars = 2000

func IsASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > 127 {
			return false
		}
	}
	return true
}

func IsEmptyFile(s string) bool {
	return strings.TrimSpace(s) == ""
}

func ExceedsLimit(s string) bool {
	return len(s) > MaxChars
}