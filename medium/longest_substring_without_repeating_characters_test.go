package medium

import (
	"strings"
	"testing"
)

/*
 * Given a string s, find the length of the longest substring without repeating characters.
 */
func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	// Example 1
	s := "abcabcbb"
	output := 3

	result := lengthOfLongestSubstring(s)

	if result != output {
		t.Error("Failed example 1")
	}

	// Example 2
	s = "bbbbb"
	output = 1

	result = lengthOfLongestSubstring(s)

	if result != output {
		t.Error("Failed example 2")
	}

	// Example 1
	s = "pwwkew"
	output = 3

	result = lengthOfLongestSubstring(s)

	if result != output {
		t.Error("Failed example 3")
	}
}

func lengthOfLongestSubstring(s string) int {
	arr := strings.Split(s, "")
	mapArr := make(map[string]bool)
	max := 0
	for len(arr) >= 1 {
		for _, value := range arr {
			if _, ok := mapArr[value]; !ok {
				mapArr[value] = true
				continue
			}
			if len(mapArr) > max {
				max = len(mapArr)
			}
			mapArr = make(map[string]bool)
			arr = arr[1:len(arr)]

			break
		}
	}
	return max
}
