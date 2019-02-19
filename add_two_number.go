package leetcode_golang

import "math"

type ListNode struct {
     Val int
     Next *ListNode
}

// 20 ms	5 MB
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

// Best1a
// 16 ms	5 MB
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{0, nil}
	current := head

	for l1 != nil || l2 != nil {
		if l1 != nil {
			current.Val += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			current.Val += l2.Val
			l2 = l2.Next
		}

		if current.Val >= 10 {
			current.Val -= 10
			current.Next = &ListNode{1, nil}
		} else if l1 != nil || l2 != nil {
			current.Next = &ListNode{0, nil}
		}
		current = current.Next
	}
	return head
}

// deckvig
// 	20 ms	4.7 MB
func addTwoNumbers3(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := l1
	for {
		if l1 == nil {
			return l3
		}
		if l2 != nil {
			l1.Val += l2.Val
			l2 = l2.Next
		}

		if l1.Val >= 10 {
			l1.Val -= 10
			if l1.Next == nil {
				l1.Next = new(ListNode)
			}
			l1.Next.Val++
		}
		if l1.Next == nil && l2 != nil {
			l1.Next = l2
			return l3
		}

		l1 = l1.Next
	}
}