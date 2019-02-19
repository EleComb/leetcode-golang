package leetcode_golang

// 4 ms	2.9 MB
func lengthOfLongestSubstring(s string) int {
	var maxSubstring = make([]int32, len(s))
	var startIndex ,endIndex, maxLength = 0, 0, 0
	for i, v := range s {
		endIndex = i+1
		maxSubstring[i] = v
		for j := startIndex; j < i; j++ {
			if maxSubstring[j] == v {
				if endIndex - startIndex > maxLength {
					maxLength = endIndex - startIndex - 1
				}
				startIndex = j+1
				break // !!!
			}
		}
	}
	if endIndex - startIndex > maxLength {
		maxLength = endIndex - startIndex
	}
	return maxLength
}

// garcez
func lengthOfLongestSubstring2(s string) int {
	var longest, curStart int
	for i, c := range s {
		for j, c2 := range(s[curStart:i]) {
			if c == c2 {
				if i - curStart > longest {
					longest = i - curStart
				}
				curStart += j + 1
				break
			}
		}
	}
	if len(s) - curStart > longest {
		longest = len(s) - curStart
	}
	return longest
}