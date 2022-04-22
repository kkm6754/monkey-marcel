package leetcode25


type ListNode struct {
	Val int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 1 {
		return head
	}

	dummy := &ListNode {
		Next : head,
	}


	hair := dummy
	cur := head
	for cur != nil {

		left := cur
		right := cur
		for i := 1; i < k; i++{
			right = right.Next
			if right == nil{
				hair.Next = cur
				break
			}
		}
		if right == nil{ // 如果没有达到一组，就不翻转
			break
		} else { // 下一轮指向开头
			cur = right.Next
		}

		left, right = reverseList(left, right)
		hair.Next = left
		right.Next = cur
		hair = right
	}

	return dummy.Next
}

// 翻转
func reverseList(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return tail, head
}