package medium

import (
	"strconv"
	"strings"
	"testing"
)

/*
 * https://leetcode.com/problems/string-to-integer-atoi/
 * Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer (similar to C/C++'s atoi function).
 */

type StringToIntegerAtoi struct {
	s      string
	output int
}

func TestStringToIntegerAtoi(t *testing.T) {
	inputs := []StringToIntegerAtoi{
		{
			s:      " +++1",
			output: 0,
		},
		{
			s:      " ++1",
			output: 0,
		},
		{
			s:      "  +  413",
			output: 0,
		},
		{
			s:      "-13+8",
			output: -13,
		},
		{
			s:      "    -88827   5655  U",
			output: -88827,
		},
		{
			s:      "42",
			output: 42,
		},
		{
			s:      "   -42",
			output: -42,
		},
		{
			s:      "4193 with words",
			output: 4193,
		},
		{
			s:      "words and 987",
			output: 0,
		},
		{
			s:      "-91283472332",
			output: -2147483648,
		},
	}

	for index, in := range inputs {
		result := myAtoi(in.s)

		if result != in.output {
			t.Errorf("Failed example %v", index+1)
		}
	}
}

func myAtoi(s string) int {
	nums := make([]string, 0)
	isMinus := false
	isPlus := false
	isLeadZero := false
	isNumber := false
	count := 0

	if strings.Index(s, "0 ") >= 0 {
		return 0
	}

	arr := strings.Split(s, "")
	if len(arr) > 0 && arr[0] == "0" {
		isLeadZero = true
	}

	for _, value := range arr {
		if value == "-" {
			count++
			if isNumber {
				break
			}
			isMinus = true
			continue
		}
		if value == "+" {
			count++
			if isNumber {
				break
			}
			isPlus = true
			continue
		}
		if value == " " {
			if isNumber || isPlus || isMinus {
				break
			}
			continue
		}
		if count >= 2 {
			break
		} else {
			count = 0
		}

		_, err := strconv.Atoi(value)

		if err == nil {
			nums = append(nums, value)
			isNumber = true
		} else {
			break
		}
	}

	if isLeadZero && (isMinus || isPlus) {
		return 0
	}

	result, _ := strconv.Atoi(strings.Join(nums, ""))

	if isMinus {
		result = 0 - result
	}

	if result > 2147483647 {
		return 2147483647
	}

	if result < -2147483648 {
		return -2147483648
	}

	return result
}
