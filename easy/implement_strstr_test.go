package easy

import (
	"strings"
	"testing"
)

/*
 * https://leetcode.com/problems/implement-strstr/
 * Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
 * Clarification:
 * What should we return when needle is an empty string? This is a great question to ask during an interview.
 * For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().
 */

func TestImplementStrstr(t *testing.T) {
	// Example 1
	haystack := "hello"
	needle := "ll"
	output := 2

	result := strStr(haystack, needle)

	if result != output {
		t.Error("Failed example 1")
	}

	// Example 2
	haystack = "aaaaa"
	needle = "bba"
	output = -1

	result = strStr(haystack, needle)

	if result != output {
		t.Error("Failed example 2")
	}

	// Example 3
	haystack = ""
	needle = ""
	output = 0

	result = strStr(haystack, needle)

	if result != output {
		t.Error("Failed example 3")
	}
}

func strStr(haystack string, needle string) int {
	index := strings.Index(haystack, needle)
	return index
}
