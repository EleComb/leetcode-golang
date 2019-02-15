package leetcode_golang

// Brute Force
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// Two-pass Hash Table
func twoSumHashTable(nums []int, target int) []int {
	var numsHashTable = make(map[int]int)
	for i, v := range nums {
		numsHashTable[v] = i
	}
	for v, i := range numsHashTable {
		need := target - v
		if k2, ok := numsHashTable[need]; ok {
			return []int{ i, k2 }
		}
	}
	return nil
}

// One-pass Hash Table
