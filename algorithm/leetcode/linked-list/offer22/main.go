package offer22

type ListNode struct {
    Val int
    Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast, slow := head, head

	for i := 1; i < k; i++ {
		if fast.Next != nil {
			fast = fast.Next
		}else {
			return nil
		}
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}