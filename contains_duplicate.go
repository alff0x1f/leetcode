package leetcode

func containsDuplicate(nums []int) bool {
	// Given an integer array nums, return true if any value appears at least twice
	// in the array, and return false if every element is distinct.
	numsSet := make(map[int]bool)
	for _, num := range nums {
		_, isExist := numsSet[num]
		if isExist {
			return true
		}
		numsSet[num] = true
	}
	return false
}
