package logs

import (
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, char := range log {
		switch {
		case char == '❗':
			return "recommendation"
		case char == '🔍':
			return "search"
		case char == '☀':
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	// strings.ReplaceAll(log, string(oldRune), string(newRune))
	/*
	// Replace replaces all occurrences of old with new, returning the modified log
	// to the caller.
	func Replace(log string, oldRune, newRune rune) string {
		runes := []rune(log)
		for index, char := range runes {
			if char == oldRune {
				runes[index] = newRune
			}
		}
		return string(runes)
	}
	*/
	var result string
	for _, char := range log {
		if char == oldRune {
			result += string(newRune)
		} else {
			result += string(char)
		}
	}
	return result
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
