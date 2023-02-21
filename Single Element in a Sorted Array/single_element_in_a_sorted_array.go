package Single_Element_in_a_Sorted_Array

func singleNonDuplicate(nums []int) int {
	// You are given a sorted array consisting of only integers where every element
	// appears exactly twice, except for one element which appears exactly once.
	//
	// Return the single element that appears only once.
	//
	// Your solution must run in O(log n) time and O(1) space.
	start := 0
	end := len(nums) - 1
	for end-start > 0 {
		middle := (end-start)/2 + start
		middle = middle + middle%2 // make even
		if nums[middle] == nums[middle-1] {
			end = middle - 2
		} else {
			start = middle
		}
	}
	return nums[start]
}
