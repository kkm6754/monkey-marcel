package leetcode19

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 引入dummyNode，处理head只有一个节点的情况
	dummy := &ListNode{
		Next: head,
		Val: 0,
	}
	fast, slow := dummy, dummy
	for n > 0  && fast.Next != nil{ // fast先走n步
		fast = fast.Next
		n--
	}
	for fast.Next != nil { // 一起走，直到fast到达最后
		fast, slow = fast.Next, slow.Next
	}
	slow.Next = slow.Next.Next // 删除slow的后一个

	return dummy.Next // 注意最后返回
}
