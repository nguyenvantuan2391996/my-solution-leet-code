package medium

import (
	"my-solution-leet-code/common"
	"my-solution-leet-code/dto"
	"testing"
)

/*
 * https://leetcode.com/problems/add-two-numbers/
 * You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
 * You may assume the two numbers do not contain any leading zero, except the number 0 itself.
 */

func TestAddTwoNumbers(t *testing.T) {
	// Example 1
	l1 := &dto.ListNode{
		Val: 2,
		Next: &dto.ListNode{
			Val: 4,
			Next: &dto.ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 := &dto.ListNode{
		Val: 5,
		Next: &dto.ListNode{
			Val: 6,
			Next: &dto.ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	output := &dto.ListNode{
		Val: 7,
		Next: &dto.ListNode{
			Val: 0,
			Next: &dto.ListNode{
				Val:  8,
				Next: nil,
			},
		},
	}

	result := addTwoNumbers(l1, l2)

	if !common.ListNodeArrayEquals(result, output) {
		t.Error("Failed example 1")
	}

	// Example 2
	l1 = &dto.ListNode{
		Val:  0,
		Next: nil,
	}

	// Example 2
	l2 = &dto.ListNode{
		Val:  0,
		Next: nil,
	}
	output = &dto.ListNode{
		Val:  0,
		Next: nil,
	}

	result = addTwoNumbers(l1, l2)

	if !common.ListNodeArrayEquals(result, output) {
		t.Error("Failed example 2")
	}

	// Example 3
	l1 = &dto.ListNode{
		Val: 9,
		Next: &dto.ListNode{
			Val: 9,
			Next: &dto.ListNode{
				Val: 9,
				Next: &dto.ListNode{
					Val: 9,
					Next: &dto.ListNode{
						Val: 9,
						Next: &dto.ListNode{
							Val: 9,
							Next: &dto.ListNode{
								Val:  9,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}

	l2 = &dto.ListNode{
		Val: 9,
		Next: &dto.ListNode{
			Val: 9,
			Next: &dto.ListNode{
				Val: 9,
				Next: &dto.ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}
	output = &dto.ListNode{
		Val: 8,
		Next: &dto.ListNode{
			Val: 9,
			Next: &dto.ListNode{
				Val: 9,
				Next: &dto.ListNode{
					Val: 9,
					Next: &dto.ListNode{
						Val: 0,
						Next: &dto.ListNode{
							Val: 0,
							Next: &dto.ListNode{
								Val: 0,
								Next: &dto.ListNode{
									Val:  1,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}

	result = addTwoNumbers(l1, l2)

	if !common.ListNodeArrayEquals(result, output) {
		t.Error("Failed example 3")
	}
}

func addTwoNumbers(l1 *dto.ListNode, l2 *dto.ListNode) *dto.ListNode {
	result := dto.ListNode{}
	sum, carry := 0, 0
	p := &result

	for l1 != nil || l2 != nil || carry != 0 {
		sum = carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		p.Next = &dto.ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		p = p.Next
	}

	return result.Next
}
