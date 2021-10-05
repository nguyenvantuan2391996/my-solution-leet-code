package hard

import (
	"testing"
)

/*
 * https://leetcode.com/problems/median-of-two-sorted-arrays/
 * Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
 * The overall run time complexity should be O(log (m+n)).
 */
func TestMedianOfTwoSortedArrays(t *testing.T) {
	// Example 1
	nums1 := []int{1, 3}
	nums2 := []int{2}
	output := 2.00000

	result := findMedianSortedArrays(nums1, nums2)

	if result != output {
		t.Error("Failed example 1")
	}

	// Example 2
	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	output = 2.50000

	result = findMedianSortedArrays(nums1, nums2)

	if result != output {
		t.Error("Failed example 2")
	}

	// Example 3
	nums1 = []int{0, 0}
	nums2 = []int{0, 0}
	output = 0.00000

	result = findMedianSortedArrays(nums1, nums2)

	if result != output {
		t.Error("Failed example 3")
	}

	// Example 4
	nums1 = []int{}
	nums2 = []int{1}
	output = 1.00000

	result = findMedianSortedArrays(nums1, nums2)

	if result != output {
		t.Error("Failed example 4")
	}

	// Example 5
	nums1 = []int{2}
	nums2 = []int{}
	output = 2.00000

	result = findMedianSortedArrays(nums1, nums2)

	if result != output {
		t.Error("Failed example 5")
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sumLen := len(nums1)
	index1 := 0
	index2 := 0
	if sumLen%2 != 0 {
		index1 = (sumLen / 2)
	} else {
		index1 = (sumLen / 2)
		index2 = (sumLen / 2) - 1
	}

	temp := 0
	for i := 0; i < len(nums1); i++ {
		for j := i + 1; j < len(nums1); j++ {
			if nums1[i] > nums1[j] {
				temp = nums1[i]
				nums1[i] = nums1[j]
				nums1[j] = temp
			}
		}
		if i > index1 {
			break
		}
	}
	if sumLen%2 != 0 {
		return float64(nums1[index1])
	} else {
		return (float64(nums1[index1]+nums1[index2]) / float64(2))
	}
}
