// Package sum provides custom functions for working with strings and words.
package sum

import "strings"

// UseCount returns the number of times each word is used in a string.
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns the number of words in a string.
func Count(s string) int {
	x := strings.Fields(s)
	return len(x)
}
