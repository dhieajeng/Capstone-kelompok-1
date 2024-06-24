package util

import (
	"strings"
	"unicode"
)

func Abbreviate(str string) string {
	words := strings.Fields(str)
	var abbrev string

	for _, word := range words {
		for _, char := range word {
			if unicode.IsLetter(char) {
				abbrev += string(char)
				break
			}
		}
	}

	return strings.ToUpper(abbrev)
}
