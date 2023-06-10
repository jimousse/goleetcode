package main

func IsValidParentheses(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	start := map[rune]bool{
		'(': true,
		'[': true,
		'{': true,
	}
	stack := []rune{}
	var last rune

	for _, v := range s {
		_, isStart := start[v]
		if isStart {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}
			//pop the first on the stack
			last, stack = stack[len(stack)-1], stack[:len(stack)-1]
			if last == '(' && v != ')' ||
				last == '[' && v != ']' ||
				last == '{' && v != '}' {
				return false
			}
		}
	}
	return len(stack) == 0
}
