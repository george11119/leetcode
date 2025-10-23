package encoding_and_decoding

import (
	"fmt"
	"strconv"
)

func Encode(strs []string) string {
	res := ""

	for _, s := range strs {
		separator := strconv.Itoa(len(s)) + "#"
		res += separator + s
	}

	fmt.Println(res)
	return res
}

func Decode(s string) []string {
	sLen := len(s)
	res := make([]string, 0)
	i := 0

	for i < sLen {
		numStr := ""
		for s[i] != '#' {
			numStr += string(s[i])
			i++
		}

		i++
		n, err := strconv.Atoi(numStr)
		if err != nil {
			return []string{}
		}
		res = append(res, s[i:i+n])
		i += n
	}

	return res
}
