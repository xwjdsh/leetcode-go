package main

// 34.58%
func hammingDistance_1(x int, y int) int {
	count := 0
	for ; x > 0 || y > 0; x, y = x>>1, y>>1 {
		if x&1 != y&1 {
			count++
		}
	}
	return count
}

// 34.58%
func hammingDistance(x int, y int) int {
	count := 0
	x ^= y
	for ; x > 0; x >>= 1 {
		count += x & 1
	}
	return count
}
