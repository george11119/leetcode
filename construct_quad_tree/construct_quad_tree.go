package construct_quad_tree

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	return createNode(grid, 0, len(grid), 0, len(grid))
}

func createNode(grid [][]int, wStart, wEnd, hStart, hEnd int) *Node {
	hasOne := false
	hasZero := false
	for i := wStart; i < wEnd; i++ {
		for j := hStart; j < hEnd; j++ {
			if grid[i][j] == 1 {
				hasOne = true
			} else if grid[i][j] == 0 {
				hasZero = true
			}
		}
	}

	// grid is either all ones or all zeros
	if hasOne && !hasZero {
		return &Node{
			Val:    true,
			IsLeaf: true,
		}
	}

	if !hasOne && hasZero {
		return &Node{
			Val:    false,
			IsLeaf: true,
		}
	}

	if wEnd-wStart == 1 && hEnd-hStart == 1 {
		return &Node{
			Val:    grid[hStart][wStart] == 1,
			IsLeaf: true,
		}
	}

	wMid := (wStart + wEnd) / 2
	hMid := (hStart + hEnd) / 2

	return &Node{
		IsLeaf:      false,
		TopLeft:     createNode(grid, wStart, wMid, hStart, hMid),
		BottomLeft:  createNode(grid, wMid, wEnd, hStart, hMid),
		TopRight:    createNode(grid, wStart, wMid, hMid, hEnd),
		BottomRight: createNode(grid, wMid, wEnd, hMid, hEnd),
	}
}
