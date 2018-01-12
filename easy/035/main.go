package main

func searchInsert(nums []int, target int) int {
	lenth := len(nums)
	if lenth == 0 {
		return 0
	}

	for i := 0; i < lenth; i++ {
		if target <= nums[i] {
			return i
		}
	}
	return lenth
}
