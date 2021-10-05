package easy

import (
	"testing"
)

/*
 * https://leetcode.com/problems/remove-duplicates-from-sorted-array/
 * Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same.
 * Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.
 * Return k after placing the final result in the first k slots of nums.
 * Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.
 */

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	// Example 1
	nums := []int{1, 1, 2}
	output := 2

	result := removeDuplicates(nums)

	if result != output {
		t.Error("Failed example 1")
	}

	// Example 2
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	output = 5

	result = removeDuplicates(nums)

	if result != output {
		t.Error("Failed example 2")
	}
}

func removeDuplicates(nums []int) int {
	unique := make(map[int]bool)
	arr := make([]int, 0)
	for _, value := range nums {
		if _, ok := unique[value]; !ok {
			unique[value] = true
			arr = append(arr, value)
		}
	}

	for k, v := range arr {
		nums[k] = v
	}
	return len(arr)
}
