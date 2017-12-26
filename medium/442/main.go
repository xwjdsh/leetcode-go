package main

func findDuplicates_1(nums []int) []int {
	check := map[int]int{}
	result := []int{}
	for i := 0; i < len(nums); i++ {
		check[nums[i]]++
		if check[nums[i]] == 2 {
			result = append(result, nums[i])
		}
	}
	return result
}

func findDuplicates(nums []int) []int {
	result := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		if inx := nums[i] - 1; nums[i] == nums[inx] {
			if i != inx {
				result = append(result, nums[i])
				nums[i] = 0
			}
		} else {
			nums[i], nums[inx] = nums[inx], nums[i]
			i--
		}
	}
	return result
}
