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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	n1, n2, carry, current := 0, 0, 0, head
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		current = current.Next
		carry = (n1 + n2 + carry) / 10
	}
	return head.Next
}

func myAddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var nowNode = &ListNode{}
	var result = nowNode
	var tempSum int

	for l1 != nil || l2 != nil || tempSum > 0 {
		//新建一个链表
		nowNode.Next = new(ListNode)
		nowNode = nowNode.Next

		//如果l1存在
		if l1 != nil {
			tempSum = tempSum + l1.Val
			//l1重新赋值
			l1 = l1.Next
		}

		//如果l2存在
		if l2 != nil {
			tempSum = tempSum + l2.Val
			//l2重新赋值
			l2 = l2.Next
		}

		//赋值
		nowNode.Val = tempSum % 10
		tempSum = (tempSum - nowNode.Val) / 10
	}

	return result.Next
}
