package oxocli

import (
	"fmt"
)

func (b Grid) String() string {

	l1 := string(b[0]) + string(b[1]) + string(b[2])
	l2 := string(b[3]) + string(b[4]) + string(b[5])
	l3 := string(b[6]) + string(b[7]) + string(b[8])
	return fmt.Sprintf("\n%s\n%s\n%s\n\n", l1, l2, l3)
}
func (t Turn) String() string {

	l1 := string(t.Board[0]) + string(t.Board[1]) + string(t.Board[2])
	l2 := string(t.Board[3]) + string(t.Board[4]) + string(t.Board[5])
	l3 := string(t.Board[6]) + string(t.Board[7]) + string(t.Board[8])
	l4 := t.Status
	return fmt.Sprintf("\n%s\n%s\n%s\n%s\n\n", l1, l2, l3, l4)
}

func (g Game) String() string {
	var l1, l2, l3, l4 string

	for i, _ := range g {
		l1 = l1 + string(g[i].Board[0]) + string(g[i].Board[1]) + string(g[i].Board[2]) + "    "
		l2 = l2 + string(g[i].Board[3]) + string(g[i].Board[4]) + string(g[i].Board[5]) + "    "
		l3 = l3 + string(g[i].Board[6]) + string(g[i].Board[7]) + string(g[i].Board[8]) + "    "
		l4 = l4 + g[i].Status + "   "
	}
	return fmt.Sprintf("\n%s\n%s\n%s\n\n%s\n", l1, l2, l3, l4)
}
