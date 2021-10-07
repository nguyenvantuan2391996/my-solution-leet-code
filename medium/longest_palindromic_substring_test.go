package medium

import (
	"strings"
	"testing"
)

/*
 * https://leetcode.com/problems/longest-palindromic-substring/
 * Given a string s, return the longest palindromic substring in s.
 */

type LongestPalindromicSubstring struct {
	s      string
	output string
}

func TestLongestPalindromicSubstring(t *testing.T) {
	inputs := []LongestPalindromicSubstring{
		{
			s:      "tattarrattat",
			output: "tattarrattat",
		},
		{
			s:      "bb",
			output: "bb",
		},
		{
			s:      "babbbad",
			output: "abbba",
		},
		{
			s:      "babad",
			output: "bab",
		},
		{
			s:      "cbbd",
			output: "bb",
		},
		{
			s:      "a",
			output: "a",
		},
		{
			s:      "ac",
			output: "a",
		},
		{
			s:      "aaaaaab",
			output: "aaaaaa",
		},
	}

	for index, in := range inputs {
		result := longestPalindrome(in.s)

		if result != in.output {
			t.Errorf("Failed example %v", index+1)
		}
	}
}

func longestPalindrome(s string) string {
	arr := strings.Split(s, "")
	arrSubEquals := make([][]string, 0)
	arrSubEqual := make([]string, 0)
	arrSubPalindromic := make([][]string, 0)

	for _, value := range arr {
		if len(arrSubEqual) == 0 || (len(arrSubEqual)-1 >= 0 && arrSubEqual[len(arrSubEqual)-1] == value) {
			arrSubEqual = append(arrSubEqual, value)
		} else {
			arrSubEquals = append(arrSubEquals, arrSubEqual)
			arrSubEqual = make([]string, 0)
			arrSubEqual = append(arrSubEqual, value)
		}
	}
	arrSubEquals = append(arrSubEquals, arrSubEqual)

	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i-j-1 >= 0 && i+j+1 < len(arr) {
				if arr[i-j-1] == arr[i+j+1] {
					arrSubPalindromic = append(arrSubPalindromic, arr[i-j-1:j+i+2])
				} else {
					break
				}
			}
		}
	}

	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i-j >= 0 && i+j+1 < len(arr) {
				if arr[i-j] == arr[i+j+1] {
					arrSubPalindromic = append(arrSubPalindromic, arr[i-j:j+i+2])
				} else {
					break
				}
			}
		}
	}

	arrSubEquals = append(arrSubEquals, arrSubPalindromic...)
	return max(arrSubEquals)
}

func max(arr [][]string) string {
	maxIndex := 0
	for index, value := range arr {
		if index != 0 && len(value) > len(arr[maxIndex]) {
			maxIndex = index
		}
	}
	return strings.Join(arr[maxIndex], "")
}
