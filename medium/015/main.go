package main

func threeSum(nums []int) [][]int {
	numsMap := map[int]int{}
	result := [][]int{}
	for _, i := range nums {
		numsMap[i]++
	}

	twoMap := map[int]int{}
	newNums := []int{}
	for k, v := range numsMap {
		if v > 1 {
			if k == 0 && v >= 3 {
				result = append(result, []int{0, 0, 0})
			} else if k != 0 {
				twoMap[k] = -k * 2
			}
		}
		newNums = append(newNums, k)
		numsMap[k] = len(newNums) - 1
	}
	for k, v := range twoMap {
		if _, ok := numsMap[v]; ok {
			result = append(result, []int{k, k, v})
		}
	}
	if len(newNums) > 2 {
		repeatMap := map[[3]int]bool{}
		for i := 0; i < len(newNums)-1; i++ {
			for j := i + 1; j < len(newNums); j++ {
				three := -newNums[i] - newNums[j]
				if v, ok := numsMap[three]; ok && v != i && v != j && !repeatMap[[3]int{i, j, v}] {
					result = append(result, []int{newNums[i], newNums[j], three})
					repeatMap[[3]int{i, j, v}] = true
					repeatMap[[3]int{i, v, j}] = true
					repeatMap[[3]int{j, i, v}] = true
					repeatMap[[3]int{j, v, i}] = true
					repeatMap[[3]int{v, i, j}] = true
					repeatMap[[3]int{v, j, i}] = true
				}
			}
		}
	}
	return result
}
