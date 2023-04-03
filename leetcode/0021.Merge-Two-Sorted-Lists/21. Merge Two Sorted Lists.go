package leetcode

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// ListNode define
type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}

func myMergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var prefixNode = &ListNode{}
	var result = prefixNode
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prefixNode.Next = l1
			//重新给l1赋值
			l1 = l1.Next
		} else {
			prefixNode.Next = l2
			//重新给l2赋值
			l2 = l2.Next
		}

		prefixNode = prefixNode.Next
	}

	if l1 != nil {
		prefixNode.Next = l1
	}

	if l2 != nil {
		prefixNode.Next = l2
	}

	return result.Next
}
