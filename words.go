package hades

import (
	"strings"
	"unicode/utf8"
)

// TotalWords counts total words of article.
func TotalWords(s string) int {
	wordCount := 0

	plainWords := strings.Fields(s)
	for _, word := range plainWords {
		runeCount := utf8.RuneCountInString(word)
		if len(word) == runeCount {
			// english word
			wordCount++
		} else {
			// cjk
			wordCount += runeCount
		}
	}

	return wordCount
}
