package common

import "my-solution-leet-code/dto"

func IntArrayEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func ListNodeArrayEquals(l1 *dto.ListNode, l2 *dto.ListNode) bool {
	listValue1 := make([]int, 0)
	listValue2 := make([]int, 0)
	for l1 != nil {
		listValue1 = append(listValue1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		listValue2 = append(listValue2, l2.Val)
		l2 = l2.Next
	}

	return IntArrayEquals(listValue1, listValue2)
}
