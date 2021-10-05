package easy

import "testing"

/*
 * https://leetcode.com/problems/palindrome-number/
 * Given an integer x, return true if x is palindrome integer.
 * An integer is a palindrome when it reads the same backward as forward. For example, 121 is palindrome while 123 is not.
 */

func TestPalindromeNumber(t *testing.T) {
	// Example 1
	x := 121
	output := true

	result := isPalindrome(x)

	if result != output {
		t.Error("Failed example 1")
	}

	// Example 2
	x = -121
	output = false

	result = isPalindrome(x)

	if result != output {
		t.Error("Failed example 2")
	}

	// Example 3
	x = 10
	output = false

	result = isPalindrome(x)

	if result != output {
		t.Error("Failed example 3")
	}

	// Example 4
	x = -101
	output = false

	result = isPalindrome(x)

	if result != output {
		t.Error("Failed example 4")
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
