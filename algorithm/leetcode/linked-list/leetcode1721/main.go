package leetcode1721


type ListNode struct {
	Val int
	Next *ListNode
}


func swapNodes(head *ListNode, k int) *ListNode {

	l1 := head
	for i := 1; i < k; i++ {
		
		l1 = l1.Next
	}

	l2, fast := head, l1
	for fast.Next != nil {
		l2 = l2.Next
		fast = fast.Next
	}
	l1.Val, l2.Val = l2.Val, l1.Val

	return head
}
