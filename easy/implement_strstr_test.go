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

type InputImplementStrstr struct {
	haystack string
	needle   string
	output   int
}

func TestImplementStrstr(t *testing.T) {
	inputs := []InputImplementStrstr{
		{
			haystack: "hello",
			needle:   "ll",
			output:   2,
		},
		{
			haystack: "aaaaa",
			needle:   "bba",
			output:   -1,
		},
		{
			haystack: "",
			needle:   "",
			output:   0,
		},
	}

	for index, in := range inputs {
		result := strStr(in.haystack, in.needle)

		if result != in.output {
			t.Errorf("Failed example %v", index+1)
		}
	}
}

func strStr(haystack string, needle string) int {
	index := strings.Index(haystack, needle)
	return index
}
