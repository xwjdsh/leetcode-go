package main

func twoSum_1(nums []int, target int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum_2(nums []int, target int) []int {
	numsMap := map[int]int{}
	for i, num := range nums {
		if j, ok := numsMap[target-num]; ok {
			return []int{i, j}
		}
		numsMap[num] = i
	}
	return nil
}

func twoSum(nums []int, target int) []int {
	numsMap := map[int]int{}
	for i, num := range nums {
		if j, ok := numsMap[num]; ok {
			return []int{i, j}
		}
		numsMap[target-num] = i
	}
	return nil
}
