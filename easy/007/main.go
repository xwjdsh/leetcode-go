package main

import (
	"math"
)

func reverse(x int) int {
	result := 0
	isMinus := x < 0
	if isMinus {
		x = 0 - x
	}
	for i := 10; i <= x*10; i *= 10 {
		result = result*10 + x%i/(i/10)
	}
	if isMinus {
		result = 0 - result
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}
