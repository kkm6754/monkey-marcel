package main



// 双指针（快慢指针）

// 链表中间值
// 移除链表中的第N个节点
// 环形链表
// 反转链表



// 给定一个链表，判断链表中是否有环。
// 快慢指针的经典题。 快指针每次走两步，慢指针一次走一步。 在慢指针进入环之后，快慢指针之间的距离每次缩小1，所以最终能相遇。


type ListNode struct {
	Val int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	fast, slow := head.Next, head
	for fast != slow {
		if fast == nil || fast.Next == nil { // 环不可能指向nil
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next

	}

	return true
}