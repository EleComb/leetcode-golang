package medium

type zigzagRow struct {
	s     string
	index int
}

func zigzagConvert(s string, numRows int) string {
	var result = ""
	var sList = make([]zigzagRow, len(s))
	plusFlag := true
	index := 1
	for i := 0; i < len(s); i++ {
		sList[i].s = string(s[i])
		if index < numRows && index > 0 {

		}
	}
	return result
}