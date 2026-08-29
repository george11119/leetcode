package course_schedule_iv

import (
	"maps"
)

// super slow DFS on every node solution
//func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
//	graph := make(map[int][]int)
//	for _, e := range prerequisites {
//		u := e[0]
//		v := e[1]
//
//		graph[u] = append(graph[u], v)
//	}
//
//	visited := make(map[int]bool)
//
//	var dfs func(i, target int) bool
//	dfs = func(i, target int) bool {
//		if i == target {
//			return true
//		}
//
//		if visited[i] {
//			return false
//		}
//
//		for _, j := range graph[i] {
//			if dfs(j, target) {
//				return true
//			}
//		}
//
//		return false
//	}
//
//	res := make([]bool, 0)
//	for _, q := range queries {
//		res = append(res, dfs(q[0], q[1]))
//	}
//
//	return res
//}

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	graph := make(map[int][]int)
	for _, e := range prerequisites {
		u := e[0]
		v := e[1]

		graph[u] = append(graph[u], v)
	}

	canReach := make(map[int]map[int]bool)

	var dfs func(i int) map[int]bool
	dfs = func(i int) map[int]bool {
		if c, ok := canReach[i]; ok {
			return c
		}

		dst := make(map[int]bool)
		for _, j := range graph[i] {
			dst[j] = true
			maps.Insert(dst, maps.All(dfs(j)))
		}
		canReach[i] = dst
		return dst
	}

	for i := range numCourses {
		dfs(i)
	}

	res := make([]bool, 0, len(queries))
	for _, q := range queries {
		res = append(res, canReach[q[0]][q[1]])
	}
	return res
}
