package open_the_lock

import (
	"strconv"
)

func openLock(deadends []string, target string) int {
	deadendsMap := make(map[string]bool)
	for _, d := range deadends {
		deadendsMap[d] = true
	}

	q := make([]string, 0)
	visited := make(map[string]bool)
	if !deadendsMap["0000"] {
		q = append(q, "0000")
		visited["0000"] = true
	}

	res := 0
	for len(q) != 0 {
		qLen := len(q)

		for qLen != 0 {
			cur := q[0]
			q = q[1:]

			if cur == target {
				return res
			}

			for i := range len(cur) {
				c, _ := strconv.Atoi(string(cur[i]))
				upGuess := cur[0:i] + strconv.Itoa((c+1)%10) + cur[i+1:]
				downGuess := cur[0:i] + strconv.Itoa((c-1+10)%10) + cur[i+1:]

				if !visited[upGuess] && !deadendsMap[upGuess] {
					q = append(q, upGuess)
					visited[upGuess] = true
				}

				if !visited[downGuess] && !deadendsMap[downGuess] {
					q = append(q, downGuess)
					visited[downGuess] = true
				}
			}

			qLen--
		}

		res += 1
	}

	return -1
}
