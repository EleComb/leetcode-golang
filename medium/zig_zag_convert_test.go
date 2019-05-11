package medium

import (
	"testing"
)

func TestConvert(t *testing.T) {
	for _, unit := range []struct {
		s       string
		numRows int
		expect  string
	}{
		//{ "PAYPALISHIRING", 3, "PAHNAPLSIIGYIR" },
		//{ "PAYPALISHIRING", 4, "PINALSIGYAHRPI" },
		{"AB", 1, "AB" },
	}{
		if actually := zigzagConvert(unit.s, unit.numRows); actually != unit.expect {
			t.Errorf("zigzagConvert: [%v], actually: [%v]", unit, actually)
		}
	}
}