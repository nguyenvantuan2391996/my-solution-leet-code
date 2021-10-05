package easy

import "testing"

/*
 * https://leetcode.com/problems/remove-element/
 * Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The relative order of the elements may be changed.
 * Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.
 * Return k after placing the final result in the first k slots of nums.
 * Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.
 */

type RemoveElement struct {
	nums   []int
	value  int
	output int
}

func TestRemoveElement(t *testing.T) {
	inputs := []RemoveElement{
		{
			nums:   []int{3, 2, 2, 3},
			value:  3,
			output: 2,
		},
		{
			nums:   []int{0, 1, 2, 2, 3, 0, 4, 2},
			value:  2,
			output: 5,
		},
	}

	for index, in := range inputs {
		result := removeElement(in.nums, in.value)

		if result != in.output {
			t.Errorf("Failed example %v", index+1)
		}
	}
}

func removeElement(nums []int, val int) int {
	unique := make(map[int]bool)
	unique[val] = true
	arr := make([]int, 0)
	for _, value := range nums {
		if _, ok := unique[value]; !ok {
			arr = append(arr, value)
		}
	}

	for k, v := range arr {
		nums[k] = v
	}
	return len(arr)
}
