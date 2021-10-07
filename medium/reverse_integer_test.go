package medium

import (
	"testing"
)

/*
 * https://leetcode.com/problems/reverse-integer/
 * Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
 * Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
 */

type ReverseInteger struct {
	x      int
	output int
}

func TestReverseInteger(t *testing.T) {
	inputs := []ReverseInteger{
		{
			x:      123,
			output: 321,
		},
		{
			x:      -123,
			output: -321,
		},
		{
			x:      120,
			output: 21,
		},
		{
			x:      0,
			output: 0,
		},
	}

	for index, in := range inputs {
		result := reverse(in.x)

		if result != in.output {
			t.Errorf("Failed example %v", index+1)
		}
	}
}

func reverse(x int) int {
	rev := 0
	for x != 0 {
		rev = rev*10 + x%10
		x = x / 10
	}
	if rev > 2147483647 || rev < -2147483648 {
		return 0
	}
	return rev
}
