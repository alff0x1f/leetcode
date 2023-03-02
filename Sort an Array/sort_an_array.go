package Sort_an_Array

import "math/rand"

func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	right, left := len(nums)-1, 0
	pivot := rand.Int() % len(nums)
	nums[right], nums[pivot] = nums[pivot], nums[right]
	for i := range nums {
		if nums[i] < nums[right] {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}
	nums[left], nums[right] = nums[right], nums[left]
	sortArray(nums[:left])
	sortArray(nums[left+1:])
	return nums
}
