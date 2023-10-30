package main

import (
	"fmt"
)

// board represents the state of a turn of the game on a 3 x 3 grid
type Grid [9]byte

// a turn has a board and status string which will be one of these strings XWIN,OWIN,DRAW,PLAY
type Turn struct {
	board  Grid
	status string
}
type Game []Turn

// This makes the board type conform to the Stringer interface so it can be used with fmt.

func (board Grid) String() string {

	l1 := string(board[0]) + string(board[1]) + string(board[2])
	l2 := string(board[3]) + string(board[4]) + string(board[5])
	l3 := string(board[6]) + string(board[7]) + string(board[8])
	return fmt.Sprintf("\n%s\n%s\n%s\n\n", l1, l2, l3)
}

// This makes the game type conform to the Stringer interface so it can be used with fmt.
// It loops over the game slice building three strings from the contents of each board.
func (g Game) String() string {
	var l1, l2, l3, l4 string

	for i, _ := range g {
		l1 = l1 + string(g[i].board[0]) + string(g[i].board[1]) + string(g[i].board[2]) + "    "
		l2 = l2 + string(g[i].board[3]) + string(g[i].board[4]) + string(g[i].board[5]) + "    "
		l3 = l3 + string(g[i].board[6]) + string(g[i].board[7]) + string(g[i].board[8]) + "    "
		l4 = l4 + g[i].status + "   "
	}
	return fmt.Sprintf("\n%s\n%s\n%s\n\n%s\n", l1, l2, l3, l4)
}

func main() {

	// Here we are building a set of literally defined turns.

	var t1 = Turn{
		board:  Grid([]byte("X........")),
		status: "PLAY",
	}
	var t2 = Turn{
		board:  Grid([]byte("X.O......")),
		status: "PLAY",
	}
	var t3 = Turn{
		board:  Grid([]byte("X.OX.....")),
		status: "PLAY",
	}
	var t4 = Turn{
		board:  Grid([]byte("X.OXO....")),
		status: "PLAY",
	}
	var t5 = Turn{
		board:  Grid([]byte("X.OXO.X..")),
		status: "XWIN",
	}
	var g Game

	// Here we append all the turns into a game slice.

	g = append(g, t1, t2, t3, t4, t5)

	// Thanks to the String method, it is easy to print out the game
	fmt.Println(g)

	// So far we have used literals to build a game slice.  Good for testing the String method
	// The next development will be to get the game loop working to build the game slice
	// go run oxo2.go

	// This is what it prints:

	//  X..    X.O    X.O    X.O    X.O
	//  ...    ...    X..    XO.    XO.
	//  ...    ...    ...    ...    X..
	//
	//  PLAY   PLAY   PLAY   PLAY   XWIN
}
