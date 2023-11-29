package main

import (
	"fmt"

	"github.com/finecoding/oxopfork/internal/oxocli"
)

func main() {

	board := oxocli.Grid([]byte("X........"))
	fmt.Println(board)
}
