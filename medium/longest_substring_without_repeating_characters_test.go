package medium

import (
	"strings"
	"testing"
)

/*
 * https://leetcode.com/problems/longest-substring-without-repeating-characters/
 * Given a string s, find the length of the longest substring without repeating characters.
 */

type LongestSubstringWithoutRepeatingCharacters struct {
	s      string
	output int
}

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	inputs := []LongestSubstringWithoutRepeatingCharacters{
		{
			s:      "abcabcbb",
			output: 3,
		},
		{
			s:      "bbbbb",
			output: 1,
		},
		{
			s:      "pwwkew",
			output: 3,
		},
	}

	for index, in := range inputs {
		result := lengthOfLongestSubstring(in.s)

		if result != in.output {
			t.Errorf("Failed example %v", index+1)
		}
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
