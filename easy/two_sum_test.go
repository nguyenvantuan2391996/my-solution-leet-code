package easy

import (
	"my-solution-leet-code/common"
	"testing"
)

/*
 * https://leetcode.com/problems/two-sum/
 * Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
 * You may assume that each input would have exactly one solution, and you may not use the same element twice.
 * You can return the answer in any order.
 */

type TwoSum struct {
	nums   []int
	target int
	output []int
}

func TestTwoSum(t *testing.T) {
	inputs := []TwoSum{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			output: []int{0, 1},
		},
		{
			nums:   []int{3, 2, 4},
			target: 6,
			output: []int{1, 2},
		},
		{
			nums:   []int{3, 3},
			target: 6,
			output: []int{0, 1},
		},
	}

	for index, in := range inputs {
		result := twoSum(in.nums, in.target)

		if !common.IntArrayEquals(result, in.output) {
			t.Errorf("Failed example %v", index+1)
		}
	}
}

func twoSum(nums []int, target int) []int {
	numbers := make(map[int]int)
	numbersSub := make(map[int]int)
	output := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		numbers[i] = nums[i]
		numbersSub[target-nums[i]] = i
	}

	for key, value := range numbers {
		if v, ok := numbersSub[value]; ok {
			if key == v && target != 0 {
				continue
			}
			if key < v {
				output = append(output, key)
				output = append(output, v)
			}
			if key > v {
				output = append(output, v)
				output = append(output, key)
			}
			break
		}
	}

	return output
}
