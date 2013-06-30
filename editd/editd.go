package editd

type cell struct {
	cost   int
	parent action
}

func (c *cell) set(cost int, parent action) {
    c.cost = cost
    c.parent = parent
}

type action string

const (
	Substitution action = "Substitution"
	Insertion    action = "Insertion"
	Deletion     action = "Deletion"
)

func EditDistance(s1, s2 string) int {
	length := max(len(s1), len(s2))
	cells := makeCells(length + 1)
	return editd(" "+s1, " "+s2, cells)
}

func editd(s1, s2 string, cells [][]cell) int {
    for x := 0; x < len(s1); x++ {
        cells[x][0].set(x, Deletion)
    }

    for x := 0; x < len(s2); x++ {
        cells[0][x].set(x, Insertion)
    }
	
	for i := 1; i < len(s1); i++ {
		for j := 1; j < len(s2); j++ {
			sub := cells[i-1][j-1].cost + subCost(rune(s1[i]), rune(s2[j]))
			ins := cells[i][j-1].cost + insCost(rune(s2[j]))
			del := cells[i-1][j].cost + delCost(rune(s1[i]))

			if sub <= ins && sub <= del {
			    cells[i][j].set(sub, Substitution)
			} else if ins <= sub && ins <= del {
			    cells[i][j].set(ins, Insertion)
			} else {
			    cells[i][j].set(del, Deletion)
			}
		}
	}

	return cells[len(s1) - 1][len(s2) - 1].cost
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
