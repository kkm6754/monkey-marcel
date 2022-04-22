package main



//给你一个链表，每k个节点一组进行翻转，请你返回翻转后的链表。
//
//k是一个正整数，它的值小于或等于链表的长度。
//
//如果节点总数不是k的整数倍，那么请将最后剩余的节点保持原有顺序。

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
		if right == nil{
			break
		} else {
			cur = right.Next
		}

		left, right = reverseList(left, right)
		hair.Next = left
		right.Next = cur
		hair = right
	}

	return dummy.Next
}


// 反转链表 有首、尾指针
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


func main(){

}
