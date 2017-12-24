package main

func selfDividingNumbers(left int, right int) []int {
	result := []int{}
	for x := left; x <= right; {
		for i := x; i != 0; i /= 10 {
			if j := i % 10; j == 0 || x%j != 0 {
				goto L1
			}
		}
		result = append(result, x)
	L1:
		x++
	}
	return result
}
