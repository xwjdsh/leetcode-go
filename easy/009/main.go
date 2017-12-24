package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	l := 1
	for ; x/l >= 10; l *= 10 {
	}
	for x > 0 {
		if x/l != x%10 {
			return false
		}
		x %= l
		x /= 10

		l /= 100
	}
	return true
}
