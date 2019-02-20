package hard

import (
	"math"
)


// 24 ms	5.9 MB
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var sortNums = make([]int, len(nums1)+len(nums2))
	i := 0
	j := 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			sortNums[i+j] = nums2[j]
			j++
		} else {
			sortNums[i+j] = nums1[i]
			i++
		}
	}
	if i == len(nums1) {
		for ;j <len(nums2);j++ {
			sortNums[i+j] = nums2[j]
		}
	}
	if j == len(nums2) {
		for ; i < len(nums1); i++ {
			sortNums[i+j] = nums1[i]
		}
	}
	tmp := float64(len(sortNums))/2
	if len(sortNums) & 0x1 == 0 { // len(sortNums) % 2
		return float64(sortNums[int(tmp)-1]+sortNums[int(tmp)]) / 2
	} else {
		return float64(sortNums[int(math.Floor(tmp))])
	}
}

type Record struct {
	last [2]float64
	this [2]float64  // [left 1;right 2 , value ]
	// this iRecord{ direction{ init:0, left:1, right:2 }, index:int }
	middleIndex int
	isOdd bool
}

func (r *Record) update(direction, value float64) {
	if r.this[0] != 0 { // if there has value in r.this, update last
		r.last[0] = r.this[0]
		r.last[1] = r.this[1]
	}
	r.this[0] = direction
	r.this[1] = value
}

func (r *Record)checkValues(thisIndex int) float64 {
	if thisIndex == r.middleIndex {
		return r.getAverage()
	}
	return -1
}

func (r *Record)getAverage() float64 {
	if r.isOdd {
		return r.this[1]
	}
	return (r.last[1] + r.this[1]) / 2
}

// 24 ms	5.4 MB
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	var rec Record
	rec.isOdd = true
	if (len(nums1) + len(nums2)) & 0x1 == 0 {
		rec.middleIndex = (len(nums1) + len(nums2)) / 2
		rec.isOdd = false
	} else {
		rec.middleIndex = (len(nums1) + len(nums2)) / 2
	}

	var (
		i, j = 0, 0
		thisIndex = 0
		checkValue float64 = -1
	)
	for ; i < len(nums1) && j < len(nums2) ; {
		if nums1[i] < nums2[j] {
			rec.update(1, float64(nums1[i]))
			checkValue = rec.checkValues(thisIndex)
			if checkValue != -1 {
				return checkValue
			}
			i++
		} else {
			rec.update(2, float64(nums2[j]))
			checkValue = rec.checkValues(thisIndex)
			if checkValue != -1 {
				return checkValue
			}
			j++
		}
		thisIndex++
	}
	if i == len(nums1) {
		for ;j <len(nums2);j++ {
			rec.update(2, float64(nums2[j]))
			checkValue = rec.checkValues(thisIndex)
			if checkValue != -1 {
				return checkValue
			}
			thisIndex++
		}
	}
	if j == len(nums2) {
		for ; i < len(nums1); i++ {
			rec.update(1, float64(nums1[i]))
			checkValue = rec.checkValues(thisIndex)
			if checkValue != -1 {
				return checkValue
			}
			thisIndex++
		}
	}

	return checkValue
}


