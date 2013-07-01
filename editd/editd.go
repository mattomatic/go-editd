package editd

type cell struct {
	cost   int
	parent action
	letter rune
}

func (c *cell) set(cost int, parent action, letter rune) {
	c.cost = cost
	c.parent = parent
	c.letter = letter
}

type action string

const (
	Substitute action = "Substitute"
	Insert     action = "Insert"
	Delete     action = "Delete"
)

func EditDistance(s1, s2 string) int {
	length := max(len(s1), len(s2))
	cells := makeCells(length + 1)
	return editd(" "+s1, " "+s2, cells)
}

func editd(s1, s2 string, cells [][]cell) int {
	for x := 0; x < len(s1); x++ {
		cells[x][0].set(x, Delete, rune(s1[x]))
	}

	for x := 0; x < len(s2); x++ {
		cells[0][x].set(x, Insert, rune(s2[x]))
	}

	for r := 1; r < len(s1); r++ {
		for c := 1; c < len(s2); c++ {
			sub := cells[r-1][c-1].cost + subCost(rune(s1[r]), rune(s2[c]))
			ins := cells[r][c-1].cost + insCost(rune(s2[c]))
			del := cells[r-1][c].cost + delCost(rune(s1[r]))

			if sub <= ins && sub <= del {
				cells[r][c].set(sub, Substitute, rune(s1[r]))
			} else if ins <= sub && ins <= del {
				cells[r][c].set(ins, Insert, rune(s2[c]))
			} else {
				cells[r][c].set(del, Delete, rune(s1[r]))
			}
		}
	}

	return cells[len(s1)-1][len(s2)-1].cost
}

func makeCells(length int) [][]cell {
	cells := make([][]cell, length)

	for i, _ := range cells {
		cells[i] = make([]cell, length)
	}

	return cells
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func insCost(r rune) int {
	return 1
}

func delCost(r rune) int {
	return 1
}

func subCost(r1, r2 rune) int {
	if r1 == r2 {
		return 0
	}

	return 1
}
