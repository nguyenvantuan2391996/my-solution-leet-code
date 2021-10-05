package easy

import "testing"

/*
 * https://leetcode.com/problems/search-insert-position/
 * Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
 * You must write an algorithm with O(log n) runtime complexity.
 */

func TestSearchInsertPosition(t *testing.T) {
	// Example 1
	nums := []int{1, 3, 5, 6}
	target := 5
	output := 2

	result := searchInsert(nums, target)

	if result != output {
		t.Error("Failed example 1")
	}

	// Example 2
	nums = []int{1, 3, 5, 6}
	target = 2
	output = 1

	result = searchInsert(nums, target)

	if result != output {
		t.Error("Failed example 2")
	}

	// Example 3
	nums = []int{1, 3, 5, 6}
	target = 7
	output = 4

	result = searchInsert(nums, target)

	if result != output {
		t.Error("Failed example 3")
	}

	// Example 4
	nums = []int{1, 3, 5, 6}
	target = 0
	output = 0

	result = searchInsert(nums, target)

	if result != output {
		t.Error("Failed example 4")
	}

	// Example 5
	nums = []int{1}
	target = 0
	output = 0

	result = searchInsert(nums, target)

	if result != output {
		t.Error("Failed example 5")
	}
}

func searchInsert(nums []int, target int) int {
	index := 0

	if target > nums[len(nums)/2] {
		index = len(nums) / 2
	}

	for i := index; i < len(nums); i++ {
		if target <= nums[i] {
			return i
		} else if i == len(nums)-1 && target > nums[i] {
			return i + 1
		}
	}

	return 0
}
