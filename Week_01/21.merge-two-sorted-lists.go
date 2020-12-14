package leetcode

/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{}
	cur := pre
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			cur.Next = l2
			l2 = l2.Next
		} else {
			cur.Next = l1
			l1 = l1.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}

	return pre.Next
}

func BuildCyclelinkedList(items []int, cycleIndex int) *ListNode {
	var head *ListNode
	var cycleRef *ListNode
	var tail *ListNode

	for i, v := range items {
		nodeRef := &ListNode{v, nil}
		if head == nil {
			head = nodeRef
			continue
		}

		if cycleIndex == i {
			cycleRef = nodeRef
		}

		temp := head
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = nodeRef
		tail = nodeRef
	}
	tail.Next = cycleRef
	return head
}

func WalkThroughList(head *ListNode) []int {
	results := []int{}
	tmp := head
	for tmp != nil {
		results = append(results, tmp.Val)
		tmp = tmp.Next
	}
	return results
}

// @lc code=end
