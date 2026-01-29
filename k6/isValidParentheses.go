package k6

//"()[]"     → true
//"({[]})"   → true
//"(]"       → false
//"([)]"     → false
//"("        → false

func isValidParentheses(s string) bool {
	stack := make([]rune, 0)

	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, r := range s {
		if r == '(' || r == '[' || r == '{' {
			stack = append(stack, r)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top != pairs[r] {
			return false
		}
	}

	return len(stack) == 0
}
