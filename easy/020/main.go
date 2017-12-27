package main

func isValid(s string) bool {
	var expect []byte
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			expect = append(expect, ')')
		case '{', '[':
			expect = append(expect, s[i]+2)
		default:
			l := len(expect)
			if len(expect) > 0 && expect[l-1] == s[i] {
				expect = expect[:l-1]
			} else {
				return false
			}
		}
	}
	return len(expect) == 0
}
