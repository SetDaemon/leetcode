package problems

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	h := &ListNode{}
	l, carry := h, 0
	for {
		a := 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		b := 0
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		l.Val = (carry + a + b) % 10
		carry = (carry + a + b) / 10
		if l1 == nil && l2 == nil && carry == 0 {
			break
		}
		l.Next = &ListNode{}
		l = l.Next
	}
	return h
}
