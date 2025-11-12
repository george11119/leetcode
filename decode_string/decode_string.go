package decode_string

import (
	"strconv"
	"strings"
	"unicode"
)

// stack
func decodeString(s string) string {
	n := len(s)
	st := make([]string, 0)

	i := 0
	for i < n {
		c := rune(s[i])

		if unicode.IsDigit(c) {
			j := i + 1
			for s[j] != '[' {
				j++
			}
			num := s[i:j]
			st = append(st, num)
			i = j
		} else if c == ']' {
			subStr := st[len(st)-1]
			st = st[:len(st)-1]

			num, _ := strconv.Atoi(st[len(st)-1])
			st = st[:len(st)-1]

			subStr = strings.Repeat(subStr, num)
			for len(st) > 0 {
				if _, err := strconv.Atoi(st[len(st)-1]); err == nil {
					break
				}
				subStr = st[len(st)-1] + subStr
				st = st[:len(st)-1]
			}
			st = append(st, subStr)
		} else {
			subStr := string(c)
			for len(st) > 0 {
				if _, err := strconv.Atoi(st[len(st)-1]); err == nil {
					break
				}
				subStr = st[len(st)-1] + subStr
				st = st[:len(st)-1]
			}
			st = append(st, subStr)
		}
		i++
	}

	return strings.Join(st, "")
}

//func helper(s string, i int) (string, int) {
//	res := ""
//	k := 0
//
//	for i < len(s) {
//		c := s[i]
//		if unicode.IsDigit(rune(c)) {
//			k = k*10 + int(c-48)
//		} else if c == '[' {
//			i++
//			subStr, j := helper(s, i)
//			i = j
//			res += strings.Repeat(subStr, k)
//			k = 0
//		} else if c == ']' {
//			return res, i
//		} else {
//			res += string(c)
//		}
//		i++
//	}
//
//	return res, i
//}
//
//// recursion
//func decodeString(s string) string {
//	res, _ := helper(s, 0)
//	return res
//}
