package leetcode206

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	dummy := &ListNode{}

	for head != nil {
		tmp := head
		head = head.Next
		tmp.Next = dummy.Next
		dummy.Next = tmp
	}

	return dummy.Next
}
