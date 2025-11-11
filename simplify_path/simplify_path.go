package simplify_path

import (
	"strings"
)

// using strings.Split
func simplifyPath(path string) string {
	s := make([]string, 0)

	for _, str := range strings.Split(path, "/") {
		if str == "" {
			continue
		}

		if str == ".." {
			if len(s) > 0 {
				s = s[:len(s)-1]
			}
		} else if str != "." {
			s = append(s, str)
		}
	}

	return "/" + strings.Join(s, "/")
}

// iterate string manually
//func simplifyPath(path string) string {
//	s := make([]string, 0)
//	i := 0
//
//	for path != "" {
//		for i < len(path) && path[i] == '/' {
//			path = path[1:]
//		}
//
//		j := i
//		for j < len(path) && path[j] != '/' {
//			j++
//		}
//
//		subStr := path[i:j]
//
//		if subStr == ".." {
//			if len(s) > 0 {
//				s = s[:len(s)-1]
//			}
//		} else if subStr != "" && subStr != "." {
//			s = append(s, subStr)
//		}
//
//		path = path[j:]
//	}
//
//	return "/" + strings.Join(s, "/")
//}
