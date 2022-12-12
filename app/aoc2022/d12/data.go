package d12

type Grid struct {
	rows []string
}

func (g *Grid) Get(row int, col int) string {
	return g.rows[row][col : col+1]
}
