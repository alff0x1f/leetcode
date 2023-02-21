package Single_Element_in_a_Sorted_Array

func singleNonDuplicate(nums []int) int {
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
