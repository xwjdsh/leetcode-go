package main

func singleNumber_1(nums []int) int {
	m := map[int]int{}
	for _, n := range nums {
		m[n]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}

func singleNumber(nums []int) int {
	r := 0
	for _, n := range nums {
		r ^= n
	}
	return r
}
