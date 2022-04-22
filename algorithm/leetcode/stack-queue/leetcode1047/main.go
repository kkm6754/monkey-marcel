package main

func removeDuplicates(s string) string {
	if len(s) < 2 {
		return s
	}

	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack) - 1] == s[i] {
			stack = stack[: len(stack) - 1]
			continue
		}
		stack = append(stack, s[i])
	}

	return string(stack)
}

