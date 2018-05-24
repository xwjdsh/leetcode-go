package main

func lengthOfLongestSubstring(s string) int {
	m := map[rune]bool{}
	max := 0
	start := 0
	for _, si := range s {
		for m[si] {
			delete(m, rune(s[start]))
			start++
		}
		m[si] = true
		if len(m) > max {
			max = len(m)
		}
	}
	return max
}
