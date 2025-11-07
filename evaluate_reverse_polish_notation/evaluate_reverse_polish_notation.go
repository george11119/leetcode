package evaluate_reverse_polish_notation

import "strconv"

func isOperator(op string) bool {
	return op == "*" || op == "+" || op == "-" || op == "/"
}

func evalRPN(tokens []string) int {
	stack := make([]string, 0)

	for _, op := range tokens {
		if isOperator(op) {
			val2, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			val1, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			switch op {
			case "+":
				stack = append(stack, strconv.Itoa(val1+val2))
			case "-":
				stack = append(stack, strconv.Itoa(val1-val2))
			case "/":
				stack = append(stack, strconv.Itoa(val1/val2))
			case "*":
				stack = append(stack, strconv.Itoa(val1*val2))
			default:
			}
		} else {
			stack = append(stack, op)
		}
	}

	res, _ := strconv.Atoi(stack[len(stack)-1])
	return res
}
