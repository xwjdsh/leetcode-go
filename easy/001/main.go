package main

import "log"

// 76ms, 21.25%
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

// 9ms, 61.56%
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

// 9ms, 61.56%
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

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	log.Printf("twoSum_1: nums:%v, target:%d, result:%v\n", nums, target, twoSum_1(nums, target))
	log.Printf("twoSum_2: nums:%v, target:%d, result:%v\n", nums, target, twoSum_2(nums, target))
	log.Printf("twoSum: nums:%v, target:%d, result:%v\n", nums, target, twoSum(nums, target))
}
