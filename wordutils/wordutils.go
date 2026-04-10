package wordutils

import "strings"

func IsSubWord(short, long string, continuous bool) bool {
	if continuous {
		return strings.Contains(long, short)
	}
	longChars := []rune(long)
	shortChars := []rune(short)
	shortLen := len(shortChars)
	if shortLen == 0 {
		return true
	}
	shortIndex := 0
	for _, char := range longChars {
		if char == shortChars[shortIndex] {
			shortIndex++
		}
		if shortIndex == shortLen {
			return true
		}
	}
	return false
}
