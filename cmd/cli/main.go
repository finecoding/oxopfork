package main

import (
	"fmt"

	"github.com/finecoding/oxopfork/internal/oxocli"
)

func main() {

	// This is a quick test to see if the internal import is working
	// oxocli package contains types.go which defines Grid as [9]byte
	b := oxocli.Grid([]byte("X........"))
	fmt.Println("This is a board", b)
	t := oxocli.Turn{
		Board:  b,
		Status: "XWIN"}
	fmt.Println("This is a turn", t)
	g := oxocli.Game{t, t, t, t, t}
	fmt.Println("This is a game", g)
}
