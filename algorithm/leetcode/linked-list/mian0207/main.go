package mian0207

type ListNode struct {
	Val int
	Next *ListNode
}


func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := 0, 0
	curA, curB := headA, headB
	for curA != nil {
		curA = curA.Next
		lenA++
	}
	for curB != nil {
		curB = curB.Next
		lenB++
	}

	// 排除前面的一段，后面2者长度一致
	if lenA > lenB {
		for i := 0; i< lenA - lenB; i++ {
			headA = headA.Next
		}
	}else {
		for i := 0; i < lenB-lenA; i++ {
			headB = headB.Next
		}
	}

	for headA != nil {
		if headA == headB { // 相交是节点相同，不是节点值相等
			return headA
		}
		headA, headB = headA.Next, headB.Next
	}

	return nil
}
