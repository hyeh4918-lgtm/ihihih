package main

import "strings"

func makeUpperCase(words []string, count int) []string {
    if count > len(words) {
        count = len(words)
    }
    start := len(words) - count
    for i := start; i < len(words); i++ {
        words[i] = strings.ToUpper(words[i])
    }
    return words
}