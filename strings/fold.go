package strings

import "strings"

// HasPrefixFold is a wrapper for strings.HasPrefix but case-insensitive based on strings.EqualFold.
func HasPrefixFold(s, prefix string) bool {
	return len(s) >= len(prefix) && strings.EqualFold(s[0:len(prefix)], prefix)
}

// HasSuffixFold is a wrapper for strings.HasSuffix but case-insensitive based on strings.EqualFold.
func HasSuffixFold(s, suffix string) bool {
	return len(s) >= len(suffix) && strings.EqualFold(s[len(s)-len(suffix):], suffix)
}

// EqualFold is a wrapper for strings.EqualFold.
func EqualFold(s, t string) bool {
	return strings.EqualFold(s, t)
}
