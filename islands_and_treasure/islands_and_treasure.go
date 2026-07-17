package islands_and_treasure

const INF = 2147483647

type Coords struct {
	row int
	col int
}

// BFS once, with all treasure tiles in queue initially
func islandsAndTreasure(grid [][]int) {
	row := len(grid)
	col := len(grid[0])

	visited := make(map[Coords]bool)

	q := make([]Coords, 0)

	for r := range row {
		for c := range col {
			if grid[r][c] == 0 {
				q = append(q, Coords{r, c})
				visited[Coords{r, c}] = true
			}
		}
	}

	var bfs func()
	bfs = func() {
		directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

		distance := 1
		for len(q) != 0 {
			qLen := len(q)
			for qLen != 0 {
				cur := q[0]
				q = q[1:]

				for _, d := range directions {
					r := cur.row + d[0]
					c := cur.col + d[1]

					if r < 0 || c < 0 || r >= row || c >= col || grid[r][c] != INF {
						continue
					}

					if _, exists := visited[Coords{r, c}]; !exists {
						grid[r][c] = distance
						visited[Coords{r, c}] = true
						q = append(q, Coords{r, c})
					}
				}

				qLen--
			}
			distance++
		}
	}

	bfs()
}

//// without hash map
//func islandsAndTreasure(grid [][]int) {
//	m, n := len(grid), len(grid[0])
//	q := [][2]int{}
//
//	for r := range m {
//		for c := range n {
//			if grid[r][c] == 0 {
//				q = append(q, [2]int{r, c})
//			}
//		}
//	}
//
//	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
//
//	for len(q) != 0 {
//		cur := q[0]
//		q = q[1:]
//
//		for _, d := range dirs {
//			r, c := cur[0]+d[0], cur[1]+d[1]
//
//			if r < 0 || c < 0 || r >= m || c >= n || grid[r][c] != INF {
//				continue
//			}
//
//			q = append(q, [2]int{r, c})
//			grid[r][c] = grid[cur[0]][cur[1]] + 1
//		}
//	}
//}

//// BFS on every node 🔥🔥🔥
//func islandsAndTreasure(grid [][]int) {
//	row := len(grid)
//	col := len(grid[0])
//
//	var bfs func(coords Coords)
//	bfs = func(coords Coords) {
//		directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
//		visited := make(map[Coords]bool)
//		visited[coords] = true
//
//		q := make([]Coords, 0)
//		q = append(q, coords)
//
//		distance := 1
//		for len(q) != 0 {
//			qLen := len(q)
//
//			for qLen != 0 {
//				cur := q[0]
//				q = q[1:]
//
//				for _, d := range directions {
//					r := cur.row + d[0]
//					c := cur.col + d[1]
//
//					if r < 0 || c < 0 || r >= row || c >= col || grid[r][c] == -1 {
//						continue
//					}
//
//					if grid[r][c] == 0 {
//						grid[coords.row][coords.col] = distance
//						return
//					}
//
//					if _, exists := visited[Coords{r, c}]; !exists {
//						visited[Coords{r, c}] = true
//						q = append(q, Coords{r, c})
//
//					}
//				}
//
//				qLen--
//			}
//
//			distance++
//		}
//	}
//
//	for r := range row {
//		for c := range col {
//			if grid[r][c] == INF {
//				bfs(Coords{r, c})
//			}
//		}
//	}
//}
