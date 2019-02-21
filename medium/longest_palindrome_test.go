package medium

import "testing"

func TestLongestPalindrome(t *testing.T) {
	for _, unit := range []struct {
		m        string
		expected string
	}{
		{"babad", "bab" }, // or "aba"
		{"cbbd", "bb" },
		{"a", "a"},
	}{
		if actually := longestPalindrome(unit.m); actually != unit.expected{
			t.Errorf("longestPalindrome: [%v], actually: [%v]", unit, actually)
		}
		if actually := longestPalindrome2(unit.m); actually != unit.expected{
			t.Errorf("longestPalindrome2: [%v], actually: [%v]", unit, actually)
		}
		if actually := longestPalindrome3(unit.m); actually != unit.expected{
			t.Errorf("longestPalindrome3: [%v], actually: [%v]", unit, actually)
		}
	}
}