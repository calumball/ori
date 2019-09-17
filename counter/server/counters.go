package server

import (
	"strings"
	"unicode"
)

func countLines(text string) uint64 {
	if text == "" {
		return 0
	}
	count := strings.Count(text, "\n") + 1
	return uint64(count)
}

func countWords(text string, caps bool) map[string]uint64 {
	if !caps {
		text = strings.ToLower(text)
	}

	text = strings.ReplaceAll(text, "'", "")

	isntAlphanumeric := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	words := strings.FieldsFunc(text, isntAlphanumeric)

	counts := make(map[string]uint64)
	for _, word := range words {
		counts[word]++
	}
	return counts
}
