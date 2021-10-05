package easy

import (
	"my-solution-leet-code/common"
	"my-solution-leet-code/dto"
	"testing"
)

/*
 * https://leetcode.com/problems/merge-two-sorted-lists/
 * Merge two sorted linked lists and return it as a sorted list. The list should be made by splicing together the nodes of the first two lists.
 */

func TestMergeTwoSortedLists(t *testing.T) {
	// Example 1
	l1 := &dto.ListNode{
		Val: 1,
		Next: &dto.ListNode{
			Val: 2,
			Next: &dto.ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	l2 := &dto.ListNode{
		Val: 1,
		Next: &dto.ListNode{
			Val: 3,
			Next: &dto.ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	output := &dto.ListNode{
		Val: 1,
		Next: &dto.ListNode{
			Val: 1,
			Next: &dto.ListNode{
				Val: 2,
				Next: &dto.ListNode{
					Val: 3,
					Next: &dto.ListNode{
						Val: 4,
						Next: &dto.ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
		},
	}

	result := mergeTwoLists(l1, l2)

	if !common.ListNodeArrayEquals(result, output) {
		t.Error("Failed example 1")
	}

	// Example 2
	l1 = nil

	l2 = nil
	output = nil

	result = mergeTwoLists(l1, l2)

	if !common.ListNodeArrayEquals(result, output) {
		t.Error("Failed example 2")
	}

	// Example 3
	l1 = nil

	l2 = &dto.ListNode{
		Val:  0,
		Next: nil,
	}
	output = &dto.ListNode{
		Val:  0,
		Next: nil,
	}

	result = mergeTwoLists(l1, l2)

	if !common.ListNodeArrayEquals(result, output) {
		t.Error("Failed example 3")
	}
}

func mergeTwoLists(l1 *dto.ListNode, l2 *dto.ListNode) *dto.ListNode {
	result := dto.ListNode{}
	p := &result

	for l1 != nil || l2 != nil {
		if l1 != nil && (l2 == nil || l1.Val <= l2.Val) {
			p.Next = &dto.ListNode{
				Val:  l1.Val,
				Next: nil,
			}
			l1 = l1.Next
		} else {
			p.Next = &dto.ListNode{
				Val:  l2.Val,
				Next: nil,
			}
			l2 = l2.Next
		}
		p = p.Next
	}

	return result.Next
}
