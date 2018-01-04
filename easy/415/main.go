package main

func addStrings(num1 string, num2 string) string {
	result := ""
	for i1, i2, delta := len(num1)-1, len(num2)-1, 0; i1 >= 0 || i2 >= 0 || delta > 0; i1, i2 = i1-1, i2-1 {
		var j1, j2 int
		if i1 >= 0 {
			j1 = int(num1[i1] - '0')
		}
		if i2 >= 0 {
			j2 = int(num2[i2] - '0')
		}
		sum := j1 + j2 + delta
		delta = sum / 10
		result = string(sum%10+'0') + result
	}
	return result
}
