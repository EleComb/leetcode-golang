package hard

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	for _, unit := range []struct {
		m        []int
		n        []int
		expected float64
	}{
		{[]int{1, 3}, []int{2}, 2.0 },
		{[]int{1, 2}, []int{3, 4}, 2.5 },
		{[]int{3}, []int{-2, -1}, -1.0 },
	}{
		if actually := findMedianSortedArrays(unit.m, unit.n); actually != unit.expected{
			t.Errorf("findMedianSortedArrays: [%v], actually: [%v]", unit, actually)
		}
		if actually := findMedianSortedArrays2(unit.m, unit.n); actually != unit.expected{
			t.Errorf("findMedianSortedArrays2: [%v], actually: [%v]", unit, actually)
		}
	}
}