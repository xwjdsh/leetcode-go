package main

func plusOne(digits []int) []int {
	result := []int{}
	n := 0
	digits[len(digits)-1] += 1
	for i := len(digits) - 1; i >= 0; i-- {
		r := digits[i] + n
		n = 0
		if r == 10 {
			r = 0
			n = 1
		}
		result = append([]int{r}, result...)
	}
	if n != 0 {
		result = append([]int{n}, result...)
	}
	return result
}
