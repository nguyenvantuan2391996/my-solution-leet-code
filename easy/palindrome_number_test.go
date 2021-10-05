package easy

import "testing"

/*
 * https://leetcode.com/problems/palindrome-number/
 * Given an integer x, return true if x is palindrome integer.
 * An integer is a palindrome when it reads the same backward as forward. For example, 121 is palindrome while 123 is not.
 */

type PalindromeNumber struct {
	x      int
	output bool
}

func TestPalindromeNumber(t *testing.T) {
	inputs := []PalindromeNumber{
		{
			x:      121,
			output: true,
		},
		{
			x:      -121,
			output: false,
		},
		{
			x:      10,
			output: false,
		},
		{
			x:      -101,
			output: false,
		},
	}

	for index, in := range inputs {
		result := isPalindrome(in.x)

		if result != in.output {
			t.Errorf("Failed example %v", index+1)
		}
	}
}

func isPalindrome(x int) bool {
	temp := x
	if x < 0 {
		return false
	} else {
		rev := 0
		for x != 0 {
			rev = rev*10 + x%10
			x = x / 10
		}

		if rev != temp {
			return false
		}

		return true
	}
}
