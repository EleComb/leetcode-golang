package leetcode_golang

import "math"

type ListNode struct {
     Val int
     Next *ListNode
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var p1 = l1
	var p2 = l2
	var ret = &ListNode{0, nil}
	var curr = ret
	var carry = 0

	for p1 != nil || p2 != nil {
		x, y := 0, 0
		if p1 != nil {
			x = p1.Val
		}
		if p2 != nil {
			y = p2.Val
		}
		sum := x + y + carry
		carry = sum/10
		curr.Next = &ListNode{int(math.Mod(float64(sum), 10)), nil}
		curr = curr.Next
		if p1 != nil {
			p1 = p1.Next
		}
		if p2 != nil {
			p2 = p2.Next
		}
	}
	if carry > 0 {
		curr.Next = &ListNode{carry, nil}
	}

	return ret.Next
}