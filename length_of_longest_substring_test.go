package leetcode_golang

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, unit := range []struct {
		m        string
		expected int
	}{
		{"abcabcbb", 3 }, // abc
		{"bbbbb", 1 }, // b
		{"pwwkew", 3 }, // wke
	}{
		if actually := lengthOfLongestSubstring(unit.m); actually != unit.expected{
			t.Errorf("lengthOfLongestSubstring: [%v], actually: [%v]", unit, actually)
		}
	}
}