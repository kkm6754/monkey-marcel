package main



// 给定一个排序链表，删除所有重复的元素，只留下原链表中没有重复的元素。
// 输入：
// linked list = 1->2->3->3->4->4->5->null
// 输出：
// 1->2->5->null


type ListNode struct {
     Val int
     Next *ListNode
}

// 1->2->3->3->4->4->5->null
// 0=> 1->2->3->3->4->4->5->null


func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyNode := ListNode{
		Val: 0,
		Next: head,
	}
	head = &dummyNode

	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			val := head.Next.Val
			for head.Next != nil && head.Next.Val == val { // 找到相同节点，删除节点
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}

	return dummyNode.Next
}