package medium

func reverse(s string) string {
	var reverseS = make([]byte, len(s))
	for i, t := len(s)-1, 0; i >= 0; i-- {
		reverseS[t] = s[i]
		t++
	}
	return string(reverseS)
}

// find one common substring
// go get -v github.com/jpillora/longestcommon
// string = longestcommon.Suffix([]string{s1,s2})
// https://stackoverflow.com/questions/34805488/finding-all-the-common-substrings-of-given-two-strings
func findCommonSubstring(s1, s2 string) []string {
	var dynamicTable = make([][]int, len(s1))
	var result = make([]string, 0)
	for i := 0; i < len(s1); i++ {
		dynamicTable[i] = make([]int, len(s2))
		for j := 0; j < len(s2); j++ {
			if s1[i] != s2[j] {
				continue
			}
			if i == 0 || j == 0 {
				dynamicTable[i][j] = 1
			} else {
				dynamicTable[i][j] = 1 + dynamicTable[i-1][j-1]
			}

			if dynamicTable[i][j] > 0 {
				result = append(result, s1[i-dynamicTable[i][j]+1:i+1])
			}
		}
	}
	return result
}

// = =
// Runtime: 7768 ms, faster than 5.09% of Go online submissions for Longest Palindromic Substring.
// Memory Usage: 277.3 MB, less than 5.22% of Go online submissions for Longest Palindromic Substring.
func longestPalindrome(s string) string {
	var reverseS = reverse(s)
	strs := findCommonSubstring(s, reverseS)
	var longestString string
	for _, v := range strs {
		if v == reverse(v) && len(v) > len(longestString) {
			longestString = v
		}
	}
	return longestString
}

// jiangjin863041
// Runtime: 4 ms, faster than 91.06% of Go online submissions for Longest Palindromic Substring.
// Memory Usage: 2.3 MB, less than 93.91% of Go online submissions for Longest Palindromic Substring.
func longestPalindrome2(s string) string {
	maxLen, curLen, middle, length := 0, 0, 0, len(s)
	sub := ""
	for ; middle+maxLen/2 < length; middle++ {
		palindromeLength := 1
		if s[middle-maxLen/2] == s[middle+maxLen/2] ||
			(middle+maxLen/2+1 < length && s[middle-maxLen/2] == s[middle+maxLen/2+1]) {
		} else {
			continue
		}
		for curLen = 1; middle-palindromeLength >= 0 && middle+palindromeLength < length; {
			if s[middle-palindromeLength] == s[middle+palindromeLength] {
				curLen += 2
				palindromeLength++
			} else {
				palindromeLength--
				break
			}
		}
		if curLen > maxLen {
			maxLen = curLen
			if middle-palindromeLength < 0 || middle+palindromeLength >= length {
				palindromeLength--
			}
			sub = s[middle-palindromeLength : middle+palindromeLength+1]
		}
		palindromeLength = 0
		for curLen = 0; middle-palindromeLength >= 0 && middle+palindromeLength+1 < length; {
			if s[middle-palindromeLength] == s[middle+palindromeLength+1] {
				curLen += 2
				palindromeLength ++
			} else {
				palindromeLength--
				break
			}
		}
		if curLen > maxLen {
			maxLen = curLen
			if middle-palindromeLength < 0 || middle+palindromeLength+1 >= length {
				palindromeLength--
			}
			sub = s[middle-palindromeLength : middle+palindromeLength+2]
		}
	}
	return sub
}

// sthadka
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Palindromic Substring.
// Memory Usage: 2.6 MB, less than 59.13% of Go online submissions for Longest Palindromic Substring.
type palin struct {
	left, right int
}

func longestPalindrome3(s string) string {
	max := &palin{}
	maxSize := 0
	l := len(s)

	// Edge cases
	if l == 0 || l == 1 {
		return s
	}

	// Check if whole string is palindrome
	isPalindrome := true
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			isPalindrome = false
			break
		}
	}
	if isPalindrome {
		return s
	}

	// Do it the hard way
	for i := 0; i < l; i++ {
		// Add consecutive equals
		if i < l-1 && s[i] == s[i+1] {
			p := &palin{
				left:  i,
				right: i + 1,
			}
			p, size := getPalin(s, l, p)
			if size > maxSize {
				maxSize = size
				max = p
			}
		}
		// Add surrounding equals
		if i > 0 && i < l-1 && s[i-1] == s[i+1] {
			p := &palin{
				left:  i - 1,
				right: i + 1,
			}
			p, size := getPalin(s, l, p)
			if size > maxSize {
				maxSize = size
				max = p
			}
		}
	}

	return s[max.left : max.right+1]
}

func getPalin(s string, l int, p *palin) (*palin, int) {
	j := 1
	for p.left-j >= 0 && p.right+j < l && s[p.left-j] == s[p.right+j] {
		j++
	}
	p.left = p.left - j + 1
	p.right = p.right + j - 1

	return p, p.right - p.left + 1
}
