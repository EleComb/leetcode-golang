package leetcode_golang

import (
	"testing"
)

func TestTwoSum(t *testing.T) {

	groupEqual := func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}
		for index, _ := range a {
			if a[index] != b[index] {
				return false
			}
		}
		return true
	}

	for _, unit := range []struct {
		m        []int
		n          int
		expected []int
	}{
		{[]int{2, 7, 11, 5}, 9, []int{0, 1} },
		{[]int{2, 7, 11, 5}, 7, []int{0, 3} },
		{[]int{2, 7, 11, 5}, 16, []int{2, 3} },
		{[]int{2, 7, 11, 5}, 18, []int{1, 2} },
	} {
		if actually := twoSum(unit.m, unit.n); !groupEqual(actually, unit.expected){
			t.Errorf("twoSum: [%v], actually: [%v]", unit, actually)
		}
	}
}
