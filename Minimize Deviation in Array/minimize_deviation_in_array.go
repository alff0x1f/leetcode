package Minimize_Deviation_in_Array

import (
	"sort"
)

func sortLast(nums []int) {
	// nums must be sorted slice, except last element
	if len(nums) < 2 {
		return
	}
	for i := len(nums) - 1; i > 0 && nums[i] < nums[i-1]; i-- {
		nums[i], nums[i-1] = nums[i-1], nums[i]
	}
}

func minimumDeviation(nums []int) int {
	numsSet := make(map[int]bool)
	for _, e := range nums {
		if e%2 != 0 {
			numsSet[e*2] = true
		} else {
			numsSet[e] = true
		}
	}

	evenNums := make([]int, len(numsSet))
	pos := 0

	minNumber := nums[0]
	if minNumber%2 != 0 {
		minNumber *= 2
	}

	for num := range numsSet {
		evenNums[pos] = num
		pos++
		if num < minNumber {
			minNumber = num
		}
	}
	for i := range evenNums {
		// optimization, not need
		for evenNums[i] > 2*minNumber && evenNums[i]%2 == 0 {
			evenNums[i] = evenNums[i] / 2
		}
	}
	nums = evenNums

	sort.Ints(nums)

	deviation := nums[len(nums)-1] - nums[0]
	// decrease deviation by divide biggest values
	for nums[len(nums)-1]%2 == 0 {
		nums[len(nums)-1] = nums[len(nums)-1] / 2
		sortLast(nums)
		if nums[len(nums)-1]-nums[0] < deviation {
			deviation = nums[len(nums)-1] - nums[0]
		}
	}
	return deviation
}
