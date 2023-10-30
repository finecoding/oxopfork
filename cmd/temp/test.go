package main

import (
	"fmt"
)

type grid []byte

func main() {

	var board grid

	board = []byte("OO XXX OO")
	fmt.Printf("%s", board)
	a, b := board.InPlay()
	fmt.Printf("%v, %v", a, b)

	board = []byte("OO0XXXXOO")
	fmt.Printf("%s", board)
	a, b = board.InPlay()

	fmt.Printf("%v, %v", a, b)

}

//String prints the OXO board to the terminal
func (board grid) String() string {

	l1 := string(board[0]) + string(board[1]) + string(board[2])
	l2 := string(board[3]) + string(board[4]) + string(board[5])
	l3 := string(board[6]) + string(board[7]) + string(board[8])
	return fmt.Sprintf("\n%s\n%s\n%s\n\n", l1, l2, l3)
}

//InPlay takes a board byte slice and returns a byte slice of empty positions
//It also returns true if there are positions that contain a space indicating that the game is still in play
func (board grid) InPlay() ([]byte, bool) {
	var spc []byte
	for k, v := range board {
		if v == 32 {
			spc = append(spc, byte(k))
		}
	}
	return spc, len(spc) != 0
}
