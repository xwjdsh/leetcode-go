package main

func lengthOfLastWord(s string) int {
	var length int
	for i, _ := range s {
		if s[i] == ' ' {
			if len(s) > i+1 && s[i+1] != ' ' {
				length = 0
			}
		} else {
			length++
		}
	}
	return length
}
