package Search_Insert_Position

func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for end-start > 1 {
		middle := (end-start)/2 + start
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			end = middle
		} else {
			start = middle
		}
	}
	if target <= nums[start] {
		return start
	}
	if target <= nums[end] {
		return end
	}
	return end + 1
}
