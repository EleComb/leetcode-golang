package medium

import "testing"

func addTowNumberEqual(a, b *ListNode) bool {
	var p1 = a
	var p2 = b
	for p1 != nil || p2 != nil {
		// in [p1 != nil || p2 != nil] state
		// filter eg: p1 == nil && p2 != nil
		if p1 == nil || p2 == nil {
			return false
		}
		// and now p1 != nil && p2 != nil
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

func TestAddTwoNumber(t *testing.T) {
	for _, unit := range []struct {
		m        *ListNode
		n        *ListNode
		expected *ListNode
	}{
		{
			&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
			&ListNode{5, &ListNode{6, &ListNode{4, nil}}},
			&ListNode{7, &ListNode{0, &ListNode{8, nil}}},
		},
		{
			&ListNode{9, &ListNode{9, nil}},
			&ListNode{1, nil},
			&ListNode{0, &ListNode{0, &ListNode{1, nil}}},
		},
		{
			&ListNode{},
			&ListNode{0, &ListNode{1, nil}},
			&ListNode{0, &ListNode{1, nil}},
		},
		{
			&ListNode{0, &ListNode{1, nil}},
			&ListNode{0, &ListNode{1, &ListNode{2, nil}}},
			&ListNode{0, &ListNode{2, &ListNode{2, nil}}},
		},
	}{
		if actually := addTwoNumbers(unit.m, unit.n); !addTowNumberEqual(actually, unit.expected){
			t.Errorf("addTwoNumbers: [%v], actually: [%v]", unit, actually)
		}
		if actually := addTwoNumbers2(unit.m, unit.n); !addTowNumberEqual(actually, unit.expected){
			t.Errorf("addTwoNumbers: [%v], actually: [%v]", unit, actually)
		}
		if actually := addTwoNumbers3(unit.m, unit.n); !addTowNumberEqual(actually, unit.expected){
			t.Errorf("addTwoNumbers: [%v], actually: [%v]", unit, actually)
		}
	}
}
