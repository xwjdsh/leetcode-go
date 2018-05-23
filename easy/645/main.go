package main

func findErrorNums(nums []int) []int {
	numsMap := map[int]bool{}
	result := []int{}
	realSum, expectSum := 0, 0
	for i, j := range nums {
		realSum += j
		expectSum += i + 1

		numsMap[j] = true
	}
	for i, _ := range nums {
		if j := i + 1; !numsMap[j] {
			result = append(result, realSum+j-expectSum, j)
		}
	}
	return result
}
