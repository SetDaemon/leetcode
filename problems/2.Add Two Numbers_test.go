package problems

import (
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	l1 := Ints2List([]int{2, 4, 3})
	l2 := Ints2List([]int{5, 6, 4})
	addTwoNumbers(l1, l2)

}

func Ints2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}
