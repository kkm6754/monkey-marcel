package leetcode328

//328. 奇偶链表



type ListNode struct {
	Val int
	Next *ListNode
}


func oddEvenList(head *ListNode) *ListNode {

	if head == nil {
		return head
	}

	odd := head
	even := head.Next
	evenHead := even

	for odd.Next != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead

	return head
}
