package main

var validParentheses = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, bracket := range s {
		if _, ok := validParentheses[bracket]; ok {
			stack = append(stack, bracket)
		} else {
			if len(stack) == 0 {
				return false
			}
			open := stack[len(stack)-1]
			if validParentheses[open] == bracket {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
