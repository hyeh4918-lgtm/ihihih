package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ProcessText(text string) string {
	if IsEmptyFile(text) {
		return "Error: empty file"
	}
	if !IsASCII(text) {
		return "Error: file contains non-ASCII characters"
	}
	if ExceedsLimit(text) {
		return fmt.Sprintf("Error: file exceeds %d characters", MaxChars)
	}

	lines := strings.Split(text, "\n")

	for i, line := range lines {
		processed, errMsg := processLine(line)
		if errMsg != "" {
			return errMsg
		}
		lines[i] = processed
	}

	return strings.Join(lines, "\n")
}

func processLine(line string) (string, string) {
	if strings.TrimSpace(line) == "" {
		return line, ""
	}

	if bad := FindMalformedTag(line); bad != "" {
		return "", "Error: malformed tag: " + bad
	}

	tokenized := TokenizeInput(line)
	rawWords := strings.Fields(tokenized)

	var words []string

	for i := 0; i < len(rawWords); i++ {
		word := rawWords[i]

		if word == "(hex)" {
			if len(words) == 0 {
				return "", "Error: no word before (hex)"
			}
			converted, ok := ConvertHex(words[len(words)-1])
			if !ok {
				return "", "Error: invalid hex value: " + words[len(words)-1]
			}
			words[len(words)-1] = converted
			continue
		}
		if word == "(bin)" {
			if len(words) == 0 {
				return "", "Error: no word before (bin)"
			}
			converted, ok := ConvertBin(words[len(words)-1])
			if !ok {
				return "", "Error: invalid binary value: " + words[len(words)-1]
			}
			words[len(words)-1] = converted
			continue
		}

		switch word {
		case "(up)":
			if len(words) == 0 {
				return "", "Error: no word before (up)"
			}
			words = makeUpperCase(words, 1)
			continue
		case "(low)":
			if len(words) == 0 {
				return "", "Error: no word before (low)"
			}
			words = makeLowerCase(words, 1)
			continue
		case "(cap)":
			if len(words) == 0 {
				return "", "Error: no word before (cap)"
			}
			words = makeCapitalized(words, 1)
			continue
		}

		if (word == "(up," || word == "(low," || word == "(cap,") && i+1 < len(rawWords) {
			tag := strings.Trim(word, "(,")
			numStr := strings.Trim(rawWords[i+1], ")")

			count, err := strconv.Atoi(numStr)
			if err != nil {
				return "", "Error: invalid count in tag: (" + tag + ", " + numStr + ")"
			}
			if count <= 0 {
				return "", fmt.Sprintf("Error: invalid usage, count must be at least 1: (%s, %d)", tag, count)
			}
			if count > len(words) {
				return "", fmt.Sprintf("Error: not enough words before (%s, %d)", tag, count)
			}

			switch tag {
			case "up":
				words = makeUpperCase(words, count)
			case "low":
				words = makeLowerCase(words, count)
			case "cap":
				words = makeCapitalized(words, count)
			}
			i++
			continue
		}

		words = append(words, word)
	}

	for i := 0; i < len(words)-1; i++ {
		words[i] = FixVowel(words[i], words[i+1])
	}

	output := strings.Join(words, " ")
	output = FormatPunctuation(output)
	output = FormatQuotes(output)

	return output, ""
}