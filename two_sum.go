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

// One-pass Hash Table
func twoSumOnePass(nums []int, target int) []int {
	var numsHashTable = make(map[int]int)
	for i, v := range nums {
		if j, ok := numsHashTable[target - v]; ok {
			return []int{ j, i }
		}
		numsHashTable[v] = i
	}
	return nil
}

// One-pass Hash Table
