package hard

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	for _, unit := range []struct {
		m        []int
		n        []int
		expected float64
	}{
		{[]int{1, 3}, []int{2}, 2.0 },
	}{
		if actually := findMedianSortedArrays(unit.m, unit.n); actually != unit.expected{
			t.Errorf("findMedianSortedArrays: [%v], actually: [%v]", unit, actually)
		}
	}
}