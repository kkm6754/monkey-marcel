package leetcode150

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, s := range tokens {
		val, err := strconv.Atoi(s)
		if err == nil { // 转换成数字成功
			stack = append(stack, val)
		} else { // 转换数字失败，
			num1, num2 := stack[len(stack) - 2], stack[len(stack) - 1] // 注意顺序，栈顶在算数运算符后
			stack = stack[: len(stack) - 2]
			switch s {
			case "+":
				stack = append(stack, num1 + num2)
			case "-":
				stack = append(stack, num1 - num2)
			case "*":
				stack = append(stack, num1 * num2)
			case "/":
				stack = append(stack, num1 / num2)
			}
		}
	}
	return stack[0]
}