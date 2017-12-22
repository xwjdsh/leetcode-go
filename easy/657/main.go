package main

func judgeCircle(moves string) bool {
	up, left := 0, 0
	for i := 0; i < len(moves); i++ {
		switch moves[i] {
		case 'U':
			up++
		case 'D':
			up--
		case 'L':
			left++
		case 'R':
			left--
		default:
			return false
		}
	}
	return up == 0 && left == 0
}
