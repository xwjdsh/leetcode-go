package main

func missingNumber(nums []int) int {
	expect, real := 0, 0
	for i, j := range nums {
		expect += i
		real += j
	}
	return expect + len(nums) - real
}
