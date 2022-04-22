package leetcode20


func isValid(s string) bool {
	ha := map[byte]byte{')':'(' , ']':'[', '}':'{'}
	stack := make([]byte, 0)
	if s == "" {
		return false
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 && stack[len(stack) - 1] == ha[s[i]] {
			stack = stack[: len(stack) - 1]
		} else {
			return false
		}
	}

	// 左括号可能有多余情况
	return len(stack) == 0
}