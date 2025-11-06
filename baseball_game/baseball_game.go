package baseball_game

import "strconv"

func calPoints(operations []string) int {
	res := 0
	stack := make([]int, 0, len(operations))

	for _, op := range operations {
		switch op {
		case "+":
			val1 := stack[len(stack)-1]
			val2 := stack[len(stack)-2]
			sum := val1 + val2
			stack = append(stack, sum)
			res += sum
		case "D":
			val := stack[len(stack)-1]
			sum := val * 2
			stack = append(stack, sum)
			res += sum
		case "C":
			res -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		default:
			val, err := strconv.Atoi(op)
			if err != nil {
				panic(err)
			}
			stack = append(stack, val)
			res += val
		}
	}

	return res
}
